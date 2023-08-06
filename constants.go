package asn

import "regexp"

const (
	// BYTE_SIZE is the size of all ASN objects in bytes.
	BYTE_SIZE int = 4
	// MAX_32 is the maximum value allowed for a 32-bit (4-byte) ASN.
	MAX_32 uint32 = 4294967295
)

var (
	asdotPattern     = regexp.MustCompile(`^(AS|as)?(\d+\.\d+)$`)
	asDecimalPattern = regexp.MustCompile(`^(AS|as)?(\d+)$`)
)
