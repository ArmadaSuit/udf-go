package main

import (
	/*
		#include <stdlib.h>
		#include <string.h>
		#include <mysql.h>
	*/
	"C"
	"unsafe"

	"github.com/ArmadaSuit/udf-go/converter"
)

var converters []func(<-chan converter.KanaConverterRune) <-chan converter.KanaConverterRune

//export udf_convert_kana_init
func udf_convert_kana_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 2 {
		m := C.CString("2 arguments expected")
		defer C.free(unsafe.Pointer(m))
		C.strcpy(message, m)
		return C.bool(true)
	}

	argsTypes := unsafe.Slice(args.arg_type, args.arg_count)

	if argsTypes[0] != C.STRING_RESULT || argsTypes[1] != C.STRING_RESULT {
		m := C.CString("2 arguments must be string")
		defer C.free(unsafe.Pointer(m))
		C.strcpy(message, m)
		return C.bool(true)
	}

	argsArgs := unsafe.Slice(args.args, args.arg_count)

	var err error
	converters, err = converter.NewKanaConverters(C.GoString(argsArgs[1]))
	if err != nil {
		m := C.CString(err.Error())
		defer C.free(unsafe.Pointer(m))
		C.strcpy(message, m)
		return C.bool(true)
	}

	return C.bool(false)
}

//export udf_convert_kana
func udf_convert_kana(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, isNull *C.char, error *C.char) *C.char {
	argsArgs := unsafe.Slice(args.args, args.arg_count)
	s := C.GoString(argsArgs[0])
	in := converter.GenerateForKanaConverter(s)
	for _, c := range converters {
		in = c(in)
	}
	str := converter.StringForKanaConverter(in)
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.strcpy(result, cstr)
	*length = C.ulong(len(str))

	return result
}

func main() {
}
