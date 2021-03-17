package main

import (
	"fmt"
	"strings"

	"code.unitiwireless.com/uniti-wireless/harepd/controllers"
	"code.unitiwireless.com/uniti-wireless/harepd/models"
)

func logic(conf *models.Config) (*models.IKnow, error) {
	for _, n := range conf.Harepd.Grpc.Neighbours {
		Logs.Info("Connecting to Neighbour: ", n)
		aware, err := controllers.GrpcClient(n, conf)
		if err != nil {
			Logs.Error(fmt.Sprintf("Could not connected to %s. Consulting witness at %s", n, conf.Harepd.Grpc.Witness))
			Logs.Error(err)
		} else {
			/*
				In case Slave required to handle RO transaction
			*/
			if models.TestAwareness(aware, conf) {
				Logs.Info(fmt.Sprintf("Everything is synced according to the neighbour: %s", n))
				for _, v := range aware.Master {
					//IF Master treat R and W transaction
					if strings.TrimSpace(v) == strings.TrimSpace(conf.Harepd.PrimaryIP) {
						/*
							If a node should be a master and agrees with their immediate neighbour they may allow read/write traffic
						*/
						err := controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Allow, conf.Harepd.AuthModes.Allow)
						if err != nil {
							Logs.Error(err)
						}
						// Break if this node is a master
						break
					} else {
						/*
							If a node should be a slave and agrees with the witness they may allow readonly traffic
						*/
						err := controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Deny, conf.Harepd.AuthModes.Allow)
						if err != nil {
							Logs.Error(err)
						}
					}
				}
				return aware, nil
			}
			Logs.Error(fmt.Sprintf("This node is not synced with: %s. Consulting witness at %s", n, conf.Harepd.Grpc.Witness))
			break
		}
	}

	// Connecting to the witness
	Logs.Info("Connecting to Witness: ", conf.Harepd.Grpc.Witness)
	aware, err := controllers.GrpcClient(conf.Harepd.Grpc.Witness, conf)

	if err != nil || aware == nil {
		/*
			If the neighbour nor witness unreachable node is unhealthy
		*/
		Logs.Error(fmt.Sprintf("Could not connected to witness at %s", conf.Harepd.Grpc.Witness))
		err = controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Deny, conf.Harepd.AuthModes.Deny)
		if err != nil {
			Logs.Error(err)
		}
		return nil, err
	}
	if models.TestAwareness(aware, conf) {
		for _, v := range aware.Master {
			if strings.TrimSpace(v) == strings.TrimSpace(conf.Harepd.PrimaryIP) {
				/*
					If a node should be a master and agrees with the witness they may allow read/write traffic
				*/
				err = controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Allow, conf.Harepd.AuthModes.Allow)
				if err != nil {
					Logs.Error(err)
				}
				// Break if this node is a master
				break
			} else {
				/*
					If a node should be a slave and agrees with the witness they may allow readonly traffic
				*/
				err := controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Deny, conf.Harepd.AuthModes.Allow)
				if err != nil {
					Logs.Error(err)
				}
			}
		}
		Logs.Info(fmt.Sprintf("Everything is synced according to the witness at %s", conf.Harepd.Grpc.Witness))
		return aware, nil
	}
	Logs.Error("This node is not synced.")
	/*
		If a node with node synced with either their immediate neighbour nor witness node should consider as their unhealthy
	*/
	err = controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Deny, conf.Harepd.AuthModes.Deny)
	if err != nil {
		Logs.Error(err)
	}
	return nil, nil
}
