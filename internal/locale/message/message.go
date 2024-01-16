package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
    currentDir := filepath.Dir(currentFile)
    filePath := filepath.Join(currentDir, "messages.en.json")
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(filePath)
}

func GetMessage(key string, args ...string) (string, error) {
	arguments := make(map[string]string)
	
	for i, arg := range args {
		arguments[fmt.Sprintf("Arg%d", i)] = arg
	}

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: key,
		TemplateData: arguments,
	}
	localizer = i18n.NewLocalizer(bundle, "en")  
	localizedMessage, err := localizer.Localize(&localizeConfigWelcome)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return localizedMessage, nil
}