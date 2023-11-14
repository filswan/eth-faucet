package chain

import (
	"math/big"
	"reflect"
	"testing"
)

func TestIsValidAddress(t *testing.T) {
	type args struct {
		address     string
		checksummed bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid address", args: args{address: "0xab5801a7d398351b8be11c439e05c5b3259aec9b", checksummed: false}, want: true},
		{name: "invalid address", args: args{address: "invalid address", checksummed: false}, want: false},
		{name: "address without 0x", args: args{address: "ab5801a7d398351b8be11c439e05c5b3259aec9b", checksummed: false}, want: true},
		{name: "valid checksum address", args: args{address: "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B", checksummed: true}, want: true},
		{name: "invalid checksum address", args: args{address: "0xab5801a7d398351b8be11c439e05c5b3259aec9b", checksummed: true}, want: false},
		{name: "checksum address without 0x", args: args{address: "Ab5801a7D398351b8bE11C439e05C5B3259aeC9B", checksummed: true}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAddress(tt.args.address, tt.args.checksummed); got != tt.want {
				t.Errorf("IsValidAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEtherToWei(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
		want   *big.Int
	}{
		{name: "1ether", amount: 1, want: new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EtherToWei(tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtherToWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringEthToWei(t *testing.T) {
	type args struct {
		amount string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 0.01 Ether",
			args: args{amount: "0.01"},
			want: big.NewInt(10000000000000000), // 0.01 ether in wei
		},
		{
			name: "Test 1 Ether",
			args: args{amount: "1"},
			want: big.NewInt(1000000000000000000), // 1 ether in wei
		},
		{
			name: "Test 1.5 Ether",
			args: args{amount: "1.5"},
			want: big.NewInt(1500000000000000000), // 1.5 ether in wei
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringEthToWei(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringEthToWei() = %v, want %v", got, tt.want)
			}
		})
	}
}
