package main

import (
	"contentService/cmd/api"
	"contentService/cmd/rpc"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "content-serv",
	Short: "文章服务",
	Long:  `练习文章服务`,
}

func init() {
	rootCmd.AddCommand(api.ApiCmd)
	rootCmd.AddCommand(rpc.RpcCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
