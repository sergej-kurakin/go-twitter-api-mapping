package directmessages

import "github.com/sergej-kurakin/go-twitter-api-mapping/entities"

type MessageData struct {
	Text       string            `json:"text"`
	Entities   entities.Entities `json:"entities"`
	Attachment Attachment        `json:"attachment"`
}
