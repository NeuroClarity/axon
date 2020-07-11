package database

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/lib/pq"
	//"github.com/NeuroClarity/axon/pkg/domain/gateway"
	"github.com/NeuroClarity/axon/pkg/domain/core"
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

/* 
* Adds the new user to the database and returns an error on failure
*/
func (repo database) NewReviewer(uid, firstName, lastName, email string, dem core.Demographics) error {
  // check if demographics exist
  var demographicsId int
  query := "SELECT id FROM demographics WHERE gender = $1 AND race = $2 AND age = $3"
  err := repo.dbClient.QueryRow(query, dem.Gender, dem.Race, dem.Age).Scan(&demographicsId)
  if err == sql.ErrNoRows {
    repo.dbClient.Exec("INSERT INTO demographics VALUES($1, $2, $3)", dem.Age, dem.Gender, dem.Race)
    demographicsId, err := repo.getCurrentIdFromTable("demographics_id_seq")
    if err != nil {
      return errors.New(fmt.Sprintf("Error when adding demographics to database: %s", err))
    }
  }

  // put the user in the database
  query = "INSERT INTO reviewer VALUES($1, $2, $3, $4, $5)"
  _, err = repo.dbClient.Exec(query, uid, firstName, lastName, email, demographicsId)
  if err != nil {
    return errors.New(fmt.Sprintf("Error when adding user to the database: %s", err))
  }

	return nil
}

/*
* Gets the user's profile given the uid of that user
*/
func (repo database) GetReviewer(uid string) (*core.Reviewer, error) {
  var firstName, lastName, email string
  var demographicsId int
  query := "SELECT first_name, last_name, email, demographics_id FROM reviewers WHERE uid = $1"
  err := repo.dbClient.QueryRow(query, uid).Scan(&firstName, &lastName, &email, &demographicsId)
  if err == sql.ErrNoRows {
    return nil, errors.New(fmt.Sprintf("Unable to find user with uid %s", uid))
  } else if err != nil {
    return nil, errors.New(fmt.Sprintf("Failed to get reviewer with uid %s: %s", uid, err))
  }

  var age int
  var gender, race string
  query = "SELECT age, gender, race FROM demographics WHERE id = $1"
  err = repo.dbClient.QueryRow(query, demographicsId).Scan(&age, &gender, &race)
  if err == sql.ErrNoRows {
    return nil, errors.New(fmt.Sprintf("Demographics with id %d does not exist", demographicsId))
  } else if err != nil {
    return nil, errors.New(fmt.Sprintf("Error when getting demographics with id %d: %s", demographicsId, err))
  }

  reviewer := &core.Reviewer{
    UID: uid,
    FirstName: firstName,
    LastName: lastName,
    Email: email,
    Demographics: core.Demographics{
      Age: age,
      Gender: gender,
      Race: race,
    },
  }
  // populare the reviewer struct
	return reviewer, nil
}

/* Adds a new creator profile to the database */
func (repo database) NewCreator(uid, firstName, lastName, email, company string) error {
  var temp string
  query := "SELECT first_name FROM creator WHERE uid = $1"
  err := repo.dbClient.QueryRow(query, uid).Scan(&temp)
  if err != sql.ErrNoRows {
    return errors.New(fmt.Sprintf("Creator with uid %s already exists", uid))
  } else if err != nil {
    return errors.New(fmt.Sprintf("Failed to get reviewer with uid %s: %s", uid, err))
  }
  // put the user in the database
  query = "INSERT INTO creator VALUES($1, $2, $3, $4, $5)"
  _, err = repo.dbClient.Exec(query, uid, firstName, lastName, email, company)
  if err != nil {
    return errors.New(fmt.Sprintf("Error when adding user to the database: %s", err))
  }

	return nil
}

func (repo database) GetCreator(uid string) (*core.Creator, error) {
	// database logic
  var firstName, lastName, email, company string
  query := "SELECT first_name, last_name, email, company FROM creator WHERE uid = $1"
  err := repo.dbClient.QueryRow(query, uid).Scan(&firstName, &lastName, &email, &company)
  if err == sql.ErrNoRows {
    return nil, errors.New(fmt.Sprintf("Unable to find user with uid %s", uid))
  } else if err != nil {
    return nil, errors.New(fmt.Sprintf("Failed to get reviewer with uid %s: %s", uid, err))
  }

  creator := &core.Creator{
    UID: uid,
    FirstName: firstName,
    LastName: lastName,
    Email: email,
    Company: company,
  }
	return creator, nil
}

/*
* Creates a new study in the database and return the unique id of that study
*/
func (repo database) NewStudy(creatorId int, videoKey string, reviewCount, ageMax, ageMin int,
                              gender, race string, eegHeadset, eyeTracking bool) error {
  // check if the study already exists
  var temp int
  query := "SELECT creator_id FROM study WHERE creatorId = $1 AND videoKey = $2"
  err := repo.dbClient.QueryRow(query, creatorId, videoKey).Scan(&temp)
  if err != sql.ErrNoRows {
    return errors.New("A study with that creator and video already exists")
  } else if err != nil {
    return errors.New("An error occured when querying the database")
  }

  query = "INSERT INTO study VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
  _, err = repo.dbClient.Exec(query, creatorId, videoKey, reviewCount, reviewCount, ageMax, ageMin, gender,
                               race, eegHeadset, eyeTracking)
  if err != nil {
    return errors.New(fmt.Sprintf("Error when adding study to the database: %s", err))
  }

	return nil
}

func (repo database) GetStudy(creatorId, videoKey string) (*core.Study, error) {
  var gender, race string
  var reviewCount, reviewsRemaining, ageMax, ageMin int
  var eeg, eyeTracking bool

  query := "SELECT review_count, reviews_remaining, age_max, age_min, gender, race, eeg_headset, eye_tracking FROM study WHERE creator_id = $1 AND video_key = $2"
  err := repo.dbClient.QueryRow(query, creatorId, videoKey).Scan(&reviewCount, &reviewsRemaining, &ageMax, &ageMin, &gender, &race, &eeg, &eyeTracking)
  if err == sql.ErrNoRows {
    return nil, errors.New(fmt.Sprintf("There are no records with creator_id %s and video_key %s", creatorId, videoKey))
  } else if err != nil {
    return nil, errors.New(fmt.Sprintf("Error occured when querying the study table"))
  }

  // get the associated reviews
  reviews, err := repo.GetStudyReviews(creatorId, videoKey)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Error when retrieving study reviews: %s", err))
  }
  // get the associated creator profile
  creator, err := repo.GetCreator(creatorId)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Error when retrieving user profile: %s", err))
  }

  study := core.Study{
    NumParticipants: reviewCount,
    NumRemaining: reviewsRemaining,
    StudyRequest: core.StudyRequest{
      MinAge: ageMin,
      MaxAge: ageMax,
      Gender: gender,
      Race: race,
      Eeg: eeg,
      EyeTracking: eyeTracking,
    },
    Reviews: reviews,
    Creator: *creator,
    Content: core.Content{
      VideoLocation: videoKey,
    },
  }

  return &study, nil
}

func (repo database) GetAllStudies(creatorId string) ([]*core.Study, error) {
  var results []*core.Study

  query := "SELECT video_key, review_count, reviews_remaining, age_max, age_min, gender, race, eeg_headset, eye_tracking FROM study WHERE creator_id = $1"
  rows, err := repo.dbClient.Query(query, creatorId)
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
    reviews, err := repo.GetStudyReviews(creatorId, videoKey)
    if err != nil {
      return nil, errors.New(fmt.Sprintf("Error when retrieving study reviews: %s", err))
    }
    // get the associated creator profile
    creator, err := repo.GetCreator(creatorId)
    if err != nil {
      return nil, errors.New(fmt.Sprintf("Error when retrieving user profile: %s", err))
    }

    study := core.Study{
      NumParticipants: reviewCount,
      NumRemaining: reviewsRemaining,
      StudyRequest: core.StudyRequest{
        MinAge: ageMin,
        MaxAge: ageMax,
        Gender: gender,
        Race: race,
        Eeg: eeg,
        EyeTracking: eyeTracking,
      },
      Reviews: reviews,
      Creator: *creator,
      Content: core.Content{
        video: videoKey,
      },
    }
    results = append(results, &study)
  }

	return results, nil
}

func (repo database) NewReview(reviewerId, videoKey, creatorId, eeg core.EEGData, webcam core.WebcamData) error {
  // check if the study already exists
  var temp int
  query := "SELECT reviewer_id FROM review WHERE reviewerId = $1 AND videoKey = $2"
  err := repo.dbClient.QueryRow(query, reviewerId, videoKey).Scan(&temp)
  if err != sql.ErrNoRows {
    return errors.New("A review with that reviewer and video already exists")
  } else if err != nil {
    return errors.New("An error occured when querying the database")
  }

  query = "INSERT INTO review VALUES($1, $2, $3, $4, $5)"
  _, err = repo.dbClient.Exec(query, reviewerId, videoKey, creatorId, eeg.s3Key, webcam.s3Key)
  if err != nil {
    return errors.New(fmt.Sprintf("Error when adding review to the database: %s", err))
  }

	return nil
}

// TODO: Implement for Milestone #2 (user dashboard)
func (repo database) GetReviewerReviews(reviewerId string) ([]*core.Review, error) {
  return nil, nil
}

func (repo database) GetStudyReviews(creatorId, videoKey string) ([]*core.Review, error) {
  var results []*core.Review

  query := "SELECT reviewer_id, eeg_data_key, eye_data_key FROM review WHERE creator_id = $1 AND video_key = $2"
  rows, err := repo.dbClient.Query(query, creatorId, videoKey)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Error occured when querying the reviews table"))
  }
  defer rows.Close()

  for rows.Next() {
    var reviewerId, eegKey, eyeKey string
    err = rows.Scan(&reviewerId, &eegKey, &eyeKey)
    // get reviewer
    reviewer, err := repo.GetReviewer(reviewerId)
    if err != nil {
      return nil, errors.New(fmt.Sprintf("Error when retrieving reviewer data: %s", err))
    }

    review := &core.Review{
      Reviewer: reviewer,
      Insights: core.Insights{
        Eeg: eegKey,
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

func (repo database) getCurrentIdFromTable(primaryKey string) (int, error) {
  // get the newly generated uid
  var uid int
  err := repo.dbClient.QueryRow("SELECT curval($1)", primaryKey).Scan(&uid)
  if err != nil {
    return -1, errors.New(fmt.Sprintf("Error occured when getting the unique id: %s", err))
  }
  return uid, nil
}
