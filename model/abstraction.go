package model

const (
	// FAILED :
	FAILED int = 400
	// SUCCESS :
	SUCCESS int = 200
)

// ResponseMessage :
type ResponseMessage struct {
	Status       int      `json:"status"`
	Message      string   `json:"message"`
	UserResp     User     `json:"user_response"`
	DoucmentResp Document `json:"document_response"`
}

// AbstractCleaner :
type AbstractCleaner interface {
	Clean()
}

// RemoveSpaces :
func RemoveSpaces(o AbstractCleaner) {
	o.Clean()
}
