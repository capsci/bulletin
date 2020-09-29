package markdown

// GetH1 gets H1 string
func GetH1(title string) string {
	return getMarkdown("# ", title, "\n")
}

// GetH2 gets H2 string
func GetH2(title string) string {
	return getMarkdown("## ", title, "\n")
}

// GetH3 gets H3 string
func GetH3(title string) string {
	return getMarkdown("### ", title, "\n")
}

// GetH4 gets H4 string
func GetH4(title string) string {
	return getMarkdown("#### ", title, "\n")
}

// GetH5 gets H5 string
func GetH5(title string) string {
	return getMarkdown("##### ", title, "\n")
}

// GetH6 gets H6 string
func GetH6(title string) string {
	return getMarkdown("###### ", title, "\n")
}

// GetBold gets bold string
func GetBold(text string) string {
	return getMarkdown("**", text, "**")
}

// GetItalic gets italic string
func GetItalic(text string) string {
	return getMarkdown("__", text, "__")
}

// GetStrikethrough gets italic string
func GetStrikethrough(text string) string {
	return getMarkdown("--", text, "--")
}

// GetLink gets link string
func GetLink(displayText, address string) string {
	return "[" + displayText + "](" + address + ")"
}

// GetListItem gets list item
func GetListItem(text string) string {
	return getMarkdown("- ", text, "\n")
}

func getMarkdown(prepend, text, append string) string {
	return prepend + matchAndReplace(text) + append
}
