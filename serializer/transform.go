package serializer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

	case bool:
		i, _ := strconv.ParseBool(value)
		ret = bool(i)
	}
	return ret.(T)
}

func TimestampToPath(timestamp *timestamppb.Timestamp) string {
	if timestamp == nil {
		timestamp = timestamppb.Now()
	}
	year := fmt.Sprintf("%02d", timestamp.AsTime().Year())
	month := fmt.Sprintf("%02d", timestamp.AsTime().Month())
	day := fmt.Sprintf("%02d", timestamp.AsTime().Day())

	return fmt.Sprintf("%s/%s/%s", year, month, day)
}

func XTimeFromNToNow(fromTime time.Time) int64 {
	return (time.Now().UnixNano() - fromTime.UnixNano()) / (int64(time.Millisecond) / int64(time.Nanosecond))
}
