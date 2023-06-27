package main_test

import (
	"testing"

	"github.com/edakhmetgareev/sumup"
)

func TestIsValidCardNumber(t *testing.T) {
	type args struct {
		cardNum string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid card number",
			args: args{
				cardNum: "5237 2516 2477 8133",
			},
			want: true,
		},
		{
			name: "invalid card number",
			args: args{
				cardNum: "5237 2516 2477 8134",
			},
			want: false,
		},
		{
			name: "empty card number",
			args: args{
				cardNum: "",
			},
			want: false,
		},
		{
			name: "wrong size card number",
			args: args{
				cardNum: "5237 2516",
			},
			want: false,
		},
		{
			name: "card number with letters",
			args: args{
				cardNum: "5237 2516 abcd eeee",
			},
			want: false,
		},
		{
			name: "zero card",
			args: args{
				cardNum: "0000 0000 0000 0000",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.IsValidCardNumber(tt.args.cardNum); got != tt.want {
				t.Errorf("IsValidCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
