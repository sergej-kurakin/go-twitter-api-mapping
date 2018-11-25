package directmessages

type MessageCreate struct {
	Target      Target      `json:"target"`
	SenderID    string      `json:"sender_id"`
	MessageData MessageData `json:"message_data"`
}
