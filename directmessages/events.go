package directmessages

type Events struct {
	Events     []Message `json:"events"`
	NextCursor string    `json:"next_cursor"`
}
