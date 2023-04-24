package main

import (
	"flag"
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/db"
	"github.com/ethereum/sync-monitor/log"
	"github.com/ethereum/sync-monitor/services"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const CONTRACTLEN = 42

var (
	conffile string
	env      string
)

func init() {
	flag.StringVar(&conffile, "conf", "config.yaml", "conf file")
	flag.StringVar(&env, "env", "prod", "Deploy environment: [ prod | test ]. Default value: prod")
}

func main() {
	var err error
	if config.Conf, err = config.LoadConfig("./conf"); err != nil {
		logrus.Error("🚀 Could not load environment variables")
		panic(err)
	}

	flag.Parse()

	err = log.Init("sync-monitor", &config.Conf)
	if err != nil {
		log.Fatal(err)
	}
	btc_dbConnection, err := db.NewBTCMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	bsc_dbConnection, err := db.NewBSCMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	eth_dbConnection, err := db.NewETHMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	hui_dbConnection, err := db.NewHUIMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	tron_dbConnection, err := db.NewTRONMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	//listen kill signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	//setup scheduler
	scheduler, err := services.NewServiceScheduler(&config.Conf, btc_dbConnection, bsc_dbConnection, eth_dbConnection, hui_dbConnection, tron_dbConnection, sigCh)
	if err != nil {
		return
	}
	scheduler.Start()
}
