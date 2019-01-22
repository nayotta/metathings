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
	MAJOR_SESSION_FLAG             = 0x40 << 24
	MINOR_SESSION_FLAG             = 0x60 << 24
	TEMP_SESSION_FLAG              = 0x10 << 24
	STARTUP_SESSION_MASK           = 0x7fffffff00000000
	CONNECTION_SESSION_MASK        = 0x000000007fffffff
	x_CONNECTION_SESSION_DATA_MASK = 0x000000000fffffff
	MAJOR_SESSION_MASK             = MAJOR_SESSION_FLAG | x_CONNECTION_SESSION_DATA_MASK
	MINOR_SESSION_MASK             = MINOR_SESSION_FLAG | x_CONNECTION_SESSION_DATA_MASK
)

func GenerateStartupSession() int32 {
	return rand_helper.Int31()
}

func GenerateMajorSession() int32 {
	return rand_helper.Int31() & MAJOR_SESSION_MASK
}

func GenerateMinorSession() int32 {
	return rand_helper.Int31() & MINOR_SESSION_MASK
}

func GenerateTempSession() int32 {
	return GenerateMinorSession() & TEMP_SESSION_FLAG
}

func IsMajorSession(s int64) bool {
	return GetConnectionSession(s)&MAJOR_SESSION_FLAG == MAJOR_SESSION_FLAG
}

func IsMinorSession(s int64) bool {
	return GetConnectionSession(s)&MINOR_SESSION_FLAG == MINOR_SESSION_FLAG
}

func IsTempSession(s int64) bool {
	return GetConnectionSession(s)&TEMP_SESSION_FLAG == TEMP_SESSION_FLAG
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
