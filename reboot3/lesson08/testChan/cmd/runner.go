// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// runnerCmd represents the runner command
var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runner called")

		ch := make(chan int, 0)
		wg.Add(1)
		fmt.Println("start runner()")
		go Runner(ch)
		ch <- 1
		wg.Wait()
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(runnerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runnerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runnerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Runner for race
func Runner(ch chan int) {
	var newRunner int

	runner := <-ch
	fmt.Printf("\nrunner %d running with Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		// next runner prepared to running
		go Runner(ch)
	}

	// rand sleep time
	n := rand.Intn(4)
	time.Sleep(time.Second * time.Duration(n))
	fmt.Printf("runner %d use %d seconds to the line\n", runner, n)

	if runner == 4 {
		fmt.Printf("runner %d finish. Race over!\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("runner %d exchange with runner %d\n", runner, newRunner)
	ch <- newRunner
}
