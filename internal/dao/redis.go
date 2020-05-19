package dao

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func NewRedis() (r *redis.Redis, err error) {
	var cfg struct {
		RedisPass struct {
			Password string
		}
		Client *redis.Config
	}
	if err = paladin.Get("redis.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewRedis(cfg.Client, redis.DialPassword(cfg.RedisPass.Password))
	if _, err = r.Do(context.Background(), "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	if value, err := redis.String(r.Do(context.Background(), "GET", "ping")); err != nil {
		log.Error("conn.get(PING) error(%v)", err)
		return nil, err
	} else {
		log.Info("", value)
	}
	return
}

func (d *Dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

// args ...  key , value
// 设置超时时间  args...  key, value , EX, 120
func (d *Dao) RedisSet(ctx context.Context, args ...interface{}) error {
	if b, err := redis.String(d.redis.Do(ctx, "SET", args...)); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
		return err
	} else {
		if b != "OK" {
			log.Error("Redis Key 值更新失败", ctx)
			return fmt.Errorf("redis key 更新失败")
		}
	}
	return nil
}

func (d *Dao) RedisGet(ctx context.Context, key interface{}) (value interface{}, err error) {
	if value, err = d.redis.Do(ctx, "GET", key); err != nil {
		log.Error("conn.get(PING) error(%v)", err)
		return nil, err
	}
	return
}

func (d *Dao) RedisDel(ctx context.Context, key interface{}) (i int64, err error) {
	if i, err = redis.Int64(d.redis.Do(ctx, "DEL", key)); err != nil {
		log.Error("del key err: ", err)
	}
	return
}

// 判断某个key是否存在
func (d *Dao) RedisExists(ctx context.Context, key interface{}) (is_key_exit bool, err error) {
	if is_key_exit, err = redis.Bool(d.redis.Do(ctx, "EXISTS", key)); err != nil {
		log.Error("redis exists err :", err)
		return false, err
	}
	return
}

// 返回一个redis的连接 使用完要close
func (d *Dao) RedisDo(ctx context.Context) redis.Conn {
	return d.redis.Conn(ctx)
}
