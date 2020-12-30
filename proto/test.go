package msg

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
)

func test() {
	p := &sync.Pool{
		//New: connectToService,
	}
	p.Get()
}

func testMarshal(pb proto.Message, outpb proto.Message) {
	// encode  msg
	data, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	msginfo := &NetworkMsg{}
	msginfo.Module = 1
	msginfo.Method = 2
	msginfo.MsgBytes = data
	msgdata, err := proto.Marshal(msginfo)
	if err != nil {
		fmt.Printf("msg encode error[%s]\n", err.Error())
	}

	// decode  msg
	msginfo2 := &NetworkMsg{}
	err = proto.Unmarshal(msgdata, msginfo2)
	if err != nil {
		fmt.Printf("msg decode error[%s]\n", err.Error())
		return
	}

	err = proto.Unmarshal(msginfo2.MsgBytes, outpb)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", err.Error())
		return
	}

}

// func ProtobufTest() {
// 	searchRequest := &SearchRequest{}
// 	searchRequest.Query = "select *from query"
// 	searchRequest.PageNumber = 4
// 	searchRequest.ResultPerPage = 4

// 	searchRequest2 := &SearchRequest{}
// 	testMarshal(searchRequest, searchRequest2)

// 	fmt.Printf("proto decode searchRequest: [%s]\n", searchRequest2)
// }
