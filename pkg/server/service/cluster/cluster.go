package cluster

import (
	"kubernetes_management_system/common"
	"kubernetes_management_system/models/cluster"
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
