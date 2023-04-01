package service

import (
	"encoding/xml"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"log"
	"os"
	"strconv"
	"time"
)

type (
	PushNotificationServiceImpl struct {
		projectRootPath string
		repo            domain.PushNotificationRepo
		ctxTimeout      time.Duration
	}
)

func NewPushNotificationService(path string, repo domain.PushNotificationRepo, t time.Duration) PushNotificationServiceImpl {
	return PushNotificationServiceImpl{
		projectRootPath: path,
		repo:            repo,
		ctxTimeout:      t,
	}
}

func (notify PushNotificationServiceImpl) SendNotifyMessage(sender string, receptor []string, msg string) error {
	return notify.repo.SendMessage(sender, receptor, msg)
}

func (notify PushNotificationServiceImpl) GetSenderNotifyMessage(info domain.AccountInfoOutput, templatePath string) string {
	var data []byte
	data = ReadTemplateFile(notify.projectRootPath + templatePath)
	t := domain.SenderMessageTemplate{}
	xml.Unmarshal(data, &t)
	t.Amount = info.TransactionAmount
	t.FromCard = info.CardNum
	t.Balance = strconv.FormatInt(info.Balance, 10)

	return fmt.Sprintf("%s%s%s%s%s%s", t.Texts[0], t.FromCard, t.Texts[1], t.Amount, t.Texts[2], t.Balance)
}

func (notify PushNotificationServiceImpl) GetReceiverNotifyMessage(info domain.AccountInfoOutput, templatePath string) string {
	var data []byte
	data = ReadTemplateFile(notify.projectRootPath + templatePath)
	t := domain.ReceiverMessageTemplate{}
	xml.Unmarshal(data, &t)
	t.Amount = info.TransactionAmount
	t.ToCard = info.CardNum
	t.Balance = strconv.FormatInt(info.Balance, 10)

	return fmt.Sprintf("%s%s%s%s%s%s", t.Texts[0], t.ToCard, t.Texts[1], t.Amount, t.Texts[2], t.Balance)
}

func ReadTemplateFile(filePath string) []byte {
	body, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	//t := domain.SenderMessageTemplate{}
	//xml.Unmarshal(body, &t)
	//fmt.Println(t, string(body))
	return body
}
