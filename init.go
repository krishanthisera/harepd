package main

import (
	"encoding/json"
	"fmt"
	"time"

	"code.unitiwireless.com/uniti-wireless/harepd/controllers"
	"code.unitiwireless.com/uniti-wireless/harepd/models"
)

func serverOnlyMode(conf *models.Config) {
	Logs.Info("Server Only Mode.")
	err := controllers.GrpcServer(*conf)
	if err != nil {
		Logs.Fatal(err)
		panic(err)
	}
}
func clientOnlyMode(conf *models.Config) {
	Logs.Info("Client Only Mode.")
	dbg, err := logic(conf)
	if err != nil {
		Logs.Error(err)
	}
	Logs.Debug(json.Marshal(dbg))
	Logs.Debug(json.Marshal(conf))
	Logs.Info(conf)
}

func dualMode(conf *models.Config) {
	Logs.Info("Dual mode Activated.")
	go controllers.GrpcServer(*conf)

	Logs.Info(fmt.Sprintf("Client will be Activated: in every %ds", conf.Harepd.WatchDog))
	for range time.Tick(time.Second * time.Duration(conf.Harepd.WatchDog)) {
		dbg, err := logic(conf)
		if err != nil {
			Logs.Error(err)
		}

		Logs.Debug(json.Marshal(dbg))
		Logs.Debug(json.Marshal(conf))
	}
}

func dbInit(conf *models.Config) {
	if err := models.DbConnect(conf); err != nil {
		Logs.Error(err)
		/*
			If the DB initiation is failed due to any reason node is considered as unhealthy
		*/
		err = controllers.AlterRuleLegacy(conf, conf.Harepd.AuthModes.Deny, conf.Harepd.AuthModes.Deny)
		if err != nil {
			Logs.Error(err)
		}
	}
}
