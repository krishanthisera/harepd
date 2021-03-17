package controllers

import (
	"context"
	"time"

	"code.unitiwireless.com/uniti-wireless/harepd/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//GrpcClient to get Info
func GrpcClient(neighbour string, conf *models.Config) (*models.IKnow, error) {

	var opts []grpc.DialOption
	if conf.Harepd.Grpc.TLS.Enabled {
		caFile := conf.Harepd.Grpc.TLS.Ca
		creds, err := credentials.NewClientTLSFromFile(caFile, "")
		if err != nil {

			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.FailOnNonTempDialError(true))
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(neighbour, opts...)

	if err != nil {

		return nil, err
	}
	defer conn.Close()
	client := models.NewClusterInfoClient(conn)

	// Get the awareness
	var myInfo models.WhatYouKnow
	err = fetchLocal(conf, &myInfo)
	if err != nil {

		return nil, err
	}
	info, err := GetInfo(client, &myInfo, conf)
	if err != nil {

		return nil, err
	}
	return info, nil
}

//GenInfo Dial the gRPC upstream server and ge the information
func GetInfo(client models.ClusterInfoClient, myinfo *models.WhatYouKnow, conf *models.Config) (*models.IKnow, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Harepd.Grpc.ConnectionDeadline)*time.Second)

	// ctx, cancel = context.WithTimeout(ctx, time.Duration(conf.Harepd.Grpc.ConnectionDeadline))

	defer cancel()

	info, err := client.GetClusterInfo(ctx, myinfo)
	if err != nil {

		return nil, err
	}
	return info, nil
}

//Fetch Local Node Info

func fetchLocal(conf *models.Config, local *models.WhatYouKnow) error {
	// SQL Queries
	// Node ID
	nodeID := models.DB.Raw("SELECT node_id FROM repmgr.nodes WHERE node_name = ?", conf.Harepd.NodeName).Row()

	// Node IP
	local.Ip = conf.Harepd.Grpc.BindAddress

	// Node ID
	var r int32
	err := nodeID.Scan(&r)
	if err != nil {
		Logs.Warn("Node Id could not be found on the database appending : -1")
		local.NodeId = -1
	} else {
		local.NodeId = r
	}
	return nil
}
