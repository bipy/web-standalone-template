package schedules

import (
	"time"
	"web-standalone-template/pkg/repository"
)

func StartHelloSchedule() {
	go updateId()
}

func updateId() {
	for {
		repository.ID++
		time.Sleep(100 * time.Second)
	}
}
