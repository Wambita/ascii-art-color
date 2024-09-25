package asciiart

func IsAltered(checksum uint32) bool {
	const (
		standardChecksum   = 0x9ffd59bc
		shadowChecksum     = 0x2f465361
		thinkertoyChecksum = 0x6ee86a07
		testFileChecksum   = 0x597778CF
	)

	if !(checksum == standardChecksum || checksum == shadowChecksum || checksum == thinkertoyChecksum || checksum == testFileChecksum) {
		return true
	}
	return false
}
