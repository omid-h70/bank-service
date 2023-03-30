package notification

import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)

const (
	BASE = "sms"
)

type KaveNegarNotifyMsg struct {
	apiKey string
}

func NewKaveNegarNotifyMsg(apikey string) KaveNegarNotifyMsg {
	return KaveNegarNotifyMsg{apikey}
}

func (notify KaveNegarNotifyMsg) SendMessage(sender string, receptor []string, message string) error {
	api := kavenegar.New(" your apikey ")
	//sender := ""
	//receptor := []string{"", ""}
	//message := "Hello Go!"
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}
	return nil
}
