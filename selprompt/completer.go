package selprompt

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"strings"
)

type Completer struct {
	
}

func NewCompleter() *Completer {
	return &Completer{}
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	var suggests []prompt.Suggest

	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	w := d.GetWordBeforeCursor()
	words := strings.Split(d.TextBeforeCursor(), " ")

	currentArgPos := 0
	for i := range words {
		if words[i] == " " {
			continue
		}
		currentArgPos ++
	}

	if currentArgPos == 1 {
		if len(w) >= 2 {
			suggests = commandSuggest
			suggests = prompt.FilterFuzzy(suggests, d.GetWordBeforeCursor(), true)
			return suggests
		}
	}
	fileCompleter := completer.FilePathCompleter{
		IgnoreCase: true,
		Filter: nil,
	}

	if currentArgPos == 2 {
		return fileCompleter.Complete(d)
	}

	return []prompt.Suggest{}
}