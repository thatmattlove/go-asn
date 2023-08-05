package asn

type iterRange struct {
	start ASN
	stop  ASN
	value ASN
}

type iter struct {
	start ASN
	value ASN
}

// Continue returns true if the current iteration value is less than the high end of the range.
func (i *iterRange) Continue() bool {
	return i.value.LessThan(i.stop)
}

// Next returns the next ASN in the range.
func (i *iterRange) Next() ASN {
	i.value = i.value.Next()
	return i.value
}

// Range returns an iterator object that can be used to iterate through a range of ASNs starting
// with the current object and ending with the input ASN.
func (asn ASN) Range(high ASN) iterRange {
	return iterRange{
		start: asn,
		stop:  high,
		value: asn,
	}
}

// Continue returns true if the current iteration value is less than the highest possible ASN.
func (i *iter) Continue() bool {
	return i.value.LessThan(ASN{255, 255, 255, 255})
}

// Next returns the next ASN.
func (i *iter) Next() ASN {
	next := i.value.Next()
	i.value = next
	return next
}

// Iter returns an iterator object that can be used to iterate through ASNs, starting from the
// current ASN and ending with the highest possible ASN.
func (asn ASN) Iter() iter {
	return iter{
		start: asn,
		value: asn,
	}
}
