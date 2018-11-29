package g

import "../config"

type DaemonCfg struct {
	config.DaemonConf

	CfgFile string
}

type ServerCfg struct {
	config.ServerConf

	CfgFile string
}

var (
	Version         string = "0.3"
	GlobalDaemonCfg *DaemonCfg
	GlobalServerCfg *ServerCfg
)

func init() {
	GlobalDaemonCfg = &DaemonCfg{
		DaemonConf: *config.GetDefaultDaemonConf(),
	}
	GlobalServerCfg = &ServerCfg{
		ServerConf: *config.GetDefaultServerConf(),
	}
}
