package response

// Burada benim response yapısında generic kullanmış olmam swagger'da sorun yarattı.
// Bu yüzden bir workaround'la sorunu bu şekilde çözdüm.
// task içeriğinde olmayan ek bir yapı için ekstra vakit harcamak istemedim.
type ExampleMessage struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	ToPhone   string `json:"to_phone"`
	IsSent    bool   `json:"is_sent"`
	MessageID string `json:"message_id"`
	CreatedAt string `json:"created_at"`
	SentAt    string `json:"sent_at"`
}

type MessageListResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []ExampleMessage `json:"data"`
}

type EmptyResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
