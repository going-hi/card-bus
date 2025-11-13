package menu

import (
	"card-bus/map-step"
)

func GetMapFirstStep() *mapstep.MapStep {
	return mapstep.CreateMapStep(
		[]string{
			"Первый раунд. Какая будет следующая карта?",
			"1. Красная",
			"2. Черная",
			"3. Выйти"},
		map[string]any{
			"1": "Red",
			"2": "Black",
			"3": "Exit",
		})
}

func GetMapSecondStep() *mapstep.MapStep {
	return mapstep.CreateMapStep(
		[]string{
			"Второй раунд. Следующая карту будет выше или ниже предыдущей?",
			"1. Выше",
			"2. Ниже",
			"3. Выйти",
		},
		map[string]any{
			"1": -1,
			"2": 1,
		},
	)
}

func GetMapThirthStep() *mapstep.MapStep {
	return mapstep.CreateMapStep([]string{
		"Третий	раунд. Следующая карта будет между или за пределами?",
		"1. Между",
		"2. За пределами",
		"3. Выйти",
	}, map[string]any{
		"1": true,
		"2": false,
	},
	)
}

func GetMapFourthStep() *mapstep.MapStep {
	return mapstep.CreateMapStep([]string{
		"Четвертый раунд. Какая масть будет у следующей карты?",
		"1. Черви ♥",
		"2. Крести ♣",
		"3. Бубны ♦",
		"4. Пики ♠",
		"5. Выйти",
	},
		map[string]any{
			"1": "hearts",
			"2": "clubs",
			"3": "diamonds",
			"4": "spades"},
	)
}
