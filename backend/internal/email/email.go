package email

type Email struct {
	SenderName string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	RecipientEmail string `json:"recipient_email"`
	Subject string `json:"subject"`
	Attachment string `json:"attachment"`
	Content string `json:"content"`

}