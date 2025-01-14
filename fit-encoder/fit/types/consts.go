// Copyright 2021 Artem Mikheev

package types

import (
	"errors"
	"reflect"
)

const (
	FIT_TYPE_BYTE_SIZE    FitUint8 = 1
	FIT_TYPE_ENUM_SIZE    FitUint8 = FIT_TYPE_BYTE_SIZE
	FIT_TYPE_UINT8_SIZE   FitUint8 = 1
	FIT_TYPE_SINT8_SIZE   FitUint8 = FIT_TYPE_UINT8_SIZE
	FIT_TYPE_UINT16_SIZE  FitUint8 = 2
	FIT_TYPE_SINT16_SIZE  FitUint8 = FIT_TYPE_UINT16_SIZE
	FIT_TYPE_UINT32_SIZE  FitUint8 = 4
	FIT_TYPE_SINT32_SIZE  FitUint8 = FIT_TYPE_UINT32_SIZE
	FIT_TYPE_FLOAT32_SIZE FitUint8 = 4
	FIT_TYPE_FLOAT64_SIZE FitUint8 = 8
	FIT_TYPE_UINT64_SIZE  FitUint8 = 8
	FIT_TYPE_SINT64_SIZE  FitUint8 = FIT_TYPE_UINT64_SIZE
)

const (
	FIT_TYPE_INVALID FitBaseType = 0xFF
	FIT_TYPE_ENUM    FitBaseType = 0
	FIT_TYPE_SINT8   FitBaseType = 1
	FIT_TYPE_UINT8   FitBaseType = 2
	FIT_TYPE_SINT16  FitBaseType = 131
	FIT_TYPE_UINT16  FitBaseType = 132
	FIT_TYPE_SINT32  FitBaseType = 133
	FIT_TYPE_UINT32  FitBaseType = 134
	FIT_TYPE_STRING  FitBaseType = 7
	FIT_TYPE_FLOAT32 FitBaseType = 136
	FIT_TYPE_FLOAT64 FitBaseType = 137
	FIT_TYPE_UINT8Z  FitBaseType = 10
	FIT_TYPE_UINT16Z FitBaseType = 139
	FIT_TYPE_UINT32Z FitBaseType = 140
	FIT_TYPE_BYTE    FitBaseType = 13
	FIT_TYPE_SINT64  FitBaseType = 142
	FIT_TYPE_UINT64  FitBaseType = 143
	FIT_TYPE_UINT64Z FitBaseType = 144
	FIT_TYPE_COUNT   FitBaseType = 17
)

var (
	ErrUnknownFitType      = errors.New("unknown fit type")
	ErrFitBaseTypeMismatch = errors.New("fit base type and actual type of value differ")
)

var FitTypeSize = map[FitBaseType]FitUint8{
	FIT_TYPE_ENUM:    FIT_TYPE_ENUM_SIZE,
	FIT_TYPE_SINT8:   FIT_TYPE_SINT8_SIZE,
	FIT_TYPE_UINT8:   FIT_TYPE_UINT8_SIZE,
	FIT_TYPE_SINT16:  FIT_TYPE_SINT16_SIZE,
	FIT_TYPE_UINT16:  FIT_TYPE_UINT16_SIZE,
	FIT_TYPE_SINT32:  FIT_TYPE_SINT32_SIZE,
	FIT_TYPE_UINT32:  FIT_TYPE_UINT32_SIZE,
	FIT_TYPE_FLOAT32: FIT_TYPE_FLOAT32_SIZE,
	FIT_TYPE_FLOAT64: FIT_TYPE_FLOAT64_SIZE,
	FIT_TYPE_SINT64:  FIT_TYPE_SINT64_SIZE,
	FIT_TYPE_UINT64:  FIT_TYPE_UINT64_SIZE,
}

var FitTypeMap = map[FitBaseType]reflect.Type{
	FIT_TYPE_ENUM:   reflect.TypeOf((*FitEnum)(nil)).Elem(),
	FIT_TYPE_SINT8:  reflect.TypeOf((*FitSint8)(nil)).Elem(),
	FIT_TYPE_UINT8:  reflect.TypeOf((*FitUint8)(nil)).Elem(),
	FIT_TYPE_SINT16: reflect.TypeOf((*FitSint16)(nil)).Elem(),
	FIT_TYPE_UINT16: reflect.TypeOf((*FitUint16)(nil)).Elem(),
	FIT_TYPE_SINT32: reflect.TypeOf((*FitSint32)(nil)).Elem(),
	FIT_TYPE_UINT32: reflect.TypeOf((*FitUint32)(nil)).Elem(),
	FIT_TYPE_SINT64: reflect.TypeOf((*FitSint64)(nil)).Elem(),
	FIT_TYPE_UINT64: reflect.TypeOf((*FitUint64)(nil)).Elem(),
	FIT_TYPE_STRING: reflect.TypeOf((*FitString)(nil)).Elem(),
}

var implementedFitTypes = map[reflect.Type]bool{
	FitTypeMap[FIT_TYPE_ENUM]:   true,
	FitTypeMap[FIT_TYPE_SINT8]:  true,
	FitTypeMap[FIT_TYPE_UINT8]:  true,
	FitTypeMap[FIT_TYPE_SINT16]: true,
	FitTypeMap[FIT_TYPE_UINT16]: true,
	FitTypeMap[FIT_TYPE_SINT32]: true,
	FitTypeMap[FIT_TYPE_UINT32]: true,
	FitTypeMap[FIT_TYPE_SINT64]: true,
	FitTypeMap[FIT_TYPE_UINT64]: true,
	FitTypeMap[FIT_TYPE_STRING]: true,
}
