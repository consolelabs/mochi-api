package util

import (
	"testing"
)

func TestConvertInputToChainId(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid input chain",
			args: args{
				input: "eth",
			},
			want: "1",
		},
		{
			name: "valid input chainId",
			args: args{
				input: "1",
			},
			want: "1",
		},
		{
			name: "valid input chainId",
			args: args{
				input: "128",
			},
			want: "128",
		},
		{
			name: "valid input chain",
			args: args{
				input: "heco",
			},
			want: "128",
		},
		{
			name: "valid input marketplace",
			args: args{
				input: "quixotic",
			},
			want: "10",
		},
		{
			name: "valid input marketplace",
			args: args{
				input: "paintswap",
			},
			want: "250",
		},
		{
			name: "valid input marketplace",
			args: args{
				input: "opensea",
			},
			want: "1",
		},
		{
			name: "invalid input chain",
			args: args{
				input: "invalid",
			},
			want: "",
		},
		{
			name: "invalid input chainId",
			args: args{
				input: "9999",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertInputToChainId(tt.args.input)
			if got != tt.want {
				t.Errorf("ConvertInputToChainId() = %v, want %v", got, tt.want)
			}
		})
	}
}
