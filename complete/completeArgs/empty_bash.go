package completeargs

import (
	"complete/internal"
	"strings"
)

func (ce CTempty) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	return nil,nil
}
