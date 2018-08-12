package model

// Dictionary represents a table in the databse
type Dictionary struct {
	ID          int    `json:"id"`
	Term        string `json:"term"`
	Frequency   int    `json:"frequency"`
	PostingList []int  `json:"posting_list"`
}
