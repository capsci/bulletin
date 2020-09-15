package git

import "time"

// LogOptions helps user to specify structure of log Message
type LogOptions struct {
	CommitID bool
	From     string
	To       string
}

// Commit structure (Check https://git-scm.com/docs/pretty-formats)
type Commit struct {
	CommitHash hash
	Author     user
	Subject    string
	Body       string
}

type user struct {
	Name  string
	Email string
	Time  time.Time
}

type hash struct {
	Abbreviated string
	Full        string
}
