package common

type ResponseData struct {
	ID      int    `json:id`
	Message string `json:message`
}

type Filter struct {
	Key      string
	Value    interface{}
	Operator string
}
