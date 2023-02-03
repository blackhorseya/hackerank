package main

import (
	"testing"
)

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[1,3,5,7,9] then 16 24",
			args: args{arr: []int32{1, 3, 5, 7, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			miniMaxSum(tt.args.arr)
		})
	}
}
