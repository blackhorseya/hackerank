package main

import (
	"testing"
)

func Test_findMedian(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "[0,1,2,4,6,5,3] then 3",
			args: args{arr: []int32{0, 1, 2, 4, 6, 5, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedian(tt.args.arr); got != tt.want {
				t.Errorf("findMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
