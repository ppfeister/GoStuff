package main

import (
	"flag"

	"github.com/ppfeister/GoStuff/debugkit"
)

func main() {
	var num_workers uint
	var manifest_uri_simple string
	var manifest_uri_batteringram string
	var debug bool
	var verbosity uint
	flag.UintVar(&num_workers, "t", 20, "set number of concurrent threads")
	flag.StringVar(&manifest_uri_simple, "s", "", "use a simple manifest")
	flag.StringVar(&manifest_uri_batteringram, "b", "", "use a grouped manifest (not recommended)")
	flag.BoolVar(&debug, "d", false, "dump debugging info to stdout")
	flag.UintVar(&verbosity, "verbosity", 1, "")
	flag.Parse()
	//var single_target []string = flag.Args()

	debugkit.SetVerbosity(uint8(verbosity))

	/*if len(single_target) > 0 {
		fmt
	}*/

	debugkit.Out(debugkit.Severity["warn"], "Something fucked up!")
}
