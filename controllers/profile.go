package controllers

import (
	"code.unitiwireless.com/uniti-wireless/harepd/models"
)

//TryMaster !!!This has been deprecated no longer used
func TryMaster(conf *models.Config) (*models.Master, error) {
	var master models.Master
	if err := models.DB.Table("pg_stat_replication").First(&master).Error; err != nil {
		Logs.Info(err)
		return nil, err
	}
	return &master, nil

}

//TrySlave !!!!!This has been deprecated no longer used
func TrySlave(conf *models.Config) (*models.Slave, error) {
	var slave models.Slave
	if err := models.DB.Table("pg_stat_wal_receiver").First(&slave).Error; err != nil {
		return nil, err
	}
	return &slave, nil
}
