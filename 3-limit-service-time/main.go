//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

func (u *User) AddTime() {
	u.TimeUsed += 1
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	isDone := make(chan bool)

	go func() {
		process()
		isDone <- true
		close(isDone)
	}()

	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			u.AddTime()
			if u.TimeUsed > 9 {
				return false
			}
		case <-isDone:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
