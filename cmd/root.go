package cmd

import (
	"fmt"
	"log"

	"github.com/sawilkhan/gophercises-cli-tool/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "task",
	Short: "simple cli tool",
	Long: "simple cli task manager tool built with go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chal gaya mai")
	},
}

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		fmt.Println("in Execute() function")
		log.Fatal(err)
	}
}

func init(){
	cobra.OnInitialize(initDB)
}

func initDB() {
    if err := db.InitDB(); err != nil {
        log.Fatalf("Could not initialize database: %v", err)
    }
	fmt.Println("db initialized")
}