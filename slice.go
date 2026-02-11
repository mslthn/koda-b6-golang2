package main

import (
	"fmt"
)

// 1 print string "Koda" dgn we.are.the.best
type the struct {
	best string
}

type are struct {
	the the
}

type we struct {
	are are
}

// 2
type hello struct {
	world string
}

//3 obj.str[3][1][2].man[0].tech.academy -> Tech Academy

type tech struct {
	academy string
}
type man [0]string

type str [3][1][2]string

type obj struct {
	str str
}

//4 akses array my[0].favourite[3].fruit.is
//5 print numebr 32 ketika kalkulasi num.first[1] + num.second[2]

func slice() {
	//3. obj.str[3][1][2].man[0].tech.academy -> Tech Academy

	// obj:=obj{
	// 	str{
	// 		{},
	// 		{},
	// 		{

	// 			[2]struct{
	// 				{},
	// 				{
	// 					man[
	// 						tech:tech{
	// 							academy: "Tech Academy",
	// 						}],
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	//2
	hello := hello{
		world: "Hello World",
	}

	//1
	we := we{
		are: are{
			the: the{
				best: "Koda",
			},
		},
	}

	fmt.Println(we.are.the.best)
	fmt.Println(hello.world)
	// fmt.Println(obj.str[3][1][2].man[0].tech.academy)
}
