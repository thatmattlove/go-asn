package asn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn"
)

func TestASN_Equal(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		one := asn.ASN{0, 0, 0x38, 0xbd}
		two := asn.ASN{0, 0, 0x38, 0xbd}
		assert.True(t, one.Equal(two))
	})
	t.Run("not equal", func(t *testing.T) {
		t.Parallel()
		one := asn.ASN{0, 0, 0x38, 0xbd}
		two := asn.ASN{1, 2, 3, 4}
		assert.False(t, one.Equal(two))
	})
	t.Run("mismatching lengths", func(t *testing.T) {
		one := asn.ASN{1, 2, 3, 4}
		two := asn.ASN{1, 2}
		assert.False(t, one.Equal(two))
	})
	t.Run("not equal with next", func(t *testing.T) {
		t.Parallel()
		one := asn.ASN{255, 255, 255, 255}
		next := one.Next()
		expected := asn.ASN{254, 255, 255, 255}
		assert.False(t, next.Equal(expected))
	})
	t.Run("not equal with previous", func(t *testing.T) {
		t.Parallel()
		one := asn.ASN{255, 255, 255, 255}
		prev := one.Previous()
		expected := asn.ASN{254, 255, 255, 255}
		assert.False(t, prev.Equal(expected))
	})
}
func TestASN_GreaterThan(t *testing.T) {
	t.Run("greater", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbc}
		assert.True(t, a.GreaterThan(o))
	})

	t.Run("less", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbe}
		assert.False(t, a.GreaterThan(o))
	})
}

func TestASN_LessThan(t *testing.T) {
	t.Run("less", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbe}
		assert.True(t, a.LessThan(o))
	})

	t.Run("greater", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbc}
		assert.False(t, a.LessThan(o))
	})
}

func TestASN_GEqual(t *testing.T) {
	t.Run("less", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbc}
		assert.True(t, a.GEqual(o))
	})

	t.Run("greater", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbe}
		assert.False(t, a.GEqual(o))
	})

	t.Run("equal", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbd}
		assert.True(t, a.GEqual(o))
	})
}

func TestASN_LEqual(t *testing.T) {
	t.Run("less", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbe}
		assert.True(t, a.LEqual(o))
	})

	t.Run("greater", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbc}
		assert.False(t, a.LEqual(o))
	})

	t.Run("equal", func(t *testing.T) {
		t.Parallel()
		a := asn.ASN{0, 0, 0x38, 0xbd}
		o := asn.ASN{0, 0, 0x38, 0xbd}
		assert.True(t, a.LEqual(o))
	})
}
