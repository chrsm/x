package indexof

// generate code
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=u.go gen "ValueType=uint FnName=U"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=u8.go gen "ValueType=uint8 FnName=U8"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=u16.go gen "ValueType=uint16 FnName=U16"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=u32.go gen "ValueType=uint32 FnName=U32"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=u64.go gen "ValueType=uint64 FnName=U64"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=i.go gen "ValueType=int FnName=I"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=i8.go gen "ValueType=int8 FnName=I8"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=i16.go gen "ValueType=int16 FnName=I16"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=i32.go gen "ValueType=int32 FnName=I32"
//go:generate genny -pkg=indexof -in=indexof.go_gen -out=i64.go gen "ValueType=int64 FnName=I64"

// generate tests
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=u_test.go gen "ValueType=uint FnName=U"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=u8_test.go gen "ValueType=uint8 FnName=U8"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=u16_test.go gen "ValueType=uint16 FnName=U16"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=u32_test.go gen "ValueType=uint32 FnName=U32"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=u64_test.go gen "ValueType=uint64 FnName=U64"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=i_test.go gen "ValueType=int FnName=I"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=i8_test.go gen "ValueType=int8 FnName=I8"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=i16_test.go gen "ValueType=int16 FnName=I16"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=i32_test.go gen "ValueType=int32 FnName=I32"
//go:generate genny -pkg=indexof -in=indexof_test.go_gen -out=i64_test.go gen "ValueType=int64 FnName=I64"
