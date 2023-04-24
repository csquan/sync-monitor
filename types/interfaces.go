package types

import "github.com/go-xorm/xorm"

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	GetBTCHeight(name string) (int, error)
	GetBSCHeight(name string) (int, error)
	GetETHHeight(name string) (int, error)
	GetHUIHeight(name string) (int, error)
	GetTRONHeight(name string) (int, error)
}

type IWriter interface {
	GetSession() *xorm.Session
	GetEngine() *xorm.Engine
	CommitWithSession(db IDB, executeFunc func(*xorm.Session) error) (err error)
}

type IDB interface {
	IReader
	IWriter
}

type IAsyncService interface {
	Name() string
	Run() error
}
