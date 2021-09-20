package responses

type Response struct {
	Error   string `json:"-"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    Meta   `json:"meta"`
}

type Meta struct {
	UnixTime int64 `json:"time"`
}
