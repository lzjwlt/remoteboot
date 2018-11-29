package config

type DaemonConf struct {
	ServerAddr string `json:"serverAddr"`
	ServerPort int    `json:"ServerPort"`
	LogFile    string `json:"logFile"`
}

func GetDefaultDaemonConf() *DaemonConf {
	return &DaemonConf{
		ServerAddr: "lzjwlt.cn",
		ServerPort: 30000,
		LogFile:    "console",
	}
}
