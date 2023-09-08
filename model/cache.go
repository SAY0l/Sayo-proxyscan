package model

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/patrickmn/go-cache"
	"github.com/sayol/Sayo-proxyscan/log"
	"github.com/urfave/cli"
)

func init() {
	gob.Register(ProxyInfo{})
}

func SaveProxies(isProxy bool, proxyinfo ProxyInfo, err error) {
	if err == nil && isProxy {
		k := fmt.Sprintf("%v://%v:%v", proxyinfo.Protocol, proxyinfo.Addr, proxyinfo.Port)
		CACHE_PROXIES.Set(k, true, cache.NoExpiration)
	}
} //常规go-cache键值存储

func SaveProxiesToFile() error {
	return CACHE_PROXIES.SaveFile("Sayo-proxyscan.db")
} // 持久化键值对

func CacheStatus() (count int, items map[string]cache.Item) {
	count = CACHE_PROXIES.ItemCount()
	items = CACHE_PROXIES.Items()
	return count, items
} // 快速读取键值

func ProxiesNum() {
	log.Log.Infof("Total right proxies: %v", CACHE_PROXIES.ItemCount())
}

func LoadProxiesFromFile() {
	CACHE_PROXIES.LoadFile("Sayo-proxyscan.db")
	ProxiesNum()
} //释放持久化，将数据放入缓存

func Dump(ctx *cli.Context) (err error) {
	LoadProxiesFromFile()

	if ctx.IsSet("file") {
		DUMP_FILENAME = ctx.String("file")
	}
	err = DumpToFile(DUMP_FILENAME)
	if err != nil {
		log.Log.Fatalf("Dump proxies to file err, Err: %v", err)
	}
	return err
}

func DumpToFile(filename string) (err error){
	file,err :=os.Create(filename)
	if err ==nil {
		_ , items := CacheStatus()
		for k := range items {
			file.WriteString(fmt.Sprintf("%v\n",k))
		}
	}
	return err
} //常规情况下的go-cache转储方式