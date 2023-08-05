package asn

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/thatmattlove/go-asn/internal/util"
)

// Decimal formats the ASN in asplain format. See RFC5396.
func (asn ASN) Decimal() string {
	d := binary.BigEndian.Uint32(asn)
	return fmt.Sprint(d)
}

// ASPlain formats the ASN in asplain format. See RFC5396.
func (asn ASN) ASPlain() string {
	return asn.Decimal()
}

// ASDotPlus formats the ASN in asdot+ format. See RFC5396.
func (asn ASN) ASDotPlus() string {
	high := binary.BigEndian.Uint32(util.ZeroPad(asn[0:2], BYTE_SIZE))
	low := binary.BigEndian.Uint32(util.ZeroPad(asn[2:4], BYTE_SIZE))
	return fmt.Sprintf("%d.%d", high, low)
}

// ASDot formats the ASN in asdot format. See RFC5396.
func (asn ASN) ASDot() string {
	if asn.Size() == 2 {
		return asn.Decimal()
	}
	return asn.ASDotPlus()
}

// String formats the ASN in decimal/asplain format. See RFC5396.
func (asn ASN) String() string {
	return asn.Decimal()
}

// Uint32 returns the ASN as a 32-bit unsigned integer.
func (asn ASN) Uint32() uint32 {
	return binary.BigEndian.Uint32(asn)
}

// ByteString returns a string representation of each ASN byte.
func (asn ASN) ByteString() string {
	bs := []string{}
	for _, b := range asn {
		bs = append(bs, fmt.Sprint(b))
	}
	return fmt.Sprintf("{%s}", strings.Join(bs, ","))
}
