package response

type Response struct {
	Code uint `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
