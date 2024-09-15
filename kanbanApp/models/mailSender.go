package models

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"time"
)

func SendMailSimple(subject string, body string, to []string) {

	auth := smtp.PlainAuth(
		"",
		"berkayveysel1@gmail.com",
		"chev cfyo utzv kozr",
		"smtp.gmail.com",
	)
	msg := "Subject:" + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"berkayveysel1@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}
func SendMailSimpleHTML(subject string, templatePath string, to []string, name string, who string, whichProject string, projectid uint) {
	//get html

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct {
		Name         string
		Who          string
		WhichProject string
		ProjectID    uint
	}{Name: name, Who: who, WhichProject: whichProject, ProjectID: projectid})
	auth := smtp.PlainAuth(
		"",
		"berkayveysel1@gmail.com",
		"chev cfyo utzv kozr",
		"smtp.gmail.com",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"berkayveysel1@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}
func SendMailSimpleHTMLforRole(subject string, templatePath string, to []string, name string, who string, whichProject string, projectid uint, role string) {
	//get html

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct {
		Name         string
		Who          string
		WhichProject string
		ProjectID    uint
		Role         string
	}{Name: name, Who: who, WhichProject: whichProject, ProjectID: projectid, Role: role})
	auth := smtp.PlainAuth(
		"",
		"berkayveysel1@gmail.com",
		"chev cfyo utzv kozr",
		"smtp.gmail.com",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"berkayveysel1@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}
func SendMailSimpleHTMLForAssignment(subject string, templatePath string, to []string, name string, who string, whichProject, whichIssue string, projectid uint, date time.Time) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct {
		Name         string
		Who          string
		WhichProject string
		ProjectID    uint
		WhichIssue   string
		Date         time.Time
	}{Name: name, Who: who, WhichIssue: whichIssue, ProjectID: projectid, Date: date, WhichProject: whichProject})
	auth := smtp.PlainAuth(
		"",
		"berkayveysel1@gmail.com",
		"chev cfyo utzv kozr",
		"smtp.gmail.com",
	)
	if err != nil {
		fmt.Println("şablon yüklenirken bir hata oluştu")
		return
	}

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"berkayveysel1@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}
func SendMailSimpleHTMLforDiscard(subject string, templatePath string, to []string, name string, who string, whichProject string, projectid uint64) {
	//get html

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct {
		Name         string
		Who          string
		WhichProject string
		ProjectID    uint64
	}{Name: name, Who: who, WhichProject: whichProject, ProjectID: projectid})
	auth := smtp.PlainAuth(
		"",
		"berkayveysel1@gmail.com",
		"chev cfyo utzv kozr",
		"smtp.gmail.com",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"berkayveysel1@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}
