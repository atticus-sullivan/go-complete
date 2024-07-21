package positional

import (
	"complete/complete/completeArgs"
)

type Positional struct {
	Idx      uint
	Help     string
	Arg      completeargs.CompleteType
	Optional bool
}
