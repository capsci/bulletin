package git

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/capsci/bulletin/utils"
)

// GetLogs gets logs
func GetLogs(from string, to string) *CommitIterator {
	output, err := runGitCli([]string{"log", "--pretty=format:\"%h\",\"%H\",\"%aN\",\"%aE\",\"%at\",\"%s\",\"%b\"", from + "..." + to})
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(strings.NewReader(output))
	// TODO: Smarter way to handle "wrong number of fields error"
	reader.FieldsPerRecord = -1

	ci := GetCommitIterator()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		commit := new(Commit)
		commit.Set(record)
		ci.Push(*commit)
	}
	return ci
}

// Thanks https://stackoverflow.com/a/40770011
func runGitCli(args []string) (string, error) {
	outStr, errStr, _ := utils.RunCommand("git", args)
	if len(errStr) != 0 {
		return outStr, errors.New(errStr)
	}
	return outStr, nil
}
