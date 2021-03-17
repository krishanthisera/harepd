package main

import (
	"flag"

	"code.unitiwireless.com/uniti-wireless/harepd/controllers"
	"code.unitiwireless.com/uniti-wireless/harepd/models"
	"github.com/sirupsen/logrus"
)

// func main() {

// 	// ARGS
// confFile := flag.String("f", "conf.yaml", "Path to config file")
// preflight := flag.Bool("preflight", false, "Run a Preflight check")
// serverOnly :=  flag.Bool("server-only", false, "Run the server only")
// clientOnly	:= flag.Bool("client-only", false, "Run the client only")
// flag.Parse()

// 	// Read Configurations
// 	var conf models.Config
// 	controllers.ReadConf(*confFile, &conf)

// 	// preflight
// 	if *preflight {
// 		controllers.PreflightCheck(&conf)
// 		return
// 	}

// 	master, _ := controllers.TryMaster(&conf)
// 	slave, _ := controllers.TrySlave(&conf)

// 	// Alow HAProxy Trafic
// 	var rules models.Rules
// 	if master != nil {
// 		conf.Harepd.AuthModes.Deny = "md5"
// 		controllers.AlterRule(&rules, &conf)
// 	}
// 	if slave != nil {
// 		conf.Harepd.AuthModes.Deny = "reject"
// 		controllers.AlterRuleLegacy(&rules, &conf)
// 	}

// }
var rules models.Rules

//Logs for loggin with loggers
var Logs *logrus.Logger

func main() {
	var conf models.Config

	//Flags
	confFile := flag.String("f", "conf.yaml", "Path to config file")
	// preflight := flag.Bool("preflight", false, "Run a Preflight check")
	serverOnly := flag.Bool("server-only", false, "Run the server only")
	clientOnly := flag.Bool("client-only", false, "Run the client only")

	flag.Parse()

	//Conf File init
	controllers.ReadConf(*confFile, &conf)

	// Initiating logs
	Logs = models.NewLogger(&conf)

	//DB Init
	dbInit(&conf)

	// Initiating modes
	if *serverOnly {
		serverOnlyMode(&conf)
	} else if *clientOnly {
		clientOnlyMode(&conf)
	} else {
		dualMode(&conf)
	}

}

//fmt.Println(controllers.GrpcServer(conf))

// for x := range time.Tick(10 * time.Second) {
// 	aware, _ := controllers.GrpcClient("10.21.57.62", &conf)
// 	fmt.Println(x)
// 	fmt.Println(aware)
// }
