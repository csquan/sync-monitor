package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/sirupsen/logrus"
	"time"
)

type HUIMonitorService struct {
	db     types.IDB
	config *config.Config
}

func NewHUIMonitorService(db types.IDB, cfg *config.Config) *HUIMonitorService {
	return &HUIMonitorService{
		db:     db,
		config: cfg,
	}
}

// 读取不同的数据库，在一定时间内是否连续，否则报警
func (a *HUIMonitorService) Run() (err error) {
	logrus.Info("sync-monitor run at ")

	for {
		height, err := a.db.GetHUIHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		time.Sleep(time.Duration(a.config.MonitorTime.HUI) * time.Second)
		afterHeight, err := a.db.GetHUIHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		if height == afterHeight {
			logrus.Error("hui 高度在配置的期限内没有变化")
		} else {
			logrus.Info("hui 高度在配置的期限内正常变化")
		}
	}
	return
}

func (a *HUIMonitorService) Name() string {
	return "HUI monitor"
}
