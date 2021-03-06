package requests

type CreateShortRequest struct {
	URL string `json:"url" example:"http://example.com/asd"`
}

type CreateShortBatchRequest []CreateShortBatchItem

type CreateShortBatchItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type DeleteShortBatchRequest []string
