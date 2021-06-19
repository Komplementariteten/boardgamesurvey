package model

type Vote struct {
	Title       string
	Owner       string
	ObjectId    int
	Description string
}

type VoteCount struct {
	Title       string
	Description string
	Votes       int
	ObjectId    int
	Voters      []string `json:"-"`
}

type Poll struct {
	Votes []*VoteCount
}
