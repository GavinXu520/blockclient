package entity

type Result struct {
	Status int      `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}
