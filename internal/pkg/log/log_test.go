package log

import (
	"context"
	"testing"
	"time"

	"github.com/bilibili/kratos/pkg/log"
)

func TestLog(m *testing.T) {
	Log("./log/test", InfoLevel, log.KVString("name", "xiaoming"))
	Log("./log/test", DebugLevel, log.KVString("name", "xiaoming"))
	Log("./log/test", ErrorLevel, log.KVString("name", "xiaoming"))
	Log("./log/test", WarnLevel, log.KVString("name", "xiaoming"))
	Log("./log/test", FatalLevel, log.KVString("name", "xiaoming"))

	conf := &log.Config{
		Stdout: true,
		Dir:    "./log/default",
	}
	log.Init(conf)

	log.Info("hi:%s", "kratos")
	log.Infoc(context.TODO(), "hi:%s", "kratos")
	log.Infov(context.TODO(), log.KVInt("key1", 100), log.KVString("key2", "test value"))
	log.Infow(context.TODO(), "i", "like", "dog", "ds")
	log.Error("hi:%s", "kratos")
	log.Errorc(context.TODO(), "hi:%s", "kratos")
	log.Errorv(context.TODO(), log.KVInt("key1", 100), log.KVString("key2", "test value"))
	log.Errorw(context.TODO(), "i", "like", "dog", "dog")
	log.Warn("hi:%s", "kratos")
	log.Warnc(context.TODO(), "hi:%s", "kratos")
	log.Warnv(context.TODO(), log.KVInt("key1", 100), log.KVString("key2", "test value"))
	log.Warnw(context.TODO(), "i", "like", "a", "dog")

	time.Sleep(time.Second * 1)
}
