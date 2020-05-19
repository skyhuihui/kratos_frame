package dao

import (
	"context"
	"time"

	"github.com/casbin/casbin/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// dao dao.
type Dao struct {
	db         []*gorm.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	enforcer   *casbin.Enforcer
	demoExpire int32
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, dbs []*gorm.DB, enforcer *casbin.Enforcer) (d Dao, err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = Dao{
		db:         dbs,
		redis:      r,
		mc:         mc,
		cache:      fanout.New("cache"),
		enforcer:   enforcer,
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.DbClose()
	d.cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.PingRedis(ctx); err != nil {
		return err
	}
	if err = d.PingMC(ctx); err != nil {
		return err
	}

	return nil
}
