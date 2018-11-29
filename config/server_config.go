package config

type ServerConf struct {
	BindAddr string `json:"bindAddr"`
	BindPort int    `json:"bindPort"`
	ApiPort  int    `json:"ApiPort"`
	LogFile  string `json:"logFile"`
}

func GetDefaultServerConf() *ServerConf {
	return &ServerConf{
		BindAddr: "0.0.0.0",
		BindPort: 30000,
		ApiPort:  30001,
		LogFile:  "console",
	}
}
