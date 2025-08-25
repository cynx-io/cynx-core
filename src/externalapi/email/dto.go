package email

type SendEmailRequest struct {
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	IsHTML  bool     `json:"is_html"`
}
