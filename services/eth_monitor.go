package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/ethereum/sync-monitor/util"
	"github.com/sirupsen/logrus"
	"time"
)

type ETHMonitorService struct {
	db     types.IDB
	config *config.Config
}

func NewETHMonitorService(db types.IDB, cfg *config.Config) *ETHMonitorService {
	return &ETHMonitorService{
		db:     db,
		config: cfg,
	}
}

// 读取不同的数据库，在一定时间内是否连续，否则报警
func (a *ETHMonitorService) Run() (err error) {
	logrus.Info("sync-monitor run at ")
	for {
		height, err := a.db.GetETHHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		time.Sleep(time.Duration(a.config.MonitorTime.ETH) * time.Second)
		afterHeight, err := a.db.GetETHHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		if height == afterHeight {
			util.TgAlert("eth 高度在配置的期限内没有变化")
		} else {
			//logrus.Info("eth 高度在配置的期限内正常变化")
		}
	}
	return
}

func (a *ETHMonitorService) Name() string {
	return "ETH monitor"
}
