package asn

// Equal determines if this ASN is equal to the input ASN.
func (asn ASN) Equal(other ASN) bool {
	if len(asn) != len(other) {
		return false
	}
	for i := range asn {
		if asn[i] != other[i] {
			return false
		}
	}
	return true
}

// GreaterThan determines if this ASN is greater than (higher number) the input ASN.
func (asn ASN) GreaterThan(other ASN) bool {
	return asn.Uint32() > other.Uint32()
}

// LessThan determines if this ASN is less than (lower number) the input ASN.
func (asn ASN) LessThan(other ASN) bool {
	return asn.Uint32() < other.Uint32()
}

// GEqual determines if this ASN is greater than (higher number) or equal to the input ASN.
func (asn ASN) GEqual(other ASN) bool {
	return asn.GreaterThan(other) || asn.Equal(other)
}

// LEqual determines if this ASN is less than (lower number) or equal to the input ASN.
func (asn ASN) LEqual(other ASN) bool {
	return asn.LessThan(other) || asn.Equal(other)
}
