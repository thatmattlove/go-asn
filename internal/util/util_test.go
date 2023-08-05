package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/go-asn/internal/util"
)

func Test_ZeroPad(t *testing.T) {
	t.Run("2-4", func(t *testing.T) {
		t.Parallel()
		in := []byte{1, 2}
		expected := []byte{0, 0, 1, 2}
		result := util.ZeroPad(in, len(expected))
		assert.Equal(t, expected, result)
	})
	t.Run("4-4", func(t *testing.T) {
		t.Parallel()
		in := []byte{1, 2, 3, 4}
		expected := in
		result := util.ZeroPad(in, len(expected))
		assert.Equal(t, expected, result)
	})

	t.Run("4-2", func(t *testing.T) {
		t.Parallel()
		in := []byte{1, 2, 3, 4}
		expected := []byte{3, 4}
		result := util.ZeroPad(in, 2)
		assert.Equal(t, expected, result)
	})
}
