package entities

type Entities struct {
	HashTags     []HashTags    `json:"hashtags"`
	Symbols      []Symbols     `json:"symbols"`
	UserMentions []UserMention `json:"user_mentions"`
	Urls         []URL         `json:"urls"`
}
