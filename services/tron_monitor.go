package services

import (
	"fmt"
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/ethereum/sync-monitor/util"
	"github.com/sirupsen/logrus"
	"time"
)

type TRONMonitorService struct {
	db     types.IDB
	config *config.Config
}

func NewTRONMonitorService(db types.IDB, cfg *config.Config) *TRONMonitorService {
	return &TRONMonitorService{
		db:     db,
		config: cfg,
	}
}

// 读取不同的数据库，在一定时间内是否连续，否则报警
func (a *TRONMonitorService) Run() (err error) {
	logrus.Info("sync-monitor run at ")

	for {
		height, err := a.db.GetTRONHeight("TRC")
		if err != nil {
			logrus.Error(err)
		}
		time.Sleep(time.Duration(a.config.MonitorTime.TRON) * time.Second)
		afterHeight, err := a.db.GetTRONHeight("TRC")
		if err != nil {
			logrus.Error(err)
		}
		if height == afterHeight {
			str := fmt.Sprintf("tron 高度在配置的期限内没有变化,均为%d", afterHeight)
			util.TgAlert(str)
		} else {
			//logrus.Info("tron 高度在配置的期限内正常变化")
		}
	}

	return
}

func (a *TRONMonitorService) Name() string {
	return "TRON monitor"
}
