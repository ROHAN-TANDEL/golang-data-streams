package repository

import "database/sql"
import "log"
type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	ListUsers() []User
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) ListUsers() []User {
  rows, err := r.db.Query("SELECT id, first_name as name FROM users")

  log.Printf("query error:", err)
  if err != nil {
      // optionally log or return a fallback
      log.Printf("result is null");
      return []User{}
  }

	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	return users
}
