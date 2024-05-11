package main

import (
	"fmt"
	"go/format"
	"strings"
)

func main() {
	//num := 5
	//
	//for i := range make([]int, num) {
	//	n := i + 1
	//
	//	t := makeType(n)
	//	makeFunc(n)
	//
	//	fmt.Println(t)
	//}

	out := makeType(3) + "\n" + makeFunc(3)

	fmt.Println(out)

	code, err := format.Source([]byte(out))

	if err != nil {
		text := string(code)

		fmt.Println(text)
	} else {
		fmt.Println(err)
		fmt.Println("BAD!!!!!")
	}

	//makeFunc(3)
}

func makeType(num int) string {
	r := make([]int, num)

	typeArgs := "\n"
	var inner string

	for i := range r {
		n := i + 1
		typeArgs += fmt.Sprintf("T%d any,\n", n)
		inner += fmt.Sprintf("	T%d T%d\n", n, n)
	}

	code := fmt.Sprintf("type And%dResult[%s] struct {\n	%s\n}", num, typeArgs, strings.TrimSpace(inner))

	return code
}

func makeFunc(num int) string {
	name := fmt.Sprintf("And%d", num)

	typeArg := "["
	args := "("
	resType := name + "Result["

	block := ""

	for i := range make([]int, num) {
		n := i + 1

		typeArg += fmt.Sprintf("T%d any, ", n)
		args += fmt.Sprintf("p%d Parser[T%d], ", n, n)
		resType += fmt.Sprintf("T%d, ", n)

		block += fmt.Sprintf(
			`		res%d, ok%d := p%d(in)
		if ok%d {
`,
			n, n, n, n,
		)
	}

	typeArg = strings.TrimRight(typeArg, ", ")
	typeArg += "]"
	args = strings.TrimRight(args, ", ")
	args += ")"
	resType = strings.TrimRight(resType, ", ")
	resType += "]"

	success := "return " + resType + "{\n"
	closing := ""

	for i := range make([]int, num) {
		n := i + 1

		success += fmt.Sprintf("T%d: res%d,\n", n, n)
		closing += "}"
	}

	success += "}, true"

	res := fmt.Sprintf(
		`func %s%s%s Parser[%s] {
	return func(in *Input) (%s, bool) {
		start := in.Position

%s
%s
%s
	in.SetPosition(start)
	return %s{}, false
	}
}`,
		name, typeArg, args, resType, resType, block, success, closing, resType,
	)

	return res
}
