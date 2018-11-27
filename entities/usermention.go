package entities

type UserMention struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	IDStr      string `json:"id_str"`
	ID         uint64 `json:"id"`
}
