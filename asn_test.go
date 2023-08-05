package asn_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn"
)

func TestASN_Uint32(t *testing.T) {
	a := asn.ASN{0, 0, 253, 232}
	n := a.Uint32()
	assert.Equal(t, uint32(65000), n)
}

func TestASN_Size(t *testing.T) {
	type caseT struct {
		size int
		asn  asn.ASN
	}
	cases := []caseT{
		{2, asn.ASN{0, 0, 255, 255}},
		{4, asn.ASN{255, 255, 255, 255}},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprint(c.size), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.size, c.asn.Size())
		})
	}
}

func TestASN_Next(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		expected := asn.ASN{0, 0, 0x38, 0xbe}
		next := a.Next()
		assert.True(t, next.Equal(expected), "next=%s, expected=%s", next.ByteString(), expected.ByteString())
	})
	t.Run("basic2", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{255, 255, 255, 254}
		e := asn.ASN{255, 255, 255, 255}
		next := a.Next()
		assert.True(t, next.Equal(e), "next=%s, expected=%s", next.ByteString(), e.ByteString())
	})
}

func TestASN_Previous(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		expected := asn.ASN{0, 0, 0x38, 0xbc}
		prev := a.Previous()
		assert.True(t, prev.Equal(expected), "prev=%s, expected=%s", prev.ByteString(), expected.ByteString())
	})
	t.Run("basic2", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{255, 255, 255, 0}
		e := asn.ASN{255, 255, 254, 255}
		next := a.Previous()
		assert.True(t, next.Equal(e), "next=%s, expected=%s", next.ByteString(), e.ByteString())
	})
}

type privateCasesT struct {
	asn  string
	want bool
}

var privateCases = []privateCasesT{
	{"65000", true},
	{"65534", true},
	{"64512", true},
	{"64600", true},
	{"14525", false},
	{"13335", false},
	{"4200000000", true},
	{"4200000005", true},
	{"4200090000", true},
	{"4294967294", true},
	{"4294967293", true},
	{"4204967293", true},
	{"4194967294", false},
	{"395077", false},
	{"4199999999", false},
}

func TestASN_IsPrivate(t *testing.T) {
	for _, c := range privateCases {
		c := c
		t.Run(c.asn, func(t *testing.T) {
			t.Parallel()
			a := asn.MustParse(c.asn)
			assert.Equal(t, c.want, a.IsPrivate())
		})
	}
}

func TestASN_IsGlobal(t *testing.T) {
	for _, c := range privateCases {
		c := c
		t.Run(c.asn, func(t *testing.T) {
			t.Parallel()
			a := asn.MustParse(c.asn)
			assert.Equal(t, !c.want, a.IsGlobal())
		})
	}
}
