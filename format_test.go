package asn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn"
)

type decimalFmtT struct {
	from asn.ASN
	want string
}

var decimalCases = []decimalFmtT{
	{asn.ASN{0, 0, 0, 1}, "1"},
	{asn.ASN{0, 0, 56, 189}, "14525"},
}

func TestASN_Decimal(t *testing.T) {
	for _, c := range decimalCases {
		c := c
		t.Run(c.want, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.want, c.from.Decimal())
		})
	}
}

func TestASN_ASPlain(t *testing.T) {
	for _, c := range decimalCases {
		c := c
		t.Run(c.want, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.want, c.from.ASPlain())
		})
	}
}

func TestASN_String(t *testing.T) {
	for _, c := range decimalCases {
		c := c
		t.Run(c.want, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.want, c.from.String())
		})
	}
}

func TestASN_ASDot(t *testing.T) {
	t.Run("4B", func(t *testing.T) {
		t.Parallel()
		a := asn.MustDecimal("4200000000")
		assert.Equal(t, "64086.59904", a.ASDot())
	})
	t.Run("2B", func(t *testing.T) {
		t.Parallel()
		a := asn.MustDecimal("65000")
		assert.Equal(t, "65000", a.ASDot())
	})
}

func TestASN_ASDotPlus(t *testing.T) {
	t.Run("2B", func(t *testing.T) {
		t.Parallel()
		a := asn.MustParse("65000")
		assert.Equal(t, "0.65000", a.ASDotPlus())
	})
	t.Run("4B", func(t *testing.T) {
		t.Parallel()
		a := asn.MustParse("4200000000")
		assert.Equal(t, "64086.59904", a.ASDotPlus())
	})
}

func TestASN_ByteString(t *testing.T) {
	asn := asn.MustParse("65000")
	expected := "{0,0,253,232}"
	assert.Equal(t, expected, asn.ByteString())
}
