package service

import (
	"encoding/json"
	"fmt"
	"heptaber/notification/app/initializers"
	"heptaber/notification/domain/model"
	"log"
	"net/smtp"
	"os"
	"sync"
)

func init() {
	initializers.LoadEnvVariables()
}

const (
	HTMLMIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func SendEmail(eventType string, data []byte) {
	wg := &sync.WaitGroup{}
	switch eventType {
	case string(model.VERIFICATION):
		wg.Add(1)
		done := make(chan struct{}, 2)
		var vEmail model.VerificationEmail

		err := json.Unmarshal(data, &vEmail)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %s", err.Error())
		}
		go sendVerificationEmail(vEmail.Recipient, vEmail.VerificationCode, done, wg)
		wg.Wait() // block until wg calls done
		close(done)
	case string(model.SUBSCRIPTION):
		// on blog-service implementation
		wg.Add(1)
		done := make(chan struct{}, 2)
		go sendSubscriptionEmail(done, wg)
		wg.Wait()
		close(done)
	}
}

func sendVerificationEmail(recipient string, verificationCode string, respch chan struct{}, wg *sync.WaitGroup) {
	senderEmail := os.Getenv("EMAIL_SENDER")
	senderPassword := os.Getenv("EMAIL_PSW")
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_HOST_PORT")
	blogLink := os.Getenv("BLOG_LINK")

	body := "<html><body><h1>Welcome to Heptaber-world!</h1>" +
		"<p>You can verify your account following " +
		// TODO: change links on publishing
		"<a href=\"https://" + blogLink + ":10010/api/v1/auth/verify/" + verificationCode + "\">this link</a>" +
		"</body></html>"
	msg := generateHTMLEmailMessage("Heptaber Email verification", body)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+fmt.Sprint(smtpPort),
		auth,
		senderEmail,
		[]string{recipient},
		msg,
	)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Panic at the disco: ", r)
			}
		}()
		log.Panic("failed to send email:", err)
	}
	respch <- struct{}{}
	wg.Done()
}

func sendSubscriptionEmail(respch chan struct{}, wg *sync.WaitGroup) {
	// TODO: impl

	respch <- struct{}{}
	wg.Done()
}

func generateHTMLEmailMessage(subjectText string, body string) []byte {
	return []byte("Subject: " + subjectText + "\n" + HTMLMIME + body)
}
