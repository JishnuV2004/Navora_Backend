package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//-> OTP senting third party wesite

type BrevoEmail struct {
	Sender struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"sender"`
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	Subject     string `json:"subject"`
	HTMLContent string `json:"htmlContent"`
}

func SentToEmail(TOemail, Subject, htmlContent string) error {
	url := "https://api.brevo.com/v3/smtp/email"

	payload := BrevoEmail{}
	payload.Sender.Email = os.Getenv("BREVO_FROM_EMAIL")
	payload.Sender.Name = os.Getenv("BREVO_FROM_NAME")
	payload.Subject = Subject
	payload.HTMLContent = htmlContent

//"Your CRYTINOX OTP"
//"<h2>Your OTP</h2><b>" + otp + "</b><p>Valid for 5 min</p>"

	payload.To = append(payload.To, struct {
		Email string `json:"email"`
	}{Email: TOemail})

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", os.Getenv("BREVO_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
    return fmt.Errorf(
        "brevo error: status=%d, body=%s",
        resp.StatusCode,
        string(respBody),
	)
	}

	return nil

}