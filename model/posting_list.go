package model

// PostingList representes the list of documents IDs
// that contains a given term identified by DicID
type PostingList struct {
	DocID int `json:"doc_id"`
	DicID int `json:"dic_id"`
}
