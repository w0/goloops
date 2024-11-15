package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/w0/goloops/internal/audiocontent"
	"github.com/w0/goloops/internal/client"
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
		return "", err
	}

	if _, err := os.Stat(absPath); errors.Is(err, os.ErrNotExist) {
		return "", errors.New("Path to file doesn't exist. Try providing the absolute path.")
	}

	return absPath, nil
}

func HandleCommands(c Commands) {
	switch {
	case c.Get != nil:
		GetCommand(c.Get)

	case c.List != nil:
		ListCommand(c.List)
	}
}

func GetCommand(get *Choices) {
	fp, _ := ResolvePath(get.Plist)
	ac, _ := audiocontent.NewAudioContent(fp)

	urls := ac.GetMandatory()

	limit := make(chan int, 3)

	for _, url := range urls {
		limit <- 1
		go func() {
			fmt.Printf("Downloading: %s\n", url)
			client.DownloadFile(url)
			<-limit
		}()
	}
}

func ListCommand(list *Choices) {
	fp, _ := ResolvePath(list.Plist)
	ac, _ := audiocontent.NewAudioContent(fp)

	if list.Mandatory {
		ac.ListMandatory()
	}

	if list.Optional {
		ac.ListOptional()
	}

	if list.All {
		ac.ListAll()
	}
}
