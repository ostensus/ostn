package entropy

// #include "md5.h"
// #cgo CFLAGS: -w
import "C"
import (
	"errors"
)

func autoload() error {
	if ret := C.auto_extension(); ret != 0 {
		return errors.New("Could not autoload md5 function")
	}
	return nil
}
