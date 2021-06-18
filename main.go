package main

import (
	"gin-admin-cli/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fusion-gin-cli"
	app.Description = "Fusion web项目辅助工具，提供创建项目、快速生成功能模块的功能"
	app.Version = "4.0.1"
	app.Commands = []cli.Command{
		cmd.NewCommand(),
		cmd.GenerateCommand(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
