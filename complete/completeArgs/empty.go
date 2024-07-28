package completeargs

type CTempty struct{}

func (ce CTempty) SetIdxNS(uint) CompleteTypeNoSub {return ce}
func (ce CTempty) SetIdxS(uint) CompleteTypeSub {return ce}

func (CTempty) isNoSub() {}
