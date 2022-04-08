package pagination

import (
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	type T struct {
		str string
	}
	type args struct {
		tar   []*T
		index uint
		size  uint
	}
	data := []*T{{str: "1"}, {str: "2"}, {str: "3"}, {str: "4"}, {str: "5"}}
	tests := []struct {
		name  string
		args  args
		want  []*T
		want1 uint
	}{
		{
			name: "normal page test",
			args: args{
				tar:   data,
				index: 2,
				size:  2,
			},
			want:  []*T{{str: "3"}, {str: "4"}},
			want1: 3,
		},
		{
			name: "out of range",
			args: args{
				tar:   data,
				index: 9,
				size:  2,
			},
			want:  []*T{{str: "5"}},
			want1: 3,
		},
		{
			name: "zero test",
			args: args{
				tar:   []*T{},
				index: 1,
				size:  25,
			},
			want:  []*T{},
			want1: 0,
		},
		{
			name: "nil array",
			args: args{
				tar:   nil,
				index: 1,
				size:  25,
			},
			want:  []*T{},
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Page(tt.args.tar, tt.args.index, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Page() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Page() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
