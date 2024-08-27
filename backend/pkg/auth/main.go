package auth

import (
	"fmt"
	"log"
	"sade-backend/api/models"
	"sade-backend/db/cmd"
	"sade-backend/pkg/utility"
	"time"
)

func New(userTable, linkTable *cmd.DataTable, timeout time.Duration) *Auth {
	return &Auth{
		userTable: userTable,
		linkTable: linkTable,
		timeout:   timeout,
	}
}

func (a *Auth) LoginUser(email, password, link string) error {
	result, err := a.userTable.CmdRead("email", email)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return fmt.Errorf("email does not exist")
	}
	verified := result[0]["verified"].(bool)
	if !verified {
		return fmt.Errorf("email is not verified")
	}
	role := result[0]["role"].(string)
	if role == "admin" {
		storedPass := result[0]["password"].(string)
		if !utility.CheckPass(password, storedPass) {
			return fmt.Errorf("password is incorrect")
		}
	}
	data := map[string]interface{}{
		"email":  email,
		"link":   link,
		"expiry": time.Now().Add(a.timeout).Unix(),
	}
	err = a.linkTable.CmdInsert(data)
	if err != nil {
		return fmt.Errorf("unable to insert link: %v", err)
	}
	return nil
}

func (a *Auth) ValidateLink(link string) (*models.User, bool, error) {
	result, err := a.linkTable.CmdRead("link", link)
	if err != nil {
		return nil, false, fmt.Errorf("error reading link from database: %v", err)
	}
	if len(result) == 0 {
		return nil, false, fmt.Errorf("link not found in database")
	}

	email, ok := result[0]["email"].(string)
	if !ok {
		return nil, false, fmt.Errorf("invalid or missing email in database result")
	}
	expiry, ok := result[0]["expiry"].(int64)
	if !ok {
		return nil, false, fmt.Errorf("invalid or missing expiry in database result")
	}

	if time.Now().Unix() > expiry {
		err = a.linkTable.CmdDelete("link", link)
		if err != nil {
			log.Println("Error deleting expired link from database:", err)
		}
		return nil, false, nil
	}

	userResult, err := a.userTable.CmdRead("email", email)
	if err != nil {
		return nil, false, fmt.Errorf("error reading user from database: %v", err)
	}
	if len(userResult) == 0 {
		return nil, false, fmt.Errorf("user not found for email: %v", email)
	}

	user := &models.User{
		ID:       userResult[0]["id"].(int64),
		Email:    userResult[0]["email"].(string),
		Role:     userResult[0]["role"].(string),
		Verified: userResult[0]["verified"].(bool),
	}

	if firstName, ok := userResult[0]["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := userResult[0]["last_name"].(string); ok {
		user.LastName = lastName
	}

	err = a.userTable.CmdUpdate(map[string]interface{}{"verified": true}, int(user.ID))
	if err != nil {
		log.Println("Error updating user in database:", err)
		return nil, false, fmt.Errorf("error updating user in database: %v", err)
	}

	err = a.linkTable.CmdDelete("link", link)
	if err != nil {
		log.Println("Error deleting link from database:", err)
	}
	return user, true, nil
}

func (a *Auth) RegisterUser(email, link, role, firstName, lastName, password string) error {
	result, err := a.userTable.CmdRead("email", email)
	if err != nil {
		return err
	}
	if len(result) != 0 {
		return fmt.Errorf("user already exists")
	}
	userData := map[string]interface{}{
		"email": email,
		"role":  role,
	}

	if firstName != "" && lastName != "" && password != "" {
		hashedPass, err := utility.Hash(password)
		if err != nil {
			return err
		}
		userData["password"] = hashedPass
		userData["first_name"] = firstName
		userData["last_name"] = lastName
	}

	err = a.userTable.CmdInsert(userData)
	if err != nil {
		return fmt.Errorf("unable to insert user: %v", err)
	}
	linkData := map[string]interface{}{
		"email":  email,
		"link":   link,
		"expiry": time.Now().Add(a.timeout).Unix(),
	}
	err = a.linkTable.CmdInsert(linkData)
	if err != nil {
		return fmt.Errorf("unable to insert link: %v", err)
	}
	return nil
}
