/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sawilkhan/gophercises-cli-tool/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil{
			log.Fatal("error parsing arguments")
		}
		err = db.DeleteTasks(taskId)
		if err != nil{
			log.Fatalf("could not delete task %v", err)
		}
		fmt.Println("task deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
