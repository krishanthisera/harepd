package models

import (
	"fmt"
	reflect "reflect"
	"sort"
)

// TestAwareness Test Whether everithing is synced
func TestAwareness(upstreamData *IKnow, conf *Config) bool {

	var localData IKnow

	// // SQL Queries
	// // Masters
	masters, err := DB.Raw("SELECT conninfo FROM repmgr.nodes WHERE type = ? AND active = ?", "primary", "t").Rows()
	defer masters.Close()
	if err != nil {
		Log.Error(err)
		return false
	}

	//Slaves
	slaves, err := DB.Raw("SELECT conninfo FROM repmgr.nodes WHERE type = ? AND active = ?", "standby", "t").Rows()
	defer slaves.Close()
	if err != nil {
		Log.Error(err)
		return false
	}

	// Master Node
	for masters.Next() {
		var r string
		err := masters.Scan(&r)
		if err != nil {
			Log.Error(err)
			return false
		}

		localData.Master = append(localData.Master, IpScrapper(r)[0])
	}

	// Slaves
	for slaves.Next() {
		var r string
		err := slaves.Scan(&r)
		if err != nil {
			Log.Error(err)
			return false
		}
		localData.Slaves = append(localData.Slaves, IpScrapper(r)[0])
	}

	// Sort the arrays
	sort.Strings(localData.Master)
	sort.Strings(localData.Slaves)

	sort.Strings(upstreamData.Master)
	sort.Strings(upstreamData.Slaves)

	if reflect.DeepEqual(localData.Master, upstreamData.Master) {
		if reflect.DeepEqual(localData.Slaves, upstreamData.Slaves) {
			Log.Info(err)
			return true
		}
		Log.Warn(fmt.Sprintf("Complication on the slave records: [%s]:[%s]", localData.Slaves, upstreamData.Slaves))
	}
	Log.Warn(fmt.Sprintf("Complication on the master records: [%s]:[%s]", localData.Master, upstreamData.Master))
	return false
}
