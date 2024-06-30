package challenge_one

import "testing"

func TestHexToBase64(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
		{
			"3fbc39ae69f620cd554b",
			"P7w5rmn2IM1VSw",
		},
	}
	for _, tt := range tests {
		t.Run("Converts hex to base 64", func(t *testing.T) {
			res, err := HexToBase64(tt.input)
			if err != nil {
				t.Errorf("err %s", err)
			}
			if res != tt.want {
				t.Errorf("got %s want %s", res, tt.want)
			}
		})
	}
}

func TestHexToBase64Failure(t *testing.T) {
	s := "nonHex"
	want := "encoding/hex: invalid byte: U+006E 'n'"
	res, err := HexToBase64(s)
	if res != "" {
		t.Errorf("expected empty res, got %s", res)
	}
	if err == nil {
		t.Errorf("expected error")
	}
	if err.Error() != want {
		t.Errorf("expected %s got %s", want, err.Error())
	}
}
