package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/sawilkhan/gophercises-cli-tool/db"
	"github.com/spf13/cobra"
)
var addCmd = &cobra.Command{
	Use: "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string){
		taskName := strings.Join(args, " ")
        task := db.Task{
            Name: taskName,
        }
        if err := db.AddTask(task); err != nil {
            log.Fatalf("Could not add task: %v", err)
        }
        fmt.Println("Task added:", taskName)
	},

}

func init(){
	rootCmd.AddCommand(addCmd)
}