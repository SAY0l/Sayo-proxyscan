# Sayo-proxyscan
轻量快速代理扫描器。

渗透、爬虫再也不怕没代理资源了

## Help
NAME:
   Sayo-proxyscan - SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS fast proxy scanner

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   1.0

AUTHOR:
   Sayol <sayol@github.com>

COMMANDS:
   scan     let's scan proxy
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                 debug mode
   --scan_num value, -n value  scan num (default: 1000)
   --timeout value, -t value   timeout (default: 5)
   --filename value, -f value  filename (default: "./input/input_proxyaddr.txt")
   --output_file value, -o value  output_file (default: "sayo-proxyscan.txt")
   --help, -h                  show help
   --version, -v               print the version

## Use
当前支持 SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS 代理扫描

编辑代理验证配置见 ./proxy/proxy.go
## Ver 1.0.0 优化
1. 优化代理验证方案。
2. 解耦代码架构。

## Ver 1.1.0 待优化
1. 支持多代理验证方案，提高代理实时性与准确性
2. 优化并发算法
3. 输出更加友好，可显示代理地理位置