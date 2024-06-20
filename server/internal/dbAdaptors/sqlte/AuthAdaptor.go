package sqlte

import (
	"database/sql"
	"fmt"
	"security418/internal/auth"
)

// AddUser(u *User) (*User,error)
// AddUser inserts a new user into the database and returns the created user.
func (this *adaptor) AddUser( u *auth.User) (*auth.User, error) {
	query := `
	INSERT INTO users ( first_name, last_name, role, password, email)
	VALUES (?, ?, ?, ?, ?, ?)
	RETURNING id;
	`
	stmt, err := this.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	var lastInsertID int
	err = stmt.QueryRow( u.FirstName, u.LastName, u.Role, u.Password, u.Email).Scan(&lastInsertID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	u.ID = lastInsertID
	return u, nil
}


  // GetUser(userId int) (*User, error)
// GetUser retrieves a user by their userId.
func (this *adaptor) GetUser( userId int) (*auth.User, error) {
	query := `SELECT id,  first_name, last_name, role, email FROM users WHERE id = ?`
	stmt, err := this.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	var user auth.User
	err = stmt.QueryRow(userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Role, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &user, nil
}


  // ChangeRole(userId int, newRole string) error 
  // LogAction(l *Log) 
func (this *adaptor) ChangeRole( userId int, newRole string) error {
	query := `UPDATE users SET role = ? WHERE id = ?`
	stmt, err := this.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(newRole, userId)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}


func (this *adaptor) LogAction( l *auth.Log) error {
	query := `INSERT INTO logs (user_id, action) VALUES (?, ?)`
	stmt, err := this.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.UserID, l.Action)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

