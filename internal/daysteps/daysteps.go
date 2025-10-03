package daysteps

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"errors"
	"log"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	if duration <= 0{
		return 0, 0, errors.New("Lkbntkmyjcnm ifuf ljk;yf ,snm ,jkmit 0")
    }
	return steps, duration, nil//Алгоритм 6.
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data) //Алгоритм 1.
	if err != nil{
		log.Println("Ошибка ", err)
		return ""
	} 
	if steps <= 0 {  //Алгоритм 2.
		return  ""
	}
 

	distanceM := float64(steps) * stepLength //Алгоритм 3.
	distanceKm := distanceM / mInKm //Алгоритм 4.

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration) //Алгоритм 5.
    if err != nil{
		log.Println("Ошибка ", err)
		return ""
	} 
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, calories)//Алгоритм 6.
	return result
}
