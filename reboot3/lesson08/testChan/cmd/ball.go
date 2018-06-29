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

// ballCmd represents the ball command
var ballCmd = &cobra.Command{
	Use:   "ball",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ball called")

		ch := make(chan int, 0)
		var wg sync.WaitGroup

		fmt.Println("start playing!!")

		wg.Add(2)
		go player("song", ch, &wg)
		go player("chen", ch, &wg)

		// setup play
		ch <- 0
		wg.Wait()
	},
}

func init() {
	// rand seed for playing
	rand.Seed(time.Now().UnixNano())

	rootCmd.AddCommand(ballCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ballCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ballCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// player for two players
func player(name string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s starting!\n", name)
	var n int
	for {
		ball, ok := <-ch
		if !ok {
			// 对方失败，关了通道
			fmt.Printf("%s won!!!\n", name)
			break
		}
		n = rand.Intn(100)
		if n%19 == 0 {
			// 自己失败
			fmt.Printf("%s miss, the number is %d\n", name, n)
			close(ch)
			break
		}
		ball++
		fmt.Printf("Player %s hit ball %d with rand %d\n", name, ball, n)
		ch <- ball
	}
}
