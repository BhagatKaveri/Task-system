package db

import (
	"GST/billlingSystem/model"
	"log"
)

func (s *pgStore) Login(user model.User) (bool, error) {
	query := `SELECT EXISTS(
		SELECT 1 FROM users WHERE username=$1 and password_hash=$2
	) AS is_valid_login`
	isExist := false
	err := s.db.QueryRow(query, user.Email, user.Password).Scan(&isExist)
	if err != nil {
		log.Fatal("err db", err)
		return false, err
	}

	return isExist, err
}
