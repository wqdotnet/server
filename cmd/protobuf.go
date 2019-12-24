package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// protobufCmd build protobuf
var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "Short",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("protoc --proto_path=e:/worke/proto  --go_out=pro/  *.pro")

	},
}

func init() {
	rootCmd.AddCommand(protobufCmd)

}
