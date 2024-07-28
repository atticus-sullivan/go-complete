package completeargs

type CTfile struct{
	Glob string
	OnlyDirs bool
}

func (cf CTfile) SetIdxNS(uint) CompleteTypeNoSub {return cf}
func (cf CTfile) SetIdxS(uint) CompleteTypeSub {return cf}

func (CTfile) isNoSub() {}
