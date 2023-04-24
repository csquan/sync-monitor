package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/sirupsen/logrus"
)

type BSCMonitorService struct {
	db     types.IDB
	config *config.Config
}

func NewBSCMonitorService(db types.IDB, cfg *config.Config) *BSCMonitorService {
	return &BSCMonitorService{
		db:     db,
		config: cfg,
	}
}

// 读取不同的数据库，在一定时间内是否连续，否则报警
func (a *BSCMonitorService) Run() (err error) {
	logrus.Info("sync-monitor run at ")

	height, err := a.db.GetBTCHeight("BTC")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(height)
	return
}

func (a *BSCMonitorService) Name() string {
	return "BSC monitor"
}
