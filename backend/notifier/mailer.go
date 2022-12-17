package notifier

import (
	"context"
	"encoding/json"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

type Mailer struct {
	From    string
	Subject string
	Body    string
	To      string
}

func NewMailer(from, subject, body, receiver string) *Mailer {
	return &Mailer{
		From:    from,
		Subject: subject,
		Body:    body,
		To:      receiver,
	}
}

func (m *Mailer) SendMail(mg *mailgun.MailgunImpl) error {
	var (
		message     = mg.NewMessage(m.From, m.Subject, m.Body, m.To)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	)
	defer cancel()

	r := createBody(m.To)

	if len(r) != 0 {
		message.SetTemplate(templateName)
		vars, _ := json.Marshal(r)
		message.AddHeader("X-Mailgun-Variables", string(vars))
	}

	_, _, err := mg.Send(ctx, message)

	return err
}

func CreateTemplate() error {
	var (
		mg          = mailgun.NewMailgun(domain, privateApiKey)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	)
	defer cancel()

	return mg.CreateTemplate(ctx, &mailgun.Template{
		Name: templateName,
		Version: mailgun.TemplateVersion{
			Template: `'<h1>Lịch học hôm nay của bạn:</h1>
						<ul>
							{{#each classes}}
								<li>Môn {{this.subject}} {{this.lesson}} phòng {{this.room}}</li>
							{{/each}}
						</ul>'`,
			Engine: mailgun.TemplateEngineHandlebars,
			Tag:    "v1",
		},
	})
}
