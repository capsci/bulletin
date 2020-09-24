package section

import (
	"testing"

	"github.com/capsci/bulletin/test"

	s "github.com/capsci/bulletin/section"
	"github.com/capsci/bulletin/test/utils"
)

// TestAddEmojiShortCodes tests if we are able to add emoji shortcodes to the
func TestAddEmojiShortCodes(t *testing.T) {
	const EmojiSC = ":emoji"
	sect := s.Section{}
	sect.AddEmojiShortcodes([]string{EmojiSC})
	_, found := utils.FindInSlice(sect.EmojiShortcodes, EmojiSC)
	if !found {
		t.Error("Emoji not added to section")
	}
	test.Expected(t, len(sect.EmojiShortcodes), 1, "Incorrect number of items in the section")
}

// TestAddEmojiShortCodes tests if we are able to add emoji shortcodes to the
func TestAddTextMarkers(t *testing.T) {
	const TextSC = ":text"
	sect := s.Section{}
	sect.AddTextMarkers([]string{TextSC})
	_, found := utils.FindInSlice(sect.TextShortcodes, TextSC)
	if !found {
		t.Error("Text Marker not added to section")
	}
	test.Expected(t, len(sect.TextShortcodes), 1, "Incorrect number of items in the section")
}

func TestBelongsEmoji(t *testing.T) {
	const EmojiSC = ":emoji"
	const TextSC = ":text"
	const EmojiCommit = EmojiSC + " adds a new emoji commit"
	const TextCommit = TextSC + " adds a new text marker commit"
	sect := s.Section{}
	sect.AddEmojiShortcodes([]string{EmojiSC})
	sect.AddTextMarkers([]string{TextSC})

	test.Expected(t, sect.Belongs(EmojiCommit, true), true, "Emoji commit incorrectly classified as not belongs")
	test.Expected(t, sect.Belongs(TextCommit, true), false, "Text commit incorrectly classified as belongs")

	test.Expected(t, sect.Belongs(EmojiCommit, false), false, "Emoji commit incorrectly classified as belongs")
	test.Expected(t, sect.Belongs(TextCommit, false), true, "Text commit incorrectly classified as not belongs")
}
