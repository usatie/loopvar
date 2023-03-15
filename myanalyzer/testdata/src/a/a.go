package a

import "fmt"

// for 文を見つける
func findFor() {
	for findForVar := 0; findForVar < 3; findForVar++ { // want "for found"
		fmt.Println(findForVar)
	}
}

//
