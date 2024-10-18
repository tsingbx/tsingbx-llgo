package visitor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goplus/llgo/chore/llcppg/ast"
)

type DocVisitor interface {
	VisitStart(docPath string)
	Visit(node ast.Node)
	VisitFuncDecl(funcDecl *ast.FuncDecl)
	VisitDone(docPath string)
	VisitStruct(structName *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl)
	//VisitClass(className *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl)
	//VisitMethod(className *ast.Ident, method *ast.FuncDecl, typeDecl *ast.TypeDecl)
	VisitUnion(unionName *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl)
	VisitEnumTypeDecl(enumTypeDecl *ast.EnumTypeDecl)
	VisitTypedefDecl(typedefDecl *ast.TypedefDecl)
}

type BaseDocVisitor struct {
	DocVisitor
}

func NewBaseDocVisitor(Visitor DocVisitor) *BaseDocVisitor {
	return &BaseDocVisitor{DocVisitor: Visitor}
}

func (p *BaseDocVisitor) visitNode(decl ast.Node) {
	switch v := decl.(type) {
	case *ast.FuncDecl:
		p.visitFuncDecl(v)
	case *ast.TypeDecl:
		p.visitTypeDecl(v)
	case *ast.EnumTypeDecl:
		p.visitEnumTypeDecl(v)
	case *ast.TypedefDecl:
		p.visitTypedefDecl(v)
	default:
		panic(fmt.Errorf("todo visit %v", v))
	}
}

func (p *BaseDocVisitor) Visit(node ast.Node) {
	switch v := node.(type) {
	case *ast.File:
		for _, decl := range v.Decls {
			p.visitNode(decl)
		}
	default:
		p.visitNode(v)
	}
}

func (p *BaseDocVisitor) visitFuncDecl(funcDecl *ast.FuncDecl) {
	if funcDecl == nil {
		return
	}
	p.VisitFuncDecl(funcDecl)
}

func (p *BaseDocVisitor) visitTypeDecl(typeDecl *ast.TypeDecl) {
	if typeDecl == nil || typeDecl.Name == nil || len(typeDecl.Name.Name) <= 0 {
		return
	}
	if typeDecl.Type.Tag == ast.Class {
		p.visitClass(typeDecl.Name, typeDecl.Type.Fields, typeDecl)
		for _, method := range typeDecl.Type.Methods {
			p.visitMethod(typeDecl.Name, method, typeDecl)
		}
	} else if typeDecl.Type.Tag == ast.Struct {
		p.visitStruct(typeDecl.Name, typeDecl.Type.Fields, typeDecl)
	} else if typeDecl.Type.Tag == ast.Union {
		p.visitUnion(typeDecl.Name, typeDecl.Type.Fields, typeDecl)
	}
}

func (p *BaseDocVisitor) visitClass(className *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl) {
	//p.VisitClass(className, fields, typeDecl)
}

func (p *BaseDocVisitor) visitMethod(className *ast.Ident, method *ast.FuncDecl, typeDecl *ast.TypeDecl) {
	//p.VisitMethod(className, method, typeDecl)
}

func (p *BaseDocVisitor) visitStruct(structName *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl) {
	p.VisitStruct(structName, fields, typeDecl)
}

func (p *BaseDocVisitor) visitUnion(unionName *ast.Ident, fields *ast.FieldList, typeDecl *ast.TypeDecl) {
	p.VisitUnion(unionName, fields, typeDecl)
}

func (p *BaseDocVisitor) visitEnumTypeDecl(enumTypeDecl *ast.EnumTypeDecl) {
	if enumTypeDecl == nil || enumTypeDecl.Name == nil || len(enumTypeDecl.Name.Name) <= 0 {
		return
	}
	p.VisitEnumTypeDecl(enumTypeDecl)
}

func (p *BaseDocVisitor) visitTypedefDecl(typedefDecl *ast.TypedefDecl) {
	if typedefDecl == nil {
		return
	}
	switch v := typedefDecl.Type.(type) {
	case *ast.RecordType:
		name := NewRandIdent()
		typeDecl := NewTypeDeclWithName(name, v)
		p.visitTypeDecl(typeDecl)
		typeDef := NewTypedefDecl(v.Tag, name)
		p.visitTypedefDecl(typeDef)
	case *ast.EnumType:
		name := NewRandIdent()
		enumTypeDecl := NewEnumTypeDecl(name, v)
		p.visitEnumTypeDecl(enumTypeDecl)
		typeDef := NewTypedefDecl(ast.Enum, name)
		p.visitTypedefDecl(typeDef)
	default:
		p.VisitTypedefDecl(typedefDecl)
	}
}

func NewRandIdent() *ast.Ident {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Int()
	name := "r" + fmt.Sprintf("%d", n)
	return &ast.Ident{Name: name}
}

func NewTypeDeclWithName(name *ast.Ident, recordType *ast.RecordType) *ast.TypeDecl {
	typDecl := &ast.TypeDecl{Name: name, Type: recordType}
	return typDecl
}

func NewEnumTypeDecl(name *ast.Ident, enumType *ast.EnumType) *ast.EnumTypeDecl {
	enumDecl := &ast.EnumTypeDecl{Name: name, Type: enumType}
	return enumDecl
}

func NewTagExpr(tag ast.Tag, name *ast.Ident) *ast.TagExpr {
	return &ast.TagExpr{Tag: tag, Name: name}
}

func NewTypedefDecl(tag ast.Tag, name *ast.Ident) *ast.TypedefDecl {
	tagExpr := NewTagExpr(tag, name)
	typedefDecl := &ast.TypedefDecl{Name: name, Type: tagExpr}
	return typedefDecl
}
