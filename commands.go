package main

import (
	"github.com/alexflint/go-arg"
)

type Choices struct {
	Plist     string `arg:"positional"`
	Mandatory bool   `arg:"-m"`
	Optional  bool   `arg:"-o"`
	All       bool   `arg:"-a"`
}

type Commands struct {
	Get  *Choices `arg:"subcommand:get"`
	List *Choices `arg:"subcommand:list"`
}

func ParseCommands() (Commands, error) {
	var args Commands

	err := arg.Parse(&args)

	if err != nil {
		return Commands{}, err
	}

	return args, nil
}
