package util

import (
	"reflect"
	"testing"
	"time"
)

func TestRandomString(t *testing.T) {
	prevStr := ""
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		args   args
		strLen int
		isFail bool
	}{
		{
			name: "case successfully",
			args: args{
				n: 10,
			},
			strLen: 10,
		},
		{
			name: "fail because of different length",
			args: args{
				n: 10,
			},
			strLen: 8,
			isFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.args.n)
			if !tt.isFail && len(got) != tt.strLen {
				t.Errorf("RandomString() = %v, want %v", len(got), tt.strLen)
				return
			}
			if got == prevStr {
				t.Error("RandomString() should not gen same string")
				return
			}
			prevStr = got
		})
	}
}

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "minh@gmail.com",
			},
			want: true,
		},
		{
			name: "invalid email",
			args: args{
				email: "minh.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.args.email); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSaltedPassword(t *testing.T) {
	// salted password for 123456
	saltedPass := "$1024$2YuHqMAoVUs605apb8907f5394f6cd5c82acec9bb4892c7137b80c48"

	type args struct {
		password string
		salt     string
		loops    int
	}
	tests := []struct {
		name    string
		args    args
		result  string
		isFail  bool
		wantErr bool
	}{
		{
			name: "success with pwd 123456",
			args: args{
				password: "123456",
				salt:     "2YuHqMAoVUs605ap",
				loops:    1024,
			},
			result:  saltedPass,
			isFail:  false,
			wantErr: false,
		},
		{
			name: "fail because of wrong password",
			args: args{
				password: "1245",
				salt:     "2YuHqMAoVUs605ap",
				loops:    1024,
			},
			result:  saltedPass,
			isFail:  true,
			wantErr: false,
		},
		{
			name: "fail because of wrong salt",
			args: args{
				password: "123456",
				salt:     "2YuHq",
				loops:    1024,
			},
			result:  saltedPass,
			isFail:  true,
			wantErr: false,
		},
		{
			name: "success with salt = loop + salt",
			args: args{
				password: "123456",
				salt:     "$1024$2YuHqMAoVUs605ap",
				loops:    1024,
			},
			result:  saltedPass,
			isFail:  false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSaltedPassword(tt.args.password, tt.args.salt, tt.args.loops)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSaltedPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.isFail && got != tt.result {
				t.Errorf("GenerateSaltedPassword() = %v, want %v", got, tt.result)
				return
			}
			if tt.isFail && got == tt.result {
				t.Error("GenerateSaltedPassword() should fail")
				return
			}
		})
	}
}

func TestGenRandomInRange(t *testing.T) {
	min := 0
	max := 999
	if got := GenRandomInRange(min, max); got <= min || got >= max {
		t.Errorf("GenRandomInRange() = %v, out of range %v <= x <= %v", got, min, max)
	}
}

func TestSplitAndTrimSpaceString(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "return empty string list for empty string",
			args: args{
				s:   "",
				sep: "sep",
			},
			want: nil,
		},
		{
			name: "return list with one value when sep is empty",
			args: args{
				s:   " s ",
				sep: "",
			},
			want: []string{"s"},
		},
		{
			name: "return a string list",
			args: args{
				s:   " another|string ",
				sep: "|",
			},
			want: []string{"another", "string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitAndTrimSpaceString(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitAndTrimSpaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyMap(t *testing.T) {
	src := map[string]interface{}{
		"key1":   1,
		"string": "string",
	}

	if got := CopyMap(src); !reflect.DeepEqual(got, src) {
		t.Errorf("CopyMap() = %v, want %v", got, src)
	}

}
func TestHashNumber(t *testing.T) {
	first := int64(1)
	second := int64(2)
	if HashNumber(first) == HashNumber(second) {
		t.Errorf("HashNumber() hash %v = %v: %v", first, second, HashNumber(first))
	}
}

func TestFormatDiffTimeToHumanReadable(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}
	tests := []struct {
		name   string
		args   args
		result string
	}{
		{
			name: "year - month - day - hour",
			args: args{
				a: time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
				b: time.Date(2021, 6, 2, 1, 1, 1, 1, time.UTC),
			},
			result: "1 year 1 month 1 day 1 hour ",
		},
		{
			name: "month - day - hour",
			args: args{
				a: time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC),
				b: time.Date(2021, 2, 2, 2, 0, 0, 0, time.UTC),
			},
			result: "1 month 1 day 1 hour ",
		},
		{
			name: "day - hour",
			args: args{
				a: time.Date(2021, 1, 2, 1, 0, 0, 0, time.UTC),
				b: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			result: "1 day 1 hour ",
		},
		{
			name: "hour",
			args: args{
				a: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
				b: time.Date(2021, 1, 2, 1, 0, 0, 0, time.UTC),
			},
			result: "1 hour ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDiffTimeToHumanReadable(tt.args.a, tt.args.b)
			if got != tt.result {
				t.Errorf("FormatDiffTimeToHumanReadable() = %v, want %v", got, tt.result)
				return
			}
		})
	}
}
