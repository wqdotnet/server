package msg

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

// type Message interface {
// 	Reset()
// 	String() string
// 	ProtoMessage()
// }

func testMarshal(pb Message) {

	// encode
	data, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	// decode
	searchRequest2 := &SearchRequest{}
	err = proto.Unmarshal(data, searchRequest2)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", err.Error())
		return
	}

}

func protobufTest() {
	searchRequest := &SearchRequest{}

	// encode
	data, err := proto.Marshal(searchRequest)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	// decode
	searchRequest2 := &SearchRequest{}
	err = proto.Unmarshal(data, searchRequest2)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", err.Error())
		return
	}

}
