package util

// ZeroPad takes a byte slice and ensures it is the specified size by zero-padding leading slots.
func ZeroPad(b []byte, size int) []byte {
	l := len(b)
	if l == size {
		return b
	}
	if l > size {
		return b[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], b)
	return tmp
}
