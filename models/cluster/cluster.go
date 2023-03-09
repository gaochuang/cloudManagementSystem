package cluster

import "kubernetes_management_system/models"

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
