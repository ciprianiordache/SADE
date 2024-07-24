package notifier

type Notifier struct {
	Host   string
	APIKey string
}

type emailRequest struct {
	From    map[string]string   `json:"from"`
	To      []map[string]string `json:"to"`
	Subject string              `json:"subject"`
	HTML    string              `json:"html"`
}

type MailData struct {
	Data string
}
