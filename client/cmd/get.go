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

	"github.com/skjune12/grpc-eth/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the smart contract",
	Long:  `Get a value from the smart contract`,
	Run: func(cmd *cobra.Command, args []string) {
		// setup gRPC conn
		conn, err := grpc.Dial("localhost:4567", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial: %s\n", err)
		}
		defer conn.Close()

		// setup gRPC client
		grpcClient := api.NewExampleClient(conn)

		response, err := grpcClient.Exec(context.Background(), &api.TestMsg{Method: api.GET})
		if err != nil {
			log.Fatalf("Error when calling grpcClient.Exec: %s\n", err)
		}

		log.Printf("Response from server: %s\n", response.Msg)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
