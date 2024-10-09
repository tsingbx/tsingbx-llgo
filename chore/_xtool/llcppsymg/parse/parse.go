package parse

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/clang"
	"github.com/goplus/llgo/chore/_xtool/llcppsymg/clangutils"
)

type SymbolInfo struct {
	GoName    string
	ProtoName string
}

type SymbolProcessor struct {
	TrimPrefixes []string
	ReplPrefixes []string
	SymbolMap    map[string]*SymbolInfo
	CurrentFile  string
	NameCounts   map[string]int
}

func NewSymbolProcessor(TrimPrefixes []string, ReplPrefixes []string) *SymbolProcessor {
	return &SymbolProcessor{
		TrimPrefixes: TrimPrefixes,
		ReplPrefixes: ReplPrefixes,
		SymbolMap:    make(map[string]*SymbolInfo),
		NameCounts:   make(map[string]int),
	}
}

func (p *SymbolProcessor) setCurrentFile(filename string) {
	p.CurrentFile = filename
}

func TrimPrefix(org, prefix, repl string) (string, error) {
	exp, err := regexp.Compile(prefix)
	if err != nil {
		if strings.HasPrefix(org, prefix) {
			return strings.ReplaceAll(org, prefix, repl), nil
		}
		return org, fmt.Errorf("trim prefix fail")
	}
	result := exp.ReplaceAllString(org, repl)
	if result == org {
		return org, fmt.Errorf("trim prefix fail")
	}
	fmt.Println("org:", org, "prefix:", prefix, "repl:", repl, "result:", result)
	return result, nil
}

func (p *SymbolProcessor) DoTrimPrefixes(str string) string {
	for i, prefix := range p.TrimPrefixes {
		repl := ""
		if i < len(p.ReplPrefixes) {
			repl = p.ReplPrefixes[i]
		}
		result, err := TrimPrefix(str, prefix, repl)
		if err == nil {
			return result
		}
	}
	return str
}

func toTitle(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + (s[1:])
}

func toUpperCamelCase(originName string) string {
	if originName == "" {
		return ""
	}
	subs := strings.Split(string(originName), "_")
	name := ""
	for _, sub := range subs {
		name += toTitle(sub)
	}
	return name
}

// 1. remove prefix from config
// 2. convert to camel case
func (p *SymbolProcessor) ToGoName(name string) string {
	return toUpperCamelCase(p.DoTrimPrefixes(name))
}

func (p *SymbolProcessor) GenMethodName(class, name string, isDestructor bool) string {
	prefix := "(*" + class + ")."
	if isDestructor {
		return prefix + "Dispose"
	}
	if class == name {
		return prefix + "Init"
	}
	return prefix + name
}

func (p *SymbolProcessor) genGoName(cursor clang.Cursor) string {
	funcName := cursor.String()
	defer funcName.Dispose()

	originName := c.GoString(funcName.CStr())
	isDestructor := cursor.Kind == clang.CursorDestructor
	var convertedName string
	if isDestructor {
		convertedName = p.ToGoName(originName[1:])
	} else {
		convertedName = p.ToGoName(originName)
	}

	if parent := cursor.SemanticParent(); parent.Kind == clang.CursorClassDecl {
		parentName := parent.String()
		defer parentName.Dispose()
		class := p.ToGoName(c.GoString(parentName.CStr()))
		return p.AddSuffix(p.GenMethodName(class, convertedName, isDestructor))
	}
	return p.AddSuffix(convertedName)
}

func (p *SymbolProcessor) genProtoName(cursor clang.Cursor) string {
	displayName := cursor.DisplayName()
	defer displayName.Dispose()

	scopingParts := clangutils.BuildScopingParts(cursor.SemanticParent())

	var builder strings.Builder
	for _, part := range scopingParts {
		builder.WriteString(part)
		builder.WriteString("::")
	}

	builder.WriteString(c.GoString(displayName.CStr()))
	return builder.String()
}

func (p *SymbolProcessor) AddSuffix(name string) string {
	p.NameCounts[name]++
	if count := p.NameCounts[name]; count > 1 {
		return name + "__" + strconv.Itoa(count-1)
	}
	return name
}

func (p *SymbolProcessor) collectFuncInfo(cursor clang.Cursor) {
	symbol := cursor.Mangling()
	defer symbol.Dispose()

	symbolName := c.GoString(symbol.CStr())
	if len(symbolName) >= 1 && symbolName[0] == '_' {
		symbolName = symbolName[1:]
	}
	p.SymbolMap[symbolName] = &SymbolInfo{
		GoName:    p.genGoName(cursor),
		ProtoName: p.genProtoName(cursor),
	}
}

func (p *SymbolProcessor) visitTop(cursor, parent clang.Cursor) clang.ChildVisitResult {
	switch cursor.Kind {
	case clang.CursorNamespace, clang.CursorClassDecl:
		clangutils.VisitChildren(cursor, p.visitTop)
	case clang.CursorCXXMethod, clang.CursorFunctionDecl, clang.CursorConstructor, clang.CursorDestructor:
		loc := cursor.Location()
		file, _, _, _ := clangutils.GetLocation(loc)
		filename := file.FileName()
		defer filename.Dispose()

		isCurrentFile := c.Strcmp(filename.CStr(), c.AllocaCStr(p.CurrentFile)) == 0
		isPublicMethod := (cursor.CXXAccessSpecifier() == clang.CXXPublic) && cursor.Kind == clang.CursorCXXMethod || cursor.Kind == clang.CursorConstructor || cursor.Kind == clang.CursorDestructor

		if isCurrentFile && (cursor.Kind == clang.CursorFunctionDecl || isPublicMethod) {
			p.collectFuncInfo(cursor)
		}
	}
	return clang.ChildVisit_Continue
}

func ParseHeaderFile(files []string, TrimPrefixes []string, ReplPrefixes []string, isCpp bool, isTemp bool) (map[string]*SymbolInfo, error) {
	processer := NewSymbolProcessor(TrimPrefixes, ReplPrefixes)
	index := clang.CreateIndex(0, 0)
	for _, file := range files {
		_, unit, err := clangutils.CreateTranslationUnit(&clangutils.Config{
			File:  file,
			Temp:  isTemp,
			IsCpp: isCpp,
			Index: index,
		})
		if err != nil {
			return nil, errors.New("Unable to parse translation unit for file " + file)
		}
		cursor := unit.Cursor()
		if isTemp {
			processer.setCurrentFile(clangutils.TEMP_FILE)
		} else {
			processer.setCurrentFile(file)
		}
		clangutils.VisitChildren(cursor, processer.visitTop)
		unit.Dispose()
	}
	index.Dispose()
	return processer.SymbolMap, nil
}

func GenHeaderFilePath(cflags string, files []string) ([]string, error) {
	prefixPath := strings.TrimPrefix(cflags, "-I")

	var validPaths []string
	var errs []string

	for _, file := range files {
		if file == "" {
			continue
		}
		fullPath := filepath.Join(prefixPath, file)
		if f, err := os.Open(fullPath); err != nil {
			if os.IsNotExist(err) {
				errs = append(errs, fmt.Sprintf("file not found: %s", file))
			} else {
				errs = append(errs, fmt.Sprintf("error accessing file %s: %v", file, err))
			}
		} else {
			f.Close()
			validPaths = append(validPaths, fullPath)
		}
	}

	if len(validPaths) == 0 && len(errs) == 0 {
		return nil, fmt.Errorf("no valid header files")
	}

	if len(errs) > 0 {
		return validPaths, fmt.Errorf("some files not found or inaccessible: %v", errs)
	}

	return validPaths, nil
}
