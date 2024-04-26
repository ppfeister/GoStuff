package main

import (
	"flag"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ppfeister/GoStuff/messages"
	"golang.org/x/text/language"
)

func main() {
	messages.LoadMessages(language.English)

	var num_workers uint
	var manifest_uri_simple string
	var verbosity uint
	flag.UintVar(&num_workers, "t", 20, messages.Fetch(i18n.LocalizeConfig{MessageID: "usage_concurrency"}))
	flag.StringVar(&manifest_uri_simple, "m", "", messages.Fetch(i18n.LocalizeConfig{MessageID: "usage_pairedManifest"}))
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
		if manifest_uri_simple != "" && len(single_target) != 0 {
			log.Fatal(messages.Fetch(i18n.LocalizeConfig{
				MessageID: "Err_MixedCredTypes",
			}))
		}
	}
}
