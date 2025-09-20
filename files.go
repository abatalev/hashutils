package hashutils

import (
	"path/filepath"
)

func CalcHashOfFiles(workDir string, files []string) string {
	return calcHashFiles(calcHashes(workDir, files))
}

func calcHashFiles(files []string) string {
	s := ""
	for _, file := range files {
		s += file + "\n"
	}
	return CalcHashBytes([]byte(s))
}

func calcHashes(workDir string, files []string) []string {
	filesWithHashes := make([]string, 0)
	for _, f := range files {
		filesWithHashes = appendFileAndHash(filesWithHashes, f, CalcHashFile(filepath.Join(workDir, f)))
	}
	return filesWithHashes
}

func appendFileAndHash(filesWithHashes []string, f, hash string) []string {
	return append(filesWithHashes, f+" "+hash)
}
