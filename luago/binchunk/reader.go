package binchunk

import (
	"encoding/binary"
	"fmt"
	"math"
)

type reader struct {
	data []byte
}

func (r *reader) readByte() byte {
	b := r.data[0]
	r.data = r.data[1:]
	return b
}
func (r *reader) readUint32() uint32 {
	data := binary.LittleEndian.Uint32(r.data)
	r.data = r.data[4:]
	return data
}
func (r *reader) readUint64() uint64 {
	data := binary.LittleEndian.Uint64(r.data)
	r.data = r.data[8:]
	return data
}
func (r *reader) readLuaInterger() int64 {
	return int64(r.readUint64())
}
func (r *reader) readLuaNumber() float64 {
	return math.Float64frombits(r.readUint64())
}
func (r *reader) readString() string {
	length := uint(r.readByte())
	if length == 0 {
		return ""
	}
	if length == 0xFF {
		length = uint(r.readUint64())
	}
	data := r.readBytes(length - 1)
	return string(data)
}
func (r *reader) readBytes(length uint) []byte {
	data := r.data[:length]
	r.data = r.data[length:]
	return data
}
func (r *reader) readCode() []uint32 {
	code := make([]uint32, r.readUint32())
	for i := range code {
		code[i] = r.readUint32()
	}
	return code
}
func (r *reader) readConstants() []interface{} {
	constants := make([]interface{}, r.readUint32())
	for i := range constants {
		constants[i] = r.readConstant()
	}
	return constants
}
func (r *reader) readConstant() interface{} {
	switch r.readByte() {
	case TAG_NIL:
		return nil
	case TAG_BOOLEAN:
		return r.readByte() != 0
	case TAG_INTEGER:
		return r.readLuaInterger()
	case TAG_NUMBER:
		return r.readLuaNumber()
	case TAG_SHORT_STR:
		return r.readString()
	case TAG_LONG_STR:
		return r.readString()
	default:
		panic("corrupted!")
	}
}
func (r *reader) readUpvalues() []Upvalue {
	upvalues := make([]Upvalue, r.readUint32())
	for i := range upvalues {
		upvalues[i] = Upvalue{
			Instack: r.readByte(),
			Idx:     r.readByte(),
		}
	}
	return upvalues
}
func (r *reader) readProtos(data string) []*ProtoType {
	protos := make([]*ProtoType, r.readUint32())
	for i := range protos {
		protos[i] = r.readProto(data)
	}
	return protos
}
func (r *reader) readLineInfo() []uint32 {
	lineInfo := make([]uint32, r.readUint32())
	for i := range lineInfo {
		lineInfo[i] = r.readUint32()
	}
	return lineInfo
}
func (r *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, r.readUint32())
	for i := range locVars {
		locVars[i] = LocVar{
			VarName: r.readString(),
			StartPC: r.readUint32(),
			EndPC:   r.readUint32(),
		}
	}
	return locVars
}
func (r *reader) readUpvalueNames() []string {
	names := make([]string, r.readUint32())
	for i := range names {
		names[i] = r.readString()
	}
	return names
}
func (r *reader) checkHeader() {
	if string(r.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk!")
	} else if r.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	} else if r.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	} else if string(r.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	} else if r.readByte() != CINT_SIZE {
		panic("int size mismatch!")
	} else if r.readByte() != CSZIET_SIZE {
		panic("size_t size mismatch!")
	} else if data := r.readByte(); data != INSTRUCTION_SIZE {
		fmt.Println("data", data, INSTRUCTION_SIZE)
		panic("instruction size mismatch!")
	} else if r.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Interger size mismatch!")
	} else if r.readByte() != LUA_NUMBER_SIZE {
		panic("lua_Number size mismatch!")
	} else if r.readLuaInterger() != LUAC_INT {
		panic("endianness mismatch!")
	} else if r.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}
func (r *reader) readProto(src string) *ProtoType {
	data := r.readString()
	if data == "" {
		data = src
	}
	return &ProtoType{
		Source:          data,
		LineDefined:     r.readUint32(),
		LastLineDefined: r.readUint32(),
		NumParams:       r.readByte(),
		IsVararg:        r.readByte(),
		MaxStackSize:    r.readByte(),
		Code:            r.readCode(),
		Constants:       r.readConstants(),
		Upvalues:        r.readUpvalues(),
		Protos:          r.readProtos(src),
		LineInfo:        r.readLineInfo(),
		LocVars:         r.readLocVars(),
		UpvaluesNames:   r.readUpvalueNames(),
	}
}
