package completeargs

import (
	"complete/internal"
	"strings"
)

func (ce CTempty) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(": : ")
	return nil
}
