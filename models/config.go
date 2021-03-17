package models

//Config map the user configuration as a yaml file
type Config struct {
	Harepd struct {
		AllowRO   bool   `yaml:"allowRO"`
		PrimaryIP string `yaml:"primaryIP"`
		NodeName  string `yaml:"nodeName"`
		HbaConfig string `yaml:"hbaConfig"`
		WatchDog  int16  `yaml:"watchDog"`
		Logs      struct {
			FilePath     string `yaml:"filePath"`
			MaxAge       int32  `yaml:"maxAge"`
			RotationTime int32  `yaml:"rotationTime"`
		} `yaml:"logs"`
		Haproxy struct {
			Server []string `yaml:"server"`
			Users  struct {
				ReadOnly  string `yaml:"readOnly"`
				ReadWrite string `yaml:"readWrite"`
			} `yaml:"users"`
		} `yaml:"haproxy"`
		AuthModes struct {
			Allow string `yaml:"allow"`
			Deny  string `yaml:"deny"`
		} `yaml:"authModes"`
		Repmgr struct {
			User string `yaml:"user"`
			Db   string `yaml:"db"`
		} `yaml:"repmgr"`
		Grpc struct {
			BindPort           int16    `yaml:"bindPort"`
			BindAddress        string   `yaml:"bindAddress"`
			ServerHostOverride string   `yaml:"serverHostOverride"`
			Neighbours         []string `yaml:"neighbours"`
			ConnectionDeadline int16    `yaml:"connectionDeadline"`
			Witness            string   `yaml:"witness"`
			TLS                struct {
				Enabled bool   `yaml:"enabled"`
				Ca      string `yaml:"ca"`
				Key     string `yaml:"key"`
				Cert    string `yaml:"cert"`
			} `yaml:"tls"`
		} `yaml:"gRPC"`
	} `yaml:"harepd"`
}
