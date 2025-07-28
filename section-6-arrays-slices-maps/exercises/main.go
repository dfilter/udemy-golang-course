package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"violin", "tech", "history"}
	fmt.Println("My hobbies:", hobbies)
	// 2
	firstHobby := hobbies[0]
	fmt.Println("First Element:", firstHobby)
	lastTwoHobbies := hobbies[1:3]
	// lastTwoHobbies := hobbies[1:]
	fmt.Println("Second and third hobbies", lastTwoHobbies)
	// 3
	firstAndSecondHobbies := hobbies[0:2]
	// firstAndSecondHobbies := hobbies[:2]
	fmt.Println("First and second hobbies:", firstAndSecondHobbies)
	// 4
	firstAndSecondHobbies = firstAndSecondHobbies[1:3]
	// firstAndSecondHobbies = firstAndSecondHobbies[1:cap(firstAndSecondHobbies)]
	fmt.Println("Second and last hobbies:", firstAndSecondHobbies)

	// 5
	courseGoals := []string{"Learn Go", "Have fun"}
	fmt.Println("Goals slice", courseGoals)
	// 6
	courseGoals[1] = "Be Productive"
	courseGoals = append(courseGoals, "Use go for work")
	fmt.Println("Goals slice", courseGoals)

	// 7
	products := []Product{
		{
			id:    1,
			title: "Shoes",
			price: 50.99,
		},
		{
			id:    2,
			title: "Hat",
			price: 5.99,
		},
	}
	fmt.Println("Products", products)
	products = append(products, Product{id: 3, title: "T-Shirt", price: 6.99})
	fmt.Println(products)
}
