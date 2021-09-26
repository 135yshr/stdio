package stdio

import (
	"bufio"
	"fmt"
	"io"
)

// Stdio is interface of input and output
type Stdio struct {
	Writer io.Writer
	Reader io.Reader
}

// Options is option for input api
type Options struct {
	Required bool
	DefaultValue string
}

func (o *Options) AskTextByDefaultValue() string {
	if o.DefaultValue == "" {
		return ""
	}
	return fmt.Sprintf(" [%s]", o.DefaultValue)
}

// Ask asks the user for input using the specified query.
// The response is returned as a string.
func (s *Stdio) Ask(question string, opt *Options) (string, error) {
	scanner := bufio.NewScanner(s.Reader)
	for {
		if _, err := fmt.Fprintf(s.Writer, "%s%s: ", question, opt.AskTextByDefaultValue()); err != nil {
			return "", err
		}
		scanner.Scan()
		in := scanner.Text()
		if in == "" {
			in = opt.DefaultValue
		}
		if !opt.Required {
			return in, nil
		}
		if in != "" {
			return in, nil
		}
		if _, err := fmt.Fprintln(s.Writer, ""); err != nil {
			return "", err
		}
	}
}
