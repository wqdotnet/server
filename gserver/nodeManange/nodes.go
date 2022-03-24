package nodeManange

import (
	"fmt"
	"server/gserver/commonstruct"
	"sync"

	"github.com/ergo-services/ergo/node"
)

//ergo.Node 节点管理
var (
	nodesMap sync.Map
	//remoteMap sync.Map  //远程连接节点
)

func Start(command chan string) {
	for _, v := range commonstruct.ServerCfg.StartList {
		switch v {
		case "gateway":
			gateNodeName := fmt.Sprintf("gatewayNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
			gateNode, _, gerr := StartGateSupNode(gateNodeName)
			if gerr != nil {
				panic(gerr)
			}
			nodesMap.Store(gateNode.NodeName(), gateNode)
		case "server":
			serverNodeName := fmt.Sprintf("serverNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
			serverNode, _, serr := StartGameServerSupNode(serverNodeName, command)
			if serr != nil {
				panic(serr)
			}
			nodesMap.Store(serverNode.NodeName(), serverNode)
		case "db":
			dbNodeName := fmt.Sprintf("dbNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
			dbNode, _, derr := StartDataBaseSupSupNode(dbNodeName)
			if derr != nil {
				panic(derr)
			}
			nodesMap.Store(dbNode.NodeName(), dbNode)
		}

	}

}

func GetNode(nodename string) node.Node {
	if v, ok := nodesMap.Load(nodename); ok {
		return v.(node.Node)
	}
	return nil
}

func GetNodes() map[string]node.Node {
	nodemap := map[string]node.Node{}
	nodesMap.Range(func(key, value interface{}) bool {
		nodemap[key.(string)] = value.(node.Node)
		return true
	})
	return nodemap
}
