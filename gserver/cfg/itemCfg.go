package cfg

type ItemInfoCfg struct {
	ID           uint32   `json:"id"`
	Type         uint32   `json:"type"`
	Level        uint32   `json:"level"`
	Quality      uint32   `json:"quality"`
	Name         string   `json:"name"`
	BagType      uint32   `json:"bagType"`
	OverlayLimit uint32   `json:"overlayLimit"`  //堆叠上限
	SellItemID   uint32   `json:"sell_item_id"`  //售卖物品ID
	SellItemNum  uint32   `json:"sell_item_num"` //售卖货币数量
	Usetype      uint32   `json:"usetype"`
	Data1        uint32   `json:"data1"`
	Data2        uint32   `json:"data2"`
	IsUse        uint32   `json:"isUse"` //使用后消失
	AttrCreat    []uint32 `json:"attrCreat"`
}

func GetItemCfg(id uint32) *ItemInfoCfg {
	itemlist := GetGameCfg().ItemInfo
	for _, iic := range itemlist {
		if iic.ID == id {
			return iic
		}
	}
	return nil
}
