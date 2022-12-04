/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startGolangIDECmd represents the startGolangIDE command
var startGolangIDECmd = &cobra.Command{
	Use:   "startGolangIDE",
	Short: "Command used to start Golang IDE",
	Long:  `Command used to start Golang IDE...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("StartGolangIDE Called")

		mainCommand := "cd " + fmt.Sprint(viper.Get("IDE_PATH")) + " && ./liteide"

		out, err := exec.Command("bash", "-c", mainCommand).Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(startGolangIDECmd)
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()
}
