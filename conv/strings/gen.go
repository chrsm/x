package strings

// generate code
//go:generate genny -pkg=strings -in=to_unsigned.go_gen -out=u8.go gen "ValueType=uint8 FnName=U8 Bits=8"
//go:generate genny -pkg=strings -in=to_unsigned.go_gen -out=u16.go gen "ValueType=uint16 FnName=U16 Bits=16"
//go:generate genny -pkg=strings -in=to_unsigned.go_gen -out=u32.go gen "ValueType=uint32 FnName=U32 Bits=32"
//go:generate genny -pkg=strings -in=to_unsigned.go_gen -out=u64.go gen "ValueType=uint64 FnName=U64 Bits=64"

// generate signed
//go:generate genny -pkg=strings -in=to_signed.go_gen -out=i8.go gen "ValueType=int8 FnName=I8 Bits=8"
//go:generate genny -pkg=strings -in=to_signed.go_gen -out=i16.go gen "ValueType=int16 FnName=I16 Bits=16"
//go:generate genny -pkg=strings -in=to_signed.go_gen -out=i32.go gen "ValueType=int32 FnName=I32 Bits=32"
//go:generate genny -pkg=strings -in=to_signed.go_gen -out=i64.go gen "ValueType=int64 FnName=I64 Bits=64"

// generate tests
//go:generate genny -pkg=strings -in=to_unsigned_test.go_gen -out=u8_test.go gen "ValueType=uint8 FnName=U8 Bits=8"
//go:generate genny -pkg=strings -in=to_unsigned_test.go_gen -out=u16_test.go gen "ValueType=uint16 FnName=U16 Bits=16"
//go:generate genny -pkg=strings -in=to_unsigned_test.go_gen -out=u32_test.go gen "ValueType=uint32 FnName=U32 Bits=32"
//go:generate genny -pkg=strings -in=to_unsigned_test.go_gen -out=u64_test.go gen "ValueType=uint64 FnName=U64 Bits=64"
//go:generate genny -pkg=strings -in=to_signed_test.go_gen -out=i8_test.go gen "ValueType=int8 FnName=I8 Bits=8"
//go:generate genny -pkg=strings -in=to_signed_test.go_gen -out=i16_test.go gen "ValueType=int16 FnName=I16 Bits=16"
//go:generate genny -pkg=strings -in=to_signed_test.go_gen -out=i32_test.go gen "ValueType=int32 FnName=I32 Bits=32"
//go:generate genny -pkg=strings -in=to_signed_test.go_gen -out=i64_test.go gen "ValueType=int64 FnName=I64 Bits=64"
