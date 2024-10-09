package convert_test

import (
	"log"
	"os"
	"testing"

	"github.com/goplus/llgo/chore/gogensig/cmptest"
	"github.com/goplus/llgo/chore/gogensig/config"
	"github.com/goplus/llgo/chore/gogensig/convert"
	cppgtypes "github.com/goplus/llgo/chore/llcppg/types"
)

func TestCommentSlashSlashSlash(t *testing.T) {
	cmptest.RunTest(t, "comment", false, []config.SymbolEntry{
		{
			MangleName: "ExecuteFoo",
			CppName:    "ExecuteFoo",
			GoName:     "CustomExecuteFoo",
		},
	},
		&cppgtypes.Config{},
		`
/// Foo comment
struct Foo { int a; double b; bool c; };

/// ExecuteFoo comment
int ExecuteFoo(int a,Foo b);
	`, `
// Code generated by gogen; DO NOT EDIT.

package comment

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/// Foo comment
type Foo struct {
	a c.Int
	b float64
	c c.Int
}
/// ExecuteFoo comment
//go:linkname CustomExecuteFoo C.ExecuteFoo
func CustomExecuteFoo(a c.Int, b Foo) c.Int
	`, nil)
}

func TestEnum(t *testing.T) {
	cmptest.RunTest(t, "spectrum", true, []config.SymbolEntry{}, &cppgtypes.Config{
		Cplusplus: true,
	},
		`
	enum spectrum
	{
	    red,
	    orange,
	    yello,
	    green,
	    blue,
	    violet
	};

	enum kids
	{
	    nippy,
	    slats,
	    skippy,
	    nina,
	    liz
	};

	enum levels
	{
	    low = 100,
	    medium = 500,
	    high = 2000
	};

	enum feline
	{
	    cat,
	    lynx = 10,
	    puma,
	    tiger
	};

	enum class PieceType
	{
	    King = 1,
	    Queen,
	    Rook = 10,
	    Pawn
	};`,
		`
// Code generated by gogen; DO NOT EDIT.

package spectrum

import _ "unsafe"

const (
	Spectrum_red    = 0
	Spectrum_orange = 1
	Spectrum_yello  = 2
	Spectrum_green  = 3
	Spectrum_blue   = 4
	Spectrum_violet = 5
)
const (
	Kids_nippy  = 0
	Kids_slats  = 1
	Kids_skippy = 2
	Kids_nina   = 3
	Kids_liz    = 4
)
const (
	Levels_low    = 100
	Levels_medium = 500
	Levels_high   = 2000
)
const (
	Feline_cat   = 0
	Feline_lynx  = 10
	Feline_puma  = 11
	Feline_tiger = 12
)
const (
	Piecetype_King  = 1
	Piecetype_Queen = 2
	Piecetype_Rook  = 10
	Piecetype_Pawn  = 11
)
`, nil)
}

// Test generating a basic struct, correctly converting its fields,
// and properly referencing it in a function
func TestStructDeclRef(t *testing.T) {
	cmptest.RunTest(t, "typeref", false, []config.SymbolEntry{
		{
			MangleName: "ExecuteFoo",
			CppName:    "ExecuteFoo",
			GoName:     "CustomExecuteFoo",
		},
	},
		&cppgtypes.Config{},
		`
struct Foo { int a; double b; bool c; };
int ExecuteFoo(int a,Foo b);
	`, `
// Code generated by gogen; DO NOT EDIT.

package typeref

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type Foo struct {
	a c.Int
	b float64
	c c.Int
}
//go:linkname CustomExecuteFoo C.ExecuteFoo
func CustomExecuteFoo(a c.Int, b Foo) c.Int
	`, nil)
}

// Test if function names and type names can remove specified prefixes,
// generate correct linkname, and use function names defined in llcppg.symb.json
func TestCustomStruct(t *testing.T) {
	cmptest.RunTest(t, "typeref", false, []config.SymbolEntry{
		{MangleName: "lua_close", CppName: "lua_close", GoName: "Close"},
		{MangleName: "lua_newthread", CppName: "lua_newthread", GoName: "Newthread"},
		{MangleName: "lua_closethread", CppName: "lua_closethread", GoName: "Closethread"},
		{MangleName: "lua_resetthread", CppName: "lua_resetthread", GoName: "Resetthread"},
	}, &cppgtypes.Config{
		TrimPrefixes: []string{"lua_"},
	}, `
typedef struct lua_State lua_State;
typedef int (*lua_CFunction)(lua_State *L);
LUA_API void(lua_close)(lua_State *L);
LUA_API lua_State *(lua_newthread)(lua_State *L);
LUA_API int(lua_closethread)(lua_State *L, lua_State *from);
LUA_API int(lua_resetthread)(lua_State *L);
	`, `
// Code generated by gogen; DO NOT EDIT.

package typeref

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type State struct {
	Unused [8]uint8
}
// llgo:type C
type CFunction func(*State) c.Int
//go:linkname Close C.lua_close
func Close(L *State) c.Int
//go:linkname Closethread C.lua_closethread
func Closethread(L *State, from *State) c.Int
//go:linkname Resetthread C.lua_resetthread
func Resetthread(L *State) c.Int
	`, nil)
}

// Test if it can properly skip types from packages that have already been confirmed to be mapped
// The _int8_t, _int16_t, _int32_t, _int64_t below are types that have already been confirmed to be mapped (macos).
// The corresponding header files only define these aliases. For these header files, we skip them directly.
//
// In the follow include,the follow header files are included in the stdint.
// And this sys/_types/* int header files are have mapped,so we need skip them.
// And stdint.h's other included header files are not mapped yet, so we need to gradually generate them and create mappings for them.
//
// #include <sys/_types/_int8_t.h>
// #include <sys/_types/_int16_t.h>
// #include <sys/_types/_int32_t.h>
// #include <sys/_types/_int64_t.h>

// #include <sys/_types/_u_int8_t.h>
// #include <sys/_types/_u_int16_t.h>
// #include <sys/_types/_u_int32_t.h>
// #include <sys/_types/_u_int64_t.h>
func TestSkipBuiltinTypedefine(t *testing.T) {
	cmptest.RunTest(t, "skip", false, []config.SymbolEntry{
		{MangleName: "testInt", CppName: "testInt", GoName: "TestInt"},
		{MangleName: "testUint", CppName: "testUint", GoName: "TestUint"},
	}, &cppgtypes.Config{}, `
#include <stdint.h>

void testInt(int8_t a, int16_t b, int32_t c, int64_t d);
void testUint(u_int8_t a, u_int16_t b, u_int32_t c, u_int64_t d);
	`, `
// Code generated by gogen; DO NOT EDIT.

package skip

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
//go:linkname TestInt C.testInt
func TestInt(a int8, b int16, c c.Int, d c.LongLong)
//go:linkname TestUint C.testUint
func TestUint(a int8, b uint16, c c.Uint, d c.UlongLong)
	`, func(t *testing.T, pkg *convert.Package) {
		files, err := os.ReadDir(pkg.GetOutputDir())
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range files {
			log.Println("Generated file:", file.Name())
			typeAliasMap := convert.NewBuiltinTypeMap(".", "temp", nil).GetTypeAliases()
			for _, v := range typeAliasMap {
				if file.Name() == convert.HeaderFileToGo(v.HeaderFile) {
					t.Fatal("skip file should not be output")
				}
			}
		}
	})
}

// ===========================error
func TestNewAstConvert(t *testing.T) {
	convert.NewAstConvert(&convert.AstConvertConfig{
		PkgName:  "test",
		SymbFile: "",
		CfgFile:  "",
	})
}
