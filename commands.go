package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/w0/goloops/internal/audiocontent"
)

type Choices struct {
	Plist     string `arg:"positional" help:"Path to the plist that contains audio content information."`
	Mandatory bool   `arg:"-m" help:"Only mandatory pkgs."`
	Optional  bool   `arg:"-o" help:"Only optional pkgs."`
	All       bool   `arg:"-a" help:"Both mandatory and optional pkgs."`
}

type Commands struct {
	Get  *Choices `arg:"subcommand:get" help:"Download audio content from apple."`
	List *Choices `arg:"subcommand:list" help:"Print pkg names to stdout."`
}

func ParseCommands() Commands {
	var args Commands

	arg.MustParse(&args)

	return args
}

func ResolvePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)

	if err != nil {
		return "", nil
	}

	if _, err := os.Stat(absPath); errors.Is(err, os.ErrNotExist) {
		return "", errors.New("Path to file doesn't exist. Try providing the absolute path.")
	}

	return absPath, nil
}

func HandleCommands(c Commands) {

	switch {
	case c.Get != nil:
	case c.List != nil:
		fp, _ := ResolvePath(c.List.Plist)
		ac, _ := audiocontent.NewAudioContent(fp)

		if c.List.Mandatory {
			ac.ListMandatory()
		}

		if c.List.Optional {
			ac.ListOptional()
		}

		if c.List.All {
			ac.ListAll()
		}

	}
}
