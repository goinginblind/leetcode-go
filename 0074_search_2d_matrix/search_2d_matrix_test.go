package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Square even matrix, tgt exists",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 3},
					{4, 5, 6, 7},
					{8, 9, 10, 11},
					{12, 13, 14, 15},
				},
				target: 13,
			},
			want: true,
		},
		{
			name: "Square odd matrix, tgt exists",
			args: args{
				matrix: [][]int{
					{-1, 0, 1, 2, 3},
					{4, 5, 6, 7, 8},
					{9, 9, 10, 11, 12},
					{12, 13, 14, 15, 15},
				},
				target: 13,
			},
			want: true,
		},
		{
			name: "Lots of non-decresing stuff, tgt exists",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 1, 2, 3, 4, 5, 6},
					{7, 7, 7, 7, 7, 8, 9, 9, 9},
					{10, 10, 10, 10, 11, 12, 12, 12, 13},
				},
				target: 8,
			},
			want: true,
		},
		{
			name: "Lots of non-decresing stuff, tgt absent",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 1, 1, 3, 4, 5, 6},
					{7, 7, 7, 7, 7, 8, 9, 9, 9},
					{10, 10, 10, 10, 11, 12, 12, 12, 13},
				},
				target: 2,
			},
			want: false,
		},
		{
			name: "Very short rows",
			args: args{
				matrix: [][]int{
					{1},
					{7},
					{10},
				},
				target: 10,
			},
			want: true,
		},
		{
			name: "Very short cols",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 2, 3, 4, 5, 6, 6, 6, 7, 8, 10, 12},
				},
				target: 10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
