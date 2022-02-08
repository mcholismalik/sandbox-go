package replace

import (
	"fmt"
	"strings"
)

func New() {
	fmt.Println("string")

	str1 := "Hai kamu siapa ?, nama aku"
	str2 := strings.Replace(str1, "[NAME]", "Malik", 2)
	fmt.Println(str2)
}

func QuestionMark() {
	query := `insert into users (id, name, age) values (?,?,?)`
	str := strings.Replace(query, "?", "test", 1)
	fmt.Println(str)
}
