package model

type Vote struct {
	Title         string
	Owner         string
	ObjectId      int
	Description   string
	Admin         string
	Disable       bool
	DisableReason string
}

type VoteCount struct {
	Title         string
	Description   string
	Votes         int
	ObjectId      int
	Disabled      bool
	DisableReason string
	Voters        []string `json:"-"`
}

type Poll struct {
	Votes []*VoteCount
}
