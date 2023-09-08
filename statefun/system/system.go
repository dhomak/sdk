

// Foliage primary statefun system package.
// Provides shared system functions for statefun packages
package system

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func MsgOnErrorReturn(err error) {
	fmt.Printf("ERROR: %s\n", err)
}

func GetEnvMustProceed[T any](key string, defaultVal T) T {
	v, _ := GetEnv(key, defaultVal)
	return v
}

func GetEnv[T any](key string, defaultVal T) (value T, err error) {
	value = defaultVal
	err = nil

	if strValue, exists := os.LookupEnv(key); exists {
		switch any(defaultVal).(type) {
		case string:
			value = any(strValue).(T)
		case uint:
			v, e := strconv.ParseUint(strValue, 10, 0)
			value = any(uint(v)).(T)
			err = e
		case uint8:
			v, e := strconv.ParseUint(strValue, 10, 8)
			value = any(uint8(v)).(T)
			err = e
		case uint16:
			v, e := strconv.ParseUint(strValue, 10, 16)
			value = any(uint16(v)).(T)
			err = e
		case uint32:
			v, e := strconv.ParseUint(strValue, 10, 32)
			value = any(uint32(v)).(T)
			err = e
		case uint64:
			v, e := strconv.ParseUint(strValue, 10, 0)
			value = any(v).(T)
			err = e
		case int:
			v, e := strconv.ParseInt(strValue, 10, 0)
			value = any(int(v)).(T)
			err = e
		case int8:
			v, e := strconv.ParseInt(strValue, 10, 8)
			value = any(int8(v)).(T)
			err = e
		case int16:
			v, e := strconv.ParseInt(strValue, 10, 16)
			value = any(int16(v)).(T)
			err = e
		case int32:
			v, e := strconv.ParseInt(strValue, 10, 32)
			value = any(int32(v)).(T)
			err = e
		case int64:
			v, e := strconv.ParseInt(strValue, 10, 64)
			value = any(v).(T)
			err = e
		case bool:
			v, e := strconv.ParseBool(strValue)
			value = any(v).(T)
			err = e
		case float32:
			v, e := strconv.ParseFloat(strValue, 32)
			value = any(float32(v)).(T)
			err = e
		case float64:
			v, e := strconv.ParseFloat(strValue, 64)
			value = any(v).(T)
			err = e
		}
	}
	return
}

func Str2Int(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return value
	}
	return 0
}

func MergeMaps[T any](m1 map[string]T, m2 map[string]T) map[string]T {
	merged := make(map[string]T)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

func Int64ToBytes(v int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return b
}

func BytesToInt64(v []byte) int64 {
	if len(v) < 8 {
		return 0
	}
	return int64(binary.LittleEndian.Uint64(v))
}

func GetCurrentTimeNs() int64 {
	return time.Now().UnixNano()
}

func GetUniqueStrID() string {
	baseStr := fmt.Sprintf("%d-%f", GetCurrentTimeNs(), rand.Float64())
	data := []byte(baseStr)
	hash := md5.Sum(data)
	id := hex.EncodeToString(hash[:])
	return id
}

func GetHashStr(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	id := hex.EncodeToString(hash[:])
	return id
}
