package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juanrod1028/Transactions/src/usecase/models"
)

func CalculateSummary(transactions []models.Transactions) (float64, map[string]int, float64, float64) {
	totalBalance := 0.0
	transactionsPerMonth := make(map[string]int)
	totalCredit := 0.0
	totalDebit := 0.0
	creditCount := 0
	debitCount := 0

	for _, tx := range transactions {
		dateParts := strings.Split(tx.Date, "/")
		month := dateParts[0]
		year := "2024"
		monthYear := fmt.Sprintf("%s/%s", month, year)

		amount, _ := strconv.ParseFloat(tx.Transactions, 64)
		totalBalance += amount

		transactionsPerMonth[monthYear]++

		if amount >= 0 {
			totalCredit += amount
			creditCount++
		} else {
			totalDebit += amount
			debitCount++
		}
	}

	averageCredit := 0.0
	if creditCount > 0 {
		averageCredit = totalCredit / float64(creditCount)
	}

	averageDebit := 0.0
	if debitCount > 0 {
		averageDebit = totalDebit / float64(debitCount)
	}

	return totalBalance, transactionsPerMonth, averageDebit, averageCredit
}
