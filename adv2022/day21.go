package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day21() {
	sc := bufio.NewScanner(os.Stdin)
	type aExpr struct {
		op   string
		args []string
	}
	expr := make(map[string]*aExpr)
	val := make(map[string]float64)
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		monkey := items[0][:len(items[0])-1]
		if len(items) == 2 {
			val[monkey], _ = strconv.ParseFloat(items[1], 64)
			continue
		} else {
			args := []string{items[1], items[3]}
			expr[monkey] = &aExpr{items[2], args}
		}
	}
	var f func(string) float64
	f = func(node string) float64 {
		e := expr[node]
		if e == nil {
			return float64(val[node])
		}
		v := make([]float64, 2)
		for i, arg := range e.args {
			v[i] = f(arg)
			val[arg] = v[i]
		}
		var ve float64
		switch e.op {
		case "+":
			ve = v[0] + v[1]
		case "-":
			ve = v[0] - v[1]
		case "*":
			ve = v[0] * v[1]
		case "/":
			ve = v[0] / v[1]
		}
		val[node] = ve
		return ve
	}
	fmt.Println(int(f("root")))
	r := expr["root"]
	r.op = "-"
	xStr, yStr := r.args[0], r.args[1]
	y := f(yStr)
	// value of humn only influences lhs of root operation
	g := func(i float64) float64 {
		val["humn"] = i
		return f(xStr) - y
	}
	x := lib.SecantSearch(g, -1e6, 1e6, 0.000001)
	fmt.Println(int(x))
}
