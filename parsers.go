package asn

import (
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
)

// MustDecimal parses an ASN in decimal/asplain format and panics if the input is invalid.
func MustDecimal(d string) ASN {
	a, err := FromDecimal(d)
	if err != nil {
		panic(err)
	}
	return a
}

// MustASDot parses an ASN in asdot format and panics if the input is invalid.
func MustASDot(d string) ASN {
	a, err := FromASDot(d)
	if err != nil {
		panic(err)
	}
	return a
}

// MustParse parses an ASN in any valid format and panics if the input is invalid.
func MustParse(d string) ASN {
	a, err := Parse(d)
	if err != nil {
		panic(err)
	}
	return a
}

// FromDecimal parses an ASN in decimal/asplain format and returns an error if the input is invalid.
func FromDecimal(d string) (ASN, error) {
	n64, err := strconv.ParseUint(d, 10, 32)
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			return nil, ErrOutOf4ByteRange
		}
		return nil, ErrInvalidInput
	}
	n := uint32(n64)
	a := make(ASN, BYTE_SIZE)
	binary.BigEndian.PutUint32(a, n)
	return a, nil
}

// FromASDot parses an ASN in asdot format and returns an error if the input is invalid.
func FromASDot(i string) (ASN, error) {
	parts := strings.Split(strings.TrimSpace(i), ".")

	if len(parts) != 2 {
		return nil, ErrInvalidInput
	}

	asn := make(ASN, 4)

	high, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, ErrInvalidInput
	}

	low, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, ErrInvalidInput
	}

	if high < 0 || high > 65535 || low < 0 || low > 65535 {
		return nil, ErrOutOf2ByteRange
	}

	asn[0] = byte(high >> 8)
	asn[1] = byte(high)
	asn[2] = byte(low >> 8)
	asn[3] = byte(low)

	return asn, nil
}

// FromUint32 parses an ASN from an unsigned 32-bit integer.
func FromUint32(n uint32) ASN {
	a := make(ASN, BYTE_SIZE)
	binary.BigEndian.PutUint32(a, n)
	return a
}

// FromUint64 parses an ASN from an unsigned 64-bit integer and returns an error if the value is
// greater than 32 bits.
func FromUint64(n uint64) (ASN, error) {
	if n > uint64(MAX_32) {
		return nil, ErrOutOf4ByteRange
	}
	a := make(ASN, BYTE_SIZE)
	binary.BigEndian.PutUint32(a, uint32(n))
	return a, nil
}

// From4Bytes creates an ASN object from 4 bytes.
func From4Bytes(one, two, three, four byte) ASN {
	return ASN{one, two, three, four}
}

// From2Bytes creates an ASN object from 2 bytes.
func From2Bytes(one, two byte) ASN {
	return ASN{0, 0, one, two}
}

// FromBytes creates an ASN object from either 2 or 4 bytes. An error is returned if the number of
// bytes provided is not 2 or 4.
func FromBytes(bytes ...byte) (ASN, error) {
	if len(bytes) != 2 && len(bytes) != 4 {
		return nil, ErrOutOf4ByteRange
	}
	if len(bytes) == 2 {
		return From2Bytes(bytes[0], bytes[1]), nil
	}
	return From4Bytes(bytes[0], bytes[1], bytes[2], bytes[3]), nil
}

// Parse parses and validates an ASN from an input string. An error is returned if the input is
// invalid.
func Parse(in string) (ASN, error) {
	if asdotPattern.MatchString(in) {
		return FromASDot(in)
	}
	if asDecimalPattern.MatchString(in) {
		return FromDecimal(in)
	}
	return nil, ErrInvalidInput
}
