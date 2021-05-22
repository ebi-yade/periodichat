package zoom

type BotNotification struct {
	Event   string `json:"event"`
	Payload struct {
		UserId    string `json:"userId"`
		Cmd       string `json:"cmd"`
		Timestamp int    `json:"timestamp"`
	} `json:"payload"`
}
