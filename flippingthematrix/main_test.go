package main

import (
	"testing"
)

func Test_flippingMatrix(t *testing.T) {
	type args struct {
		matrix [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "[[1,2],[3,4]]",
			args: args{[][]int32{
				{1, 2},
				{3, 4},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippingMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("flippingMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
