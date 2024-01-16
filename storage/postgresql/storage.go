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
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", i.Host, i.Port, i.User, i.Password, i.Database)
}

type postgres struct {
	db     *sql.DB
	signal chan struct{}
}

func (p *postgres) FindUser(email string) (model.User, bool) {
	u := model.User{}
	query := "SELECT * FROM users WHERE email=$1"
	rows, err := p.db.Query(query, email)
	if rows.Err() != nil {
		fmt.Printf("FindUser query error: %v\n", rows.Err)
		return model.User{}, false
	}
	for rows.Next() {
		err = rows.Scan(&u)
		if err != nil {
			return model.User{}, false
		}
	}
	return u, true
}

func (p *postgres) GetUser(id model.UserIdentifier) (model.User, error) {
	u := model.User{}
	query := "SELECT * FROM users WHERE uid=$1;"
	rows, err := p.db.Query(query, id.Uid)
	if err != nil {
		fmt.Printf("GetUser query error: %v\n", err)
		return model.User{}, err
	}
	rows.Next()
	err = rows.Scan(&u)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (p *postgres) SetUser(user model.User) error {
	query := "INSERT INTO users (email, password, nickname) VALUES ($1, $2, $3); "
	result, err := p.db.Exec(query, user.Email, user.Password, user.Nickname)
	if err != nil {
		fmt.Printf("SetUser query error: %v\n", err)
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
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

func (p *postgres) Close() error {
	return p.db.Close()
}

func Storage(info DBInfo) (storage.Storage, error) {
	s := new(postgres)

	db, err := sql.Open("postgres", info.connString())
	if err != nil {
		return nil, err
	}
	s.db = db

	return s, nil
}
