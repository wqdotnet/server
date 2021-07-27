package tools

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	str := "这里是 www\n.runoob\n.com"
	fmt.Println(StringReplace(str))
}
