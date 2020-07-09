package database

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/lib/pq"
	//"github.com/NeuroClarity/axon/pkg/domain/gateway"
	//"github.com/NeuroClarity/axon/pkg/domain/core"
)

func NewDatabase(username, password, endpoint, port, dbName string) (*database, error) {
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

func (repo database) GetReviewerByUID(uid int) (string, error) {
  var name string
  sql := "SELECT name FROM reviewers WHERE uid = $1"
  err := repo.dbClient.QueryRow(sql, uid).Scan(&name)
  if err != nil {
    return "", errors.New(fmt.Sprintf("Failed to get reviewer with uid %d: %s", uid, err))
  }
  // populare the reviewer struct
  return name, nil
}

//func (repo database) GetReviewJob() core.ReviewJob {
	//// database logic
	//return core.ReviewJob{}
//}

//func (repo database) NewClient(name string) (int, error) {

//}

//func (repo database) GetClient(name string) (*core.Client, error) {

//}

//func (repo database) GetStudy() (*core.Study, error) {

//}
