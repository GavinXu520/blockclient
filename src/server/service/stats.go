package service

import (
"encoding/json"
"expvar"
"runtime"
"time"
)

// 程序开始运行时间
var start = time.Now()

// 计算运行时间
func calculateUptime() interface{} {
	return time.Since(start).String()
}

// 当前 Golang 版本
func currentGoVersion() interface{} {
	return runtime.Version()
}

// 获取 CPU 核心数量
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// 当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// 当前 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// 调用cgo次数
func getNumCgoCall() interface{} {
	return runtime.NumCgoCall()
}


var lastPause uint32

// 获取上次 GC 的暂停时间
func getLastGCPauseTime() interface{} {
	var gcPause uint64
	ms := new(runtime.MemStats)
	statString := expvar.Get("memstats").String()
	if statString != "" {
		json.Unmarshal([]byte(statString), ms)

		if lastPause == 0 || lastPause != ms.NumGC {
			gcPause = ms.PauseNs[(ms.NumGC+255)%256]
			lastPause = ms.NumGC
		}
	}
	return gcPause
}

func init() {
	expvar.Publish("运行时间", expvar.Func(calculateUptime))
	expvar.Publish("版本", expvar.Func(currentGoVersion))
	expvar.Publish("cup数量", expvar.Func(getNumCPUs))
	expvar.Publish("系统", expvar.Func(getGoOS))
	expvar.Publish("cgo调用次数", expvar.Func(getNumCgoCall))
	expvar.Publish("协程数量", expvar.Func(getNumGoroutins))
	expvar.Publish("上次gc时间", expvar.Func(getLastGCPauseTime))
}
