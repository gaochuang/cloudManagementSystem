// Package app /*
package app

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/middleware"
	"github.com/gaochuang/cloudManagementSystem/api/routers"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/options"
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

func NewServerCommand() *cobra.Command {
	opt := options.NewOptions()
	if nil == opt {
		log.Logger.LogError("create options failed")
	}
	cmd := &cobra.Command{
		Use:   "managerServer",
		Short: "cloud management system",
		Long:  `welcome to https://github.com/gaochuang/cloudManagementSystem`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opt.Initialize(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err := run(opt); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},

		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	opt.BindConfigurationFlag(cmd)
	return cmd
}

func run(opt *options.Options) error {
	cms.Setup(opt)
	initRouters(opt)
	if err := opt.RunHttpServer(); err != nil {
		return err
	}
	return nil
}

func initRouters(opt *options.Options) {
	routers.InitializePublicRoutes(opt.GinEngine)
	middleware.InitMiddlewares(opt.GinEngine)
	routers.InitUserRouter(opt.GinEngine)
	routers.InitSystem(opt.GinEngine)
	routers.InitRoleRouter(opt.GinEngine)
}
