/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/sawilkhan/gophercises-cli-tool/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasksList, err := db.ListTasks()
		if err != nil{
			log.Fatal(err)
		}
		
		for _, task := range tasksList{
			fmt.Printf("%d. %s", task.ID, task.Name)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
