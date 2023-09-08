package task

import (
	"sync"

	"github.com/sayol/Sayo-proxyscan/log"
	"github.com/sayol/Sayo-proxyscan/model"
	"github.com/sayol/Sayo-proxyscan/proxy"
)

func DistributeTask(sliceProxyAddr []model.Input_ProxyAddr) {
	proxyNum := len(sliceProxyAddr)
	log.Log.Infof("now %v proxies will be check", proxyNum)
	scanBatch := proxyNum / proxy.SCAN_NUM
	for i := 0; i < scanBatch; i++ {
		proxies := sliceProxyAddr[i*proxy.SCAN_NUM : (i+1)*proxy.SCAN_NUM]
		CheckProxy(proxies)
	}

	log.Log.Debugf("Scanning the last batch(%v)", scanBatch)
	if proxyNum%proxy.SCAN_NUM > 0 {
		proxies := sliceProxyAddr[scanBatch*proxy.SCAN_NUM : proxyNum]
		CheckProxy(proxies)
	}
}

func CheckProxy(proxies []model.Input_ProxyAddr) {
	var wg sync.WaitGroup
	wg.Add(len(proxies) * (len(proxy.HttpProxyProtocol) + len(proxy.SockProxyProtocol)))

	for _, addr := range proxies {
		for _, httppro := range proxy.HttpProxyProtocol {
			go func(ip string, port int, protocol string) {
				defer wg.Done()
				model.SaveProxies(proxy.CheckHttpProxy(ip, port, protocol)) //err丢失
			}(addr.IP, addr.Port, httppro)
		}

		for sockproint := range proxy.SockProxyProtocol {
			go func(ip string, port int, protocol int) {
				defer wg.Done()
				model.SaveProxies(proxy.CheckSocksProxy(ip, port, protocol))
			}(addr.IP, addr.Port, sockproint)
		}
	}
	wg.Wait()
}
