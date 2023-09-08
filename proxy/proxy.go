package proxy

import "time"

var (
	DEBUG_MODE = false
	TIMEOUT = time.Second * 5
	SCAN_NUM = 1000 //线程数
	INPUT_FILE = "./input/input_proxyaddr.txt"

	WebUrl  = "https://cn.bing.com/" //用于测试的网址
	Proxy_Confirm = "<title>Bing" //测试成功标识
)
