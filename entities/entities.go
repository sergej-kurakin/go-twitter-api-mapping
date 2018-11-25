package entities

type Entities struct {
	HashTags     []HashTags     `json:"hashtags"`
	Symbols      []Symbols      `json:"symbols"`
	UserMentions []UserMentions `json:"user_mentions"`
	Urls         []URL          `json:"urls"`
}
