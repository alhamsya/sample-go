package utils

import (
	"reflect"
	"testing"
)

func TestBubbleSorting(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "When_arrayEmpty_expectNil",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
		{
			name: "When_arrayOneNumber_expectNumber",
			args: args{
				arr: []int{
					1,
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "When_arrayTwoNumber_expectSortingNumber",
			args: args{
				arr: []int{
					2,
					1,
				},
			},
			want: []int{
				1,
				2,
			},
		},
		{
			name: "When_arrayThereNumber_expectSortingNumber",
			args: args{
				arr: []int{
					3,
					1,
					4,
					2,
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
			//i = 0 -> 1, 3, 2, 4
			//i = 1 -> 1, 2, 3, 4
			//i = 2 -> 1, 2, 3, 4
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSorting(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSorting() = %v, want %v", got, tt.want)
			}
		})
	}
}
