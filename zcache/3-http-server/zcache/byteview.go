package zcache

type ByteView struct {
	b []byte
}

// Len  returns the length of the byte array
func (bv ByteView) Len() int {
	return len(bv.b)
}

func cloneBytes(b []byte) []byte {
	if b == nil {
		return nil
	}
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// ByteSlice 只读 返回复制对象
func (bv ByteView) ByteSlice() []byte {
	return cloneBytes(bv.b)
}

func (bv ByteView) String() string {
	return string(bv.b)
}
