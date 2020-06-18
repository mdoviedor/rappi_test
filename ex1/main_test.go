package main

import "testing"

func TestHandler(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case x = 1 and y = 2",
			args: args{
				x: 1,
				y: 2,
			},
			want: 4,
		},
		{
			name: "case x = 4 and y = 2",
			args: args{
				x: 4,
				y: 2,
			},
			want: 16,
		},
		{
			name: "case x = 2 and y = 4",
			args: args{
				x: 2,
				y: 4,
			},
			want: 12,
		},
		{
			name: "case x = -2 and y = -4",
			args: args{
				x: -2,
				y: -4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Handler(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
