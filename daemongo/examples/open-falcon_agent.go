package main

import (
	"flag"
	"fmt"
	"github.com/open-falcon/agent/cron"
	"github.com/open-falcon/agent/funcs"
	"github.com/open-falcon/agent/g"
	"github.com/open-falcon/agent/http"
	daemon "github.com/rongyungo/sdk/daemongo"
	"os"
)

func init() {
	daemon.AppName = "openfalcon-agent"
	//启动守护进程，监控app
	daemon.Warden = true
}

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if g.Config().Daemon {
		//变成daemon模式
		switch isDaemon, err := daemon.Daemonize(); {
		case !isDaemon:
			return
		case err != nil:
			fmt.Printf("main(): could not start daemon, reason -> %s", err.Error())
		}
	}

	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()

	funcs.BuildMappers()

	go cron.InitDataHistory()

	cron.ReportAgentStatus()
	cron.SyncMinePlugins()
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()
	cron.Collect()

	go http.Start()

	select {}

}
