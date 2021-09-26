package stdio

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStdio_Ask(t *testing.T) {
	type args struct {
		question string
		opt      *Options
		input    string
	}
	type want struct {
		out   string
		value string
	}
	tests := map[string]struct {
		args    args
		want    want
		wantErr bool
	}{
		"Should displayed `What is your name`": {
			args: args{question: "What is your name:", opt: &Options{}, input: "\n"},
			want: want{out: "What is your name:", value: ""},
		},
		"Should return `Your name`": {
			args: args{question: "", opt: &Options{}, input: "Your name\n"},
			want: want{out: "", value: "Your name"},
		},
		"Should return `My name`": {
			args: args{question: "", opt: &Options{Required: true}, input: "\nMy name"},
			want: want{out: "\n", value: "My name"},
		},
		"Should displayed `What is your name:_nWhat is your name:`": {
			args: args{question: "What is your name:", opt: &Options{Required: true}, input: "\nMy name"},
			want: want{out: "What is your name:\nWhat is your name:", value: "My name"},
		},
		"Should return `Default value`": {
			args: args{question: "", opt: &Options{DefaultValue: "Default value"}, input: "\n"},
			want: want{out: "", value: "Default value"},
		},
	}
	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			var in, out bytes.Buffer
			in.Write([]byte(tt.args.input))

			s := &Stdio{
				Writer: &out,
				Reader: &in,
			}
			got, err := s.Ask(tt.args.question, tt.args.opt)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.value, got)
			assert.Equal(t, tt.want.out, out.String())
		})
	}
}
