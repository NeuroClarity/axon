package database

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/lib/pq"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
	"github.com/NeuroClarity/axon/pkg/domain/core"
)

func NewDatabase(username, password, endpoint, port, dbName string) (*gateway.Database, error) {
  dbDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, endpoint, port, dbName)
  db, err := sql.Open("postgres", dbDsn)
  if err != nil {
    return nil, err
  }
  return &database{dbClient : db}, nil
}

type database struct {
	dbClient  *sql.DB
}

func (repo database) NewReviewer(firstName, lastName, email string, demographics core.Demographics) (int, error) {
  // query for demographics, if it doesn't exist create it and get the key

  // put the user in the database
  sql := "INSERT INTO REVIEWER VALUES()"
  _, err := repo.dbClient.Exec(sql, firstName, lastName, email)
	return 0, nil
}

func (repo database) GetReviewer() (*core.Reviewer, error) {
	// database logic
  var name string
  sql := "SELECT name FROM reviewers WHERE uid = $1"
  err := repo.dbClient.QueryRow(sql, uid).Scan(&name)
  if err != nil {
    return "", errors.New(fmt.Sprintf("Failed to get reviewer with uid %d: %s", uid, err))
  }
  // populare the reviewer struct
	return &core.Reviewer{}, nil
}

func (repo database) GetReviewJob() (core.ReviewJob, error) {
	// database logic
	return core.ReviewJob{}, nil
}

func (repo database) NewCreator(name string) (int, error) {
	// database logic
	return 0, nil
}

func (repo database) GetCreator() (*core.Creator, error) {
	// database logic
	return &core.Creator{}, nil
}

func (repo database) GetStudy() (*core.Study, error) {
	// database logic
	return &core.Study{}, nil
}

