// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"crypto/rand"
	"fmt"
	"io"
)

// uuid generates new random subscription UUID
func uuid() (string, error) {
	// prep container
	uuid := make([]byte, 16)

	// try to read random numbers
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random)
	uuid[6] = uuid[6]&^0xf0 | 0x40

	// format
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
