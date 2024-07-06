package config

import "github.com/spf13/cobra"

type cmdArgs struct {
	// 是否发开项目选项，会校验用户项目权限和cookie中项目
	//Project bool
}

var CmdArgs cmdArgs

func InitCmdArgs(cmd *cobra.Command) {
	//cmd.Flags().BoolVar(&CmdArgs.Project, "p", false, "是否启用项目")
}
