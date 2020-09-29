package markdown

import (
	"strings"
)

var matcher = make(map[string]string)

func init() {
	addReplacer("README.md", "https://github.com/capsci/bulletin/blob/master/README.md")
	addReplacer("for", "rof")
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
