package session_helper

import (
	"time"

	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
)

var (
	STARTUP_SESSION_EXPIRE = 120 * time.Second
)

// session protocol:
// session contains startup session and connection session(major session or minor session).
// startup session is device start up and pick a random number to become startup session.
// connection session is when create connection to pick a random number to become connection session.
// connection session dive into two type session: major session and minor session.
// first byte is flag byte and rest three bytes is session data bytes.
// 0x[ flags 1 byte ] [ data 3 bytes ]
// session flags:
// [ 0 reserverd ] [ 1 session code ] [ x major/minor ] [ y temp ] [ 0000 reserverd ]
//                                      0 = major
//                                      1 = minor
// major session design for major connection.
// minor session design for minor connection.

const (
	x_CONNECTION_SESSION_CODE_MASK = int32(0x40 << 24) // 0100 0000
	x_CONNECTION_SESSION_TYPE_MASK = int32(0x20 << 24) // 0010 0000
	x_CONNECTION_SESSION_TEMP_MASK = int32(0x10 << 24) // 0001 0000

	MAJOR_SESSION_FLAG             = int32(0x00)       // 0000 0000
	MINOR_SESSION_FLAG             = int32(0x20 << 24) // 0010 0000
	TEMP_SESSION_FLAG              = int32(0x10 << 24) // 0001 0000
	STARTUP_SESSION_MASK           = int64(0x7fffffff00000000)
	CONNECTION_SESSION_MASK        = int64(0x000000007fffffff)
	x_CONNECTION_SESSION_DATA_MASK = int64(0x0000000000ffffff)
)

func GenerateStartupSession() int32 {
	return rand_helper.Int31()
}

func generateConnectionSessionData() int32 {
	return rand_helper.Int31() & int32(x_CONNECTION_SESSION_DATA_MASK)
}

func GenerateMajorSession() int32 {
	return generateConnectionSessionData() | MAJOR_SESSION_FLAG
}

func GenerateMinorSession() int32 {
	return generateConnectionSessionData() | MINOR_SESSION_FLAG
}

func GenerateTempSession() int32 {
	return GenerateMinorSession() | TEMP_SESSION_FLAG
}

func isConnectionSession(cs int32) bool {
	return cs&x_CONNECTION_SESSION_CODE_MASK == x_CONNECTION_SESSION_CODE_MASK
}

func isMajorSession(cs int32) bool {
	return isConnectionSession(cs) && cs&x_CONNECTION_SESSION_TYPE_MASK == MAJOR_SESSION_FLAG
}

func IsMajorSession(s int64) bool {
	cs := GetConnectionSession(s)
	return isMajorSession(cs)
}

func isMinorSession(cs int32) bool {
	return isConnectionSession(cs) && cs&x_CONNECTION_SESSION_TYPE_MASK == MINOR_SESSION_FLAG
}

func IsMinorSession(s int64) bool {
	cs := GetConnectionSession(s)
	return isMinorSession(cs)
}

func IsTempSession(s int64) bool {
	cs := GetConnectionSession(s)
	return isMinorSession(cs) && cs&x_CONNECTION_SESSION_TEMP_MASK == TEMP_SESSION_FLAG
}

func GetStartupSession(s int64) int32 {
	return int32((s & STARTUP_SESSION_MASK) >> 32)
}

func GetConnectionSession(s int64) int32 {
	return int32((s & CONNECTION_SESSION_MASK))
}

func NewSession(startup, conn int32) int64 {
	return int64(startup)<<32 + int64(conn)
}
