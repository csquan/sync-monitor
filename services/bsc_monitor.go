package services

import (
	"fmt"
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"github.com/ethereum/sync-monitor/util"
	"github.com/sirupsen/logrus"
	"time"
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
	logrus.Info("bsc monitor run ")

	for {
		height, err := a.db.GetBSCHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		time.Sleep(time.Duration(a.config.MonitorTime.BSC) * time.Second)
		//当前获取一个高度，sleep后再获取一个高度，看是否变化
		afterHeight, err := a.db.GetBSCHeight("erc20_tx")
		if err != nil {
			logrus.Error(err)
		}
		if height == afterHeight {
			str := fmt.Sprintf("bsc 高度在配置的期限内没有变化,均为%d", afterHeight)
			util.TgAlert(str)
		} else {
			//logrus.Info("bsc 高度在配置的期限内正常变化")
		}
	}
	return
}

func (a *BSCMonitorService) Name() string {
	return "BSC monitor"
}
