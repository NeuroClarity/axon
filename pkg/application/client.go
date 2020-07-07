package app

import "fmt"

func ClientRegister() string {
	return "Client Register. \n"
}

func ClientLogin(cid int) string {
	return fmt.Sprintf("Client Login uid: %d.\n", cid)
}

func CreateStudy() string {
	return "Study Created. \n"
}

func ViewStudy(sid int) string {
	return fmt.Sprintf("Study sid: %d.\n", sid)
}
