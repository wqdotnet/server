package cfg

//TroopsCfg 部队配置
// type TroopsCfg struct {
// 	CfgList []struct {
// 		ID           int    `json:"Id"`
// 		Name         string `json:"Name"`
// 		Pic          string `json:"Pic"`
// 		Intro        string `json:"Intro"`
// 		Type         int    `json:"Type"`
// 		Quality      int    `json:"Quality"`
// 		Leader       int    `json:"Leader"`
// 		Strength     int    `json:"Strength"`
// 		Politics     int    `json:"Politics"`
// 		Intelligence int    `json:"Intelligence"`
// 		OriginAtt    int    `json:"OriginAtt"`
// 		OriginDef    int    `json:"OriginDef"`
// 		OriginBlood  int    `json:"OriginBlood"`
// 		Trop         string `json:"Trop"`
// 		SkillID      string `json:"SkillId"`
// 		TalentID     string `json:"TalentId"`
// 		StrategyID   string `json:"StrategyId"`
// 	} `json:"CfgList"`
// }

//TroopsCfg 部队配置
type TroopsCfg struct {
	ID           int    `json:"Id"`
	Name         string `json:"Name"`
	Pic          string `json:"Pic"`
	Intro        string `json:"Intro"`
	Type         int    `json:"Type"`
	Quality      int    `json:"Quality"`
	Leader       int    `json:"Leader"`
	Strength     int    `json:"Strength"`
	Politics     int    `json:"Politics"`
	Intelligence int    `json:"Intelligence"`
	OriginAtt    int    `json:"OriginAtt"`
	OriginDef    int    `json:"OriginDef"`
	OriginBlood  int    `json:"OriginBlood"`
	Trop         string `json:"Trop"`
	SkillID      int    `json:"SkillId"`
	TalentID     string `json:"TalentId"`
	StrategyID   []int  `json:"StrategyId"`
}

//GetTroopsCfg TroopsCfg struct
func GetTroopsCfg(id int32) *TroopsCfg {
	for _, v := range GameCfg.Troops.CfgList {
		if id == int32(v.ID) {
			return &v
		}
	}
	return nil
}
