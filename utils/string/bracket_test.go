package utils

import "testing"

func TestFindFirstStringInBracketV2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "When_reqEmpty_expectEmpty",
			args: args{
				str: "",
			},
			want: "",
		},
		{
			name: "When_reqSingleCharacter_expectEmpty",
			args: args{
				str: "a",
			},
			want: "",
		},
		{
			name: "When_reqFirstSingleBracket_expectEmpty",
			args: args{
				str: "(a",
			},
			want: "",
		},
		{
			name: "When_reqLastSingleBracket_expectEmpty",
			args: args{
				str: "a)",
			},
			want: "",
		},
		{
			name: "When_reqFirstNestedBracket_expectEmpty",
			args: args{
				str: "((a",
			},
			want: "",
		},
		{
			name: "When_reqLastNestedBracket_expectEmpty",
			args: args{
				str: "a))",
			},
			want: "",
		},
		{
			name: "When_reqFirstDoubleBracket_expectSuccess",
			args: args{
				str: "((a)",
			},
			want: "a",
		},
		{
			name: "When_reqLastDoubleBracket_expectSuccess",
			args: args{
				str: "(a))",
			},
			want: "a",
		},
		{
			name: "When_reqNestedDoubleBracket_expectSuccess",
			args: args{
				str: "((a))",
			},
			want: "a",
		},
		{
			name: "When_reqDoubleBracket_expectSuccess",
			args: args{
				str: "al(alham)al(tes)",
			},
			want: "alham",
		},
		{
			name: "When_reqTwoDoubleBracket_expectSuccess",
			args: args{
				str: "al((alham))al((tes))",
			},
			want: "alham",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstStringInBracketV2(tt.args.str); got != tt.want {
				t.Errorf("FindFirstStringInBracketV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
