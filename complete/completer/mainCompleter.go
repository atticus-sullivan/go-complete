package completer

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/atticus-sullivan/go-complete/complete/types"
	"github.com/riywo/loginshell"
)

type MainCompleter struct {
	types.Completer
	Opts CompleterOpts
}

type CompleterOpts struct {
	OptionStacking     bool
	OptionStackingArgs bool
	DashSep            bool
	ToFile             bool
}

func (mc *MainCompleter) GenerateAuto(builder *strings.Builder, indent string) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return fmt.Errorf("couldn't determine user's shell: %w", err)
	}
	shell = filepath.Base(shell)
	switch shell {
	case "zsh": mc.GenerateZsh(builder, indent)
	case "bash": mc.GenerateBash(builder, indent)
	default: return fmt.Errorf("Shell %s is unsupported\n", shell)
	}
	return nil
}
