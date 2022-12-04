/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startBrowserCmd represents the startBrowser command
var startBrowserCmd = &cobra.Command{
	Use:   "startBrowser",
	Short: "Command used to start google-chrome",
	Long:  `Command used to start google-chrome.....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("StartBrowser Called")

		out, err := exec.Command("google-chrome").Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(startBrowserCmd)
}
