package database

import (
  "fmt"
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

//func (repo database) GetReviewerByUID(uid int) (*core.Reviewer, error) {
  //var name string
  //sql := "SELECT name FROM reviewers WHERE uid = $1"
  //err := db.QueryRow(sql, uid).Scan(&name)
  //if err != nil {
    //return nil, fmt.Sprintf("Failed to get reviewer with uid %d: %s", uid, err)
  //}
  //// populare the reviewer struct
//}

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
