package model

import "github.com/patrickmn/go-cache"

var (
	CACHE_PROXIES *cache.Cache
	DUMP_FILENAME = "sayo-proxyscan.txt"
)

type ProxyInfo struct {
	Addr     string
	Port     int
	Protocol string
}

type Input_ProxyAddr struct {
	IP   string
	Port int
}

func init(){
	CACHE_PROXIES = cache.New(cache.NoExpiration,cache.DefaultExpiration)
}