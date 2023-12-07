package main

import (
	"testing"
)

func Test_evalString(t *testing.T) {
	type args struct {
		text    string
		reverse bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two74119onebtqgnine",
			args: args{
				text:    "two74119onebtqgnine",
				reverse: false,
			},
			want: 2,
		},
		{
			name: "two74119onebtqgnine",
			args: args{
				text:    "two74119onebtqgnine",
				reverse: true,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalString(tt.args.text, tt.args.reverse); got != tt.want {
				t.Errorf("evalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bneightwo8kpkkgbxgnineqjkt4fiveninesevennine",
			args: args{
				line: "bneightwo8kpkkgbxgnineqjkt4fiveninesevennine",
			},
			want: 89,
		},
		{
			name: "fivethreeonezblqnsfk1",
			args: args{
				line: "fivethreeonezblqnsfk1",
			},
			want: 51,
		},
		{
			name: "two74119onebtqgnine",
			args: args{
				line: "two74119onebtqgnine",
			},
			want: 29,
		},
		{
			name: "jrjh5vsrxbhsfour3",
			args: args{
				line: "jrjh5vsrxbhsfour3",
			},
			want: 53,
		},
		{
			name: "tn5eightfncnzcdtthree8",
			args: args{
				line: "tn5eightfncnzcdtthree8",
			},
			want: 58,
		},
		{
			name: "kpmrk5flx",
			args: args{
				line: "kpmrk5flx",
			},
			want: 55,
		},
		{
			name: "fkxxqxdfsixgthreepvzjxrkcfk6twofour",
			args: args{
				line: "fkxxqxdfsixgthreepvzjxrkcfk6twofour",
			},
			want: 64,
		},
		{
			name: "dqbx6six5twoone",
			args: args{
				line: "dqbx6six5twoone",
			},
			want: 61,
		},
		{
			name: "glmsckj2bvmts1spctnjrtqhmbxzq",
			args: args{
				line: "glmsckj2bvmts1spctnjrtqhmbxzq",
			},
			want: 21,
		},
		{
			name: "7sixthreerzmpbffcx",
			args: args{
				line: "7sixthreerzmpbffcx",
			},
			want: 73,
		},
		{
			name: "zhss9gfxfgmrmzthreefivevpkljfourtwoeight",
			args: args{
				line: "zhss9gfxfgmrmzthreefivevpkljfourtwoeight",
			},
			want: 98,
		},
		{
			name: "6tfzvrbkfour",
			args: args{
				line: "6tfzvrbkfour",
			},
			want: 64,
		},
		{
			name: "86pbrnhhhpn",
			args: args{
				line: "86pbrnhhhpn",
			},
			want: 86,
		},
		{
			name: "two87",
			args: args{
				line: "two87",
			},
			want: 27,
		},
		{
			name: "jqnzhnzsdgg3eightfournbgsdxbfg31xrnv",
			args: args{
				line: "jqnzhnzsdgg3eightfournbgsdxbfg31xrnv",
			},
			want: 31,
		},
		{
			name: "4three3sxvnnzqvhcxj9",
			args: args{
				line: "4three3sxvnnzqvhcxj9",
			},
			want: 49,
		},
		{
			name: "mpjmgpmninekffbkgprkb9tzzdznine8rtq",
			args: args{
				line: "mpjmgpmninekffbkgprkb9tzzdznine8rtq",
			},
			want: 98,
		},
		{
			name: "xhpfptdmpkeightdkqvpzscx7jndgsxxtqh",
			args: args{
				line: "xhpfptdmpkeightdkqvpzscx7jndgsxxtqh",
			},
			want: 87,
		},
		{
			name: "17vvktdvdgpthreeone1gqnb",
			args: args{
				line: "17vvktdvdgpthreeone1gqnb",
			},
			want: 11,
		},
		{
			name: "fzpbhtmdvseven2one",
			args: args{
				line: "fzpbhtmdvseven2one",
			},
			want: 71,
		},
		{
			name: "sixzkkbjmnxnxm3gcvpsddqqgzktp63h",
			args: args{
				line: "sixzkkbjmnxnxm3gcvpsddqqgzktp63h",
			},
			want: 63,
		},
		{
			name: "nineonexkqkfszg8qdzrlqszpltwo9cseven",
			args: args{
				line: "nineonexkqkfszg8qdzrlqszpltwo9cseven",
			},
			want: 97,
		},
		{
			name: "7xfhtdfdldzplmrfivegpkgfseventwo",
			args: args{
				line: "7xfhtdfdldzplmrfivegpkgfseventwo",
			},
			want: 72,
		},
		{
			name: "xvone5two7",
			args: args{
				line: "xvone5two7",
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.line); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
