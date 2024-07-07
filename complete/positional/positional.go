package positional

import (
	"complete/internal"
)

type Positional struct {
	Idx      uint
	Help     string
	Arg      internal.CompleteType
	Optional bool
}
