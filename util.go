package astra

import (
	"fmt"
	"go/ast"
	astparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"io/ioutil"

	"github.com/vetcher/go-astra/types"
)

// Opens and parses file by name and return information about it.
func ParseFile(filename string) (*types.File, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can not filepath.Abs: %v", err)
	}
	fset := token.NewFileSet()
	tree, err := astparser.ParseFile(fset, path, nil, astparser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error when parse file: %v", err)
	}
	info, err := ParseAstFile(tree)
	if err != nil {
		return nil, fmt.Errorf("error when parsing info from file: %v", err)
	}
	return info, nil
}

func ParseFileWithoutGOPATH(filename string) (*types.File, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can not filepath.Abs: %v", err)
	}
	fset := token.NewFileSet()
	tree, err := astparser.ParseFile(fset, path, nil, astparser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error when parse file: %v", err)
	}
	info, err := ParseAstFile(tree)
	if err != nil {
		return nil, fmt.Errorf("error when parsing info from file %s: %v", filename, err)
	}
	return info, nil
}

func MergeFiles(files []*types.File) (*types.File, error) {
	targetFile := &types.File{}
	for _, file := range files {
		if file == nil {
			continue
		}
		// do not merge documentation.
		targetFile.Base.Name = file.Base.Name
		targetFile.Imports = mergeImports(targetFile.Imports, file.Imports)
		targetFile.Constants = append(targetFile.Constants, file.Constants...)
		targetFile.Vars = append(targetFile.Vars, file.Vars...)
		targetFile.Interfaces = append(targetFile.Interfaces, file.Interfaces...)
		targetFile.Structures = append(targetFile.Structures, file.Structures...)
		targetFile.Methods = append(targetFile.Methods, file.Methods...)
		targetFile.Types = append(targetFile.Types, file.Types...)
	}
	err := linkMethodsToStructs(targetFile)
	if err != nil {
		return nil, err
	}
	return targetFile, nil
}

func ParsePackage(path string) ([]*types.File, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("can not filepath.Abs: %v", err)
	}
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, fmt.Errorf("can not read dir: %v", err)
	}
	var parsedFiles []*types.File
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		f, err := ParseFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf("can not parse %s: %v", file.Name(), err)
		}
		parsedFiles = append(parsedFiles, f)
	}
	return parsedFiles, nil
}

func ResolvePackagePath(outPath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", ErrGoPathIsEmpty
	}

	absOutPath, err := filepath.Abs(filepath.Dir(outPath))
	if err != nil {
		return "", err
	}

	gopathSrc := filepath.Join(gopath, "src")
	if !strings.HasPrefix(absOutPath, gopathSrc) {
		return "", ErrNotInGoPath
	}

	return absOutPath[len(gopathSrc)+1:], nil
}

func namesOfIdents(idents []*ast.Ident) (res []string) {
	for i := range idents {
		if idents[i] != nil {
			res = append(res, idents[i].Name)
		}
	}
	return
}

func mergeStringSlices(slices ...[]string) []string {
	if len(slices) == 0 {
		return nil
	}
	return append(slices[0], mergeStringSlices(slices[1:]...)...)
}

func parseCommentFromSources(opt Option, groups ...*ast.CommentGroup) []string {
	temp := make([][]string, len(groups))
	for i := range groups {
		temp[i] = parseComments(groups[i], opt)
	}
	return mergeStringSlices(temp...)
}
