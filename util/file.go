package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sayol/Sayo-proxyscan/log"
	"github.com/sayol/Sayo-proxyscan/model"
)

func ReadProxyAddr(filename string) (sliceProxyAddr []model.Input_ProxyAddr) {
	proxyFile, err := os.Open(filename)
	if err != nil {
		log.Log.Fatalf("Open proxy file err, %v", err)
	}

	defer proxyFile.Close()

	scanner := bufio.NewScanner(proxyFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipPort := scanner.Text()
		t := strings.Split(ipPort, ":")
		ip := t[0]
		port, err := strconv.Atoi(t[1])
		if err == nil {
			proxyAddr := model.Input_ProxyAddr{
				IP:   ip,
				Port: port,
			}
			sliceProxyAddr = append(sliceProxyAddr, proxyAddr)
		}
	}
	return sliceProxyAddr
}
