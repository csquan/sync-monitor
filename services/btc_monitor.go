package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/sirupsen/logrus"
	"time"
)

type BTCMonitorService struct {
	db     types.IDB
	config *config.Config
}

func NewBTCMonitorService(btc_db types.IDB, cfg *config.Config) *BTCMonitorService {
	return &BTCMonitorService{
		db:     btc_db,
		config: cfg,
	}
}

// 读取不同的数据库，在一定时间内是否连续，否则报警
func (a *BTCMonitorService) Run() (err error) {
	for {
		height, err := a.db.GetBTCHeight("BTC")
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(height)
		time.Sleep(time.Duration(a.config.MonitorTime.BTC) * time.Second)
		logrus.Info("BTC 经过sleep")
	}
	return
}

func (a *BTCMonitorService) Name() string {
	return "BTC monitor"
}
