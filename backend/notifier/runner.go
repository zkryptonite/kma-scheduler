package notifier

import (
	"log"
	"runtime"
	"github.com/prprprus/scheduler"
)

func Run() {
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		log.Fatalln(err)
	}

	task := func() {
		BatchSending("Admin <non-reply@example.com>", 
					 "Thông báo lịch học", 
					 "Hôm nay bạn không có lịch học")
	}

	s.Every().Second(10).Do(task)
	runtime.Goexit()
}