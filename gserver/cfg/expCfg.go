package cfg

//ExpCfg exp
type ExpCfg struct {
	Level     int   `json:"Level"`
	RoleExp   int64 `json:"RoleExp"`
	TroopsExp int64 `json:"TroopsExp"`
	UnlockID  int   `json:"UnlockID"`
}

//GetExp level up exp
func GetExp(level int) *ExpCfg {
	for _, v := range GameCfg.RoleExp.ExpList {
		if level == v.Level {
			return &v
		}
	}
	return nil
}

//AddRoleExp add exp
func AddRoleExp(level, exp, addexp int64) (newlevel int64, newExp int64) {
	//log.Info(level, exp, addexp)
	expCfg := GetExp(int(level))
	if expCfg == nil {
		return level, exp
	}
	if exp+addexp >= expCfg.RoleExp {
		return AddRoleExp(level+1, 0, (exp+addexp)-expCfg.RoleExp)
	}
	return level, exp + addexp
}

//AddTroopsExp add troops exp
func AddTroopsExp(level, exp, addexp int64) (newlevel int64, newExp int64) {
	expCfg := GetExp(int(level))
	if expCfg == nil {
		return level, exp
	}
	if exp+addexp >= expCfg.TroopsExp {
		return AddTroopsExp(level+1, 0, (exp+addexp)-expCfg.TroopsExp)
	}
	return level, exp + addexp
}
