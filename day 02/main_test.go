package main

import (
	"reflect"
	"testing"
)

func Test_getGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Game 1",
			args: args{
				line: "Game 1: 4 blue, 7 red, 5 green; 3 blue, 4 red, 16 green; 3 red, 11 green",
			},
			want: 0,
		},
		{
			name: "Game 2",
			args: args{
				line: "Game 2: 20 blue, 8 red, 1 green; 1 blue, 2 green, 8 red; 9 red, 4 green, 18 blue; 2 green, 7 red, 2 blue; 10 blue, 2 red, 5 green",
			},
			want: 0,
		},
		{
			name: "Game 3",
			args: args{
				line: "Game 3: 2 red, 5 green, 1 blue; 3 blue, 5 green; 8 blue, 13 green, 2 red; 9 green, 3 blue; 12 green, 13 blue; 3 green, 3 blue, 1 red",
			},
			want: 3,
		},
		{
			name: "Game 4",
			args: args{
				line: "Game 4: 1 red, 6 green, 4 blue; 3 green, 1 blue, 1 red; 7 blue, 1 red, 2 green",
			},
			want: 4,
		},
		{
			name: "Game 5",
			args: args{
				line: "Game 5: 2 green, 9 blue, 1 red; 3 green, 1 blue, 3 red; 1 red, 4 blue, 9 green",
			},
			want: 5,
		},
		{
			name: "Game 56",
			args: args{
				line: "Game 56: 1 blue; 3 red, 2 blue; 1 red, 2 green",
			},
			want: 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIDPower(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGame(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
