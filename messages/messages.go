package messages

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var langpref string = "English"
var localizer *i18n.Localizer

// SetLocale not yet boundary-safe
// TODO: Add validation against available values
func SetLanguage(language string) {
	langpref = language
}

func LoadMessages() {
	var bundle *i18n.Bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("en.toml")
	localizer = i18n.NewLocalizer(bundle, language.English.String())
}

func Fetch(conf i18n.LocalizeConfig) string {
	var result string = localizer.MustLocalize(&conf)
	return result
}
