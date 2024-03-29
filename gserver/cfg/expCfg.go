package cfg

type ExpLvInfoCfg struct {
	Level            int     `json:"level"`
	NeedExp          int64   `json:"needExp"`
	FailExp          int64   `json:"failExp"`
	MaxExp           int64   `json:"maxExp"`
	BigLevel         string  `json:"bigLevel"`
	HeadSize         int     `json:"headSize"`
	CycleEXP         int64   `json:"cycleEXP"`
	ExpWeight        []int   `json:"expWeight"`
	ExpWeightValue   [][]int `json:"expWeightValue"`
	ExpUp            int     `json:"expUp"`
	PropertiesID     []int   `json:"propertiesId"`
	AttributeValues  []int   `json:"attributeValues"`
	Times            int     `json:"times"`
	Properties       []int   `json:"properties"`
	PropertiesWeight []int   `json:"propertiesWeight"`
	CycleProperties  []int   `json:"cycleProperties"`
	PropertiesMax    []int   `json:"propertiesMax"`
	TribulationID    int     `json:"tribulationId"`
	ShowNum          []int   `json:"showNum"`
	IDGroup          int     `json:"idGroup"`
	MailID           int     `json:"mailId"`
}

func GetLvExpInfo(lv int32) *ExpLvInfoCfg {
	for _, v := range GetGameCfg().ExpXiufaInfo {
		if v.Level == int(lv) {
			return v
		}
	}
	return nil
}
