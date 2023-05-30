package store

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
)

type PostgresStore struct {
	db *sql.DB
}

func (p *PostgresStore) ConnectToDatabase() error {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres53"
		dbname   = "spo_users"
	)

	zap.S().Infof("Connecting to database ...")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	p.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer p.db.Close()
	zap.S().Infof("Connection established successfully with user: %v, db: %v", user, dbname)
	return err
}

func (p *PostgresStore) UserSignup(username string, hash string) error {
	sqlStatement := `
       INSERT INTO users (username, password)
       VALUES($1, $2)`
	zap.S().Infof("Query string for insert or update is : " + sqlStatement)
	_, err := p.db.Exec(sqlStatement, username, hash)
	if err != nil {
		zap.S().Errorf("Unable to insert the given post with error : %v", err)
	}
	zap.S().Infof("Post  inserted successfully !")
	return err
}

func (p *PostgresStore) UserLogin(username string, hash string) bool {
	sqlStatement := `
       SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE username = $1 AND password = $2
) AS row_exists;
;`
	var checkUser bool
	err := p.db.QueryRow(sqlStatement, username, hash).Scan(&checkUser)
	if err != nil {
		zap.S().Errorf("Error while calling Query : %v", err)
	}
	zap.S().Infof("result of queryRow is : %v", checkUser)
	return checkUser
}
