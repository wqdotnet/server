package messageroute

import (
	"server/msgproto/account"

	"google.golang.org/protobuf/proto"
)

type ExecFunc func(proto.Message)

var execMap map[int]ExecFunc

func init() {
	execMap = make(map[int]ExecFunc)
}

func RegisterRouteFunc(messageID int, execfunc ExecFunc) {
	execMap[messageID] = execfunc
}

func ExecMsgRount(messageID int, pb proto.Message) {
	if fun, ok := execMap[messageID]; ok {
		fun(pb)
	}
}

func test() {

	//RegisterRouteFunc(123, pbtest)
	ExecMsgRount(123, &account.C2S_CreateRole{})

}

func pbtest(pb *account.C2S_CreateRole) {

}
