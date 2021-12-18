package repl

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"

	parser "github.com/batazor/shortlink/internal/pkg/shortdb/parser/v1"
	session "github.com/batazor/shortlink/internal/pkg/shortdb/session/v1"
)

type repl struct {
	session *session.Session
}

func New(s *session.Session) (*repl, error) {
	return &repl{
		session: s,
	}, nil
}

func (r *repl) Run() {
	r.help()

	for {
		t := prompt.Input("> ", completer)

		if t == "" {
			continue
		}

		switch t[0] {
		case '.':
			s := strings.Split(t, " ")

			switch s[0] {
			case ".close":
				return
			case ".open":
				if err := r.open(t); err != nil {
					pterm.FgRed.Println(err)
				}
			case ".help":
				r.help()
			default:
				pterm.FgRed.Println("incorrect command")
			}
		default: // if this not command then this SQL-expression
			p, err := parser.New(t)
			if err.Error() != "" {
				pterm.FgRed.Println(err)
				continue
			}

			fmt.Println(p.Query)
		}
	}
}
