package matcher

import (
	"testing"
)

func TestFieldMatcher_Matches(t *testing.T) {
	type testStruct struct {
		x string
		y int
		z uint64
		b bool
	}

	type fields struct {
		Key   string
		Value interface{}
	}
	type args struct {
		x interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "compare field",
			fields: fields{
				Key:   "x",
				Value: "val",
			},
			args: args{
				x: testStruct{
					x: "val",
				},
			},
			want: true,
		},
		{
			name: "compare int field",
			fields: fields{
				Key:   "y",
				Value: int64(1),
			},
			args: args{
				x: testStruct{
					y: 1,
				},
			},
			want: true,
		},
		{
			name: "compare uint field",
			fields: fields{
				Key:   "z",
				Value: uint64(1),
			},
			args: args{
				x: testStruct{
					z: 1,
				},
			},
			want: true,
		},
		{
			name: "compare bool field",
			fields: fields{
				Key:   "b",
				Value: false,
			},
			args: args{
				x: testStruct{
					b: false,
				},
			},
			want: true,
		},
		{
			name: "compare bool false",
			fields: fields{
				Key:   "b",
				Value: true,
			},
			args: args{
				x: testStruct{
					b: false,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fieldMatcher{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := m.Matches(tt.args.x); got != tt.want {
				t.Errorf("FieldMatcher.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFieldMatcher(t *testing.T) {
	got := NewFieldMatcher("key", "val")
	if got == nil {
		t.Errorf("NewFieldMatcher() data is nil")
	}
}

func Test_fieldMatcher_String(t *testing.T) {
	m := fieldMatcher{
		Key:   "Key",
		Value: "Value",
	}

	if got := m.String(); got != "obj.Key is equal to Value" {
		t.Errorf("fieldMatcher.String() = %v", got)
	}
}
