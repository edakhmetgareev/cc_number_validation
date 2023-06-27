package main_test

import (
	"testing"

	"github.com/edakhmetgareev/sumup"
)

func TestDetectCardSchemeName(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid american express",
			args: args{
				cardNumber: "3782 8224 6310 005",
			},
			want: "American Express",
		},
		{
			name: "invalid american express",
			args: args{
				cardNumber: "3782 8224 6310 0050",
			},
			want: "Unknown",
		},
		{
			name: "valid jcb",
			args: args{
				cardNumber: "3530 1113 3330 0000",
			},
			want: "JCB",
		},
		{
			name: "invalid jcb",
			args: args{
				cardNumber: "3530 1113 3330 0000 0000",
			},
			want: "Unknown",
		},
		{
			name: "valid maestro",
			args: args{
				cardNumber: "6304 0000 0000 0000",
			},
			want: "Maestro",
		},
		{
			name: "invalid maestro",
			args: args{
				cardNumber: "6304 0000 0000 0000 1111",
			},
			want: "Unknown",
		},
		{
			name: "valid visa",
			args: args{
				cardNumber: "4111 1111 1111 1111",
			},
			want: "Visa",
		},
		{
			name: "invalid visa",
			args: args{
				cardNumber: "4000 1111 1111 1111",
			},
			want: "Unknown",
		},
		{
			name: "valid mastercard",
			args: args{
				cardNumber: "5555 5555 5555 4444",
			},
			want: "MasterCard",
		},
		{
			name: "invalid mastercard",
			args: args{
				cardNumber: "5755 5555 5555 4444",
			},
			want: "Unknown",
		},
		{
			name: "invalid card number",
			args: args{
				cardNumber: "0000 0000 0000 0000",
			},
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.DetectCardScheme(tt.args.cardNumber); got != tt.want {
				t.Errorf("DetectCardScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
