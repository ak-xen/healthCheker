package TelegramBot

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
)

func SendMessageBot(text string) {
	err := godotenv.Load("cmd/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	chatId := os.Getenv("CHAT_ID")

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%v&text=%s", token, chatId, text)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

}
