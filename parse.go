package main

import (
	"reflect"
	"text/template"
	"text/template/parse"

	"github.com/gopherjs/gopherjs/js"
)

func format(node parse.Node) map[string]interface{} {
	if reflect.ValueOf(node).IsNil() {
		return nil
	}

	m := map[string]interface{}{
		"type": reflect.ValueOf(node).Type().Elem().Name(),
		"pos":  node.Position(),
	}

	switch node := node.(type) {
	case *parse.ActionNode:
		m["pipe"] = format(node.Pipe)
	case *parse.BoolNode:
		m["true"] = node.True
	case *parse.ChainNode:
		m["field"] = node.Field
	case *parse.CommandNode:
		args := []interface{}{}
		for _, n := range node.Args {
			args = append(args, format(n))
		}
		m["args"] = args
	case *parse.DotNode:
	case *parse.FieldNode:
		m["ident"] = node.Ident
	case *parse.IdentifierNode:
		m["ident"] = node.Ident
	case *parse.IfNode:
		m["pipe"] = format(node.Pipe)
		m["list"] = format(node.List)
		m["elseList"] = format(node.ElseList)
	case *parse.ListNode:
		nodes := []interface{}{}
		for _, n := range node.Nodes {
			nodes = append(nodes, format(n))
		}
		m["nodes"] = nodes
	case *parse.NilNode:
	case *parse.NumberNode:
		m["value"] = node.String()
	case *parse.PipeNode:
		cmds := []interface{}{}
		for _, n := range node.Cmds {
			cmds = append(cmds, format(n))
		}
		m["cmds"] = cmds
		decl := []interface{}{}
		for _, n := range node.Decl {
			decl = append(decl, format(n))
		}
		m["decl"] = decl
	case *parse.RangeNode:
		m["pipe"] = format(node.Pipe)
		m["list"] = format(node.List)
		m["elseList"] = format(node.ElseList)
	case *parse.StringNode:
		m["text"] = node.Text
	case *parse.TemplateNode:
		m["name"] = node.Name
		m["pipe"] = format(node.Pipe)
	case *parse.TextNode:
		m["text"] = string(node.Text)
	case *parse.VariableNode:
		m["ident"] = node.Ident
	case *parse.WithNode:
		m["pipe"] = format(node.Pipe)
		m["list"] = format(node.List)
		m["elseList"] = format(node.ElseList)
	}
	return m
}

func Parse(input string) interface{} {
	return format(template.Must(template.New("").Parse(input)).Root)
}

func main() {
	js.Module.Set("exports", Parse)
}
