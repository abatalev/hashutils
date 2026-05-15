package hashutils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"sort"
	"strings"
)

type FileInfo struct {
	FileName string
	Hash     string
}

func CalcHashBytes(buf []byte) string {
	h := sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func CalcHashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func CalcHashFiles(fileList []FileInfo) string {
	var b strings.Builder
	for _, file := range fileList {
		b.WriteString(file.FileName)
		b.WriteString(" ")
		b.WriteString(file.Hash)
		b.WriteString("\n")
	}
	return CalcHashBytes([]byte(b.String()))
}

func SortFiles(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		if files[i].FileName == files[j].FileName {
			return files[i].Hash < files[j].Hash
		}
		return files[i].FileName < files[j].FileName
	})
}
