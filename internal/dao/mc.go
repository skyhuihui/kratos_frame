package dao

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func NewMC() (mc *memcache.Memcache, err error) {
	var cfg struct {
		Client *memcache.Config
	}
	if err = paladin.Get("memcache.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	mc = memcache.New(cfg.Client)
	return
}

func (d *Dao) PingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func keyArt(id int64) string {
	return fmt.Sprintf("art_%d", id)
}
