package orm

import (
	"time"

	"github.com/ldd27/go-lib/zaplog"

	"github.com/jinzhu/gorm"
)

type Option struct {
	Dialect             string
	Address             string
	EnableSingularTable bool
	EnableLogMode       bool
	ConnMaxLifetime     time.Duration
	MaxIdleConns        int
	MaxOpenConns        int
	Log                 *zaplog.ExtendLogger
}

var defaultGORMConf = Option{
	Dialect:             "mysql",
	EnableSingularTable: true,
	EnableLogMode:       true,
	ConnMaxLifetime:     2 * time.Hour,
	MaxIdleConns:        16,
	MaxOpenConns:        100,
}

func NewDB(opts ...func(*Option)) *gorm.DB {
	opt := defaultGORMConf
	for _, o := range opts {
		o(&opt)
	}

	// 数据库连接
	db, err := gorm.Open(opt.Dialect, opt.Address)
	if err != nil {
		panic("connect to database fail:" + err.Error())
	}

	// 全局禁用表名复数 true则表名为user false则表名为users
	db.SingularTable(opt.EnableSingularTable)
	// 设置日志级别
	db.LogMode(opt.EnableLogMode)
	//
	db.DB().SetConnMaxLifetime(opt.ConnMaxLifetime)
	//
	db.DB().SetMaxIdleConns(opt.MaxIdleConns)
	//
	db.DB().SetMaxOpenConns(opt.MaxOpenConns)

	if opt.Log != nil {
		db.SetLogger(opt.Log)
	}

	// ping
	err = db.DB().Ping()
	if err != nil {
		panic("connect to database fail:" + err.Error())
	}

	return db
}

func AutoMigrate(db *gorm.DB, tables ...interface{}) error {
	if len(tables) == 0 {
		return nil
	}

	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables...).Error

	if err != nil {
		return err
	}

	return nil
}
