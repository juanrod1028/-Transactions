package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/juanrod1028/Transactions/src/usecase/models"
)

func HandelTransactions(r *http.Request) ([]models.Transactions, error) {
	file, err := handelCsv(r)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	transactions, err := ReadCSVRecords(file)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
func handelCsv(r *http.Request) (multipart.File, error) {
	err := ParseMultipartForm(r)
	if err != nil {
		return nil, err
	}

	file, err := GetFileFromRequest(r)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func ReadCSVRecords(file io.Reader) ([]models.Transactions, error) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV file: %v", err)
	}

	if len(records) > 0 {
		records = records[1:]
	} else {
		return nil, fmt.Errorf("CSV file has no data")
	}

	var transactions []models.Transactions
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error converting %v to int: %v\n", record[0], err)
			continue
		}
		transaction := models.NewTransaction(id, record[1], record[2])
		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
