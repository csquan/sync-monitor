package services

import (
	"github.com/ethereum/sync-monitor/config"
	"github.com/ethereum/sync-monitor/types"
	"os"
	"time"
)

type ServiceScheduler struct {
	conf *config.Config

	btc_db types.IDB

	bsc_db types.IDB

	eth_db types.IDB

	hui_db types.IDB

	tron_db types.IDB

	services []types.IAsyncService

	closeCh <-chan os.Signal
}

func NewServiceScheduler(conf *config.Config, btc_db types.IDB, bsc_db types.IDB, eth_db types.IDB, hui_db types.IDB, tron_db types.IDB, closeCh <-chan os.Signal) (t *ServiceScheduler, err error) {
	t = &ServiceScheduler{
		conf:     conf,
		closeCh:  closeCh,
		btc_db:   btc_db,
		bsc_db:   bsc_db,
		eth_db:   eth_db,
		hui_db:   hui_db,
		tron_db:  tron_db,
		services: make([]types.IAsyncService, 0),
	}

	return
}

func (t *ServiceScheduler) Start() {
	//create services
	btcMonitorService := NewBTCMonitorService(t.btc_db, t.conf)
	bscMonitorService := NewBSCMonitorService(t.bsc_db, t.conf)
	ethMonitorService := NewETHMonitorService(t.eth_db, t.conf)
	huiMonitorService := NewHUIMonitorService(t.hui_db, t.conf)
	tronMonitorService := NewTRONMonitorService(t.tron_db, t.conf)

	//t.services = []types.IAsyncService{
	//	btcMonitorService,
	//	bscMonitorService,
	//	//ethMonitorService,
	//	//huiMonitorService,
	//	//tronMonitorService,
	//}

	go btcMonitorService.Run()
	go bscMonitorService.Run()
	go ethMonitorService.Run()
	go huiMonitorService.Run()
	go tronMonitorService.Run()

	timer := time.NewTimer(2)
	for {
		select {
		case <-t.closeCh:
			return
		case <-timer.C:
		}
	}
}
