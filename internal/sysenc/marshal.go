index a6a50e9..a3bca95 100644
 	return newBuffer(buf), nil
 }
 
// optimizedUnmarshal handles common types directly without using reflection
// to improve performance for frequently used types.
func optimizedUnmarshal(data interface{}, buf []byte) (bool, error) {
	switch value := data.(type) {
	case *int32:
		if len(buf) == 4 {
			*value = int32(internal.NativeEndian.Uint32(buf))
			return true, nil
		}
	case *uint32:
		if len(buf) == 4 {
			*value = internal.NativeEndian.Uint32(buf)
			return true, nil
		}
	}
	return false, nil
}

 var bytesReaderPool = sync.Pool{
 	New: func() interface{} {
 		return new(bytes.Reader)
 // Returns an error if buf can't be unmarshalled according to the behaviour
 // of [binary.Read].
 func Unmarshal(data interface{}, buf []byte) error {
	// Try optimized path first for common types
	if handled, err := optimizedUnmarshal(data, buf); handled {
		return err
	}
	
	// Fall back to standard handling
 	switch value := data.(type) {
 	case encoding.BinaryUnmarshaler:
 		return value.UnmarshalBinary(buf)
