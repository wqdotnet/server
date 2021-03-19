package cfg

//SkillCfg 技能
type SkillCfg struct {
	ID                   int32     `json:"id"`
	TroopsName           string    `json:"troopsName"`
	SkillName            string    `json:"skillName"`
	SkillIcon            string    `json:"skillIcon"`
	SkiilIntro           string    `json:"skiilIntro"`
	Skilltype            int32     `json:"skilltype"`
	UseTimeType          int32     `json:"useTimeType"`
	UseTime              int32     `json:"useTime"`
	UseEffects1          string    `json:"useEffects1"`
	BeforeEffects2       string    `json:"beforeEffects2"`
	BallisticEffects     string    `json:"ballisticEffects"`
	StruckEffect         string    `json:"struckEffect"`
	SceneEffects         string    `json:"SceneEffects"`
	EffectType           int32     `json:"effectType"`
	EffectSubtypes1      int32     `json:"effectSubtypes1"`
	Key1                 float64   `json:"key1"`
	Value1               float64   `json:"value1"`
	CorrectionConditions float64   `json:"correctionConditions"`
	ConditionsValue      []float64 `json:"conditionsValue"`
	Correct              []float64 `json:"correct"`
	Range                int32     `json:"range"`
	AddSkillsID          []int32   `json:"addSkillsID"`
	AddbuffID            []int32   `json:"addbuffID"`
}

//SillLandformCfg 地形技能
type SillLandformCfg struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       int    `json:"type"`
	Lv         int    `json:"lv"`
	Icon       string `json:"icon"`
	Intro      string `json:"intro"`
	BaseDamage int    `json:"baseDamage"`
}

//BuffCfg buff
type BuffCfg struct {
	ID                   int32   `json:"id"`
	TroopsName           string  `json:"troopsName"`
	BuffName             string  `json:"BuffName"`
	SkillIcon            string  `json:"skillIcon"`
	SkiilIntro           string  `json:"skiilIntro"`
	Bufftype             string  `json:"bufftype"`
	Bufflv               string  `json:"bufflv"`
	RoundNumber          string  `json:"roundNumber"`
	RoleEffects          string  `json:"RoleEffects"`
	SceneEffects         string  `json:"SceneEffects"`
	EffectType           int32   `json:"effectType"`
	EffectSubtypes1      int32   `json:"effectSubtypes1"`
	Key1                 int32   `json:"key1"`
	Value1               float64 `json:"value1"`
	CorrectionConditions string  `json:"correctionConditions"`
	ConditionsValue      string  `json:"conditionsValue"`
	Correct              string  `json:"correct"`
	Range                int32   `json:"range"`
	RemoveRoundNum       int32   `json:"removeRoundNum"`
	Addbuffid            string  `json:"addbuffid"`
}

//GetSkillCfg 技能(战法)配置
func GetSkillCfg(skillid int32) *SkillCfg {
	for _, v := range GameCfg.Skill.SkillList {
		if skillid == int32(v.ID) {
			return &v
		}
	}
	return nil
}

//GetSkillTacticsCfg (战术)配置
func GetSkillTacticsCfg(tacticsid int32) *SillLandformCfg {
	for _, v := range GameCfg.Skill.SkillLandform {
		if tacticsid == int32(v.ID) {
			return &v
		}
	}
	return nil
}

//GetBuffCfg GetBuff config
func GetBuffCfg(buffid int32) *BuffCfg {
	for _, cfg := range GameCfg.Skill.BuffList {
		if cfg.ID == buffid {
			return &cfg
		}
	}
	return &BuffCfg{}
}
