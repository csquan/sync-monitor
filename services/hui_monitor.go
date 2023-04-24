package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/sirupsen/logrus"
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

	height, err := a.db.GetBTCHeight("BTC")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(height)
	return
}

func (a *HUIMonitorService) Name() string {
	return "HUI monitor"
}
