/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sawilkhan/gophercises-cli-tool/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		db.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
