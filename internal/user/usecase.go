package user

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sidaurukdedi/go-boiler/domain"
	"github.com/sidaurukdedi/go-boiler/internal/entity"
	"github.com/sidaurukdedi/go-boiler/pkg/exception"
	"github.com/sidaurukdedi/go-boiler/pkg/response"
	"github.com/sirupsen/logrus"
)

// Collection of messages
const (
	createUserUnexpectedErrMessage string = "an error occured while processing"
	createUserSuccessMessage       string = "success save"
	updateUserUnexpectedErrMessage string = "an error occured while processing"
	updateUserSuccessMessage       string = "success update"
	getUserUnexpectedErrMessage    string = "an error occured while processing"
	getUserSuccessMessage          string = "success get data"
	getUserNotFoundErrMessage      string = "user not found"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, param domain.User) response.Response
	UpdateUser(ctx context.Context, userId string, param domain.User) response.Response
	GetUser(ctx context.Context, userId string) response.Response
}

type userUsecase struct {
	serviceName    string
	logger         *logrus.Logger
	location       *time.Location
	userRepository UserRepository
}

func NewUserUsecase(property UserUsecaseProperty) UserUsecase {
	return &userUsecase{
		serviceName:    property.ServiceName,
		logger:         property.Logger,
		location:       property.Location,
		userRepository: property.UserRepository,
	}
}

// CreateUser save user
func (u *userUsecase) CreateUser(ctx context.Context, payload domain.User) response.Response {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user := entity.User{
		ID:        uuid.New().String(),
		Name:      payload.Name,
		Address:   payload.Address,
		Timestamp: time.Now().In(u.location),
	}

	err = u.userRepository.Save(ctx, user)
	if err != nil {
		return response.NewErrorResponse(exception.ErrInternalServer, http.StatusInternalServerError, response.StatUnexpectedError, response.StatUnexpectedError, "")
	}

	return response.NewSuccessResponse(user, response.StatOK, createUserSuccessMessage)
}

func (u *userUsecase) UpdateUser(ctx context.Context, userId string, payload domain.User) response.Response {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user := entity.User{
		ID:        userId,
		Name:      payload.Name,
		Address:   payload.Address,
		Timestamp: time.Now().In(u.location),
	}

	err = u.userRepository.Update(ctx, userId, user)
	if err != nil {
		return response.NewErrorResponse(exception.ErrInternalServer, http.StatusInternalServerError, response.StatUnexpectedError, response.StatUnexpectedError, "")
	}

	return response.NewSuccessResponse(user, response.StatOK, updateUserSuccessMessage)
}

func (u *userUsecase) GetUser(ctx context.Context, userId string) response.Response {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	requestId, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		requestId = "unknown"
	}
	defer cancel()

	user, err := u.userRepository.FindByID(ctx, userId)
	if err != nil {
		u.logger.WithContext(ctx).WithFields(logrus.Fields{"requestId": requestId, "payload": userId}).Error(err)
		if err == exception.ErrNotFound {
			return response.NewSuccessResponse(nil, response.StatOK, getUserNotFoundErrMessage)
		}
		return response.NewErrorResponse(exception.ErrInternalServer, http.StatusInternalServerError, response.StatUnexpectedError, response.StatUnexpectedError, "")
	}

	return response.NewSuccessResponse(user, response.StatOK, getUserSuccessMessage)
}
