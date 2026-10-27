package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает поведение для данных о тренировках и прогулках.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку dataset и выводит информацию об активности.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("ошибка парсинга данных: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
