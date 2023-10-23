package entity

type PushPayload struct {
	Interests []string `json:"interests"`
	Web       struct {
		Notification struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"notification"`
	} `json:"web"`
}
