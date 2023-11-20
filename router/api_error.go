package router

type APIError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e APIError) Error() string {
	return e.Description
}
