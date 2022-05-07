package clienconnect

import pbitem "server/proto/item"

func (c *Client) getBackpackInfo(msg *pbitem.C2S_GetBackpackInfo) {

	c.SendToClient(int32(pbitem.MSG_ITEM_Module),
		int32(pbitem.MSG_ITEM_GetBackpackInfo),
		&pbitem.S2C_GetBackpackInfo{})
}
