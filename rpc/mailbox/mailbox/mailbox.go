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

func (m *MailboxService) AddMailbox(id int, name string) error {
	if _, ok := m.boxes[id]; ok {
		return errors.New("user already exists")
	}

	m.boxes[id] = &Mailbox{UserId: id, UserName: name}
	return nil
}

func (m *MailboxService) PutMessage(id int, message string) error {
	b, ok := m.boxes[id]
	if !ok {
		return errors.New("user doesn't exist")
	}
	b.Message = message
	return nil
}

func (m *MailboxService) GetMessage(id int) (msg string, err error) {
	b, ok := m.boxes[id]
	if !ok {
		err = errors.New("user doesn't exist")
		return
	}
	msg = b.Message
	return
}
