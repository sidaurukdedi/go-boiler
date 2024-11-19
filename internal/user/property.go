package user

import (
	"time"

	"github.com/sirupsen/logrus"
)

type UserUsecaseProperty struct {
	ServiceName    string
	Logger         *logrus.Logger
	Location       *time.Location
	UserRepository UserRepository
}
