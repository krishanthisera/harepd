package controllers

import (
	"fmt"

	"code.unitiwireless.com/uniti-wireless/harepd/models"
)

/*AlterRuleLegacy is used to alter pg_hba rule withouting accessing the DB.
It is must to ensure that the correct path has been configured
*/
func AlterRuleLegacy(conf *models.Config, ruleRW, ruleRO string) error {

	address := conf.Harepd.Haproxy.Server

	// Delete Old Rules for HAProxy
	for _, addr := range address {
		delErr := models.DeleteLines(conf.Harepd.HbaConfig, addr)
		if delErr != nil {
			Logs.Error(delErr)
			return delErr
		}
		// Inject new rules RW
		injErr := models.InjectLines(conf.Harepd.HbaConfig, fmt.Sprintf("\nhost    %s    %s    %s/32    %s", conf.Harepd.Repmgr.Db, conf.Harepd.Haproxy.Users.ReadWrite, addr, ruleRW))
		if injErr != nil {
			Logs.Error(injErr)
			return injErr
		}
		// If the traffic should send to the slave
		// Inject new rules RO
		if conf.Harepd.AllowRO {
			injErr := models.InjectLines(conf.Harepd.HbaConfig, fmt.Sprintf("\nhost    %s    %s    %s/32    %s", conf.Harepd.Repmgr.Db, conf.Harepd.Haproxy.Users.ReadOnly, addr, ruleRO))
			if injErr != nil {
				Logs.Error(injErr)
				return injErr
			}
		}
	}
	// Tied Up the File
	if err := models.TideUp(conf.Harepd.HbaConfig); err != nil {
		Logs.Error(err)
		return err
	}
	// reload the config
	if err := models.DB.Exec("SELECT pg_reload_conf();").Error; err != nil {
		Logs.Error(err)
		return err
	}

	return nil
}
