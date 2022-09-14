package paramlang

import (
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type syntaxError struct {
	line   int
	column int
	msg    string
}

//nolint:errname
type syntaxErrorListener struct {
	Errors []syntaxError
	*antlr.DefaultErrorListener
}

func newSyntaxErrorListener() *syntaxErrorListener {
	return new(syntaxErrorListener)
}

func (c *syntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	// Default error message: fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
	c.Errors = append(c.Errors, syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (c *syntaxErrorListener) Error() string {
	var str strings.Builder
	sep := ""
	for _, e := range c.Errors {
		_, _ = str.WriteString(sep + "line " + strconv.Itoa(e.line) + ":" + strconv.Itoa(e.column) + " " + e.msg)
		sep = "\n"
	}
	return str.String()
}
