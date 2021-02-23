package mailbox

import (
	"errors"
)

type Mailbox struct {
	UserId   int
	UserName string
	Message  string
}

type MailboxService struct {
	boxes map[int]*Mailbox
}

func NewMailboxService() *MailboxService {
	return &MailboxService{boxes: make(map[int]*Mailbox)}
}

func (m *MailboxService) AddMailbox(req *AddMailboxRequest, resp *AddMailboxResponse) error {
	if _, ok := m.boxes[req.UserId]; ok {
		return errors.New("user already exists")
	}

	m.boxes[req.UserId] = &Mailbox{UserId: req.UserId, UserName: req.UserName}
	return nil
}

func (m *MailboxService) PutMessage(req *PutMessageRequest, resp *PutMessageResponse) error {
	b, ok := m.boxes[req.UserId]
	if !ok {
		return errors.New("user doesn't exist")
	}
	b.Message = req.Message
	return nil
}

func (m *MailboxService) GetMessage(req *GetMessageRequest, resp *GetMessageResponse) error {
	b, ok := m.boxes[req.UserId]
	if !ok {
		return errors.New("user doesn't exist")
	}
	resp.Message = b.Message
	resp.UserName = b.UserName
	return nil
}
