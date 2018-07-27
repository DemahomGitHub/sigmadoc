package model

// Dictionary represents a table in the databse
type Dictionary struct {
	ID          int    `json:"id"`
	Term        string `json:"term"`
	Frequency   int    `json:"frequency"`
	PostingList []int  `json:"posting_list"`
}

// PostingList representes the list of documents IDs
// that contains a given term identified by DicID
type PostingList struct {
	DocID int `json:"doc_id"`
	DicID int `json:"dic_id"`
}
