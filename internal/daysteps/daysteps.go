package daysteps

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"errors"

	"dzpracticum/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")//Алгоритм 1.
	if len(parts) != 2{ //Алгоритм 2.
		return 0, 0, errors.New("Некорректный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])//Алгоритм 3.
	if err != nil{
		return 0, 0, err
	}
	if steps <= 0{//Алгоритм 4.
		return 0, 0, errors.New("Кол-во шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(parts[1])//Алгоритм 5.
    if err != nil{
		return 0, 0, err
    }
	return steps, duration, nil//Алгоритм 6.
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data) //Алгоритм 1.
	if err != nil{
		fmt.Println("Ошибка ", err)
		return ""
	} 
	if steps <= 0 {  //Алгоритм 2.
		return  ""
	}

	distanceM := float64(steps)*stepLength //Алгоритм 3.
	distanceKm := distanceM/mInKm //Алгоритм 4.

	calories := spentcalories.WalkingSpentCalories(weight, height, duration, distanceKm) //Алгоритм 5.

	result := fmt.Sprintf( //Алгоритм 6.
		"Количество шагов: %d. \nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distanceKm, calories
	)
	return result
}
