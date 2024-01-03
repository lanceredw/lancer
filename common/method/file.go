package method

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ParseFileFDStats(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	parts := bytes.Split(bytes.TrimSpace(content), []byte("\u0009"))
	if len(parts) < 3 {
		return nil, fmt.Errorf("unexpected number of file stats in %q", filename)
	}

	var fileFDStat = map[string]string{}
	// The file-nr proc is only 1 line with 3 values.
	fileFDStat["allocated"] = string(parts[0])
	// The second value is skipped as it will always be zero in linux 2.6.
	fileFDStat["maximum"] = string(parts[2])

	return fileFDStat, nil
}
