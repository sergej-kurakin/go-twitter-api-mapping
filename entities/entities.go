package entities

type Entities struct {
	HashTags     []HashTag     `json:"hashtags"`
	Symbols      []Symbol      `json:"symbols"`
	UserMentions []UserMention `json:"user_mentions"`
	Urls         []URL         `json:"urls"`
}
