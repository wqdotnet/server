package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test demo",
	Long:  `test demo`,
	Run:   exectest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func exectest(cmd *cobra.Command, args []string) {

	// info := &account.C2S_CreateRole{RoleName: "sdfhowme"}
	// buf, _ := proto.Marshal(info)

}

//protobuf 解码
func decodeProto(info interface{}, buf []byte) error {
	if data, ok := info.(protoreflect.ProtoMessage); ok {
		return proto.Unmarshal(buf, data)
	}
	return nil
}
