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

func TestGetReviewJob(t *testing.T) {

}

func TestGetStudy(t *testing.T) {
	study, err := db.GetStudy(2)
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
		return
	}

	fmt.Println(study)
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
// Commented out because I don't want to deal with this shit right now
//func TestPutReview(t *testing.T) {
//user, err := db.PutReview("test-uid", "some key", "creator")
//if err != nil {
//t.Errorf(fmt.Sprintf("%s", err))
//return
//}

//fmt.Println(user)
//return

//}

//func TestPutNewStudy(t *testing.T) {
//id, err := db.NewStudy("creator-uid-test", "video", &core.StudyRequest{
//NumParticipants: 20,
//MinAge:          10,
//MaxAge:          30,
//Gender:          "male",
//Race:            "other",
//Eeg:             true,
//EyeTracking:     true,
//})
//if err != nil {
//t.Errorf(fmt.Sprintf("%s", err))
//return
//}

//fmt.Println(id)
//return
//}

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

//func TestPutNewCreator(t *testing.T) {
//err := db.NewCreator("creator-uid-test", "creator-test", "lastname", "email", "test")
//if err != nil {
//t.Errorf(fmt.Sprintf("%s", err))
//return
//}
//return

//}
