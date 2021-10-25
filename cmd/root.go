package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var rootCMD = &cobra.Command{
	Use:   "educative",
	Short: "educative is a service for managing students and their courses",
}

func init() {
	rootCMD.PersistentFlags().StringVarP(&configPath, "configPath", "c", "env.yaml", "path to config directory")
}

//Execute runs through the command tree finding appropriate matches for commands and then corresponding flags
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
