package directmessages

type AttachmentSize struct {
	W      uint   `json:"w"`
	H      uint   `json:"h"`
	Resize string `json:"resize"`
}

type AttachmentSizes struct {
	Medium AttachmentSize `json:"medium"`
	Thumb  AttachmentSize `json:"thumb"`
	Small  AttachmentSize `json:"small"`
	Large  AttachmentSize `json:"large"`
}

type VideoInfoVariant struct {
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
	BitRate     uint   `json:"bitrate"`
}

type VideoInfo struct {
	Variants []VideoInfoVariant `json:"variants"`
}

type AttachmentMedia struct {
	ID            uint64          `json:"id"`
	IDString      string          `json:"id_str"`
	MediaURL      string          `json:"media_url"`
	MediaURLHttps string          `json:"media_url_https"`
	URL           string          `json:"url"`
	DisplayURL    string          `json:"display_url"`
	ExpandedURL   string          `json:"expanded_url"`
	Type          string          `json:"type"`
	Sizes         AttachmentSizes `json:"sizes"`
	VideoInfo     VideoInfo       `json:"video_info"`
}

type Attachment struct {
	Type  string          `json:"type"`
	Media AttachmentMedia `json:"media"`
}
