package mapstep

import "errors"

type VariantMap = map[string]any

type MapStep struct {
	InputMenu       []string
	StepVariantsMap VariantMap
}

func CreateMapStep(inputMenu []string, stepVariantsMap VariantMap) *MapStep {
	return &MapStep{
		InputMenu:       inputMenu,
		StepVariantsMap: stepVariantsMap,
	}
}

func GetVariantMapStep[T any](m *MapStep, key string) (T, error) {
	raw, ok := m.StepVariantsMap[key]
	var zero T

	if !ok {
		return zero, errors.New("Такого аргумента нет")
	}

	val, ok := raw.(T)
	if !ok {
		return zero, errors.New("Тип значения не совпадает")
	}

	return val, nil
}
