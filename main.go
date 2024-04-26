package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ppfeister/gostuff/messages"
	utils "github.com/ppfeister/gostuff/utils/parsers"
	"golang.org/x/text/language"
)

func main() {
	messages.LoadMessages(language.English)

	var num_workers uint
	var verbosity uint
	flag.UintVar(&num_workers, "t", 20, messages.Fetch(i18n.LocalizeConfig{MessageID: "usage_concurrency"}))
	flag.UintVar(&verbosity, "verbosity", 1, "")
	flag.Parse()
	var single_target []string = flag.Args()

	messages.LoadMessages(language.English)

	if len(single_target) > 0 {
		if len(single_target) != 2 {
			log.Fatal(messages.Fetch(i18n.LocalizeConfig{
				MessageID: "Err_SinglePairCredQty",
			}))
		}
	}

	var testing utils.Targets = utils.FetchManifest()

	fmt.Printf("%v\n", testing)
}
