package service

import (
	"GST/billlingSystem/model"
	"GST/billlingSystem/utils"
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (b *billerService) Login(u model.User) (token string, err error) {

	var role string

	isExist, err := b.store.Login(u)
	if err != nil {
		err = errors.New("error in log in")
		return "", err
	}
	if isExist {
		userID := uuid.NewString()
		token, err = utils.GenerateToken(role, userID, u.Email)
		if err != nil {
			logrus.WithField("err", err.Error()).Error("error generating jwt token for user")
			return "", err

		}
	}
	return token, nil

}
