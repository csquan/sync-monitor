package services

import (
	"fmt"
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/ethereum/sync-monitor/util"
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
		time.Sleep(time.Duration(a.config.MonitorTime.BTC) * time.Second)
		afterHeight, err := a.db.GetBTCHeight("BTC")
		if err != nil {
			logrus.Error(err)
		}
		if height == afterHeight {
			str := fmt.Sprintf("btc 高度在配置的期限内没有变化,均为%d", afterHeight)
			util.TgAlert(str)
		} else {
			//logrus.Info("btc 高度在配置的期限内正常变化")
		}
	}
	return
}

func (a *BTCMonitorService) Name() string {
	return "BTC monitor"
}
