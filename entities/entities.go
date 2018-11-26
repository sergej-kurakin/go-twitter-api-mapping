package entities

type Entities struct {
	HashTags     []HashTag     `json:"hashtags"`
	Symbols      []Symbols     `json:"symbols"`
	UserMentions []UserMention `json:"user_mentions"`
	Urls         []URL         `json:"urls"`
}
