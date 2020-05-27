// Code generated by github.com/ridge/must/generator. DO NOT EDIT.

package must

import (
	"io"
	"net"
	"os"
	"net/url"
)

// Uint8 panics on error, returns the first argument otherwise
func Uint8(arg uint8, err error) uint8 {
	OK(err)
	return arg
}

// Uint8s panics on error, returns the first argument otherwise
func Uint8s(arg []uint8, err error) []uint8 {
	OK(err)
	return arg
}

// Int8 panics on error, returns the first argument otherwise
func Int8(arg int8, err error) int8 {
	OK(err)
	return arg
}

// Int8s panics on error, returns the first argument otherwise
func Int8s(arg []int8, err error) []int8 {
	OK(err)
	return arg
}

// Uint16 panics on error, returns the first argument otherwise
func Uint16(arg uint16, err error) uint16 {
	OK(err)
	return arg
}

// Uint16s panics on error, returns the first argument otherwise
func Uint16s(arg []uint16, err error) []uint16 {
	OK(err)
	return arg
}

// Int16 panics on error, returns the first argument otherwise
func Int16(arg int16, err error) int16 {
	OK(err)
	return arg
}

// Int16s panics on error, returns the first argument otherwise
func Int16s(arg []int16, err error) []int16 {
	OK(err)
	return arg
}

// Uint32 panics on error, returns the first argument otherwise
func Uint32(arg uint32, err error) uint32 {
	OK(err)
	return arg
}

// Uint32s panics on error, returns the first argument otherwise
func Uint32s(arg []uint32, err error) []uint32 {
	OK(err)
	return arg
}

// Int32 panics on error, returns the first argument otherwise
func Int32(arg int32, err error) int32 {
	OK(err)
	return arg
}

// Int32s panics on error, returns the first argument otherwise
func Int32s(arg []int32, err error) []int32 {
	OK(err)
	return arg
}

// Uint64 panics on error, returns the first argument otherwise
func Uint64(arg uint64, err error) uint64 {
	OK(err)
	return arg
}

// Uint64s panics on error, returns the first argument otherwise
func Uint64s(arg []uint64, err error) []uint64 {
	OK(err)
	return arg
}

// Int64 panics on error, returns the first argument otherwise
func Int64(arg int64, err error) int64 {
	OK(err)
	return arg
}

// Int64s panics on error, returns the first argument otherwise
func Int64s(arg []int64, err error) []int64 {
	OK(err)
	return arg
}

// Uint panics on error, returns the first argument otherwise
func Uint(arg uint, err error) uint {
	OK(err)
	return arg
}

// Uints panics on error, returns the first argument otherwise
func Uints(arg []uint, err error) []uint {
	OK(err)
	return arg
}

// Int panics on error, returns the first argument otherwise
func Int(arg int, err error) int {
	OK(err)
	return arg
}

// Ints panics on error, returns the first argument otherwise
func Ints(arg []int, err error) []int {
	OK(err)
	return arg
}

// Bool panics on error, returns the first argument otherwise
func Bool(arg bool, err error) bool {
	OK(err)
	return arg
}

// Bools panics on error, returns the first argument otherwise
func Bools(arg []bool, err error) []bool {
	OK(err)
	return arg
}

// Float32 panics on error, returns the first argument otherwise
func Float32(arg float32, err error) float32 {
	OK(err)
	return arg
}

// Float32s panics on error, returns the first argument otherwise
func Float32s(arg []float32, err error) []float32 {
	OK(err)
	return arg
}

// Float64 panics on error, returns the first argument otherwise
func Float64(arg float64, err error) float64 {
	OK(err)
	return arg
}

// Float64s panics on error, returns the first argument otherwise
func Float64s(arg []float64, err error) []float64 {
	OK(err)
	return arg
}

// Complex64 panics on error, returns the first argument otherwise
func Complex64(arg complex64, err error) complex64 {
	OK(err)
	return arg
}

// Complex64s panics on error, returns the first argument otherwise
func Complex64s(arg []complex64, err error) []complex64 {
	OK(err)
	return arg
}

// Complex128 panics on error, returns the first argument otherwise
func Complex128(arg complex128, err error) complex128 {
	OK(err)
	return arg
}

// Complex128s panics on error, returns the first argument otherwise
func Complex128s(arg []complex128, err error) []complex128 {
	OK(err)
	return arg
}

// Byte panics on error, returns the first argument otherwise
func Byte(arg byte, err error) byte {
	OK(err)
	return arg
}

// Bytes panics on error, returns the first argument otherwise
func Bytes(arg []byte, err error) []byte {
	OK(err)
	return arg
}

// Rune panics on error, returns the first argument otherwise
func Rune(arg rune, err error) rune {
	OK(err)
	return arg
}

// Runes panics on error, returns the first argument otherwise
func Runes(arg []rune, err error) []rune {
	OK(err)
	return arg
}

// Uintptr panics on error, returns the first argument otherwise
func Uintptr(arg uintptr, err error) uintptr {
	OK(err)
	return arg
}

// Uintptrs panics on error, returns the first argument otherwise
func Uintptrs(arg []uintptr, err error) []uintptr {
	OK(err)
	return arg
}

// String panics on error, returns the first argument otherwise
func String(arg string, err error) string {
	OK(err)
	return arg
}

// Strings panics on error, returns the first argument otherwise
func Strings(arg []string, err error) []string {
	OK(err)
	return arg
}

// Any panics on error, returns the first argument otherwise
func Any(arg interface{}, err error) interface{} {
	OK(err)
	return arg
}

// OSFile panics on error, returns the first argument otherwise
func OSFile(arg *os.File, err error) *os.File {
	OK(err)
	return arg
}

// OSFileInfo panics on error, returns the first argument otherwise
func OSFileInfo(arg os.FileInfo, err error) os.FileInfo {
	OK(err)
	return arg
}

// OSFileInfos panics on error, returns the first argument otherwise
func OSFileInfos(arg []os.FileInfo, err error) []os.FileInfo {
	OK(err)
	return arg
}

// IOReadCloser panics on error, returns the first argument otherwise
func IOReadCloser(arg io.ReadCloser, err error) io.ReadCloser {
	OK(err)
	return arg
}

// IOWriter panics on error, returns the first argument otherwise
func IOWriter(arg io.Writer, err error) io.Writer {
	OK(err)
	return arg
}

// NetIP panics on error, returns the first argument otherwise
func NetIP(arg net.IP, err error) net.IP {
	OK(err)
	return arg
}

// NetListener panics on error, returns the first argument otherwise
func NetListener(arg net.Listener, err error) net.Listener {
	OK(err)
	return arg
}

// NetURL panics on error, returns the first argument otherwise
func NetURL(arg *url.URL, err error) *url.URL {
	OK(err)
	return arg
}
