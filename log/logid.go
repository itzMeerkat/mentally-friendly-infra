package log

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// TODO: optimization could be done
func GetLogID(ip string) string {
	ip = strings.ReplaceAll(ip, ".", "")
	buf := make([]byte, 0, 64)
	buf = time.Now().AppendFormat(buf, "20060102150405")
	buf = append(buf, ip...)

	uuidBuf := make([]byte, 4)
	_, err := rand.Read(uuidBuf)
	if err != nil {
		panic(err)
	}
	uuidNum := binary.BigEndian.Uint32(uuidBuf)
	buf = append(buf, fmt.Sprintf("%05d", uuidNum)[:5]...)
	return string(buf)
}
