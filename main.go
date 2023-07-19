package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var lvl uint32
	var flvl uint32
	var expirience uint32
	var inPayDay uint32
	var payDaysCount uint32
	var output uint32

	clearConsole()

	fmt.Printf("Введите желаемый уровень: ")
	lvl = readUint32FromConsole()
	fmt.Printf("Введите ваш уровень: ")
	flvl = readUint32FromConsole()
	fmt.Printf("Введите ваше количество опыта: ")
	expirience = readUint32FromConsole()
	fmt.Printf("Сколько опыта вы получаете в PayDay?: ")
	inPayDay = readUint32FromConsole()
	output = toExp(lvl) - (toExp(flvl) + expirience)
	payDaysCount = PayDayCount(output, inPayDay)

	clearConsole()

	fmt.Println("Ваш уровень: ", flvl)
	fmt.Println("Ваше количество опыта: ", expirience)
	fmt.Println("Желаемый уровень: ", lvl)
	fmt.Println("Вам понадобится ", output, "опыта, или же ", payDaysCount, " PayDay'ев.")
}

func toExp(level uint32) uint32 {
	if level == 1 {
		return 0
	}
	return level*4 + toExp(level-1)
}

func readUint32FromConsole() uint32 {
	var inputValue int

	for {
		_, err := fmt.Scan(&inputValue)

		if err != nil {
			continue
		}

		if inputValue < 0 {
			continue
		}

		return uint32(inputValue)
	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux", "freebsd", "openbsd", "netbsd":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println()
		fmt.Println()
	}
}

func PayDayCount(exp uint32, inPayDay uint32) uint32 {
	if exp%inPayDay == 0 {
		return exp / inPayDay
	}
	return (exp / inPayDay) + 1
}
