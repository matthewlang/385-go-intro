package mailbox

// RPC request/response types.

type AddMailboxRequest struct {
	UserId   int
	UserName string
}

type AddMailboxResponse struct {
}

type PutMessageRequest struct {
	UserId  int
	Message string
}

type PutMessageResponse struct {
}

type GetMessageRequest struct {
	UserId int
}

type GetMessageResponse struct {
	UserName string
	Message  string
}
