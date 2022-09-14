package paramlang_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang"
)

// TestParseSingle is nice to manually test the param language.
// Since it's not used anywhere in the code base, you can only execute it through tests.
// Will remove at some point.
func TestParseSingle(t *testing.T) {
	t.Skip()
	params := map[string]string{
		"params.cpu":  "1",
		"params.disk": "10",
		"params.cvm":  "true",
		"params.name": "wac-workspace",
	}

	var _ = params
	v, err := paramlang.EvalInt("{{ if 5a < 2 then 5 else 4 * ( 1 + 2 ) }}", nil)
	_, _ = fmt.Println(err, v)
}

type vectorSetType string

const (
	FloatSet   vectorSetType = "Float"
	IntegerSet vectorSetType = "Integer"
	BoolSet    vectorSetType = "Bool"
	StringSet  vectorSetType = "String"
)

// testParseSet will run the set of tests of the given type.
func testParseSet(t *testing.T, setName string, setType vectorSetType, set []evalTestVector) {
	wrap := func(in string) string {
		return fmt.Sprintf("%s%s%s",
			paramlang.StartTemplate, in, paramlang.EndTemplate)
	}

	// permutations returns potentially multiple strings that should all have
	// the same value.
	permutations := func(v evalTestVector) []string {
		if v.As == 0 {
			v.As = asBoth
		}
		var perms []string
		if v.As&asRaw != 0 {
			perms = append(perms, v.Input)
		}
		if v.As&asTpl != 0 {
			perms = append(perms, wrap(v.Input))
		}

		return perms
	}

	t.Run(setName, func(t *testing.T) {
		for _, v := range set {
			v.Name = fmt.Sprintf("[%s] %s", setName, v.Name)
			perms := permutations(v)
			for pi, p := range perms {
				var val interface{}
				var err error
				switch setType {
				case IntegerSet:
					val, err = paramlang.EvalInt(p, nil)
				case FloatSet:
					val, err = paramlang.EvalFloat(p, nil)
				case BoolSet:
					val, err = paramlang.EvalBool(p, nil)
				case StringSet:
					val, err = paramlang.EvalString(p, nil)
				default:
					t.Errorf("Cannot eval %q set", setName)
					t.FailNow()
				}

				if v.ErrorRegex == nil {
					name := fmt.Sprintf("test '%s' with perm=%d, string='%s', exp='%v'", v.Name, pi, p, v.Output)
					require.NoError(t, err, name)
					require.Equal(t, v.Output, val, name)
				} else {
					if err == nil {
						t.Errorf("test '%s' should have failed, no error", v.Name)
					} else {
						match := v.ErrorRegex.MatchString(err.Error())
						if !match {
							t.Errorf("Error regex did not match:\n Exp='%s'\nFound='%s'\n", v.ErrorRegex.String(), err.Error())
						}
					}
				}
			}
		}
	})
}

// evalTestMask allows a test vector to be ran as
// 	a template: `{{ 5 }}`
//  as raw    : `5`
// 	or as both
type evalTestMask uint8

const (
	asRaw evalTestMask = 1 << iota
	asTpl

	asBoth = asRaw | asTpl
)

type evalTestVectorSet struct {
	SetType vectorSetType
	Name    string
	Set     []evalTestVector
}

type evalTestVector struct {
	Name string
	// Input will be attempted as "input" or/and "{{ input }}" depending on
	// the 'As' mask.
	Input      string
	As         evalTestMask
	Output     interface{}
	ErrorRegex *regexp.Regexp
}

// Test_SimpleEval is a tabled set of test vectors for testing basic
// language behavior.
func Test_SimpleEval(t *testing.T) {
	sets := []evalTestVectorSet{
		{
			Name:    "Integer Literals",
			SetType: IntegerSet,
			Set: []evalTestVector{
				// Literals
				{Name: "Literal, 1", Input: "1", Output: 1},
				{Name: "Literal, 2", Input: "100", Output: 100},
				{Name: "Literal, 3", Input: "0", Output: 0},
				{Name: "Literal, 4", Input: "56", Output: 56},
				{Name: "Literal, 5", Input: "102", Output: 102},

				// TODO: Failures
				{Name: "BadLit 1", As: asTpl, Input: "1a", ErrorRegex: regexp.MustCompile("extraneous")},
			},
		},
		{
			Name:    "Integer Casting",
			SetType: IntegerSet,
			Set: []evalTestVector{
				// Castings
				{Name: "Casting, 1", As: asTpl, Input: "int(1.0)", Output: 1},
				{Name: "Casting, 2", As: asTpl, Input: "int(1.9)", Output: 1},
				{Name: "Casting, 3", As: asTpl, Input: "int(1e0)", Output: 1},
				{Name: "Casting, 4", As: asTpl, Input: "int(5.1 * 2.2)", Output: 11},
				{Name: "Casting, 5", As: asTpl, Input: "int(1.9) + 1", Output: 2},
				{Name: "Casting, 5", As: asTpl, Input: "int(1.9) + int(1.1) + int(float(1))", Output: 3},
			},
		},
		{
			Name:    "Float Literals",
			SetType: FloatSet,
			Set: []evalTestVector{
				// Literals
				{Name: "Literal, 1", Input: "1.0", Output: 1.0},
				{Name: "Literal, 2", Input: "1e10", Output: 1e10},
				{Name: "Literal, 3", Input: "0.1", Output: 0.1},
				{Name: "Literal, 4", Input: "0.0", Output: 0.0},
				{Name: "Literal, 5", Input: "1e2", Output: 100.0},

				// TODO: Failures
			},
		},
		{
			Name:    "Float Casting",
			SetType: FloatSet,
			Set: []evalTestVector{
				// Castings
				{Name: "Casting, 1", As: asTpl, Input: "float(1)", Output: float64(1)},
				{Name: "Casting, 2", As: asTpl, Input: "float(1.9)", Output: 1.9},
				{Name: "Casting, 3", As: asTpl, Input: "1.0 + float(1)", Output: float64(2)},
				{Name: "Casting, 4", As: asTpl, Input: "1.2 + 1.2", Output: 2.4},
			},
		},
		{
			Name:    "Boolean Literals",
			SetType: BoolSet,
			Set: []evalTestVector{
				// Literals :: TODO: Make failure cases for {{ t }} and others
				{Name: "Literal, 1", Input: "true", Output: true},
				{Name: "Literal, 2", As: asRaw, Input: "t", Output: true},
				{Name: "Literal, 3", As: asRaw, Input: "TRUE", Output: true},
				{Name: "Literal, 4", Input: "false", Output: false},
				{Name: "Literal, 5", As: asRaw, Input: "f", Output: false},
				{Name: "Literal, 6", As: asRaw, Input: "FALSE", Output: false},

				// TODO: Failures
			},
		},
		{
			Name:    "String Literals",
			SetType: StringSet,
			Set: []evalTestVector{
				// Literals
				{Name: "Literal, 1 (tmpled)", As: asTpl, Input: `"Hello World"`, Output: "Hello World"},
				{Name: "Literal, 2 (tmpled)", As: asTpl, Input: `"Dog"`, Output: "Dog"},
				{Name: "Literal, 3 (tmpled)", As: asTpl, Input: "\"Test\"", Output: "Test"},
				{Name: "Literal, 4", As: asRaw, Input: "Hello World", Output: "Hello World"},

				// TODO: Failures
			},
		},
		{
			Name:    "Basic conditionals",
			SetType: StringSet,
			Set: []evalTestVector{
				// Conditionals (hardcode)
				{Name: "Cond, literal 1", As: asTpl, Input: `if true then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, literal 2", As: asTpl, Input: `if false then "yes" else "no"`, Output: "no"},

				{Name: "Cond, int comp y1", As: asTpl, Input: `if 1 < 2 then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, int comp y3", As: asTpl, Input: `if 1 <= 2 then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, int comp y4", As: asTpl, Input: `if 2 <= 2 then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, int comp y5", As: asTpl, Input: `if 3 == 3 then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, int comp y6", As: asTpl, Input: `if 3 >= 3 then "yes" else "no"`, Output: "yes"},
				{Name: "Cond, int comp y7", As: asTpl, Input: `if 3 != 4 then "yes" else "no"`, Output: "yes"},

				{Name: "Cond, int comp n1", As: asTpl, Input: `if 2 < 1 then "yes" else "no"`, Output: "no"},
				{Name: "Cond, int comp n2", As: asTpl, Input: `if 3 != 3 then "yes" else "no"`, Output: "no"},
			},
		},
	}

	for _, set := range sets {
		testParseSet(t, set.Name, set.SetType, set.Set)
	}
}

func TestValidate(t *testing.T) {
	vectors := []struct {
		Input    string
		Validate func(input string) error
		Error    bool
	}{
		// Exp ints
		{"1 + 1", paramlang.ValidateInt, false},
		{"params.cpu * 4", paramlang.ValidateInt, false},
		{"if params.cvn then 5 else 2", paramlang.ValidateInt, false},
		{"params.cpu * 5 + 2 -1", paramlang.ValidateInt, false},

		// Exp bools
		{"params.cpu < params.memory", paramlang.ValidateBool, false},
		{"true", paramlang.ValidateBool, false},
		{"9 < int(8.2)", paramlang.ValidateBool, false},
		{"if params.cpu then 8 == 2 else 1 < params.memory", paramlang.ValidateBool, false},
		{"!params.cvm", paramlang.ValidateBool, false},

		// Exp Floats
		{"float(1) + 1e10", paramlang.ValidateFloat, false},
		{"params.cpu * 4.0", paramlang.ValidateFloat, false},
		{"if params.cvn then 5.1 else float(2)", paramlang.ValidateFloat, false},

		// Exp Strings
		{`"Hello World"`, paramlang.ValidateString, false},
		{`params.cpu`, paramlang.ValidateString, false},
		{`if params.cvm then "Hello" else "Goodbye"`, paramlang.ValidateString, false},
	}

	for _, vect := range vectors {
		err := vect.Validate(vect.Input)
		if vect.Error && err == nil {
			t.Errorf("expected an error for input %q", vect.Input)
		} else if !vect.Error && err != nil {
			t.Errorf("got error for input %q %+v", vect.Input, err)
		}
	}
}
