package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sidaurukdedi/go-boiler/domain"
	"github.com/sidaurukdedi/go-boiler/pkg/response"
	"github.com/sidaurukdedi/go-boiler/pkg/validator"
	"github.com/sirupsen/logrus"
)

const (
	basePath = "/tester-mongodb"
)

type HTTPHandler struct {
	Logger      *logrus.Logger
	UserUsecase UserUsecase
	Validator   *validator.Validator
}

// NewUserHTTPHandler is a constructor
func NewUserHTTPHandler(logger *logrus.Logger, validator *validator.Validator, router *mux.Router, userUsecase UserUsecase) {
	handler := &HTTPHandler{
		Logger:      logger,
		Validator:   validator,
		UserUsecase: userUsecase,
	}
	router.HandleFunc(basePath+"/v1/user", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc(basePath+"/v1/user/{id}", handler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc(basePath+"/v1/user/{id}", handler.GetUser).Methods(http.MethodGet)
}

func (handler *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var payload domain.User

	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp = response.NewErrorResponse(err, http.StatusUnprocessableEntity, nil, response.StatusInvalidPayload, err.Error())
		response.JSON(w, resp)
		return
	}

	if err := handler.Validator.Validate(payload); err != nil {
		resp = response.NewErrorResponse(err, http.StatusBadRequest, nil, response.StatusInvalidPayload, err.Error())
		response.JSON(w, resp)
		return
	}

	resp = handler.UserUsecase.CreateUser(ctx, payload)
	response.JSON(w, resp)
}

func (handler *HTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var payload domain.User

	ctx := r.Context()

	pathVariable := mux.Vars(r)
	userId := pathVariable["id"]

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp = response.NewErrorResponse(err, http.StatusUnprocessableEntity, nil, response.StatusInvalidPayload, err.Error())
		response.JSON(w, resp)
		return
	}

	if err := handler.Validator.Validate(payload); err != nil {
		resp = response.NewErrorResponse(err, http.StatusBadRequest, nil, response.StatusInvalidPayload, err.Error())
		response.JSON(w, resp)
		return
	}

	resp = handler.UserUsecase.UpdateUser(ctx, userId, payload)
	response.JSON(w, resp)
}

type contextKey string

const requestIDKey = contextKey("requestID")

func (handler *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	xRequestId := r.Header.Get("X-Request-ID")
	ctx := context.WithValue(r.Context(), requestIDKey, xRequestId)

	pathVariable := mux.Vars(r)
	userId := pathVariable["id"]

	resp = handler.UserUsecase.GetUser(ctx, userId)
	response.JSON(w, resp)
}
