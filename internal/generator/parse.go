package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type StructField struct {
	Name     string
	Type     string
	Tag      string
	Children []*StructField
}

type StructInfo struct {
	PackageName string
	Name        string
	Fields      []*StructField
}

// parseStructs は指定されたディレクトリ内の構造体情報を解析します。
func parseStructs(targetDir string) ([]*StructInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, targetDir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var structs []*StructInfo

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			packageName := pkg.Name
			ast.Inspect(file, func(n ast.Node) bool {
				ts, ok := n.(*ast.TypeSpec)
				if ok {
					st, ok := ts.Type.(*ast.StructType)
					if ok {
						structInfo := &StructInfo{
							PackageName: packageName,
							Name:        ts.Name.Name,
							Fields:      parseFields(st.Fields),
						}
						structs = append(structs, structInfo)
					}
				}
				return true
			})
		}
	}

	return structs, nil
}

func parseFields(fieldList *ast.FieldList) []*StructField {
	var fields []*StructField
	for _, field := range fieldList.List {
		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			// 埋め込みフィールドの場合
			fieldName = ""
		}

		fieldType := getTypeString(field.Type)
		var children []*StructField
		if structType, ok := field.Type.(*ast.StructType); ok {
			// ネストされた構造体の場合、再帰的にフィールドを解析
			children = parseFields(structType.Fields)
		}

		tag := ""
		if field.Tag != nil {
			tag = field.Tag.Value
		}

		fields = append(fields, &StructField{
			Name:     fieldName,
			Type:     fieldType,
			Tag:      tag,
			Children: children,
		})
	}
	return fields
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getTypeString(t.X), t.Sel.Name)
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getTypeString(t.Key), getTypeString(t.Value))
	case *ast.StructType:
		return "struct"
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return ""
	}
}
