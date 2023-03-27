package repository

type KaveNegarNotifyMsg struct {
}

func NewKaveNegarNotifyMsg() KaveNegarNotifyMsg {
	return KaveNegarNotifyMsg{}
}

func (notify KaveNegarNotifyMsg) SendSuccessfulMessage(msg string) error {
	return nil
}

func (notify KaveNegarNotifyMsg) SendFailureMessage(msg string) error {
	return nil
}
