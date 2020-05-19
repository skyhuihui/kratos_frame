package log

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/bilibili/kratos/pkg/log"
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var handlers map[string]*log.FileHandler // 写文件的FileHandler 存在内存 不用重复创建
var filters map[string]struct{}          // 用于过滤

func Log(dir string, lv int, d ...log.D) {
	dir += "/" + strings.Split(time.Now().Local().String(), " ")[0]
	logHandler := getLogHandler(dir)
	for i := range d {
		if _, ok := filters[d[i].Key]; ok {
			d[i].Value = "***"
		}
	}
	fn := funcName(2)
	d = append(d, log.KVString("source", fn))
	d = append(d, log.KV("time", time.Now()), log.KVInt64("level_value", int64(lv)), log.KVString("level", log.Level.String(log.Level(lv))))

	logHandler.Log(context.TODO(), log.Level(lv), d...)
}

// 根据不同的文件输出目录， 获取不同的FileHandler
func getLogHandler(dir string) *log.FileHandler {
	if len(handlers) == 0 {
		handlers = make(map[string]*log.FileHandler, 0)
	}
	if handlers[dir] != nil {
		return handlers[dir]
	}

	fileHandler := log.NewFile(dir, 0, 0, 0)
	handlers[dir] = fileHandler
	return fileHandler
}

// funcName get func name.
func funcName(skip int) (name string) {
	if _, file, lineNo, ok := runtime.Caller(skip); ok {
		return file + ":" + strconv.Itoa(lineNo)
	}
	return "unknown:0"
}
