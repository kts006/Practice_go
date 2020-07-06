package something

import "fmt"

func MapMain() {
	person := map[string]string{"Name": "ABCD", "Age": "11"}
	// Age는 String으로 해야함, Map의 value를 string로 선언했으므로
	fmt.Println(person)

	for key, value := range person {
		fmt.Println("Key:", key, ", Value:", value)
	}
}

type personStruct struct {
	name string
	age  int
	desc []string
}

func StructMain() {
	perDesc := []string{"AAAAAAAAAAAA", "BBBBBBB"}
	// personStruture := personStruct{"Name", 18, perDesc} // 동작함, 아래가 더욱 명확함
	personStruture := personStruct{name: "Name", age: 18, desc: perDesc}
	fmt.Println(personStruture)
	fmt.Println(personStruture.name)
	fmt.Println(personStruture.age)
	fmt.Println(personStruture.desc)

}
