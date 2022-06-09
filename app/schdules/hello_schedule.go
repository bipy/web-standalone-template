package schedules

import (
	"time"
	"web-standalone-template/pkg/repository"
)

func updateId() {
	for {
		repository.ID++
		time.Sleep(100 * time.Second)
	}
}
