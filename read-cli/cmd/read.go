/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)


var reminderDuration time.Duration
// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read a file and remaind after a certain time",
	Long: `This command reads a text file and remind you if you are still reading after a given time duration. `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]

		if reminderDuration > 0 {
			go func ()  {
				time.Sleep(reminderDuration)
				fmt.Printf(`\n you have been reading for %v. Take a break\n`,reminderDuration)
			} ()
		}

		data, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Printf("could not read file %d", err)
			return
		}

		fmt.Println("file content ")
		fmt.Println(string(data))

		fmt.Println("press enter when you are done")


		fmt.Scanln()

		
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

    // Add the reminder flag (e.g. --reminder=2m)
    readCmd.Flags().DurationVarP(&reminderDuration, "reminder", "r", 0, "Duration after which to show a reminder (e.g. 2m, 30s)")

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // randomCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
