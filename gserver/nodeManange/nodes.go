package nodemanange

import (
	"fmt"
	"server/gserver/commonstruct"
	"sync"

	"github.com/halturin/ergo"
)

var serverCfg *commonstruct.ServerConfig

var nodesMap sync.Map

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

	nodesMap.Store(serverNode.FullName, serverNode)
	nodesMap.Store(gateNode.FullName, gateNode)
	nodesMap.Store(dbNode.FullName, dbNode)
}

func GetNode(nodename string) *ergo.Node {
	if v, ok := nodesMap.Load(nodename); ok {
		return v.(*ergo.Node)
	}
	return nil
}
