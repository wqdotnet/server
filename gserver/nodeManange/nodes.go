package nodeManange

import (
	"fmt"
	"server/gserver/commonstruct"
	"sync"

	"github.com/halturin/ergo"
)

var serverCfg *commonstruct.ServerConfig

//ergo.Node 节点管理
var NodeManage = &nodeManange{}

type nodeManange struct {
	nodesMap sync.Map
}

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

	NodeManage.nodesMap.Store(serverNode.FullName, serverNode)
	NodeManage.nodesMap.Store(gateNode.FullName, gateNode)
	NodeManage.nodesMap.Store(dbNode.FullName, dbNode)
}

func GetNode(nodename string) *ergo.Node {
	if v, ok := NodeManage.nodesMap.Load(nodename); ok {
		return v.(*ergo.Node)
	}
	return nil
}

func GetNodes() map[string]*ergo.Node {
	nodemap := map[string]*ergo.Node{}
	NodeManage.nodesMap.Range(func(key, value interface{}) bool {
		nodemap[key.(string)] = value.(*ergo.Node)
		return true
	})
	return nodemap
}
