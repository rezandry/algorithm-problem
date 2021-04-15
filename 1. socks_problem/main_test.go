package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_getPair(t *testing.T) {
	type args struct {
		shocks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shorted list without pair",
			args: args{
				shocks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 0,
		}, {
			name: "shorted list with pair",
			args: args{
				shocks: []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9},
			},
			want: 5,
		}, {
			name: "unshorted list with pair",
			args: args{
				shocks: []int{1, 7, 3, 6, 9, 1, 7, 3, 5, 9},
			},
			want: 4,
		}, {
			name: "unshorted list without pair",
			args: args{
				shocks: []int{1, 7, 3, 6, 9, 2, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPair(tt.args.shocks); got != tt.want {
				t.Errorf("getPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockFmt struct {
	mock.Mock
}

func (m MockFmt) Scanln(a ...interface{}) (n int, err error) {
	args := m.Called(a)
	return args.Get(0).(int), args.Error(1)
}

func (m MockFmt) Scan(a ...interface{}) (n int, err error) {
	args := m.Called(a)
	return args.Get(0).(int), args.Error(1)
}
