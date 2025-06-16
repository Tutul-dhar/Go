/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"os"
	"github.com/spf13/cobra"
)

type User struct {
	Name string   `json:"name"`
	Age  int       `json:"age"`
}

var (
	name   string
	age    int  
	format string
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format user data in text or JSON format",
	Long: `This command takes user name and age, and outputs the result 
either in plain text or JSON format, based on the --format flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := User{
			Name: name,
			Age:  age,
		}

		switch format {
		case "json":
			output, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		case "text":
			fmt.Printf("User Name: %s\nUser Age: %d\n", user.Name, user.Age)
		default:
			fmt.Println("Unknown format. Use 'json' or 'text'.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	formatCmd.Flags().StringVarP(&name, "name","n","","Your name")
	formatCmd.Flags().IntVarP(&age, "age", "a",0, "Your age")

	formatCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: json or text")

	// Mark 'name' and 'age' as required if needed
	formatCmd.MarkFlagRequired("name")
	formatCmd.MarkFlagRequired("age")
	
}
