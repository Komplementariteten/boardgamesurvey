package model

type Vote struct {
	Title    string
	Owner    string
	ObjectId int
}

type VoteCount struct {
	Title    string
	Votes    int
	ObjectId int
	Voters   []string `json:"-"`
}

type Poll struct {
	Votes []VoteCount
}
