package database

import (
	"fmt"
	"testing"
)

var db, _ = NewDatabase("neuroc", "NeuroCDB12", "nc-database.cr7v5oc2x2xe.us-west-1.rds.amazonaws.com", "5432", "postgres")

func TestConnectToDatabaseSuccessfully(t *testing.T) {
	_, err := NewDatabase("neuroc", "NeuroCDB12", "nc-database.cr7v5oc2x2xe.us-west-1.rds.amazonaws.com", "5432", "postgres")
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
		return
	}

	fmt.Println("Successfully connected to database!")
	return
}

func TestGetReviewerByUID(t *testing.T) {
	user, err := db.GetReviewer("test-uid")
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
		return
	}

	fmt.Println(user)
	return
}

// Temp testing functions without mocking
//func TestPutNewReviewer(t *testing.T) {
//err := db.NewReviewer("test-uid2", "tester", "last", "helloworld@gmail.com", core.Demographics{
//Gender: "female",
//Age:    20,
//Race:   "other",
//})

//if err != nil {
//t.Errorf(fmt.Sprintf("%s", err))
//return
//}

//fmt.Println("Successfully created user object!")
//return
//}
