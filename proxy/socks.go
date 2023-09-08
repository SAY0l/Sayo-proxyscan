package proxy

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sayol/Sayo-proxyscan/model"
	"github.com/sayol/Sayo-proxyscan/log"
	"h12.io/socks"
)

var (
	SockProxyProtocol = map[int]string{socks.SOCKS4: "SOCKS4", socks.SOCKS4A: "SOCKS4A", socks.SOCKS5: "SOCKS5"}
)

func CheckSocksProxy(ip string, port, protocol int) (isProxy bool, proxyInfo model.ProxyInfo, err error) {
	proxyInfo.Addr = ip
	proxyInfo.Port = port
	proxyInfo.Protocol = SockProxyProtocol[protocol]

	proxy := fmt.Sprintf("%v:%v", ip, port)
	dailproxy := socks.DialSocksProxy(protocol, proxy)
	tr := &http.Transport{Dial: dailproxy}
	httpclient := &http.Client{Transport: tr, Timeout: TIMEOUT}
	log.Log.Debugf("Checking proxy: %v", fmt.Sprintf("%v://%v:%v", SockProxyProtocol[protocol], ip, port))
	resp, err := httpclient.Get(WebUrl)
	if err == nil {
		if resp.StatusCode == http.StatusOK {
			respBody, err := io.ReadAll(resp.Body)
			if err == nil && strings.Contains(string(respBody), Proxy_Confirm) {
				isProxy = true
			}
		}
	}
	return isProxy, proxyInfo, err
}
