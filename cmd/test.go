package cmd

import (
	"fmt"
	"server/msgproto/account"

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

	register(101, msghand)

	info := &account.C2S_CreateRole{RoleName: "sdfhowme"}
	buf, _ := proto.Marshal(info)

	infofunc[101](buf)

	tmp := map[int]protoreflect.ProtoMessage{}
	tmp[101] = info

}

func msghand(msg *account.C2S_CreateRole) {
	fmt.Println("C2S_CreateRole: ", msg.RoleName)
}

var infofunc = map[int]func(buf []byte){}

//消息注册
func register[T any](moduleID int, execfunc func(*T)) {
	infofunc[moduleID] = func(buf []byte) {
		info := new(T)
		err := decodeProto(info, buf)
		if err != nil {
			fmt.Println("decodeProto2", err)
		} else {
			execfunc(info)
		}
	}

}

//protobuf 解码
func decodeProto(info interface{}, buf []byte) error {
	if data, ok := info.(protoreflect.ProtoMessage); ok {
		return proto.Unmarshal(buf, data)
	}
	return nil
}
