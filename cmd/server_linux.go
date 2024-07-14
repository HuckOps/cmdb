/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmdb/docs"
	"cmdb/pkg/logger"
	"cmdb/pkg/mysql"
	"cmdb/src/config"
	"cmdb/src/model"
	"cmdb/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化框架日志器
		docs.SwaggerInfo.Title = "CMDB API"
		logger.InitLogger()
		mysql.InitMySQL()
		model.Migrate(mysql.GormClient)
		srv := gin.Default()
		srv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routes.RootRoutes(srv)
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println("Error listening:", err)
			return
		}
		//srv.Run(":8081")
		defer l.Close()
		if tcpListener, ok := l.(*net.TCPListener); ok {
			// 设置SO_REUSEPORT选项
			if err := tcpListener.SetsockoptInt(net.SOL_SOCKET, net.SO_REUSEPORT, 1); err != nil {
				fmt.Println("Error setting SO_REUSEPORT:", err)
				return
			}
		} else {
			fmt.Println("Listener is not a TCP listener")
			return
		}
		server := &http.Server{
			Handler: srv,
		}
		if err := server.Serve(l); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error serving:", err)
			return
		}
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)
	config.InitCmdArgs(serverCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
