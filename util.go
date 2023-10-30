package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func randomString(n int) (string, error) {
	const letters = "0123456789"
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random string: %v", err)
		}
		res[i] = letters[num.Int64()]
	}

	return string(res), nil
}

// formatPath ensures path start with os.PathSeparator and end with no os.PathSeparator.
// And replaces all '/' with os.PathSeparator.
func formatPath(path string) string {
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))

	path = strings.TrimSuffix(path, string(os.PathSeparator))
	if !strings.HasPrefix(path, string(os.PathSeparator)) {
		path = fmt.Sprintf("%v%v", string(os.PathSeparator), path)
	}

	return path
}