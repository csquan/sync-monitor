package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/sirupsen/logrus"
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
	height, err := a.db.GetBTCHeight("BTC")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(height)
	return
}

func (a *BTCMonitorService) Name() string {
	return "BTC monitor"
}
