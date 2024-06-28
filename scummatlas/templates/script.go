package templates

import (
	"fmt"
	"html/template"
	"os"
	"scummatlas/scummatlas"
	l "scummatlas/scummatlas/condlog"
	s "scummatlas/scummatlas/script"
)

type scriptData struct {
	Title   string
	Scripts []s.Script
}

func (self scriptData) ScriptTable() (out []template.HTML) {
	for i, _ := range self.Scripts {
		if i%12 == 0 {
			out = append(out, template.HTML("</tr><tr>"))
		}
		out = append(out, template.HTML(fmt.Sprintf(
			"<td><a href='#script-%d'>Script %d</a></td>",
			i, i)))
	}
	return
}

func writeScripts(game scummatlas.Game, outdir string) {
	filename := outdir + "/scripts.html"
	file, err := os.Create(filename)
	l.Log("template", "Create "+filename)
	if err != nil {
		panic("Can't create index file")
	}

	t := template.Must(template.ParseFiles(
		htmlPath+"scripts.html",
		htmlPath+"partials.html"))
	t.Execute(file, scriptData{
		game.Name,
		game.Scripts,
	})
}
