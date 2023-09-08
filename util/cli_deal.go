package util

import (
	"time"

	"github.com/sayol/Sayo-proxyscan/log"
	"github.com/sayol/Sayo-proxyscan/model"
	"github.com/sayol/Sayo-proxyscan/proxy"
	"github.com/sayol/Sayo-proxyscan/task"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Scan(ctx *cli.Context) (err error){
	if ctx.IsSet("debug"){
		proxy.DEBUG_MODE = ctx.Bool("debug")
	}
	
	if proxy.DEBUG_MODE {
		log.Log.Logger.Level = logrus.DebugLevel
	}

	if ctx.IsSet("timeout"){
		proxy.TIMEOUT = time.Second*time.Duration((ctx.Int("timeout")))
	}

	if ctx.IsSet("scan_num"){
		proxy.SCAN_NUM = ctx.Int("scan_num")
	}

	if ctx.IsSet("filename"){
		proxy.INPUT_FILE = ctx.String("filename")
	}

	if ctx.IsSet("output_file"){
		model.DUMP_FILENAME = ctx.String("output_file")
	}

	startTime := time.Now()

	proxyAddrList := ReadProxyAddr(proxy.INPUT_FILE)
	proxyNum := len(proxyAddrList)
	log.Log.Infof("all %v proxies will be check",proxyNum)

	task.DistributeTask(proxyAddrList)

	log.Log.Debugf("distribute is fine")
	count , _ := model.CacheStatus()
	model.SaveProxiesToFile()
	err = model.DumpToFile(model.DUMP_FILENAME)
	log.Log.Infof("Scan proxies Done, Found %v proxies, used time: %v",count,time.Since(startTime))

	return err
}
