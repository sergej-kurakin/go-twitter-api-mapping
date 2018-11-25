package entities

type UserMentions struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         string `json:"id_str"`
}
