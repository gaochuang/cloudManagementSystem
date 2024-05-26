// Package app /*
package app

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/options"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"os"

	"github.com/spf13/cobra"
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
	if err := opt.RunHttpServer(); err != nil {
		return err
	}
	return nil
}
