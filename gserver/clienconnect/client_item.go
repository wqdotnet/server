package clienconnect

import "server/proto/item"

func (c *Client) getBackpackInfo(msg *item.C2S_GetBackpackInfo) {

	c.SendToClient(int32(item.MSG_Item_Module),
		int32(item.MSG_Item_GetBackpackInfo),
		&item.S2C_GetBackpackInfo{})
}
