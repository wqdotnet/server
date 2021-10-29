package nodeManange

import (
	"fmt"
	"server/gserver/commonstruct"
	"sync"

	"github.com/ergo-services/ergo/node"
)

var serverCfg *commonstruct.ServerConfig

//ergo.Node 节点管理
var (
	nodesMap sync.Map
	//remoteMap sync.Map  //远程连接节点
)

func Start(cfg *commonstruct.ServerConfig, command chan string) {
	serverCfg = cfg

	gateNodeName := fmt.Sprintf("gatewayNode_%v@127.0.0.1", serverCfg.ServerID)
	serverNodeName := fmt.Sprintf("serverNode_%v@127.0.0.1", serverCfg.ServerID)
	dbNodeName := fmt.Sprintf("dbNode_%v@127.0.0.1", serverCfg.ServerID)

	serverNode, _, serr := StartGameServerSupNode(serverNodeName, command)
	if serr != nil {
		panic(serr)
	}
	gateNode, _, gerr := StartGateSupNode(gateNodeName)
	if gerr != nil {
		panic(gerr)
	}
	dbNode, _, derr := StartDataBaseSupSupNode(dbNodeName)
	if derr != nil {
		panic(derr)
	}

	nodesMap.Store(serverNode.NodeName(), serverNode)
	nodesMap.Store(gateNode.NodeName(), gateNode)
	nodesMap.Store(dbNode.NodeName(), dbNode)
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
