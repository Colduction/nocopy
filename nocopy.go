package nocopy

import "unsafe"

// Converts b string to an immutable string
func ByteSliceToString(b []byte) string {
	length := len(b)
	if length <= 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), length)
}

// Converts s string to an immutable slice of bytes
func StringToByteSlice(s string) []byte {
	length := len(s)
	if length <= 0 {
		return unsafe.Slice(unsafe.StringData(""), 0)
	}
	return unsafe.Slice(unsafe.StringData(s), length)
}
