package helpers

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/juanrod1028/Transactions/src/usecase/constants"
	"github.com/juanrod1028/Transactions/src/usecase/models"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("COMPANY_EMAIL")
	password := os.Getenv("COMPANY_EMAIL_PASS")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	fmt.Print(auth)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf(constants.FAILED_EMAIL, err)
	}
	return nil
}

func createSummaryBody(userName string, totalBalance float64, transactionsPerMonth map[string]int, averageDebit, averageCredit float64) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(constants.HELLO_USER, userName))
	sb.WriteString(fmt.Sprintf(constants.TOTAL_BALANCE, totalBalance))

	for month, count := range transactionsPerMonth {
		sb.WriteString(fmt.Sprintf(constants.NUMBER_OF_TRANSACTIONS, month, count))
	}

	sb.WriteString(fmt.Sprintf(constants.AVERAGE_DEBIT, averageDebit))
	sb.WriteString(fmt.Sprintf(constants.AVERAGE_CREDIT, averageCredit))

	return sb.String()
}

func ProcessAndSendSummary(user models.User) error {
	totalBalance, transactionsPerMonth, averageDebit, averageCredit := CalculateSummary(user.Transactions)

	body := createSummaryBody(user.Identification, totalBalance, transactionsPerMonth, averageDebit, averageCredit)

	subject := constants.SUBJECT
	err := SendEmail(user.Email, subject, body)
	if err != nil {
		return fmt.Errorf(constants.FAILED_EMAIL, err)
	}

	return nil
}
