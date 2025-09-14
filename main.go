package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func startTraining(charName, charClass string) string {
	if charClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	}

	if charClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	}

	if charClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var playerSelect string
	for playerSelect != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &playerSelect)

		if playerSelect == "attack" {
			fmt.Println(attack(charName, charClass))
		}

		if playerSelect == "defence" {
			fmt.Println(defence(charName, charClass))
		}

		if playerSelect == "special" {
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func chooseClass() string {
	var approve_choice string
	var charClass string

	for approve_choice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)

		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("Некорректный ввод. Попробуй еще раз.")
			continue
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approve_choice)
		approve_choice = strings.ToLower(approve_choice)
	}
	return charClass
}

func attack(charName, charClass string) string {

	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randomize(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randomize(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randomize(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}
}

// обратите внимание на "if else" и на "else"
func defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randomize(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randomize(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randomize(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	default:
		return "неизвестный класс персонажа"
	}
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := chooseClass()

	fmt.Println(startTraining(charName, charClass))
}

func randomize(min, max int) int {
	return rand.Intn(max-min) + min
}
