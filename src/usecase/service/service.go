package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juanrod1028/Transactions/src/usecase/helpers"
	"github.com/juanrod1028/Transactions/src/usecase/ports"
	"github.com/juanrod1028/Transactions/src/usecase/service/utils"
)

type APIServer struct {
	listenerAddr string
	storage      ports.Storage
}

func NewApiServer(listenerAddr string, storage ports.Storage) *APIServer {
	return &APIServer{
		listenerAddr: listenerAddr,
		storage:      storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	log.Println("JSON API server running on port:", s.listenerAddr)
	router.HandleFunc("/transactions", utils.MakeHttpHandleFunc(s.handleTransaction)).Methods("POST")
	router.HandleFunc("/user/transactions/{id}", utils.MakeHttpHandleFunc(s.handleGetAccountById)).Methods("GET")

	http.ListenAndServe(s.listenerAddr, router)
}

func (s *APIServer) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handlePostTransaction(w, r)
	default:
		return utils.NewHTTPError(http.StatusMethodNotAllowed, fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func (s *APIServer) handlePostTransaction(w http.ResponseWriter, r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		return s.emailService(w, r)
	}
	return utils.NewHTTPError(http.StatusBadRequest, fmt.Errorf(http.StatusText(http.StatusBadRequest)))
}

func (s *APIServer) emailService(w http.ResponseWriter, r *http.Request) error {
	transactions, err := helpers.HandelTransactions(r)
	if err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, err)
	}
	user, err := helpers.GetUserDataFromRequest(r, transactions)
	if err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := s.storage.CreateUser(&user); err != nil {
		return utils.NewHTTPError(http.StatusInternalServerError, err)
	}
	if err := s.storage.CreateTransactions(&user); err != nil {
		return utils.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = helpers.ProcessAndSendSummary(user)
	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusOK, transactions)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid id given: %s", idStr))
	}
	account, err := s.storage.GetTransactionById(id)
	if err != nil {
		return utils.NewHTTPError(http.StatusNotFound, err)
	}
	return utils.WriteJson(w, http.StatusOK, account)
}
