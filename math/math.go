package main

import (
	"errors"
	"fmt"
	_ "github.com/alex-ant/gomath/gaussian-elimination"
	_ "github.com/alex-ant/gomath/rational"
	"github.com/crsmithdev/goexpr"
	"go/ast"
	"go/token"
	_ "log"
	"reflect"
	"strconv"
)

func main() {
	e,err:=goexpr.Parse("")
	println(e)
	println(err)

	/*y := 16.7
	s := "(x1+1)*3+10-11.5/5+x1*2"
	e, _ := goexpr.Parse(s)


	equation, err := test2(e.Ast)
	if err != nil {
		println(err.Error())
		return
	}

	result := (y - equation.B) / equation.A
	println(result)*/

}

// TODO 支持x不在分母的一元一次方程求解，即可以化简为ax+b的方程

type Equation struct {
	A float64
	B float64
}

func test2(node ast.Node) (*Equation, error) {
	switch node.(type) {
	case *ast.Ident:
		return &Equation{
			A: 1,
			B: 0,
		}, nil
	case *ast.BinaryExpr:
		n := node.(*ast.BinaryExpr)
		xe, err := test2(n.X)
		if err != nil {
			return nil, err
		}
		ye, err := test2(n.Y)
		if err != nil {
			return nil, err
		}
		switch n.Op {
		case token.ADD:
			return &Equation{
				A: xe.A + ye.A,
				B: xe.B + ye.B,
			}, nil
		case token.SUB:
			return &Equation{
				A: xe.A - ye.A,
				B: xe.B - ye.B,
			}, nil
		case token.MUL:
			if xe.A != 0 && ye.A != 0 {
				return nil, errors.New("只支持一次方程")
			}
			return &Equation{
				A: xe.A*ye.B + xe.B*ye.A,
				B: xe.B * ye.B,
			}, nil
		case token.QUO:
			if ye.A != 0 {
				return nil, errors.New("不支持分母存在变量")
			}
			return &Equation{
				A: xe.A / ye.B,
				B: xe.B / ye.B,
			}, nil
		default:
			return nil, fmt.Errorf("unsupported binary operation: %s", n.Op)
		}
	case *ast.ParenExpr:
		n := node.(*ast.ParenExpr)
		return test2(n.X)
	case *ast.BasicLit:
		n := node.(*ast.BasicLit)
		val, err := strconv.ParseFloat(n.Value, 64)
		if err != nil {
			return nil, err
		}
		return &Equation{
			A: 0,
			B: val,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported node %+v (type %+v)", node, reflect.TypeOf(node))
	}
}
