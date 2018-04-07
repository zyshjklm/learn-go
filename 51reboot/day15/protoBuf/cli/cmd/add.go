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
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add Id Name Email Mobile",
	Long: `add Person info into address book store.

Only support one Mobile.
segment order: Id Name Email Mobile`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called, args:", args)

		var phones []*rpcproto.PhoneNumber
		id, _ := strconv.Atoi(args[0])
		phone := &rpcproto.PhoneNumber{
			Number: args[3],
			Type:   rpcproto.PhoneType_MOBILE,
		}
		phones = append(phones, phone)
		req := &rpcproto.AddPersonRequest{
			Person: &rpcproto.Person{
				Id:     int32(id),
				Name:   args[1],
				Email:  args[2],
				Phones: phones,
			},
		}

		client := newClient("127.0.0.1:8021")
		resp, err := client.AddPerson(context.TODO(), req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("add ok, id:", resp.GetId())
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
