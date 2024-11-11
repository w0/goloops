package main

import "fmt"

func main() {

	args, _ := ParseCommands()

	switch {
	case args.Get != nil:
		fmt.Printf("%v", args.Get)
	}

}
