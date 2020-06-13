package book

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseCount(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "below 10k",
			args: args{"100"},
			want: 100,
		},
		{
			name: "over 10k",
			args: args{"123.45万"},
			want: 1234500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCount(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTimeAt(t *testing.T) {
	type args struct {
		s string
	}
	var at = time.Date(2020, 06, 13, 12, 00, 00, 00, TZ)
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "date",
			args: args{"2020-06-07"},
			want: time.Date(2020, 06, 07, 00, 00, 00, 00, TZ),
		},
		{
			name: "relative minutes",
			args: args{"8分钟前"},
			want: time.Date(2020, 06, 13, 11, 52, 00, 00, TZ),
		},
		{
			name: "relative hours",
			args: args{"2小时前"},
			want: time.Date(2020, 06, 13, 10, 00, 00, 00, TZ),
		},
		{
			name: "yesterday am",
			args: args{"昨日12:34"},
			want: time.Date(2020, 06, 12, 12, 34, 00, 00, TZ),
		},
		{
			name: "yesterday pm",
			args: args{"昨日23:45"},
			want: time.Date(2020, 06, 12, 23, 45, 00, 00, TZ),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeAt(tt.args.s, at)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
