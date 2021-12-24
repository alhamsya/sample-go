package number

import "testing"

func TestFibonacci(t *testing.T) {
	/* 0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946 */
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "When_numberRequest0_expect0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "When_numberRequest1_expect1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "When_numberRequest2_expect1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "When_numberRequest3_expect2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "When_numberRequest4_expect3",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "When_numberRequest5_expect5",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "When_numberRequest6_expect8",
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			name: "When_numberRequest7_expect13",
			args: args{
				n: 7,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
