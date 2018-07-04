package pkgtest

import "fmt"

func Maps(X int) {
	if X > 4 {
		fmt.Println("Not a correct index")
	}
	var makemap map[int]string
	makemap = make(map[int]string)
	makemap[0] = "Zero"
	makemap[1] = "One"
	makemap[2] = "Two"
	makemap[3] = "Three"
	makemap[4] = "Four"
	fmt.Println(makemap[4])
	fmt.Println(makemap[X])
	delete(makemap, 4)
	fmt.Println(makemap[4])
}

func Shorthandmap(str string) {
	map2 := map[string]string{
		"Jatin":   "Artwani",
		"Harshit": "Nanda",
		"Jatinn":  "Varshney",
	}
	if value, exists := map2[str]; exists {
		fmt.Println(value)
	}
}
