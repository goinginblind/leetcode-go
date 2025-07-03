package main

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple 1",
			args: args{
				piles: []int{3, 6, 7, 11},
				h:     8,
			},
			want: 4,
		},
		{
			name: "Simple 2",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				h:     5,
			},
			want: 30,
		},
		{
			name: "Simple 3",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				h:     6,
			},
			want: 23,
		},
		{
			name: "Supa small 1",
			args: args{
				piles: []int{30},
				h:     1,
			},
			want: 30,
		},
		{
			name: "Supa small 2",
			args: args{
				piles: []int{1},
				h:     1,
			},
			want: 1,
		},
		{
			name: "Supa big num",
			args: args{
				piles: []int{865_995_455},
				h:     1,
			},
			want: 865_995_455,
		},
		{
			name: "Supa big num",
			args: args{
				piles: []int{1_000_000},
				h:     2,
			},
			want: 500_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
