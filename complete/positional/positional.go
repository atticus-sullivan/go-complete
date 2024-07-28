package positional

import (
	"github.com/atticus-sullivan/go-complete/complete/completeArgs"
)

type Positional struct {
	Idx      uint
	Help     string
	Arg      completeargs.CompleteTypeSub
	Optional bool
}
