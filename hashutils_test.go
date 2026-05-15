package hashutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcHashBytes(t *testing.T) {
	require.Equal(t, "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", CalcHashBytes([]byte("a")))
}

func TestCalcHashBytesEmpty(t *testing.T) {
	require.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", CalcHashBytes([]byte{}))
}

func TestCalcHashFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	require.NoError(t, os.WriteFile(path, []byte("hello"), 0600))

	hash, err := CalcHashFile(path)
	require.NoError(t, err)
	require.Equal(t, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", hash)
}

func TestCalcHashFileNotExists(t *testing.T) {
	_, err := CalcHashFile("/nonexistent/file")
	require.Error(t, err)
}

func TestCalcHashFiles(t *testing.T) {
	data := []FileInfo{{FileName: "a", Hash: "b"}}
	require.Equal(t, "90ce62edf2fe4940e041a68b13e7b5f9d02bbf51", CalcHashFiles(data))
}

func TestCalcHashFilesEmpty(t *testing.T) {
	require.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", CalcHashFiles([]FileInfo{}))
}

func TestSortFiles(t *testing.T) {
	files := []FileInfo{{FileName: "b", Hash: "1"}, {FileName: "a", Hash: "2"}}
	SortFiles(files)
	require.Len(t, files, 2)
	require.Equal(t, "a", files[0].FileName)
	require.Equal(t, "2", files[0].Hash)
	require.Equal(t, "b", files[1].FileName)
	require.Equal(t, "1", files[1].Hash)
}

func TestSortFilesSameName(t *testing.T) {
	files := []FileInfo{{FileName: "a", Hash: "2"}, {FileName: "a", Hash: "1"}}
	SortFiles(files)
	require.Equal(t, "a", files[0].FileName)
	require.Equal(t, "1", files[0].Hash)
}
