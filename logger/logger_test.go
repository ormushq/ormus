package logger_test

import (
	"bytes"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/ormushq/ormus/logger"
)

func TestHandler(t *testing.T) {
	l := logger.L

	tests := []struct {
		f    func()
		want string
	}{
		{
			f: func() {
				l.Info("test", "key", "val")
			},
			want: `{"test": {"key": "val"}}`,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var got bytes.Buffer
			os.Stdout.Read(got.Bytes())
			test.f()
			if !reflect.DeepEqual(got, []byte(test.want)) {
				t.Fatalf("want: %+v, got: %+v", test.want, got.String())
			}
		})
	}
}
