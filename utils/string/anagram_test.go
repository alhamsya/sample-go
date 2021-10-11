package utils

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	type args struct {
		listString []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "When_reqSingleWord_expectOneWord",
			args: args{
				listString: []string{"kita"},
			},
			want: [][]string{
				{
					"kita",
				},
			},
		},
		{
			name: "When_reqSingleWord_expectOneWord",
			args: args{
				listString: []string{
					"kita",
					"atik",
					"tika",
					"aku",
					"kia",
					"makan",
					"kua",
				},
			},
			want: [][]string{
				{
					"kita",
					"atik",
					"tika",
				},
				{
					"aku",
					"kua",
				},
				{
					"kia",
				},
				{
					"makan",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagram(tt.args.listString...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
