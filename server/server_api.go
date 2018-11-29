package server

import (
	"encoding/json"
	"log"
	"net/http"

	"../g"
)

type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// api/serverinfo
type ServerInfoResp struct {
	Response
	Version     string `json:"version"`
	BindPort    int    `json:"bindPort"`
	ApiPort     int    `json:"apiPort`
	DaemonCount int    `json:"daemonCount"`
}

func apiServerInfo(w http.ResponseWriter, r *http.Request) {
	var (
		buf []byte
		res ServerInfoResp
	)
	defer func() {
		log.Println("HTTP response [%s]: [%d]", r.URL.Path, res.Code)
	}()
	log.Print("HTTP request [%s]", r.URL.Path)
	cfg := g.GlobalServerCfg
	//statServer := StatServer()
	res = ServerInfoResp{
		Version:  g.Version,
		BindPort: cfg.BindPort,
		ApiPort:  cfg.ApiPort,
		//DaemonCount
	}
	buf, _ = json.Marshal(&res)
	w.Write(buf)
}
