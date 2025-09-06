package Ui

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func DrawTable(data map[string]int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Важно: задаём стиль с цветами вручную!
	style := table.StyleColoredBright
	style.Color = table.ColorOptions{
		Header:       text.Colors{text.FgMagenta, text.Bold},
		Row:          text.Colors{text.FgGreen},
		RowAlternate: text.Colors{text.FgGreen},
		Footer:       text.Colors{text.FgMagenta, text.Bold},
	}
	t.SetStyle(style)

	t.AppendHeader(table.Row{
		text.Colors{text.FgMagenta, text.Bold}.Sprint("URL"),
		text.Colors{text.FgMagenta, text.Bold}.Sprint("Status Code"),
	})

	colorKey := text.Colors{text.FgYellow}
	colorValue := text.Colors{text.FgGreen}
	color404 := text.Colors{text.FgRed}

	for k, v := range data {
		cKey := colorKey.Sprint(k)
		if v != 0 {
			t.AppendRow(table.Row{cKey, colorValue.Sprint(v)})
		} else {
			t.AppendRow(table.Row{cKey, color404.Sprint("404 Not Found")})

		}
	}

	fmt.Print("\033[H\033[2J") // Очистка экрана (опционально)
	t.Render()
}
