package response_success

type QuoteResponse struct {
	Text string `json:"text"`
}

func FromTextResponse(t string) QuoteResponse {
	return QuoteResponse{Text: t}
}
