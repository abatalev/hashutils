# hashutils

Go utility library for computing SHA-1 hashes of files and
byte data.

## Installation

```text
go get github.com/abatalev/hashutils
```

## Usage

```go
import "github.com/abatalev/hashutils"
```

### Hash bytes

```go
hash := hashutils.CalcHashBytes([]byte("hello"))
// "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
```

### Hash a single file

```go
hash, err := hashutils.CalcHashFile("/path/to/file")
```

### Hash multiple files

```go
hash := hashutils.CalcHashFiles([]hashutils.FileInfo{
    {FileName: "a.txt", Hash: "hash1"},
    {FileName: "b.txt", Hash: "hash2"},
})
```

### Hash files in a directory

```go
hash, err := hashutils.CalcHashOfFiles("/work/dir", []string{"a.txt", "b.txt"})
```

### Walk directory with glob patterns

```go
files, err := hashutils.WalkDirWithPatterns("/work/dir", []string{"*.go", "**/*.md"})
```

### Sort files by name

```go
hashutils.SortFiles(files)
```

## API

| Function | Description |
|---|---|
| `CalcHashBytes(buf []byte) string` | SHA-1 hash of raw bytes |
| `CalcHashFile(path string) (string, error)` | SHA-1 hash of a file |
| `CalcHashFiles(fileList []FileInfo) string` | Composite hash of a file list |
| `CalcHashOfFiles(workDir string, files []string) (string, error)` | Hash multiple files by path |
| `WalkDirWithPatterns(workDir string, patterns []string) ([]string, error)` | Walk dir matching glob patterns |
| `SortFiles(files []FileInfo)` | Sort files by name (hash as tiebreaker) |
