package common

import "log"

const (
	DbTypeHome = 1
	DbTypePost = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	// GetEmail() string
	GetRole() string
}

// Use in panic() in a custom goroutine

func AppRecover() {
	err := recover()
	if err != nil {
		log.Println("Recovery error:", err)
	}
}
