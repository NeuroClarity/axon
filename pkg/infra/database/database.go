package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// NewDatabase creates a new instance of a relational database service.
func NewDatabase(username, password, endpoint, port, dbName string) (gateway.Database, error) {
	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, endpoint, port, dbName)
	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		return nil, err
	}
	return database{dbClient: db}, nil
}

type database struct {
	dbClient *sql.DB
}

// NewReviewer instantiates a new Reviewer in our database.
func (db database) NewReviewer(uid, firstName, lastName, email string, dem core.Demographics) error {
	// check if demographics exist
	var demographicsId int
	query := "SELECT uid FROM demographics WHERE gender = $1 AND race = $2 AND age = $3"
	err := db.dbClient.QueryRow(query, dem.Gender, dem.Race, dem.Age).Scan(&demographicsId)
	if err == sql.ErrNoRows {
		_, err = db.dbClient.Exec("INSERT INTO demographics VALUES(DEFAULT, $1, $2, $3)", dem.Age, dem.Gender, dem.Race)
		if err != nil {
			return errors.New(fmt.Sprintf("Error when adding demographics to database: %s", err))
		}
		demographicsId, err = db.getCurrentIdFromTable("demographics_uid_seq")
	}

	// put the user in the database
	query = "INSERT INTO reviewer VALUES($1, $2, $3, $4, $5)"
	_, err = db.dbClient.Exec(query, uid, firstName, lastName, email, demographicsId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error when adding user to the database: %s", err))
	}

	return nil
}

// GetReviewer retreives the reviewer's profile given the uid of that reviewer.
// Returns nil for both reviewer and error if the user does not exist
func (db database) GetReviewer(uid string) (*core.Reviewer, error) {
	var firstName, lastName, email string
	var demographicsId int
	query := "SELECT first_name, last_name, email, demographics_id FROM reviewer WHERE uid = $1"
	err := db.dbClient.QueryRow(query, uid).Scan(&firstName, &lastName, &email, &demographicsId)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get reviewer with uid %s: %s", uid, err))
	}

	var age int
	var gender, race string
	query = "SELECT age, gender, race FROM demographics WHERE uid = $1"
	err = db.dbClient.QueryRow(query, demographicsId).Scan(&age, &gender, &race)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("Demographics with id %d does not exist", demographicsId))
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when getting demographics with id %d: %s", demographicsId, err))
	}

	reviewer := &core.Reviewer{
		UID:       uid,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Demographics: core.Demographics{
			Age:    age,
			Gender: gender,
			Race:   race,
		},
	}
	// populare the reviewer struct
	return reviewer, nil
}

func (db database) UpdateReviewerWithReviewJob(uid string, reviewJob *core.ReviewJob) error {
	return nil
}

// NewCreator adds a new creator profile to the database.
func (db database) NewCreator(uid, firstName, lastName, email, company string) error {
	// put the user in the database
	query := "INSERT INTO creator VALUES($1, $2, $3, $4, $5)"
	_, err := db.dbClient.Exec(query, uid, firstName, lastName, email, company)
	if err != nil {
		return errors.New(fmt.Sprintf("Error when adding user to the database: %s", err))
	}

	return nil
}

// GetCreator retreives the creator's profile given the uid of that creator.
func (db database) GetCreator(uid string) (*core.Creator, error) {
	// database logic
	var firstName, lastName, email, company string
	query := "SELECT first_name, last_name, email, company FROM creator WHERE uid = $1"
	err := db.dbClient.QueryRow(query, uid).Scan(&firstName, &lastName, &email, &company)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get reviewer with uid %s: %s", uid, err))
	}

	creator := &core.Creator{
		UID:       uid,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Company:   company,
	}
	return creator, nil
}

// NewStudy adds a new study in the database and returns the unique id of that study.
func (db database) NewStudy(creatorId, videoKey string, req *core.TargetAudience) (int, error) {
	// no need to check if the study already exists because of key constraint
	reviewCount, ageMax, ageMin := req.NumParticipants, req.MaxAge, req.MinAge
	gender, race := req.Gender, req.Race
	eegHeadset, eyeTracking := req.Eeg, req.EyeTracking

	query := "INSERT INTO study VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, err := db.dbClient.Exec(query, creatorId, videoKey, reviewCount, reviewCount, ageMax, ageMin, gender,
		race, eegHeadset, eyeTracking)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Error when adding study to the database: %s", err))
	}

	studyId, err := db.getCurrentIdFromTable("study_uid_seq")
	return studyId, err
}

// GetStudy retreives a study by creatorID and videoKey
func (db database) GetStudy(uid int) (*core.Study, error) {
	var creatorId, videoKey, gender, race string
	var reviewCount, reviewsRemaining, ageMax, ageMin int
	var eeg, eyeTracking bool

	query := "SELECT creator_id, video_key, review_count, reviews_remaining, age_max, age_min, gender, race, eeg_headset, eye_tracking FROM study WHERE uid = $1"
	err := db.dbClient.QueryRow(query, uid).Scan(&creatorId, &videoKey, &reviewCount, &reviewsRemaining, &ageMax, &ageMin, &gender, &race, &eeg, &eyeTracking)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("There are no studies with uid %d", uid))
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occured when querying the study table"))
	}

	// get the associated reviews
	reviews, err := db.GetStudyReviews(creatorId, videoKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when retrieving study reviews: %s", err))
	}
	// get the associated creator profile
	creator, err := db.GetCreator(creatorId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error when retrieving user profile: %s", err))
	}

	study := core.Study{
		UID:          uid,
		NumRemaining: reviewsRemaining,
		TargetAudience: &core.TargetAudience{
			NumParticipants: reviewCount,
			MinAge:          ageMin,
			MaxAge:          ageMax,
			Gender:          gender,
			Race:            race,
			Eeg:             eeg,
			EyeTracking:     eyeTracking,
		},
		Reviews: reviews,
		Creator: creator,
		Content: core.Content{
			VideoLocation: videoKey,
		},
	}
	fmt.Println(study)

	return &study, nil
}

// GetAllStudies retreives all Studies in the database.
func (db database) GetAllStudies(creatorId string) ([]*core.Study, error) {
	var results []*core.Study

	query := "SELECT video_key, review_count, reviews_remaining, age_max, age_min, gender, race, eeg_headset, eye_tracking FROM study WHERE creator_id = $1"
	rows, err := db.dbClient.Query(query, creatorId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occured when querying the study table"))
	}
	defer rows.Close()

	for rows.Next() {
		var videoKey, gender, race string
		var reviewCount, reviewsRemaining, ageMax, ageMin int
		var eeg, eyeTracking bool
		err = rows.Scan(&videoKey, &reviewCount, &reviewsRemaining, &ageMax, &ageMin, &gender, &race, &eeg, &eyeTracking)

		// get the associated reviews
		reviews, err := db.GetStudyReviews(creatorId, videoKey)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error when retrieving study reviews: %s", err))
		}
		// get the associated creator profile
		creator, err := db.GetCreator(creatorId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error when retrieving user profile: %s", err))
		}

		study := core.Study{
			NumRemaining: reviewsRemaining,
			TargetAudience: &core.TargetAudience{
				NumParticipants: reviewCount,
				MinAge:          ageMin,
				MaxAge:          ageMax,
				Gender:          gender,
				Race:            race,
				Eeg:             eeg,
				EyeTracking:     eyeTracking,
			},
			Reviews: reviews,
			Creator: creator,
			Content: core.Content{
				VideoLocation: videoKey,
			},
		}
		results = append(results, &study)
	}

	return results, nil
}

// NewReview adds a new review to the database.
func (db database) NewReview(reviewerId, videoKey, creatorId, eeg core.EEGData, webcam core.WebcamData) error {
	query := "INSERT INTO review VALUES($1, $2, $3, $4, $5)"
	_, err := db.dbClient.Exec(query, reviewerId, videoKey, creatorId, eeg.Location, webcam.Location)
	if err != nil {
		return errors.New(fmt.Sprintf("Error when adding review to the database: %s", err))
	}

	return nil
}

func (db database) NewReviewJob(studyID int, reviewerID string, completed time.Time) error {
	return nil
}

func (db database) GetReviewJob(demo core.Demographics, hardware core.Hardware) (*core.ReviewJob, error) {
	// for build purposes
	_ = hardware
	var videoKey string
	query := "SELECT video_key FROM study WHERE gender = $1 AND race = $2 AND age_min < $3 AND age_max > $3 AND reviews_remaining > 0"
	err := db.dbClient.QueryRow(query, demo.Gender, demo.Race, demo.Age).Scan(&videoKey)
	if err == sql.ErrNoRows {
		// TODO: Log the actual demographics in the error here for debugging purposes
		return nil, errors.New(fmt.Sprintf("Unable to find a study with matching demographics"))
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occured when querying the study table: %s", err))
	}

	return &core.ReviewJob{Study: &core.Study{Content: core.Content{VideoLocation: videoKey}}}, nil
}

func (db database) GetReviewJobByStudy(study *core.Study) (*core.ReviewJob, error) {
	return nil, nil
}

func (db database) GetStudyReviews(creatorId, videoKey string) ([]*core.Review, error) {
	var results []*core.Review

	query := "SELECT reviewer_id, eeg_data_key, eye_data_key FROM review WHERE creator_id = $1 AND video_key = $2"
	rows, err := db.dbClient.Query(query, creatorId, videoKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error occured when querying the reviews table"))
	}
	defer rows.Close()

	for rows.Next() {
		var reviewerId, eegKey, eyeKey string
		err = rows.Scan(&reviewerId, &eegKey, &eyeKey)
		// get reviewer
		reviewer, err := db.GetReviewer(reviewerId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error when retrieving reviewer data: %s", err))
		}

		review := &core.Review{
			Reviewer: reviewer,
			Insights: core.Insights{
				EEG:         eegKey,
				EyeTracking: eyeKey,
			},
		}

		results = append(results, review)
	}

	return results, nil
}

/********************/
/* HELPER FUNCTIONS */
/********************/

func (db database) getCurrentIdFromTable(primaryKey string) (int, error) {
	// get the newly generated uid
	var uid int
	err := db.dbClient.QueryRow("SELECT currval($1)", primaryKey).Scan(&uid)
	if err != nil {
		// if there are no rows in the table return 1
		var temp int
		err := db.dbClient.QueryRow("SELECT uid FROM demographics LIMIT 1").Scan(&temp)
		if err == sql.ErrNoRows {
			return 1, nil
		} else {
			return -1, errors.New(fmt.Sprintf("Error occured when getting the unique id: %s", err))
		}
	}
	return uid, nil
}
