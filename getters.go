package cfg

import (
	"time"
)

func Int64(path string) int64 {
	return inst.Int64(path)
}

func MustInt64(path string) int64 {
	return inst.MustInt64(path)
}

func Int64s(path string) []int64 {
	return inst.Int64s(path)
}

func MustInt64s(path string) []int64 {
	return inst.MustInt64s(path)
}

func Int64Map(path string) map[string]int64 {
	return inst.Int64Map(path)
}

func MustInt64Map(path string) map[string]int64 {
	return inst.MustInt64Map(path)
}

func Int(path string) int {
	return inst.Int(path)
}

func MustInt(path string) int {
	return inst.MustInt(path)
}

func Ints(path string) []int {
	return inst.Ints(path)
}

func MustInts(path string) []int {
	return inst.MustInts(path)
}

func IntMap(path string) map[string]int {
	return inst.IntMap(path)
}

func MustIntMap(path string) map[string]int {
	return inst.MustIntMap(path)
}

func Float64(path string) float64 {
	return inst.Float64(path)
}

func MustFloat64(path string) float64 {
	return inst.MustFloat64(path)
}

func Float64s(path string) []float64 {
	return inst.Float64s(path)
}

func MustFloat64s(path string) []float64 {
	return inst.MustFloat64s(path)
}

func Float64Map(path string) map[string]float64 {
	return inst.Float64Map(path)
}

func MustFloat64Map(path string) map[string]float64 {
	return inst.MustFloat64Map(path)
}

func Duration(path string) time.Duration {
	return inst.Duration(path)
}

func MustDuration(path string) time.Duration {
	return inst.MustDuration(path)
}

func Time(path, layout string) time.Time {
	return inst.Time(path, layout)
}

func MustTime(path, layout string) time.Time {
	return inst.MustTime(path, layout)
}

func String(path string) string {
	return inst.String(path)
}

func MustString(path string) string {
	return inst.MustString(path)
}

func Strings(path string) []string {
	return inst.Strings(path)
}

func MustStrings(path string) []string {
	return inst.MustStrings(path)
}

func StringMap(path string) map[string]string {
	return inst.StringMap(path)
}

func MustStringMap(path string) map[string]string {
	return inst.MustStringMap(path)
}

func StringsMap(path string) map[string][]string {
	return inst.StringsMap(path)
}

func MustStringsMap(path string) map[string][]string {
	return inst.MustStringsMap(path)
}

func Bytes(path string) []byte {
	return inst.Bytes(path)
}

func MustBytes(path string) []byte {
	return inst.MustBytes(path)
}

func Bool(path string) bool {
	return inst.Bool(path)
}

func Bools(path string) []bool {
	return inst.Bools(path)
}

func MustBools(path string) []bool {
	return inst.MustBools(path)
}

func BoolMap(path string) map[string]bool {
	return inst.BoolMap(path)
}

func MustBoolMap(path string) map[string]bool {
	return inst.MustBoolMap(path)
}
