// This will alter the pg_hba.conf
package controllers

import (
	"fmt"

	"code.unitiwireless.com/uniti-wireless/harepd/models"
)

func AlterRule(rules *models.Rules, conf *models.Config) error {

	address := conf.Harepd.Haproxy.Server
	tmpTable := fmt.Sprintf("harepd_%s", conf.Harepd.NodeName)

	// Create a temporary table
	if err := models.DB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", tmpTable)).Error; err != nil {
		return err
	}
	if err := models.DB.Exec(fmt.Sprintf("CREATE TABLE %s ( lines text );", tmpTable)).Error; err != nil {
		return err
	}

	// Copy pg_hba rules to the temporary table
	if err := models.DB.Exec(fmt.Sprintf("COPY %s FROM '%s';", tmpTable, conf.Harepd.HbaConfig)).Error; err != nil {
		return err
	}

	for _, addr := range address {
		// Delete OLD Rules
		if err := models.DB.Exec(fmt.Sprintf("DELETE FROM %s WHERE lines LIKE '%%%s%%';", tmpTable, addr)).Error; err != nil {
			return err
		}

		// Inject new rules RW
		if err := models.DB.Exec(fmt.Sprintf("insert into %s (lines) values ('host    %s    %s    %s/32    %s');", tmpTable, conf.Harepd.Repmgr.Db, conf.Harepd.Haproxy.Users.ReadWrite, addr, conf.Harepd.AuthModes.Deny)).Error; err != nil {
			return err
		}

		// If the traffic should send to the slave
		// Inject new rules RO
		if conf.Harepd.AllowRO {
			if err := models.DB.Exec(fmt.Sprintf("insert into %s (lines) values ('host    %s    %s    %s/32    %s');", tmpTable, conf.Harepd.Repmgr.Db, conf.Harepd.Haproxy.Users.ReadOnly, addr, conf.Harepd.AuthModes.Allow)).Error; err != nil {
				return err
			}
		}
	}
	// Copy table to the pg_hba
	if err := models.DB.Exec(fmt.Sprintf("COPY %s to '%s';", tmpTable, conf.Harepd.HbaConfig)).Error; err != nil {
		return err
	}

	// reload the config
	if err := models.DB.Exec("SELECT pg_reload_conf();").Error; err != nil {
		return err
	}

	return nil
}
