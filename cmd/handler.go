package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func hello(){
	fmt.Println("Hello Kerim")
}

var handler = &cobra.Command{
	
	Use: 	"handler",
	Short:  "Makes handler",
	Long:	"Creates handler with custom name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world")
	},
}

func init(){
	rootCmd.AddCommand(handler)
}