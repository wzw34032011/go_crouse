package v1

type AddUserReq struct {
	name string
}

type AddUserReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
