package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/s654632396/tools/selprompt"
)
func main() {

	myEs := selprompt.NewExecutorSession()
	myCompleter := selprompt.NewCompleter()

	p := prompt.New(
		myEs.Executor,
		myCompleter.Complete,
		[]prompt.Option{
			prompt.OptionTitle("now, for testing"),
			prompt.OptionPrefix(" ➜ "),
			prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
			//prompt.OptionHistory(),
		}...
	)
	p.Run()
}


