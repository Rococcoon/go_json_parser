package interpreter

import (
	"fmt"
	"rococcoon/go_json_test/ast"
)

func ConvertToGoValue(value ast.Value) interface{} {
	switch v := value.(type) {
	case *ast.StringLiteral:
		return v.Value
	case *ast.NumberLiteral:
		return v.Value
	case *ast.BooleanLiteral:
		return v.Value
	case *ast.NullLiteral:
		return nil
	case *ast.Object:
		result := make(map[string]interface{})
		for _, prop := range v.Pairs {
			result[prop.Key] = ConvertToGoValue(prop.Value)
		}
		return result
	case *ast.Array:
		var result []interface{}
		for _, elem := range v.Elements {
			result = append(result, ConvertToGoValue(elem))
		}
		return result
	default:
		fmt.Printf("Unknown type: %T\n", v)
		return nil
	}
}
