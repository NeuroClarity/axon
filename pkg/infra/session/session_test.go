package session

import (
  "fmt"
  "testing"
)

func TestSuccessfullyEstablishSession(t *testing.T) {
  _, err := NewSession("us-west-1")
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
    return
  }
  fmt.Println("Successfully created AWS session")
  return
}
