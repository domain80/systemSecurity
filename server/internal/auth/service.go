package auth

type Service interface {
  Login(userId int, password string)(string, error)
  AddUser(u *User) (*User, error) 
  ChangeRole(userId int, newRole string) error 
  // Logout(userId string) error
  LogAction(l *Log)  error
}
