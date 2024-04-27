package probes

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ppfeister/gostuff/messages"
	utils "github.com/ppfeister/gostuff/utils/parsers"
)

func Req(target_dat utils.Target) utils.Target {
	req, err := http.NewRequest(target_dat.Method, target_dat.Url, strings.NewReader(target_dat.Data))
	if err != nil {
		log.Fatal("ahh")
	}

	fmt.Printf("%v\n", target_dat.Data)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/116.0")

	var client *http.Client = &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	target_dat.Response_Data.Status = uint8(resp.StatusCode)
	placeholder, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Info(messages.Fetch(i18n.LocalizeConfig{
			MessageID: "Err_UnableToParseResponseData",
			TemplateData: map[string]interface{}{
				"Error": err.Error(),
			},
		}))
		target_dat.Response_Data.Data = ""
	} else {
		target_dat.Response_Data.Data = string(placeholder)
	}

	fmt.Printf("%v\n", target_dat.Response_Data.Data)
	return target_dat
}
