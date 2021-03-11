package utl

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Error      error       `json:"error,omitempty"`
	Data       interface{} `json:"data"`
	Param      interface{} `json:"interface,omitempty"`
}
