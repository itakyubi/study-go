package main

import (
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"study-go/web/api"
	"study-go/web/common"
	"study-go/web/config"
	"study-go/web/context"
	"study-go/web/log"
	"study-go/web/plugin"
	_ "study-go/web/plugin/database"
	"study-go/web/server"
)

func main() {
	defer plugin.ClosePlugins()
	runtime.GOMAXPROCS(runtime.NumCPU())

	context.Run(func(ctx context.Context) error {
		// 加载配置
		var cfg config.Config
		err := ctx.LoadCustomConfig(&cfg)
		if err != nil {
			return err
		}
		ctx.Log().Debug("cloud config", log.Any("cfg", cfg))
		common.SetConfFile(ctx.ConfFile())

		// 初始化api
		a, err := api.NewAPI(&cfg)
		if err != nil {
			return err
		}

		// 初始化server
		s, err := server.NewAdminServer(&cfg)
		if err != nil {
			return err
		}
		s.SetAPI(a)
		s.InitRoute()
		defer s.Close()

		// 启动server
		go s.Run()
		ctx.Log().Info("admin server starting")

		ctx.Wait()
		return nil
	})
}
