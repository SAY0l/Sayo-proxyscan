package main

import (
	"os"

	"github.com/sayol/Sayo-proxyscan/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sayo-proxyscan"
	app.Author = "Sayol"
	app.Email = "sayol@github.com"
	app.Version = "1.0"
	app.Usage = "SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS fast proxy scanner"
	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	app.Run(os.Args)
}
