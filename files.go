package hashutils

import (
	"path/filepath"
	"strings"
)

func CalcHashOfFiles(workDir string, files []string) (string, error) {
	hashed, err := calcHashes(workDir, files)
	if err != nil {
		return "", err
	}
	return calcHashFiles(hashed), nil
}

func calcHashFiles(files []string) string {
	var b strings.Builder
	for _, file := range files {
		b.WriteString(file)
		b.WriteString("\n")
	}
	return CalcHashBytes([]byte(b.String()))
}

func calcHashes(workDir string, files []string) ([]string, error) {
	filesWithHashes := make([]string, 0, len(files))
	for _, f := range files {
		hash, err := CalcHashFile(filepath.Join(workDir, f))
		if err != nil {
			return nil, err
		}
		filesWithHashes = append(filesWithHashes, f+" "+hash)
	}
	return filesWithHashes, nil
}
