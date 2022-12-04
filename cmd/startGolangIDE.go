/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startGolangIDECmd represents the startGolangIDE command
var startGolangIDECmd = &cobra.Command{
	Use:   "startGolangIDE",
	Short: "Command used to start Golang IDE",
	Long:  `Command used to start Golang IDE...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("StartGolangIDE Called")

		out, err := exec.Command("bash", "-c", "cd /home/teo/Downloads/liteide/build/liteide/bin/ && ./liteide").Output()

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
}
