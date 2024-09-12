package genpkg_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goplus/gogen"
	"github.com/goplus/llgo/chore/gogensig/visitor/genpkg"
	"github.com/goplus/llgo/chore/llcppg/ast"
)

func TestToType(t *testing.T) {
	pkg := genpkg.NewPackage(".", "testpkg", &gogen.Config{})

	testCases := []struct {
		name     string
		input    *ast.BuiltinType
		expected string
	}{
		{"Void", &ast.BuiltinType{Kind: ast.Void}, "invalid type"},
		{"Bool", &ast.BuiltinType{Kind: ast.Bool}, "bool"},
		{"Char_S", &ast.BuiltinType{Kind: ast.Char, Flags: ast.Signed}, "int8"},
		{"Char_U", &ast.BuiltinType{Kind: ast.Char, Flags: ast.Unsigned}, "int8"},
		{"WChar", &ast.BuiltinType{Kind: ast.WChar}, "int16"},
		{"Char16", &ast.BuiltinType{Kind: ast.Char16}, "int16"},
		{"Char32", &ast.BuiltinType{Kind: ast.Char32}, "int32"},
		{"Short", &ast.BuiltinType{Kind: ast.Int, Flags: ast.Short}, "int16"},
		{"UShort", &ast.BuiltinType{Kind: ast.Int, Flags: ast.Short | ast.Unsigned}, "uint16"},
		{"Int", &ast.BuiltinType{Kind: ast.Int}, "github.com/goplus/llgo/c.Int"},
		{"UInt", &ast.BuiltinType{Kind: ast.Int, Flags: ast.Unsigned}, "github.com/goplus/llgo/c.Uint"},
		{"Long", &ast.BuiltinType{Kind: ast.Int, Flags: ast.Long}, "github.com/goplus/llgo/c.Long"},
		{"ULong", &ast.BuiltinType{Kind: ast.Int, Flags: ast.Long | ast.Unsigned}, "github.com/goplus/llgo/c.Ulong"},
		{"LongLong", &ast.BuiltinType{Kind: ast.Int, Flags: ast.LongLong}, "github.com/goplus/llgo/c.LongLong"},
		{"ULongLong", &ast.BuiltinType{Kind: ast.Int, Flags: ast.LongLong | ast.Unsigned}, "github.com/goplus/llgo/c.UlongLong"},
		{"Float", &ast.BuiltinType{Kind: ast.Float}, "float32"},
		{"Double", &ast.BuiltinType{Kind: ast.Float, Flags: ast.Double}, "float64"},
		{"ComplexFloat", &ast.BuiltinType{Kind: ast.Complex}, "complex64"},
		{"ComplexDouble", &ast.BuiltinType{Kind: ast.Complex, Flags: ast.Double}, "complex128"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := pkg.ToType(tc.input)
			if result == nil || result.String() != tc.expected {
				t.Errorf("unexpected result:%s expected:%s", result.String(), tc.expected)
			}
		})
	}
}

func TestNewPackage(t *testing.T) {
	pkg := genpkg.NewPackage(".", "testpkg", &gogen.Config{})
	if pkg == nil {
		t.Fatal("NewPackage failed")
	}
	comparePackageOutput(t, pkg, `package testpkg`)
}

func TestFuncDeclBasic(t *testing.T) {
	input := &ast.FuncDecl{
		Name: &ast.Ident{Name: "foo"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{Name: "a"},
						},
						Type: &ast.BuiltinType{
							Kind: ast.Int,
						},
					},
					{
						Names: []*ast.Ident{
							{Name: "a"},
						},
						Type: &ast.BuiltinType{
							Kind:  ast.Int,
							Flags: ast.Unsigned | ast.Long,
						},
					},
				},
			},
			Ret: &ast.BuiltinType{
				Kind:  ast.Float,
				Flags: ast.Double,
			},
		},
	}
	pkg := genpkg.NewPackage(".", "testpkg", &gogen.Config{})
	if pkg == nil {
		t.Fatal("NewPackage failed")
	}
	err := pkg.NewFuncDecl(input)
	if err != nil {
		t.Fatalf("NewFuncDecl failed: %v", err)
	}
	// todo:(zzy) update linkname & go name
	comparePackageOutput(t, pkg, `
package testpkg

import "github.com/goplus/llgo/c"

//go:linkname Foo C.foo
func Foo(a c.Int, a c.Ulong) float64`)
}

func comparePackageOutput(t *testing.T, pkg *genpkg.Package, expect string) {
	t.Helper()
	gogenPkg := pkg.GetGogenPackage()
	var buf bytes.Buffer
	err := gogenPkg.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo failed: %v", err)
	}
	actual := strings.TrimSpace(buf.String())
	expect = strings.TrimSpace(expect)
	if actual != expect {
		t.Errorf("unexpected output:\n%s\nexpected get:\n%s", actual, expect)
	}
}
