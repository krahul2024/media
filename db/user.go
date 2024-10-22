package db

import "database/sql"

type User struct {
    ID int `json:"id"`
    Name string`json:"name"`
    UserName string `json:"username"`
    Email string `json:"email"`
    City string `json:"city"`
    Country string `json:"country"`
    CreatedAt string `json:"createdAt"`
}

type UserStore interface {
    Add(name, username, email, city, country string) error
    GetAll() ([]User, error) 
    Get(id int) (User, error)
}

type SQLUserStore struct {
    DB *sql.DB
}

func (s *SQLUserStore) Add(name, username, email, city, country string) error {
    _, err := s.DB.Exec(
        "Insert Into users (name, username, email, city, country) Values ($1, $2, $3, $4, $5)",
        name, username, email, city, country,
    )
    return err
}

func (s *SQLUserStore) GetAll() ([]User, error) {
    rows, err := s.DB.Query("Select id, name, username, email, city, country, createdAt From users")

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        var user User

        if err := rows.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.City, &user.Country, &user.CreatedAt); err != nil {
            return nil, err;
        }
        users = append(users, user)
    }

    return users, nil
}

func (s *SQLUserStore) Get(id int) (User, error) {
    var user User
    err := s.DB.QueryRow(
        "Select id, name, username, email, city, country From users Where id=$1", id,
        ).Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.City, &user.Country)

    return user, err
}
