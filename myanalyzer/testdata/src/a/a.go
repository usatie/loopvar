package a

import "fmt"

// for 文を見つける
func findFor() {
	for findForVar := 0; findForVar < 3; findForVar++ {
		fmt.Println(findForVar) // want "for found"
	}
}

func pointer() {
	for findForVar := 0; findForVar < 3; findForVar++ {
		fmt.Println(&findForVar) // want "and used in for"
	}

	for findForVar := 0; findForVar < 3; findForVar++ {
		fmt.Println(findForVar) // OK
	}
}
