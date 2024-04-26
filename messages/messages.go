package messages

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer_default *i18n.Localizer
var localizer_pref *i18n.Localizer

func LoadMessages(languagepref language.Tag) {
	var bundle *i18n.Bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("messages/en.toml")
	//bundle.MustLoadMessageFile("messages/de.toml")
	//bundle.MustLoadMessageFile("messages/fr.toml")
	// etc
	localizer_default = i18n.NewLocalizer(bundle, language.English.String())
	if languagepref == language.English {
		localizer_pref = nil
	} else {
		localizer_pref = i18n.NewLocalizer(bundle, languagepref.String())
	}
}

// Attempt to fetch localized message first, falling back to English if not localized
func Fetch(conf i18n.LocalizeConfig) string {
	if localizer_pref != nil {
		result, err := localizer_pref.Localize(&conf)
		if err == nil {
			return result
		}
	}
	return localizer_default.MustLocalize(&conf)
}
