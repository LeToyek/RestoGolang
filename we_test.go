package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Toyek(t *testing.T) {
	type param struct {
		a int
		b int
	}

	tests := []struct {
		name  string
		param param
		want  int
	}{
		{
			name: "a equal b",
			param: param{
				a: 1,
				b: 1,
			},
			want: 0,
		},
		{
			name: "a less than b",
			param: param{
				a: 2,
				b: 100000,
			},
			want: -1,
		},
		{
			name: "a more than b",
			param: param{
				a: 100,
				b: 20,
			},
			want: 2000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Toyek(test.param.a, test.param.b)
			assert.Equal(t, test.want, got)
		})
	}
}
