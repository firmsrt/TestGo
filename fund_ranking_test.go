package main

import (
	"testing"
)

func TestValidateInputShouldBeReturn1D(t *testing.T) {
	result := ValidateInput("1d")
	if "1D" != result {
		t.Error("validateInput of 1d should be '1D' but have ", result)
	}
}

func TestValidateInputShouldBeReturn1W(t *testing.T) {
	result := ValidateInput("1W")
	if "1W" != result {
		t.Error("validateInput of 1d should be '1W' but have ", result)
	}
}

func TestValidateInputShouldBeReturn1M(t *testing.T) {
	result := ValidateInput("1M")
	if "1M" != result {
		t.Error("validateInput of 1d should be '1M' but have ", result)
	}
}

func TestValidateInputShouldBeReturn1Y(t *testing.T) {
	result := ValidateInput("1y")
	if "1Y" != result {
		t.Error("validateInput of 1d should be '1y' but have ", result)
	}
}

func TestValidateInputShouldBeReturnBlank(t *testing.T) {
	result := ValidateInput("2y")
	if "" != result {
		t.Error("validateInput of 1d should be ' ' but have ", result)
	}
}

func TestGetFundRankingShouldBeReturnDataMoreThanOne(t *testing.T) {
	datas := GetFundRanking("1D")
	if len(datas) <= 0 {
		t.Error("Length of datas from getFundRanking should be more than 0 but have ", len(datas))
	}
}