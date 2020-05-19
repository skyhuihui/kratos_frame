package dao

import (
	"context"
	"kratos_frame/internal/model"
	"time"

	"github.com/bilibili/kratos/pkg/net/trace"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (dbs []*gorm.DB, err error) {
	var cfg struct {
		Client *sql.Config
	}
	if err = paladin.Get("db.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db, err := gorm.Open("mysql", cfg.Client.DSN)
	if err != nil {
		return nil, err
	}
	// 创建数据库的时候名字不是复数
	db.SingularTable(true)
	// 数据库操作日志
	//db.LogMode(true)
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return model.Tb + defaultTableName
	}
	dbs = append(dbs, db)
	for _, v := range cfg.Client.ReadDSN {
		db, err := gorm.Open("mysql", v)
		if err != nil {
			return dbs, err
		}
		db.SingularTable(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.DB().SetConnMaxLifetime(time.Hour)
		dbs = append(dbs, db)
	}
	if err = CreatTable(db); err != nil {
		return nil, err
	}
	return
}

func CreatTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Banner{},
		&model.Notice{},
		&model.Protocol{},
		&model.User{},
		&model.UserCertification{},
		&model.UserInvite{},
		&model.UserCertLevel{},
		&model.UserRole{},
		&model.Version{},
		&model.Work{},
		&model.Menu{},
		&model.Role{},
		&model.GoogleVerify{},
	).Error
}
func (d *Dao) DbClose() (err error) {
	for _, v := range d.db {
		if err = v.Close(); err != nil {
			return
		}
	}
	return
}

const (
	_family = "sql_client"
	find    = "find"
	create  = "create"
	update  = "update"
	delete  = "delete"
)

// args : ctx , 操作名称， 数据库id， 操作备注（数据库操作之类）
func DbTrace(c context.Context, operationName string, dbId int, operate string) trace.Trace {
	if t, ok := trace.FromContext(c); ok {
		t = t.Fork(_family, operationName)
		t.SetTag(trace.Int(trace.TagAddress, dbId), trace.String(trace.TagComment, operate))
		return t
	}
	return nil
}

//轮询记录值
var dbId = 1

// 获取 读哪个数据库的数据库下标 轮询
func GetDbId(l int) int {
	if l == 1 {
		return 0
	}
	if dbId == l {
		dbId = 1
	}
	dbId++
	return dbId - 1
}
