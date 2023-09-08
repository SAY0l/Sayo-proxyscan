package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sayol/Sayo-proxyscan/model"
	"github.com/sayol/Sayo-proxyscan/log"
)

var (
	HttpProxyProtocol = []string{"http", "https"}
)

func CheckHttpProxy(ip string, port int, protocol string) (isProxy bool, proxyInfo model.ProxyInfo, err error) {
	proxyInfo.Addr = ip
	proxyInfo.Port = port
	proxyInfo.Protocol = protocol

	rawProxyUrl := fmt.Sprintf("%v://%v:%v", protocol, ip, port)
	proxyUrl, err := url.Parse(rawProxyUrl)
	if err == nil {
		Transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client := &http.Client{Transport: Transport, Timeout: TIMEOUT}
		log.Log.Debugf("Checking proxy: %v", rawProxyUrl)
		resp, err := client.Get(WebUrl)
		if err == nil {
			if resp.StatusCode == http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err == nil && strings.Contains(string(body), Proxy_Confirm) {
					isProxy = true
				}
			}
		}
	}
	return isProxy, proxyInfo, err
}
