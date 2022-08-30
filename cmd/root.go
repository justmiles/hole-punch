/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	src "github.com/justmiles/hole-punch/src"
)

// defaults
var (
	shell = src.DefaultShell()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hp",
	Short: "punch a hole, yo",
	Long: `Punch a hole! This tool can be used to quickly 
produce a reverse SSH tunnel on any machine.

Example usage:

  # on the target machine
  hp --key id_rsa --remote user@172.217.164.78:2455

  # on the remote
  ssh -i id_rsa -p 3333 127.0.0.1

`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := os.ReadFile(viper.GetString("key"))
		if err != nil {
			fmt.Printf("error reading private key: %v\n", err)
			os.Exit(1)
		}

		err = src.Run(
			src.WithPrivateKey(key),
			src.WithShell(viper.GetString("shell")),
			src.WithLocalEndpoint(viper.GetString("local")),
			src.WithRemoteEndpoint(viper.GetString("remote")),
			src.WithTunnelEndpoint(viper.GetString("tunnel")),
		)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringP("key", "k", "", "path to SSH private key")
	rootCmd.PersistentFlags().StringP("remote", "r", "", "remote SSH endpoint to connect to")
	rootCmd.PersistentFlags().StringP("local", "l", "127.0.0.1:2222", "local SSH endpoint to create")
	rootCmd.PersistentFlags().StringP("tunnel", "t", "127.0.0.1:3333", "tunnel to create on remote endpoint")
	rootCmd.PersistentFlags().StringP("shell", "s", shell, "shell to invoke")

	rootCmd.MarkFlagRequired("key")
	rootCmd.MarkFlagRequired("remote")

	cobra.OnInitialize(initViper)
}

func initViper() {

	viper.SetEnvPrefix("hp")
	viper.AutomaticEnv()

	// Bind Viper flags
	viper.BindPFlags(rootCmd.PersistentFlags())

	// Handle required flags
	var requiredFlags = []string{"key", "remote"}
	var missingFlags []string

	for _, flag := range requiredFlags {
		flagValue := viper.Get(flag)

		if flagValue == nil || flagValue == "" {
			missingFlags = append(missingFlags, flag)
		}
	}

	if len(missingFlags) > 0 {
		rootCmd.Help()

		rootCmd.PrintErrf("\nMissing required flags: %v\n", missingFlags)
		os.Exit(1)
	}

}
