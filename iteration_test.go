package asn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn"
)

func TestASN_Range(t *testing.T) {
	start := 1
	first := asn.ASN{0, 0, 1, byte(start)}
	last := asn.ASN{0, 0, 1, 5}
	var final asn.ASN
	for iter := first.Range(last); iter.Continue(); {
		next := iter.Next()
		start++
		expected := asn.ASN{0, 0, 1, byte(start)}
		final = expected
		if !next.Equal(expected) {
			break
		}
	}
	assert.True(t, final.Equal(last), final.ByteString())
}

func TestASN_Iter(t *testing.T) {
	start := 253
	first := asn.ASN{255, 255, 255, byte(start)}
	last := asn.ASN{255, 255, 255, 255}
	var final asn.ASN
	for iter := first.Iter(); iter.Continue(); {
		next := iter.Next()
		start++
		expected := asn.ASN{255, 255, 255, byte(start)}
		final = expected
		if !next.Equal(expected) {
			t.Logf("first=%s, next=%s, exp=%s", first.ByteString(), next.ByteString(), expected.ByteString())
			break
		}
	}
	assert.True(t, final.Equal(last), "final=%s, last=%s", final.ByteString(), last.ByteString())
}
