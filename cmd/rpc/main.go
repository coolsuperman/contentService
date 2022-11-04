package rpc

import (
	"github.com/spf13/cobra"
)

var RpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "文章系统Rpc服务",
	Long:  "练习做文章系统Rpc服务",
	Run: func(cmd *cobra.Command, args []string) {
		rpc, cleanup, err := wireRpc()
		if err != nil {
			panic(err)
		}
		defer cleanup()
		rpc.Run()
	},
}
