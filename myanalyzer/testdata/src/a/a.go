package a

import "fmt"

func notogawa() {
	{
		foo := 0
		for foo := foo; foo < 3; foo++ { // want "foo found"
			fmt.Println(foo)
		}
		for foo := foo; foo < 3; foo++ { // want "foo found"
			fmt.Println(foo)
		}
		for foo := foo; foo < 3; foo++ { // want "foo found"
			foo := foo
			fmt.Println(foo)
		}
		for foo := foo; foo < 3; foo++ { // want "foo found"
			fmt.Println(&foo) // want "unary expr found"
		}
		for foo := &foo; *foo < 3; *foo++ { // want "foo found"
			fmt.Println(*foo) // TODO: これも通る
		}
		for foo := foo; foo < 3; foo++ { // want "foo found"
			foo := foo
			fmt.Println(&foo)
		}
	}
	{
		foo := 0
		for ; foo < 3; foo++ {
			fmt.Println(foo)
		}
		foo = 0
		for ; foo < 3; foo++ {
			fmt.Println(foo)
		}
		foo = 0
		for ; foo < 3; foo++ {
			foo := foo
			fmt.Println(foo)
		}
		foo = 0
		for ; foo < 3; foo++ {
			fmt.Println(&foo)
		}
		foo = 0
		for ; foo < 3; foo++ {
			foo := foo
			fmt.Println(&foo)
		}
	}
}
