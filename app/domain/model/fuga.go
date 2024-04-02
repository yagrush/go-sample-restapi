package model

type Fuga struct {
	message string
}

func NewFuga(message string) *Fuga {
	return &Fuga{
		message: message,
	}
}

func (a *Fuga) GetMessage() string {
	return a.message
}
