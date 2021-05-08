package v1

type AddUserReq struct {
	Name string
	Age  int
}

type AddUserReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
