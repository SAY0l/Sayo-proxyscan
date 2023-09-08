package cmd

import (
	"github.com/sayol/Sayo-proxyscan/util"
	"github.com/urfave/cli"
)

var Scan = cli.Command{
	Name : "scan",
	Usage: "let's scan proxy",
	Description: "let's scan proxy",
	Action: util.Scan,
	Flags: []cli.Flag{
		boolFlag("debug, d","debug mode"),
		intFlag("scan_num, n",1000,"scan num"),
		intFlag("timeout, t",5,"timeout"),
		stringFlag("filename ,f","./input/input_proxyaddr.txt","filename"),
		stringFlag("output_file, o", "sayo-proxyscan.txt", "output_file"),
	},
}

func stringFlag(name,value,usage string)(cli.StringFlag){
	return cli.StringFlag{
		Name : name ,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name,usage string)(cli.BoolFlag){
	return cli.BoolFlag{
		Name: name ,
		Usage: usage,
	}
}

func intFlag(name string,value int, usage string)(cli.IntFlag) {
	return cli.IntFlag{
		Name: name,
		Value: value,
		Usage: usage,
	}
}