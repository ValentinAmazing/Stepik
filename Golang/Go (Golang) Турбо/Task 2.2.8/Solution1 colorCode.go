package main

import "fmt"

var colorCode = map[string]int{
    "красный": 1,
    "белый": 2,
    "желтый": 4,
    "синий": 8,
}

var mixedColor = map[int]string{
	2: "красный",
	3: "розовый",
	4: "белый",
	5: "оранжевый",
	6: "зеленый",
	8: "желтый",
	16: "синий",
}

func main() {
    var color1, color2 string
    fmt.Scan(&color1, &color2)
    
	if val, ok := mixedColor[colorCode[color1] + colorCode[color2]]; ok{
		fmt.Println(val)
	} else {
		fmt.Println("Неизвестный цвет")
	}
}
