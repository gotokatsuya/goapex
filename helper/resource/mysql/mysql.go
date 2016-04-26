package mysql

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

type DSN struct {
  Name     string
  Host     string
  Port     string
  User     string
  Password string
}

func NewDSN(name, host, port, user, password string) DSN {
  return DSN{
    Name:     name,
    Host:     host,
    Port:     port,
    User:     user,
    Password: password,
  }
}

func (d DSN) ToMySQL() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", d.User, d.Password, d.Host, d.Port, d.Name)
}

func Ping(name, host, port, user, password string) error {
  dsn := NewDSN(name, host, port, user, password)
  db, err := sql.Open("mysql", dsn.ToMySQL())
  if err != nil {
    return err
  }
  defer db.Close()
  return db.Ping()
}
