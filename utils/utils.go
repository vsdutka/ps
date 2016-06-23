// utils
package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	mrnd "math/rand"
	"net/http"
	"time"
)

var letters = []rune("1234567890")

func init() {
	mrnd.Seed(time.Now().UTC().UnixNano())
}
func Pin() string {
	p := mrnd.Intn(9999-1000) + 1000
	return fmt.Sprintf("%d", p)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
)

func SecureRandomAlphaString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func JsonOK(w http.ResponseWriter) {
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	fmt.Fprint(w, "{\"status\": \"OK\"}")
}

func JsonError(w http.ResponseWriter, err error) {
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\": \"error\", \"message\": \"%s\"}", err.Error())
}

func JsonData(w http.ResponseWriter, data interface{}) {
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	j, err := json.Marshal(data)
	if err != nil {
		JsonError(w, err)
		return
	}
	fmt.Fprint(w, string(j))
}
