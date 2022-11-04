package api

import (
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "文章系统API服务",
	Long:  "练习做文章系统API服务",
	Run: func(cmd *cobra.Command, args []string) {
		api, cleanup, err := wireApi()
		if err != nil {
			panic(err)
		}
		defer cleanup()

		if err := api.Run(); err != nil {
			panic(err)
		}
	},
}
