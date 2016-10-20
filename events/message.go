package events

// Message is a simple struct to (un-)marshal JSON messages
type Message struct {
	Text string `json:"text"`
}
