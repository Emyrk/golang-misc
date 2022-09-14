package paramlang_test

import (
	"fmt"
	"math/rand"
	"testing"

	tassert "github.com/stretchr/testify/assert"

	"coder.com/m/lib/go/crand"
	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang"
)

// TestMathConditionals tests all sorts of integer and float conditionals.
// It will also test all sorts of conditional chaining with logic (and, or).
// So this also tests and/or logic
func TestMathConditionals(t *testing.T) {
	t.Parallel()
	g := newConditionalGenerator()
	for i := 0; i < 1000; i++ {
		cond := g.Conditional(crand.MustIntn(9) + 1)
		name := cond.Str

		v, err := paramlang.EvalBool(fmt.Sprintf("{{ %s }}", cond.Str), nil)
		tassert.NoError(t, err, fmt.Sprintf("'%s' success", name))
		tassert.Equal(t, cond.Res, v, fmt.Sprintf("'%s' exp", name))
	}
}

type conditionalVector struct {
	Str string
	Res bool
}

type conditonalGenerator struct {
	conds []conditionalVector
}

// Conditional produces a random conditional vector from the generator containing a number
// of conditionals equal to "length'
// Eg: (721 == 855 or 627 != 534 or 820.4604 == 419.0978) or 328.5773 <= 701.2268)
func (g conditonalGenerator) Conditional(length int) conditionalVector {
	if length < 1 {
		panic("cannot get a conditional of length less than 1")
	}
	conds := []conditionalVector{g.cond()}
	for i := 1; i < length; i++ {
		conds = append(conds, g.cond())
	}

	// Now merge the conditionals in random groups.
	// Each group is grouped with '()', so we want a mixture of grouping types.
	// Eg:
	// 	((a and b) and c)
	//	(a and (b and c))
	//  (a and b and c)
	for len(conds) > 1 {
		var combine []conditionalVector
		// Remove up to 4 conditionals and combine at a time
		remove := (crand.MustIntn(len(conds)) % 3) + 2
		if remove > len(conds) {
			remove = len(conds)
		}
		conds, combine = removeX(conds, remove)
		combined := randomLogicJoin(combine...)
		conds = append(conds, combined)
	}
	return conds[0]
}

// randomLogicJoin will join any number of conditionals
func randomLogicJoin(vectors ...conditionalVector) conditionalVector {
	if len(vectors) == 0 {
		panic("cannot join 0 vectors")
	}

	logics := []struct {
		str string
		op  func(a, b bool) bool
	}{
		{"and", func(a, b bool) bool {
			return a && b
		}},
		{"or", func(a, b bool) bool {
			return a || b
		}},
	}

	init := vectors[0]
	for i := 1; i < len(vectors); i++ {
		vect := vectors[i]
		logic := logics[crand.MustIntn(len(logics))]
		init = conditionalVector{
			Str: fmt.Sprintf("%s %s %s", init.Str, logic.str, vect.Str),
			Res: logic.op(init.Res, vect.Res),
		}
	}

	init.Str = "(" + init.Str + ")"
	return init
}

// removeX will remove x conditionals from the slice
func removeX(slice []conditionalVector, amt int) ([]conditionalVector, []conditionalVector) {
	var removed []conditionalVector
	for i := 0; i < amt; i++ {
		r := crand.MustIntn(len(slice))
		removed = append(removed, slice[r])
		slice = removeCondVector(slice, r)
	}
	return slice, removed
}

func removeCondVector(slice []conditionalVector, s int) []conditionalVector {
	newSlice := make([]conditionalVector, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}

// cond returns a random conditional
func (g conditonalGenerator) cond() conditionalVector {
	return g.conds[crand.MustIntn(len(g.conds))]
}

func newConditionalGenerator() *conditonalGenerator {
	g := new(conditonalGenerator)
	// Generate a bunch of integer and float operations in a big shuffled list
	g.conds = append(g.conds, integerConditionalsPerOp(50)...)
	g.conds = append(g.conds, floatConditionalsPerOp(50)...)
	rand.Shuffle(len(g.conds), func(i, j int) {
		g.conds[i], g.conds[j] = g.conds[j], g.conds[i]
	})

	return g
}

// integerConditionalsPerOp generates integer conditionals
// Eg: 5 < 9
func integerConditionalsPerOp(per int) []conditionalVector {
	ops := []struct {
		op   string
		comp func(a, b int) bool
	}{
		{"<", func(a, b int) bool { return a < b }},
		{">", func(a, b int) bool { return a > b }},
		{"==", func(a, b int) bool { return a == b }},
		{"<=", func(a, b int) bool { return a <= b }},
		{">=", func(a, b int) bool { return a >= b }},
		{"!=", func(a, b int) bool { return a != b }},
	}

	res := make([]conditionalVector, 0)
	for _, op := range ops {
		for i := 0; i < per; i++ {
			a, b := crand.MustIntn(1000), crand.MustIntn(1000)
			res = append(res, conditionalVector{
				Str: fmt.Sprintf("%d %s %d", a, op.op, b),
				Res: op.comp(a, b),
			})
		}
	}

	return res
}

// floatConditionalsPerOp generates float conditionals
// Eg: 5.5 < 9.1
func floatConditionalsPerOp(per int) []conditionalVector {
	ops := []struct {
		op   string
		comp func(a, b float64) bool
	}{
		{"<", func(a, b float64) bool { return a < b }},
		{">", func(a, b float64) bool { return a > b }},
		{"==", func(a, b float64) bool { return a == b }},
		{"<=", func(a, b float64) bool { return a <= b }},
		{">=", func(a, b float64) bool { return a >= b }},
		{"!=", func(a, b float64) bool { return a != b }},
	}

	res := make([]conditionalVector, 0)
	for _, op := range ops {
		for i := 0; i < per; i++ {
			a, b := crand.MustFloat64()*1000, crand.MustFloat64()*1000
			res = append(res, conditionalVector{
				Str: fmt.Sprintf("%.4f %s %.4f", a, op.op, b),
				Res: op.comp(a, b),
			})
		}
	}

	return res
}
