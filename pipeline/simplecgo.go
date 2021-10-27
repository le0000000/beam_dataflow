// Package simpecgo contains some simple cgo functions.
package simplecgo

// #cgo CFLAGS: -g -Wall
// #include <stdbool.h>
// #include <stdlib.h>
// #include "pipeline/simple_c_bridge.h"
import (
	"C"
)

func GetValue() int {
	value := C.CGetValue()
	return int(value)
}
