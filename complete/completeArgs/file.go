package completeargs

type CTfile struct{
	Glob string
	OnlyDirs bool
}

func (CTfile) isNoSub() {}
