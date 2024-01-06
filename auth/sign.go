package auth

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	timeLen = 8
	hashLen = 32
)

var sigTtl = time.Millisecond * 200

func init() {
	sigTtlStr := os.Getenv("SIG_TTL")
	if sigTtlStr != "" {
		var err error
		sigTtl, err = time.ParseDuration(sigTtlStr)
		if err != nil {
			panic("sig ttl parse err: " + err.Error())
		}
	}
}

func Sign(secret, payload []byte, t time.Time) ([]byte, error) {
	timeBytes, err := convertTimeToBytes(t)
	if err != nil {
		return nil, fmt.Errorf("convert time to bytes err: %w", err)
	}
	// concat secret, timeBytes & payload
	data := append(secret, timeBytes...)
	data = append(data, payload...)
	h := sha256.Sum256(data)
	sum := h[:hashLen]
	// return sum + time + payload
	signed := make([]byte, 0, hashLen+timeLen+len(payload))
	signed = append(signed, sum...)
	signed = append(signed, timeBytes...)
	return append(signed, payload...), nil
}

func Verify(secret, sum, timeBytes, payload []byte) (bool, error) {
	t, err := convertBytesToTime(timeBytes)
	if err != nil {
		return false, fmt.Errorf("convert bytes to time err: %w", err)
	}
	if t.After(time.Now().Add(sigTtl)) ||
		t.Before(time.Now().Add(-sigTtl)) {
		return false, errors.New("time is outside of threshold")
	}
	data := append(secret, timeBytes...)
	data = append(data, payload...)
	h := sha256.Sum256(data)
	mySum := h[:hashLen]
	return bytes.Equal(sum, mySum), nil
}
