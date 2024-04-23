package cluster

import "github.com/gaochuang/cloudManagementSystem/models"

type Cluster struct {
	models.Mode
	ClusterName    string `json:"clusterName" gorm:"comment:clusterName" binding:"required"`
	KubeConfig     string `json:"kubeConfig" gorm:"comment:clusterCertificate;type:varchar(15360)" binding:"required"`
	NodeCount      int    `json:"nodeCount" gorm:"comment:cluster node counts"`
	ClusterVersion string `json:"clusterVersion" gorm:"cluster version"`
}

func (c *Cluster) TableName() string {
	var k models.Mode
	return k.TableName("k8s_cluster")
}

type PaginationQ struct {
	Size    int    `form:"size" json:"size"`
	Page    int    `form:"page" json:"page"`
	Total   int64  `json:"total"`
	Keyword string `form:"keyword" json:"keyword"`
}

type ClusterIds struct {
	Data interface{} `json:"clusterIds"`
}

type NodeStatus struct {
	NodeCount       int     `json:"node_count"`
	Ready           int     `json:"ready"`
	UnReady         int     `json:"unready"`
	Namespace       int     `json:"namespace"`
	Deployment      int     `json:"deployment"`
	Pod             int     `json:"pod"`
	CpuUsage        float64 `json:"cpu_usage" desc:"cpu usage"`
	CpuCore         float64 `json:"cpu_core"`
	CpuCapacityCore float64 `json:"cpu_capacity_core"`
	MemoryUsage     float64 `json:"memory_usage" desc:"memory usage"`
	MemoryUsed      float64 `json:"memory_used"`
	MemoryTotal     float64 `json:"memory_total"`
}
