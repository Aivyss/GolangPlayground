package response

type Response[T any] struct {
	Status  int  `json:"status"`
	Success bool `json:"success"`
	Data    T    `json:"data"`
}
