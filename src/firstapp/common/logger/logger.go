package logger

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"strings"
	"time"
)

// 记录临时日志，用于日志模块初始化之前
func LogTemp(format string, a ...interface{}) (n int, err error) {
	fileInfo := ":"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		tmp := strings.Split(file, "/")
		fileInfo = fmt.Sprintf("%s:%d", tmp[len(tmp)-1], line)
	}

	// "2006-01-02 15:04:05" golang 格式化时间需写死这个。据说是golang诞生的时间
	prefix := fmt.Sprintf("[%s] [%s] ", time.Now().Format("2006-01-02 15:04:05"), fileInfo)

	return fmt.Printf(prefix+format+"\n", a...)
}

type dASLog struct {
	log *logs.BeeLogger
}

// 日志模块实例，外部调用
var Log = &dASLog{}

func init() {
	Log.log = logs.NewLogger(1)
	Log.log.EnableFuncCallDepth(true)
	Log.log.SetLogFuncCallDepth(3)
}

// 初始化日志实例
func (d *dASLog) Initialize(level int, config string) error {
	err := d.SetLogger(logs.AdapterFile, config)
	if err != nil {
		LogTemp("SetLogger AdapterFile error: %s", err)
		return err
	}

	err = d.SetLogger(logs.AdapterConsole)
	if err != nil {
		LogTemp("SetLogger AdapterConsole error: %s", err)
		return err
	}

	d.SetLevel(level)

	return nil
}
func (d *dASLog) SetLogger(adapterName string, configs ...string) error {
	config := append(configs, "{}")[0]

	return d.log.SetLogger(adapterName, config)
}

func (d *dASLog) SetLevel(l int) {
	d.log.SetLevel(l)
}

func (d *dASLog) Emergency(format string, v ...interface{}) {
	d.log.Emergency(format, v...)
}

func (d *dASLog) Alert(format string, v ...interface{}) {
	d.log.Alert(format, v...)
}

func (d *dASLog) Critical(format string, v ...interface{}) {
	d.log.Critical(format, v...)
}

func (d *dASLog) Error(format string, v ...interface{}) {
	d.log.Error(format, v...)
}

func (d *dASLog) Warn(format string, v ...interface{}) {
	d.log.Warn(format, v...)
}

func (d *dASLog) Info(format string, v ...interface{}) {
	d.log.Info(format, v...)
}

func (d *dASLog) Debug(format string, v ...interface{}) {
	d.log.Debug(format, v...)
}
