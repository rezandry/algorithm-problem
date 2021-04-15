package main

import "testing"

func Test_countingValleys(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			/*
			  /\/\
			_/    \
			UUDUDD
			*/
			name: "hike climb without valley",
			args: args{
				input: []string{"U","U","D","U","D","D"},
			},
			want: 0,
		},
		{
			/*
			_/\    /\
			   \/\/
			UDDUDUUD
			*/
			name: "hike climb with one valley",
			args: args{
				input: []string{"U","D","D","U","D","U","U","D"},
			},
			want: 2,
		},
		{
			/*
			_/\      /\
			   \    /
			    \/\/
			UDDDUDUUUD
			*/
			name: "hike climb with valley bellow altitude",
			args: args{
				input: []string{"U","D","D","D","U","D","U","U","U","D"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingValleys(tt.args.input); got != tt.want {
				t.Errorf("countingValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}
