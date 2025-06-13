/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	// "log"
	// "net/http"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID    string `json :"id"`
	Joke  string `json:"joke"`
	Status  int  `json:"status"`
}

func getRandomJoke() {
	fmt.Println("Get random joke:p")
}


/*

tutul@tutul-dhar-PC:~$ cd Documents
tutul@tutul-dhar-PC:~/Documents$ cd Go
tutul@tutul-dhar-PC:~/Documents/Go$ cd Go
tutul@tutul-dhar-PC:~/Documents/Go/Go$ cd concurrency
tutul@tutul-dhar-PC:~/Documents/Go/Go/concurrency$ cd ..
tutul@tutul-dhar-PC:~/Documents/Go/Go$ mkdir dadjoke-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go$ cd dadjoke-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ go mod init github.com/Tutul-dhar/dadjoke-cli
go: creating new go.mod: module github.com/Tutul-dhar/dadjoke-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ cobra-cli init
Command 'cobra-cli' not found, but can be installed with:
sudo apt install cobra-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ sudo apt install cobra-cli
[sudo] password for tutul: 
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
The following NEW packages will be installed:
  cobra-cli
0 upgraded, 1 newly installed, 0 to remove and 21 not upgraded.
Need to get 2,276 kB of archives.
After this operation, 6,594 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu noble-updates/universe amd64 cobra-cli amd64 1.3.0-3ubuntu0.24.04.2 [2,276 kB]
Fetched 2,276 kB in 2s (963 kB/s)      
Selecting previously unselected package cobra-cli.
(Reading database ... 201404 files and directories currently installed.)
Preparing to unpack .../cobra-cli_1.3.0-3ubuntu0.24.04.2_amd64.deb ...
Unpacking cobra-cli (1.3.0-3ubuntu0.24.04.2) ...
Setting up cobra-cli (1.3.0-3ubuntu0.24.04.2) ...
Processing triggers for man-db (2.12.0-4build2) ...
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ cobra-cli init
Your Cobra application is ready at
/home/tutul/Documents/Go/Go/dadjoke-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ go run main.go
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ go run main.go
this is the first try

tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ cobra-cli add random
random created at /home/tutul/Documents/Go/Go/dadjoke-cli
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ ^M
: command not found
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ curl -H "Accept: application/json" https://icanhazdadjoke.com/j/R7UfaahVfFd
{"id":"R7UfaahVfFd","joke":"My dog used to chase people on a bike a lot. It got so bad I had to take his bike away.","status":200}
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ go run main.go random
# github.com/Tutul-dhar/dadjoke-cli/cmd
cmd/random.go:9:2: "log" imported and not used
cmd/random.go:10:2: "net/http" imported and not used
tutul@tutul-dhar-PC:~/Documents/Go/Go/dadjoke-cli$ go run main.go random
Get random joke:p


*/