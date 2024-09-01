package helpers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/juanrod1028/Transactions/src/usecase/models"
)

func GetUserDataFromRequest(r *http.Request, transactions []models.Transactions) (models.User, error) {
	id := r.FormValue("identification")
	email := r.FormValue("email")

	if id == "" || email == "" {
		return models.User{}, fmt.Errorf("missing identification or email")
	}

	return *models.NewUser(id, email, transactions), nil
}
func ParseMultipartForm(r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return fmt.Errorf(http.StatusText(http.StatusBadRequest))
	}
	return nil
}
func GetFileFromRequest(r *http.Request) (multipart.File, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf(http.StatusText(http.StatusBadRequest))
	}
	return file, nil
}
