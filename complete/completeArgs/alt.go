package completeargs

type CTalternatives struct {
	Alts []string
}

func (ca CTalternatives) SetIdxNS(uint) CompleteTypeNoSub {return ca}
func (ca CTalternatives) SetIdxS(uint) CompleteTypeSub {return ca}

func (CTalternatives) isNoSub() {}
