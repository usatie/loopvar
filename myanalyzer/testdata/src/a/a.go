package a

import "fmt"

func testFindPointerOfLoopVar() {
	/*
		{
			for foo := 0 ; foo < 3 ; foo++ {
				bar := &foo
				f(foo) // OK
				g(&foo) // NG

				f(*bar) // OK
				g(bar) // NG

				bar = nil
				g(bar) // OK
			}
		}
	*/
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
			fmt.Println(*foo)
		}
		for foo := foo; foo < 3; foo++ { // want "foo found"
			_ = &foo
			foo = -foo
			foo = +foo
			foo--
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
