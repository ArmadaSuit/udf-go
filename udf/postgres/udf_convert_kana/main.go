package main

import (
	/*
		#include <postgres.h>

		extern Datum udf_convert_kana(PG_FUNCTION_ARGS);
	*/
	"C"

	"github.com/ArmadaSuit/udf-go/converter"
)

//export udf_go_convert_kana
func udf_go_convert_kana(text *C.char, mode *C.char) (*C.char, *C.char) {
	converters, err := converter.NewKanaConverters(C.GoString(mode))
	if err != nil {
		return nil, C.CString(err.Error())
	}

	s := C.GoString(text)
	in := converter.GenerateForKanaConverter(s)
	for _, c := range converters {
		in = c(in)
	}
	str := converter.StringForKanaConverter(in)

	return C.CString(str), nil
}

func main() {
}
