package Concurency

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ak-xen/healthCheker/Query"
	"github.com/ak-xen/healthCheker/TelegramBot"
	"github.com/ak-xen/healthCheker/Ui"
)

func GetAndDrawData(urls []string) {

	var data map[string]int

	ticker := time.NewTicker(3 * time.Second)
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("Exiting...")
			go TelegramBot.SendMessageBot("Script stop execution.... Exiting")
			time.Sleep(1 * time.Second)
			return
		case <-ticker.C:
			data = Query.QueryManager(urls)
			Ui.DrawTable(data)

		}
	}
}
