package markdown

import (
	"log"
	"strings"

	"github.com/capsci/bulletin/config"
)

var matcher = make(map[string]string)

func init() {
	// TODO: Don't load config twice.
	cfg, err := config.GenerateFromYML("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	for _, replace := range cfg.Replace.Text {
		addReplacer(replace.From, replace.To)
	}
	for _, replace := range cfg.Replace.Link {
		addReplacer(replace.From, replace.To)
	}
}

func addReplacer(find string, replace string) {
	matcher[find] = replace
}

func matchAndReplace(commitMessage string) string {
	for find, replace := range matcher {
		if strings.Contains(commitMessage, find) {
			commitMessage = strings.ReplaceAll(commitMessage, find, replace)
		}
	}
	return commitMessage
}
