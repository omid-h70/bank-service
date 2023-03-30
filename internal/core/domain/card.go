package domain

type Card struct {
	id         string
	accountId  string
	cardNumber string
}

func (c *Card) Id() string {
	return c.id
}

func (c *Card) SetId(id string) {
	c.id = id
}

func (c *Card) CardNum() string {
	return c.cardNumber
}

func (c *Card) SetCardNum(cardNum string) {
	c.cardNumber = cardNum
}

func (c *Card) AccountId() string {
	return c.accountId
}

func (c *Card) SetAccountId(accId string) {
	c.accountId = accId
}

func NewCard(cardNum string) Card {
	return Card{
		cardNumber: cardNum,
	}
}
