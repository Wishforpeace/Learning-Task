package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"test/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(username string, password string) (map[string]interface{}, error) {
	sqlStr := "select user_id,user_password from users where user_name = ?"
	stmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	userId := 0
	accountPassword := ""
	err = stmt.QueryRow(username).Scan(&userId, &accountPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid username or password.\r\n")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid username or passowrd.\r\n")

	}
	sqlStr = "insert into authentication_token(user_id,aurh_token,generated_at,expires_at) values(?,?,?,?)"
	stmt, err = utils.DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	randomToken := make([]byte, 32)
	_, err = rand.Read(randomToken)
	if err != nil {
		return nil, err
	}
	authToken := base64.URLEncoding.EncodeToString(randomToken)
	const timeLayout = "2021-12-23 00:00:00"
	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)

	generatedAt := dt.Format(timeLayout)
	expiresAt := expirtyTime.Format(timeLayout)

	_, err = stmt.Exec(userId, authToken, generatedAt, expiresAt)

	if err != nil {
		return nil, err
	}

	tokenDetails := map[string]interface{}{
		"token_type":   "Bearer",
		"auth_token":   authToken,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}
	return tokenDetails, nil
}

func ValidateToken(authToken string) (map[string]interface{}, error) {
	queryString := `select
				system_users.user_id,
				username,
				generted_at,
				expires_at
			from authentication_tokens
			left join system_users
			on authentication_tokens.user_id =system_users.user_id
			where auth_token = ?`
	stmt, err := utils.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	userId := 0
	username := ""
	generatedAt := ""
	expiresAt := ""

	err = stmt.QueryRow(authToken).Scan(&userId, &username, &generatedAt, &expiresAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid access token.\r\n")
		}
		return nil, err
	}

	const timeLayout = "2021-12-23 00:00:00"

	expiryTime, _ := time.Parse(timeLayout, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The token is expired.\r\n")
	}

	userDetails := map[string]interface{}{
		"user_id":      userId,
		"username":     username,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}
	return userDetails, nil
}
