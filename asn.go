package asn

// ASN represents a single autonomous system number, a slice of bytes. Both 2-byte (16-bit) and
// 4-byte (32-bit) ASNs are supported.
type ASN []byte

// Size returns either 2 or 4, depending on if the ASN is 2-bytes or 4-bytes.
func (asn ASN) Size() int {
	if asn[0] == 0 && asn[1] == 0 {
		return 2
	}
	return 4
}

// Next returns the next ASN after the current ASN. For example, if the current ASN is 65000,
// Next() would return 65001.
func (asn ASN) Next() ASN {
	next := make(ASN, BYTE_SIZE)
	copy(next, asn)
	for i := len(next) - 1; i >= 0; i-- {
		if next[i] < 255 {
			next[i]++
			break
		} else {
			next[i]--
		}
	}
	return next
}

// Previous returns the previous ASN before the current ASN. For example, if the current ASN is 65001,
// Previous() would return 65000.
func (asn ASN) Previous() ASN {
	prev := asn
	for i := len(prev) - 1; i >= 0; i-- {
		if prev[i] > 0 {
			prev[i]--
			break
		} else {
			prev[i] = 255
		}
	}
	return prev
}

// IsGlobal returns true if the ASN is global, i.e. not private.
// See RFC6996.
func (asn ASN) IsGlobal() bool {
	return !asn.IsPrivate()
}

// IsGlobal returns true if the ASN is private, i.e. not global.
// See RFC6996.
func (asn ASN) IsPrivate() bool {
	n := asn.Uint32()
	if asn.Size() == 2 {
		return n >= 64512 && n <= 65534
	}
	return n >= 4200000000 && n <= 4294967294
}
