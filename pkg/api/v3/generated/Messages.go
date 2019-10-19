package generated

type Message struct {
	Details []struct {
		Message      string      `json:"message"`
		MetaData     interface{} `json:"metaData"`
		Resolution   interface{} `json:"resolution"`
		ShortMessage string      `json:"shortMessage"`
		Type         string      `json:"type"`
	} `json:"details"`
	Messages []string `json:"messages"`
}
