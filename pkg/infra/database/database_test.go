package database

import (
  "fmt"
	"testing"
)


func TestConnectToDatabaseSuccessfully(t *testing.T) {
  _, err := NewDatabase("neuroc", "NeuroCDB12", "nc-database.cr7v5oc2x2xe.us-west-1.rds.amazonaws.com", "5432", "nc-database")
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
    return
  }

  fmt.Println("Successfully connected to database!")
	return
}
