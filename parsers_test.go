package asn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn"
)

func Test_MustDecimal(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		assert.NotPanics(t, func() {
			asn.MustDecimal("65000")
		})
	})
	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			asn.MustDecimal("this will panic")
		})
	})
}

func Test_MustASDot(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		assert.NotPanics(t, func() {
			asn.MustASDot("0.65000")
		})
	})
	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			asn.MustASDot("this will panic")
		})
	})
}

func Test_MustParse(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		assert.NotPanics(t, func() {
			asn.MustParse("4200000000")
		})
	})
	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			asn.MustParse("this will panic")
		})
	})
}

func Test_FromDecimal(t *testing.T) {
	type caseT struct {
		from string
		want string
	}
	cases := []caseT{
		{"65000", "65000"},
		{"65546", "65546"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.want, func(t *testing.T) {
			t.Parallel()
			a, err := asn.FromDecimal(c.from)
			assert.NoError(t, err)
			assert.Equal(t, c.want, a.Decimal())
		})
	}
	t.Run("error with asdot", func(t *testing.T) {
		t.Parallel()
		_, err := asn.FromDecimal("6.1861")
		assert.ErrorIs(t, err, asn.ErrInvalidInput)
	})
	t.Run("error with number out of range", func(t *testing.T) {
		t.Parallel()
		_, err := asn.FromDecimal("5294967295")
		assert.ErrorIs(t, err, asn.ErrOutOf4ByteRange)
	})
}

func Test_FromASDot(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		asd := "6.1861"
		a, err := asn.FromASDot(asd)
		assert.NoError(t, err)
		assert.Equal(t, asd, a.ASDot())
	})

	type errorCasesT struct {
		from string
		want error
	}

	errorCases := []errorCasesT{
		{"65000", asn.ErrInvalidInput},
		{"65536.0", asn.ErrOutOf2ByteRange},
		{"0.65536", asn.ErrOutOf2ByteRange},
		{"thiswillfail", asn.ErrInvalidInput},
		{"0.thiswillfail", asn.ErrInvalidInput},
		{"thiswillfail.0", asn.ErrInvalidInput},
	}

	for _, c := range errorCases {
		c := c
		t.Run(c.from, func(t *testing.T) {
			t.Parallel()
			_, err := asn.FromASDot(c.from)
			assert.ErrorIs(t, err, c.want)
		})
	}
}

func Test_Parse(t *testing.T) {
	type caseT struct {
		from string
		want string
	}
	casesDecimal := []caseT{
		{"6.1861", "395077"},
		{"395077", "395077"},
		{"65000", "65000"},
	}
	for _, c := range casesDecimal {
		c := c
		t.Run(c.want, func(t *testing.T) {
			t.Parallel()
			a, err := asn.Parse(c.from)
			assert.NoError(t, err)
			assert.Equal(t, c.want, a.Decimal())
		})
	}
	t.Run("errors", func(t *testing.T) {
		t.Parallel()
		_, err := asn.Parse("default error")
		assert.ErrorIs(t, err, asn.ErrInvalidInput)
	})
}

func Test_FromUint32(t *testing.T) {
	e := asn.ASN{0, 0, 253, 232}
	a := asn.FromUint32(65000)
	assert.True(t, a.Equal(e), "parsed=%s", a.ByteString())
}

func Test_FromUint64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		e := asn.ASN{0, 0, 253, 232}
		a, err := asn.FromUint64(65000)
		assert.NoError(t, err)
		assert.True(t, a.Equal(e), "parsed=%s", a.ByteString())
	})
	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		_, err := asn.FromUint64(5294967295)
		assert.ErrorIs(t, err, asn.ErrOutOf4ByteRange)
	})
}

func Test_From2Bytes(t *testing.T) {
	expected := asn.ASN{0, 0, 253, 232}
	result := asn.From2Bytes(253, 232)
	assert.True(t, result.Equal(expected))
}

func Test_From4Bytes(t *testing.T) {
	expected := asn.ASN{255, 255, 253, 232}
	result := asn.From4Bytes(255, 255, 253, 232)
	assert.True(t, result.Equal(expected))
}

func Test_FromBytes(t *testing.T) {
	t.Run("2B", func(t *testing.T) {
		t.Parallel()
		expected := asn.ASN{0, 0, 253, 232}
		result, err := asn.FromBytes(253, 232)
		assert.NoError(t, err)
		assert.True(t, result.Equal(expected))
	})
	t.Run("4B", func(t *testing.T) {
		t.Parallel()
		expected := asn.ASN{255, 255, 253, 232}
		result, err := asn.FromBytes(255, 255, 253, 232)
		assert.NoError(t, err)
		assert.True(t, result.Equal(expected))
	})
	t.Run("too many bytes", func(t *testing.T) {
		t.Parallel()
		_, err := asn.FromBytes(255, 255, 253, 232, 255, 252)
		assert.ErrorIs(t, err, asn.ErrOutOf4ByteRange)
	})
	t.Run("too few bytes", func(t *testing.T) {
		t.Parallel()
		_, err := asn.FromBytes(255)
		assert.ErrorIs(t, err, asn.ErrOutOf4ByteRange)
	})
}
