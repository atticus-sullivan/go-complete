package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
)

func (ce CTempty) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(": : ")
	return nil
}
