package main

import (
	"cmdb/pkg/config"
	"cmdb/pkg/db"
	"cmdb/pkg/web_api/router"
	"cmdb/pkg/web_api/server_init"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("c", "config.yaml", "配置目录")
	flag.Parse()
	config.ParserConfig(*configPath)
	db.InitMySQL()
	db.InitRedis()
	server_init.Init()
	r := gin.Default()
	router.RegistryRouter(r)
	server_init.InitPermission(r.Routes())
	r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
}
