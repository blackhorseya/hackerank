package main

import (
	"testing"
)

func Test_lonelyinteger(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "[1,2,3,4,3,2,1] then 4",
			args: args{a: []int32{1, 2, 3, 4, 3, 2, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lonelyinteger(tt.args.a); got != tt.want {
				t.Errorf("lonelyinteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
