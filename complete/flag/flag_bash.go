package flag

import (
	"complete/internal"
	"fmt"
	"strings"
)

func (f *Flag) GenerateBash(builder *strings.Builder, indent string, progName string) []*internal.AddFuncBash {
	if f.Short != 0 {
		fmt.Fprintf(builder, `%[1]s-%[2]c)
%[1]s    ;;
`, indent, f.Short)

		for _, arg := range f.Args {
			arg.GenerateBash(builder, indent, progName)
		}
	}

	if f.Long != "" {
		fmt.Fprintf(builder, `%[1]s--%[2]s)
%[1]s    ;;
`, indent, f.Long)

		for _, arg := range f.Args {
			arg.GenerateBash(builder, indent, progName)
		}
	}
	return nil
}
