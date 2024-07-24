package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func New(host, apiKey string) *Notifier {
	return &Notifier{
		Host:   host,
		APIKey: apiKey,
	}
}

func (n *Notifier) sendEmail(recipient []string, subject, content string) error {
	to := []map[string]string{}
	for _, r := range recipient {
		to = append(to, map[string]string{"email": r})
	}

	requestBody, err := json.Marshal(emailRequest{
		From:    map[string]string{"email": "no-reply@sade.com"},
		To:      to,
		Subject: subject,
		HTML:    content,
	})
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}
	req, err := http.NewRequest("POST", n.Host, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+n.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending email: %v, %v ", err, n.Host)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to send email, status code: %d, response: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (n *Notifier) SendLink(email, link, subject, templatePath string) error {
	data := MailData{
		Data: link,
	}

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return n.sendEmail([]string{email}, subject, buf.String())
}
