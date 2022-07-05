package util

import (
	"strings"
	"testing"
)

func TestConvertToChecksumAddr(t *testing.T) {
	type args struct {
		addrStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid address",
			args: args{
				addrStr: strings.ToLower("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79"),
			},
			want:    "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
			wantErr: false,
		},
		{
			name: "valid address",
			args: args{
				addrStr: strings.ToLower("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA"),
			},
			want:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			wantErr: false,
		},
		{
			name: "invalid address",
			args: args{
				addrStr: strings.ToLower("0x7D1070fdbF0eF8752a9627a79"),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToChecksumAddr(tt.args.addrStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToChecksumAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToChecksumAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
