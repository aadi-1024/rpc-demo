package main

import "fmt"

type MailRPC struct{}
type MailPayload struct {
	Email   string
	Content string
}

func (m *MailRPC) SendMail(info MailPayload, succ *bool) error {
	err := appConfig.SendMail(info.Email, info.Content)
	if err != nil {
		*succ = false
		fmt.Println(err)
		return err
	}
	*succ = true
	return nil
}
