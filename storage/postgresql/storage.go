package postgresql

import (
	"auth-server/model"
	"auth-server/storage"
	"database/sql"
	"fmt"
)

type DBInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (i *DBInfo) connString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", info.Host, info.Port, info.User, info.Password, info.Database)
}

type postgres struct {
	db *sql.DB
}

func (p *postgres) FindUser(email string) (model.User, bool) {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) GetUser(id model.UserIdentifier) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) SetUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) RemoveUser(id model.UserIdentifier) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) RegisterToken(token string) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) UnregisterToken(token string) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) CheckToken(token string) error {
	//TODO implement me
	panic("implement me")
}

func Storage(info DBInfo) (storage.Storage, error) {
	s := new(postgres)

	db, err := sql.Open("*postgres", info.connString())
	if err != nil {
		return nil, err
	}
	s.db = db
	return s, nil
}
