package consistent_hash

import (
	"hash/crc32"
)

func buildHashId(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}