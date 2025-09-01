package api

import (
	"crypto/md5"
	"fmt"
	"io"
)

func (s *Sdk) GenerateMD5Sign(dataStr, timestamp string) string {
	strTonHash := s.config.ApiSecret + dataStr + timestamp
	hash := md5.New()

	io.WriteString(hash, strTonHash)
	hashInBytes := hash.Sum(nil)
	hashInString := fmt.Sprintf("%x", hashInBytes)

	return hashInString
}
