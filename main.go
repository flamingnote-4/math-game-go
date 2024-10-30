package main

import (
	"fmt"
	"math-game/domain"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       int = 100
	pointsPerQuestion int = 100
)

var id uint64 = 1

func main() {
	var userList []domain.User
	fmt.Println("Вітаємо у грі!")
	menu()

	for {
		choice := ""
		fmt.Scan(&choice)
		switch choice {
		case "1":
			user := play()
			userList = append(userList, user)
		case "2":
			for _, user := range userList {
				fmt.Printf("[%v]\tName: %s\tTime: %v\n",
					user.Id, user.Name, user.TimeSpent)
			}
		case "3":
			return
		default:
			fmt.Println("Оберіть опцію 1, 2 або 3.")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 3; i >= 1; i-- {
		fmt.Printf("Гра почнеться через: %v\n", i)
		time.Sleep(time.Second)
	}

	startTime := time.Now()
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Введіть число.")
			continue
		}
		if ansInt == x+y {
			myPoints += pointsPerQuestion
			fmt.Printf("Отримано %v балів. Залишилося набрати %v\n", pointsPerQuestion, totalPoints-myPoints)
		} else {
			fmt.Println("Неправильна відповідь.")
		}
	}

	timeSpent := time.Since(startTime)

	fmt.Println("Вітаю! Остаточний час: ", timeSpent)
	fmt.Println("Введіть ім'я: ")

	name := ""
	fmt.Scan(&name)
	user := domain.User{
		Id:        id,
		Name:      name,
		TimeSpent: timeSpent,
	}
	id++

	return user
}
