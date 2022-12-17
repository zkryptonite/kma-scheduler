package notifier

import (
	"backend/driver"
	"backend/repository"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"sync"
)

func BatchSending(from, subject, body string) {
	var (
		mg                = mailgun.NewMailgun(domain, privateApiKey)
		db                = driver.ConnectMongoDB().Client.Database(dbName)
		studentRepository = repository.NewStudentRepository(db)
		students, _       = studentRepository.FindStudentsWithEmail()
	)

	var mailers []Mailer

	for _, s := range students {
		m := Mailer{
			From:    from,
			Subject: subject,
			Body:    body,
			To:      s.Email,
		}

		mailers = append(mailers, m)
	}

	var (
		mailerChan = make(chan Mailer)
		wg         sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range mailerChan {
			wg.Add(1)
			go func(m Mailer) {
				defer wg.Done()
				err := m.SendMail(mg)
				log.Println(err)
			}(m)
		}
	}()

	for _, m := range mailers {
		mailerChan <- m
	}
	close(mailerChan)

	wg.Wait()
}
