package directmessages

type Message struct {
	ID               string        `json:"id"`
	Type             string        `json:"type"`
	CreatedTimestamp string        `json:"created_timestamp"`
	MessageCreate    MessageCreate `json:"message_create"`
}
