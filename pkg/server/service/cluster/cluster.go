package cluster

import (
	"context"
	"fmt"
	"github.com/prometheus/common/expfmt"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes_management_system/common"
	"kubernetes_management_system/models/cluster"
	"strings"
)

var (
	// cpuUsage
	kubePodContainerResourceRequestsCpuCores float64 = 0
	kubeNodeStatusCapacityCpuCores           float64 = 0
	// memoryUsage
	kubePodContainerResourceRequestsMemoryBytes float64 = 0
	kubeNodeStatusAllocatableMemoryBytes        float64 = 0
)

func StorageCluster(cluster cluster.Cluster) error {
	err := common.DB.Create(&cluster).Error
	return err
}

func ListCluster(p *cluster.PaginationQ, k *[]cluster.Cluster) (err error) {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Size < 1 {
		p.Size = 5
	}

	offset := p.Size * (p.Page - 1)
	tx := common.DB
	if p.Keyword != "" {
		tx = common.DB.Where("cluster_name like ?", "%"+p.Keyword+"%").Limit(p.Size).Offset(offset).Find(&k)
	} else {
		tx = common.DB.Limit(p.Size).Offset(offset).Find(&k)

	}

	var total int64
	tx.Count(&total)

	p.Total = total
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func GetCluster(id uint) (cluster cluster.Cluster, err error) {
	if err := common.DB.Where("id = ?", id).First(&cluster).Error; err != nil {
		return cluster, err
	}
	return cluster, nil
}

func DeleteCluster(clusterId cluster.ClusterIds) error {
	var k cluster.Cluster
	if err := common.DB.Delete(&k, clusterId.Data).Error; err != nil {
		return err
	}
	return nil
}

func GetClusterInfo(c *kubernetes.Clientset) (*cluster.NodeStatus, error) {
	var node cluster.NodeStatus
	if status := getNodesRunningStatus(c, &node); status == nil {
		common.LOG.Error("get node running status failed")
		return nil, fmt.Errorf("server internal error")
	}

	data, err := c.RESTClient().Get().AbsPath("/api/v1/namespaces/kube-system/services/tke-kube-state-metrics:http-metrics/proxy/metrics").DoRaw(context.TODO())
	if err != nil {
		common.LOG.Error("get metrics failed", zap.Any("err: ", err))
		return nil, err
	}
	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(string(data)))
	if err != nil {
		common.LOG.Error("parser metrics failed", zap.Any("err:", err))
		return nil, err
	}

	for index, metricValue := range metricFamilies {
		if index == "kube_pod_container_resource_requests_cpu_cores" {
			for _, metric := range metricValue.GetMetric() {
				kubePodContainerResourceRequestsCpuCores += *metric.Gauge.Value
			}
		}
		if index == "kube_node_status_capacity_cpu_cores" {
			for _, metric := range metricValue.GetMetric() {
				kubeNodeStatusCapacityCpuCores += *metric.Gauge.Value
			}
		}
		if index == "kube_pod_container_resource_requests_memory_bytes" {
			for _, metric := range metricValue.GetMetric() {
				kubePodContainerResourceRequestsMemoryBytes += *metric.Gauge.Value
			}
		}
		if index == "kube_node_status_allocatable_memory_bytes" {
			for _, metric := range metricValue.GetMetric() {
				kubeNodeStatusAllocatableMemoryBytes += *metric.Gauge.Value
			}
		}
		if index == "kube_node_info" {
			node.NodeCount = len(metricValue.Metric)
		}
	}
	node.CpuCore = common.ParseFloat2F(kubePodContainerResourceRequestsCpuCores)
	node.CpuUsage = common.ParseFloat2F(kubePodContainerResourceRequestsCpuCores / kubeNodeStatusCapacityCpuCores * 100)
	node.CpuCapacityCore = common.ParseFloat2F(kubeNodeStatusCapacityCpuCores)

	node.MemoryUsed = common.ParseFloat2F(kubePodContainerResourceRequestsMemoryBytes / 1024 / 1024 / 1024)
	node.MemoryUsage = common.ParseFloat2F(kubePodContainerResourceRequestsMemoryBytes / kubeNodeStatusAllocatableMemoryBytes * 100)
	node.MemoryTotal = common.ParseFloat2F(kubeNodeStatusAllocatableMemoryBytes / 1024 / 1024 / 1024)
	return &node, nil
}

func getNodesRunningStatus(client *kubernetes.Clientset, status *cluster.NodeStatus) *cluster.NodeStatus {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		common.LOG.Error("get node failed", zap.Any("err: ", err))
		return nil
	}

	readyNode := 0
	notReadyNode := 0
	for _, node := range nodes.Items {
		listNode, _ := client.CoreV1().Nodes().Get(context.TODO(), node.ObjectMeta.Name, v1.GetOptions{})
		if len(listNode.Status.Conditions) >= 0 {
			if string(listNode.Status.Conditions[len(listNode.Status.Conditions)-1].Status) == "True" {
				readyNode++
			} else {
				notReadyNode++
			}
		} else {
			common.LOG.Error("get node status failed")
			return nil
		}
	}
	status.Ready = readyNode
	status.UnReady = notReadyNode
	return status
}
