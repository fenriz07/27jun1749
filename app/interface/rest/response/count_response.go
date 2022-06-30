package response

type CountResponse struct {
	Count int64  `json:"count"`
	Code  string `json:"code"`
}
