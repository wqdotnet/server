package clienconnect

import (
	"server/msgproto/bigmap"

	log "github.com/sirupsen/logrus"
)

func (c *Client) bigmapModule(method int32, buf []byte) {
	log.Info("map")
	switch bigmap.MSG_BIGMAP_Module_BIGMAP {

	case bigmap.MSG_BIGMAP_C2S_GetMapInfo:

	// case account.MSG_ACCOUNT_C2S_UpdateRoleName:
	// 	upName := &account.C2S_UpdateRoleName{}
	// 	e := proto.Unmarshal(buf, upName)
	// 	if e != nil {
	// 		log.Error(e)
	// 		return
	// 	}
	// 	c.updateRole(upName)

	default:
		log.Info("loginModule null methodID:", method)
	}
}
