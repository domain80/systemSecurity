package auth


type AuthRepo interface {
  AddUser(u *User) (*User,error)
  GetUser(userID int) (*User, error)
  ChangeRole(userId int, newRole string) error 
  LogAction(l *Log)  error
}

