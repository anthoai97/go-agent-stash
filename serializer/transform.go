package serializer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func WriteArrayStringToByte(array []string) []byte {
	data := ""
	for i := 0; i < len(array); i++ {
		str := array[i]
		if i != len(array)-1 {
			str += "\n"
		}
		data += str
	}
	return []byte(data)
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func GetEnvVar[T any](key string, defaultValue T) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	var ret any
	switch any(defaultValue).(type) {
	case string:
		ret = value

	case int:
		// don't actually ignore errors
		i, _ := strconv.ParseInt(value, 10, 64)
		ret = int(i)
	}
	return ret.(T)
}

func TimestampToPath(timestamp *timestamppb.Timestamp) string {
	if timestamp == nil {
		timestamp = timestamppb.Now()
	}
	year := timestamp.AsTime().Year()
	month := timestamp.AsTime().Month()
	day := timestamp.AsTime().Day()

	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
