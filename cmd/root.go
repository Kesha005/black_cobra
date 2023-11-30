package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = cobra.Command{
	Use: "bc",
	Short: "Gin-Gonic helper and file generator",
	Long: "Gin Gonic controller model middleware and many other generator for easy coding",
}

func Execute(){
	if err := rootCmd.Execute();err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}