package auth

import (
	"errors"

	"github.com/sethvargo/go-password/password"
)

type authService struct {
  repo AuthRepo
}

func NewAuthService(authRepo AuthRepo) (authService) {
  return authService{
    repo: authRepo,
  }
}
  // AddUser(u *User) (*User,error)
func(this authService) Login(userId int, password string) (string,  error) {
  user,err := this.repo.GetUser(userId)
  if(err != nil) {
    return "", err
  }
  _password := GetMD5Hash(password)
  if(user.Password != _password) {
    return "", errors.New("Invalid password")
  }
  token, err := getJwtToken(user)
  if(err != nil ) {
    return "", errors.New("Internal server error")
  }

  return token, nil
}


  // GetUser(userId string) (*User, error)
func (this authService)  GetUser(userId int) (*User, error) {
  return this.repo.GetUser(userId)
}
  // LogAction(l *Log) 
func (this authService)  LogAction(l *Log) (error) {
  return this.repo.LogAction(l)
}

// ChangeRole(userId int, newRole string) error 
func (this authService) ChangeRole(userId int, newRole string) error {
  return this.repo.ChangeRole(userId, newRole)
}



func (this authService)  AddUser(u *User) (*User, error) {

  pwd, err:= password.Generate(10, 4, 2, false,  false)
  if(err != nil) {
    return nil, err
  }

  u.Password = pwd
  newUser, err := this.repo.AddUser(u)
  if  err != nil {
    return nil, err
  }

  // send email to user
  return newUser, err
}


