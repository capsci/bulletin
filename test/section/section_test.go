package section

import (
	"testing"

	"github.com/capsci/bulletin/test"

	s "github.com/capsci/bulletin/section"
	"github.com/capsci/bulletin/test/utils"
)

const qual1 = ":qual1"
const qual2 = ":qual2"

// TestAddEmojiShortCodes tests if we are able to add emoji shortcodes to the
func TestAddTextQualifiers(t *testing.T) {
	sect := s.Section{}
	sect.AddTextQualifiers([]string{qual1})
	sect.AddTextQualifiers(qual2)
	_, found := utils.FindInSlice(sect.TextQualifiers, qual1)
	if !found {
		t.Error("Text qualifier list not added to section")
	}
	_, found = utils.FindInSlice(sect.TextQualifiers, qual2)
	if !found {
		t.Error("Text qualifier not added to section")
	}
	test.Expected(t, len(sect.TextQualifiers), 2, "Incorrect number of items in the section")
}

func TestBelongsEmoji(t *testing.T) {
	const qual1Commit = qual1 + " adds a new commit"
	const qual2Commit = qual2 + " adds another commit"
	sect := s.Section{}
	sect.AddTextQualifiers(qual1)

	test.Expected(t, sect.Belongs(qual1), true, "Commit incorrectly classified as not belongs")
	test.Expected(t, sect.Belongs(qual2), false, "Commit incorrectly classified as belongs")
}
