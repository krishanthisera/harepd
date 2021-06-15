# Repmgr Haproxy Daemon  
Go Daemon to track Postgres switch-overs  
![harepd-connectivity][harepd-connectivity]
# To build the application  
1. Download the go binaries to local system    
    ```sh
        wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz 
    ```
2. Extract the tar file and set the path  
    ```sh
        sudo tar -xvf go1.15.3.linux-amd64.tar.gz
        sudo mv go /usr/local
        export GOROOT=/usr/local/go 
        export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
        # To verify the installation
        go version
    ```
3. Clone and Build the application  
    ```sh
        git clone https://code.unitiwireless.com/uniti-wireless/repmgr-haproxy-daemon.git
        cd repmgr-haproxy-daemon
        go build .
    ```
# To run the application (Stand alone mode)
The application can be run as one of three modes,  
    1. Server only mode `./harepd -f /conf/file/location.yml --server-only`   
    2. Client only mode `./harepd -f /conf/file/location.yml --client-only`   
    3. Dual mode (server and client) `./harepd -f /conf/file/location.yml`   
_Ensure to make a backup of pg_hba.conf before proceedingßto_   
Prior to run the application,  
1. Create the config(yml) file using which follows the below template   
    ```sh
        harepd:
            nodeName: "node_1"
            primaryIP: "10.21.57.61"
            allowRO: false
            hbaConfig: "/etc/postgresql/12/main/pg_hba.conf"
            haproxy:
                server: 
                - "10.21.57.65"
                users:
                  readOnly: "haro"
                  readWrite: "harw"
            authModes: 
                allow: "md5"
                deny: "reject"
            repmgr:
                user: "repmgr"
                db: "repmgr"
            gRPC:
                tls:
                enabled: false
                ca: "/etc/harepd/ca_cert.pem"
                cert: "/etc/harepd/server_cert.pem"
                key: "/etc/harepd/server_key.pem"
                bindPort: 10000
                bindAddress: "10.21.51.61"
                serverHostOverride: ""
                neighbours:
                - "10.21.57.61:10000"
                witness: "10.21.57.63:10000"
                connectionDeadline: 5
            watchDog: 15
            logs: 
                filePath: /var/log/harepd.log
                maxAge: 86400
                rotationTime: 604800
    ```
- `nodeName:` Exactly as `repmgr` config   
- `primaryIp:` IP address used in Postgres replication   
- `allowRO: false` If the slave nodes supposed to serve the read only traffic  
- `readOnly: "haro"` If allowRo is set, a user who is supposed test accessibility on the slave database   
- `hbaConfig:` Path to pg_hba.conf file  
- `haproxy.server:` Ip of the HAProxy server  
- `haproxy.users.readWrite: "harw"` This is a mandatory field that specifies the user who is supposed facilitate health check quarries (For writability)
- `haproxy.users.readOnly: "haro"` This is an optional field that specifies the user who is supposed facilitate health check quarries (For ReadOnly)
  _Note that this will be only used if it is required to send ReadOnly traffic to Slave nodes_  
- `haproxy.authModes.<allow/deny>:` This will be used to alter `the pg_hba.conf` file  
- `gRPC.tls.<*>:` TLS configuration for gRPC  
- `gRPC.bindPort:` gRPC server port  
- `gRPC.bindAddress:` IP to associate with gRPC server  
- `gRPC.neighbours:"<ip>:<port>"` List if neighbours who are listening to `HAREPD-Clients`  
- `gRPC.witness:"<ip>:<port>"` IP of the witness node who is listening to `HAREPD-Clients`  
- `gRPC.connectionDeadline:` How long should a client wait for a neighbour to response  
- `watchDog:` If the application should run on `Dual mode`, how frequent the client should send the inquiries  
- `logs.filePath :` Location of the log file, if you should change this path, ensure that changes are accordingly reflected on the `systemd` definitions (`harepd-client.service`, `harepd-server.service`)    
## Bootstrap sequence  
As a best practice it is recommended to bootstrap the application according to the following sequence,    
- Start *server only* mode on each node including the witness  
- Start *client only* mode on Slave node  
- Start *client only* mode on Master nodes  
_Note that since the witness node is not required to serve the traffic client is not required run on the witness_  
_In a scenario where the daemon should be stopped, prior to stop the sever `service`, it is must to stop `client`_
2. Make sure to run the application by using `postgres` user
    ```sh
        ./harepd -f /conf/file/location.yml
    ```    
    _Make sure logs are allowed written on the given location_
# To run the application in Production
Ensure that you have root access to your postgres nodes and create `systemd` service definition. It is required to create two separate services for client and server.  
Create the following *server* service definition in : `/etc/systemd/system/harepd-server.service`  
```sh
    [Unit]
    Description=harepd server daemon. To start/stop the grpc server
    After=network.target
    StartLimitIntervalSec=0

    [Service]
    Type=simple
    Restart=always
    RestartSec=1
    User=postgres
    PermissionsStartOnly=true
    ExecStart=/usr/sbin/harepd -f /etc/harepd/conf.yml --server-only
    ExecStartPre=/usr/bin/touch /var/log/harepd.log
    ExecStartPre=/bin/chown postgres /var/log/harepd.log
    StandardOutput=syslog
    StandardError=syslog
    SyslogIdentifier=harepd-server

    [Install]
    WantedBy=multi-user.target
```
Ensure that the binaries which was build earlier and config files have been copied to desired location.  
Create the following *client* service definition in : `/etc/systemd/system/harepd-client.service`  
```sh
    # Use timers if client ought to run longer than expected
    [Unit]
    Description=harepd client to discover the truth
    After=network.target
    StartLimitIntervalSec=0

    [Service]
    Type=simple
    Restart=always
    RestartSec=10s
    User=postgres
    ExecStart=/usr/sbin/harepd -f /etc/harepd/conf.yml --client-only
    PermissionsStartOnly=true
    ExecStartPre=/usr/bin/touch /var/log/harepd.log
    ExecStartPre=/bin/chown postgres /var/log/harepd.log
    StandardOutput=syslog
    StandardError=syslog
    SyslogIdentifier=harepd-server

    [Install]
    WantedBy=multi-user.target


```
Finally, start client and server service.
```sh
    systemctl start harepd-server.service
    systemctl start harepd-client.service
```
# Communication
gRPC has been to messaging. `messaging.proto` describes the specification of the message.
```sh
    syntax = "proto3";
    package models;
    service ClusterInfo {
        rpc GetClusterInfo(WhatYouKnow) returns (IKnow){}
    }

    message WhatYouKnow{
        string ip = 1;
        int32 nodeId = 2;
        string nodeName = 3;
    }

    message IKnow{
        string ip = 1;
        int32 nodeId = 2;
        string nodeName = 3;
        string witness = 4;
        repeated string slaves = 5;
        repeated string master = 6;
}
```
![gRPC-Comms][gRPC-Comms]  
As describes, any node in the cluster can make an inquiry to their neighbours about their awareness. In that case the neighbour will act as the server and send their opinion regards to the cluster. Further, clients are responsible to adjust their configuration depending on their neighbours opinions. The decision making is solely client's responsible.  
Refer `logic.go` to more information on decision making.
 # Anatomy of Altering Rules  
To regulate the traffic between master and slave nodes, a combination of `postgres` and `haproxy` configuration has been used.  
Prior to sending the traffic HAProxy runs a health check on the Postgres servers. In this case HAProxy is trying to authenticate to a dummy database which may considered as unsuccessful in case if the Postgres server straight away rejects the request. 
Simply the application alter the `pg_hba.conf` file using two different methods depending on the context.
1. For Master nodes *AlterRule()*  
  By loading the `pg_hba.conf` to a temporary table and altering the rule, and then save the changes to the file, application make alteration to the Master node's file.  
  *This method is deprecated to maintain the consistency of the program*  
  _Without accessing the filesystem directly_  
2. For Master and Slave nodes *AlterRuleLegacy()*  
  Generally, Slave nodes are supposed to be readonly, in this case application _directly access the node's file system_ and do the alterations  
## App-logic  
The behaviour of the application can be changed by altering logic.go which belongs to the package main. For the time being, the following logic has been implemented,  
![app-logic][app-logic]


[app-logic]: https://lucid.app/publicSegments/view/3a9682f4-593c-4e14-9365-2b62f8b15a78/image.png
[gRPC-Comms]: https://lucid.app/publicSegments/view/1a88fbb4-482d-4f63-b343-973c7e429fb8/image.png
[harepd-connectivity]: https://lucid.app/publicSegments/view/cb811396-d7fc-4d48-8cee-cd2322b6bd8c/image.png
