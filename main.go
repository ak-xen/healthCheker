package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	//fmt.Println("Enter the path to the file...")
	var path string = "urls.txt"
	/*_, err := fmt.Scan(&path)
	if err != nil {
		return
	}*/

	var urls []string = GetData(path)

	var data map[string]int

	ticker := time.NewTicker(3 * time.Second)
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		text.Colors{text.FgHiCyan, text.Bold}.Sprint("URL"),
		text.Colors{text.FgHiCyan, text.Bold}.Sprint("Status Code"),
	})
	t.SetStyle(table.StyleColoredBright)

	colorKey := text.Colors{text.FgYellow, text.BgWhite}
	colorValue := text.Colors{text.FgGreen, text.BgWhite}
	color404 := text.Colors{text.FgRed, text.BgWhite}
	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("Exiting...")
			return
		case <-ticker.C:
			t.ResetRows()
			data = queryManager(urls)
			for k, v := range data {
				cKey := colorKey.Sprint(k)
				if isRightCode(v) {

					cValue := colorValue.Sprint(v)
					colorValue.Sprint(v)
					t.AppendRow(table.Row{cKey, cValue})
				} else {
					c404 := color404.Sprint("404 Not Found")
					t.AppendRow(table.Row{cKey, c404})
				}

			}
			fmt.Print("\033[H\033[2J")
			t.Render()
		}
	}

}
