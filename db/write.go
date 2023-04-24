package db

import (
	"fmt"
	"github.com/ethereum/sync-monitor/types"
	"time"

	"github.com/ethereum/sync-monitor/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

type Mysql struct {
	conf   *config.Config
	engine *xorm.Engine
}

func NewBTCMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.BTC_Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func NewBSCMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.BSC_Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func NewETHMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.ETH_Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func NewTRONMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.TRON_Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func NewHUIMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.HUI_Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func (m *Mysql) GetEngine() *xorm.Engine {
	return m.engine
}

func (m *Mysql) GetSession() *xorm.Session {
	return m.engine.NewSession()
}

func (m *Mysql) CommitWithSession(db types.IDB, executeFunc func(*xorm.Session) error) (err error) {
	session := db.GetSession()
	err = session.Begin()
	if err != nil {
		logrus.Errorf("begin session error:%v", err)
		return
	}

	defer session.Close()

	err = executeFunc(session)
	if err != nil {
		logrus.Errorf("execute func error:%v", err)
		err1 := session.Rollback()
		if err1 != nil {
			logrus.Errorf("session rollback error:%v", err1)
		}
		return
	}

	err = session.Commit()
	if err != nil {
		logrus.Errorf("commit session error:%v", err)
	}

	return
}
