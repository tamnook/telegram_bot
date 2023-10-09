package bot

import "fmt"

type Bot struct {
	Token string
}

func New() *Bot {
	fmt.Println("New")
	return &Bot{}
}
