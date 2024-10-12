package convert

import (
	"fmt"
	"go/types"

	"github.com/goplus/gogen"
	"github.com/goplus/llgo/chore/llcppg/ast"
)

type BuiltinTypeMap struct {
	pkgMap         map[string]gogen.PkgRef
	builtinTypeMap map[ast.BuiltinType]types.Type
	typeAliases    map[string]typeAliasInfo
}

type typeAliasInfo struct {
	Type       types.Type
	HeaderFile string
}

func NewBuiltinTypeMapWithPkgRefS(pkgs ...gogen.PkgRef) *BuiltinTypeMap {
	builtinTypeMap := &BuiltinTypeMap{}
	builtinTypeMap.pkgMap = make(map[string]gogen.PkgRef)
	for _, pkg := range pkgs {
		builtinTypeMap.pkgMap[pkg.Types.Name()] = pkg
		builtinTypeMap.pkgMap[pkg.Path()] = pkg
	}
	builtinTypeMap.initBuiltinTypeMap()
	return builtinTypeMap
}

func NewBuiltinTypeMap(pkgPath, name string, conf *gogen.Config) *BuiltinTypeMap {
	p := gogen.NewPackage(pkgPath, name, conf)
	clib := p.Import("github.com/goplus/llgo/c")
	builtinTypeMap := NewBuiltinTypeMapWithPkgRefS(clib, p.Unsafe())
	return builtinTypeMap
}

func (p *BuiltinTypeMap) CType(typ string) types.Type {
	clib, ok := p.pkgMap["c"]
	if ok {
		return clib.Ref(typ).Type()
	}
	return nil
}

func (p *BuiltinTypeMap) IsVoidType(typ types.Type) bool {
	voidType := p.builtinTypeMap[ast.BuiltinType{Kind: ast.Void}]
	return typ == voidType
}

func (p *BuiltinTypeMap) FindBuiltinType(builtinType ast.BuiltinType) (types.Type, error) {
	t, ok := p.builtinTypeMap[builtinType]
	if ok {
		return t, nil
	}
	return nil, fmt.Errorf("%s", "not found in type map")
}

func (p *BuiltinTypeMap) FindTypeAlias(name string) (types.Type, error) {
	t, ok := p.typeAliases[name]
	if ok {
		return t.Type, nil
	}
	return nil, fmt.Errorf("%s", "not found in type alias map")
}

func (p *BuiltinTypeMap) GetTypeAliases() map[string]typeAliasInfo {
	return p.typeAliases
}

func (p *BuiltinTypeMap) initBuiltinTypeMap() {
	// todo(zzy): int128/uint128  half(float16),long double,float 128
	p.builtinTypeMap = map[ast.BuiltinType]types.Type{
		{Kind: ast.Void}:                                    types.Typ[types.UntypedNil], // For a invalid type
		{Kind: ast.Bool}:                                    types.Typ[types.Bool],       // Bool
		{Kind: ast.Char, Flags: ast.Signed}:                 p.CType("Char"),             // Char_S
		{Kind: ast.Char, Flags: ast.Unsigned}:               p.CType("Char"),             // Char_U
		{Kind: ast.WChar}:                                   types.Typ[types.Int16],      // WChar
		{Kind: ast.Char16}:                                  types.Typ[types.Int16],      // Char16
		{Kind: ast.Char32}:                                  types.Typ[types.Int32],      // Char32
		{Kind: ast.Int, Flags: ast.Short}:                   types.Typ[types.Int16],      // Short
		{Kind: ast.Int, Flags: ast.Short | ast.Unsigned}:    types.Typ[types.Uint16],     // UShort
		{Kind: ast.Int}:                                     p.CType("Int"),              // Int
		{Kind: ast.Int, Flags: ast.Unsigned}:                p.CType("Uint"),             // UInt
		{Kind: ast.Int, Flags: ast.Long}:                    p.CType("Long"),             // Long
		{Kind: ast.Int, Flags: ast.Long | ast.Unsigned}:     p.CType("Ulong"),            // Ulong
		{Kind: ast.Int, Flags: ast.LongLong}:                p.CType("LongLong"),         // LongLong
		{Kind: ast.Int, Flags: ast.LongLong | ast.Unsigned}: p.CType("UlongLong"),        // ULongLong
		{Kind: ast.Float}:                                   p.CType("Float"),            // Float
		{Kind: ast.Float, Flags: ast.Double}:                p.CType("Double"),           // Double
		{Kind: ast.Float, Flags: ast.Double | ast.Long}:     p.CType("Double"),           // Long Double (same as double,need more precision)
		{Kind: ast.Complex}:                                 types.Typ[types.Complex64],  // ComplexFloat
		{Kind: ast.Complex, Flags: ast.Double}:              types.Typ[types.Complex128], // ComplexDouble
	}

	// todo(zzy): more types & platform
	p.typeAliases = map[string]typeAliasInfo{
		"int8_t":    {Type: p.CType("Char"), HeaderFile: "sys/_types/_int8_t.h"},
		"uint8_t":   {Type: p.CType("Char"), HeaderFile: "_types/_uint8_t.h"},
		"u_int8_t":  {Type: p.CType("Char"), HeaderFile: "sys/_types/_u_int8_t.h"},
		"int16_t":   {Type: types.Typ[types.Int16], HeaderFile: "sys/_types/_int16_t.h"},
		"uint16_t":  {Type: types.Typ[types.Uint16], HeaderFile: "_types/_uint16_t.h"},
		"u_int16_t": {Type: types.Typ[types.Uint16], HeaderFile: "sys/_types/_u_int16_t.h"},
		"int32_t":   {Type: p.CType("Int"), HeaderFile: "sys/_types/_int32_t.h"},
		"uint32_t":  {Type: p.CType("Uint"), HeaderFile: "_types/_uint32_t.h"},
		"u_int32_t": {Type: p.CType("Uint"), HeaderFile: "sys/_types/_u_int32_t.h"},
		"int64_t":   {Type: p.CType("LongLong"), HeaderFile: "sys/_types/_int64_t.h"},
		"uint64_t":  {Type: p.CType("UlongLong"), HeaderFile: "_types/_uint64_t.h"},
		"u_int64_t": {Type: p.CType("UlongLong"), HeaderFile: "sys/_types/_u_int64_t.h"},
		// todo(zzy): undef type log
		// other headerfile
		//"size_t": {Type: types.Typ[types.Uintptr], HeaderFile: "sys/_types/_u_int64_t.h"},
	}
}
