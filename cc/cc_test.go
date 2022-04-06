package cc

import "testing"

func Test_convert(t *testing.T) {
	gwei, _ := Convert(1, "GWEI")
	eth, _ := Convert(1, "ETH")

	if eth-(gwei*1000000000) > .0001 {
		t.Errorf("expected %f to be %f", eth, gwei*1000000000)
	}
}
func Test_parse(t *testing.T) {
	gwei := ParseCoin("GWEI")
	eth := ParseCoin("ETH")

	if gwei.Ticker != "ETH" {
		t.Error("GWEI ticker should be ETH")
	}
	if gwei.Unit != "GWEI" {
		t.Error("GWEI Unit should be GWEI")
	}
	if gwei.Exp != -9 {
		t.Error("GWEI exp should be -9")
	}

	if eth.Unit != "ETH" {
		t.Error("ETH Unit should be ETH")
	}
	if eth.Exp != 1 {
		t.Error("ETH ticker should be ETH")
	}
	if eth.Ticker != "ETH" {
		t.Error("ETH Exp should be 1")
	}
}
