package cluster

import (
	"kubernetes_management_system/common"
	"kubernetes_management_system/models/cluster"
)

func StorageCluster(cluster cluster.Cluster) error {
	err := common.DB.Create(&cluster).Error
	return err
}
