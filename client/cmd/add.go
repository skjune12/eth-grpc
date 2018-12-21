// Copyright © 2018 Kohei Suzuki <jingle@sfc.wide.ad.jp>
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
	"log"
	"os"
	"strconv"

	"github.com/skjune12/grpc-eth/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Set a value to the smart contract",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) < 2 {
			log.Fatal("invalid arguments")
		}

		// setup gRPC conn
		conn, err := grpc.Dial("localhost:4567", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial: %s\n", err)
		}
		defer conn.Close()

		// setup gRPC client
		grpcClient := api.NewExampleClient(conn)

		value, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		response, err := grpcClient.Exec(context.Background(), &api.TestMsg{Method: api.ADD, Value: int32(value)})
		if err != nil {
			log.Fatalf("Error when calling grpcClient.Exec: %s\n", err)
		}

		log.Printf("Response from server: %s\n", response.Msg)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
