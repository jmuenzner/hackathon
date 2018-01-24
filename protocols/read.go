package protocols

// Read returns the protocol matching the given resource name, or an error
func Read(resourceName string) ([]byte, error) {
	return Asset(resourceName + ".proto")
}

// MustRead returns the protocol matching the given resource name, or panics
func MustRead(resourceName string) []byte {
	return MustAsset(resourceName + ".proto")
}
