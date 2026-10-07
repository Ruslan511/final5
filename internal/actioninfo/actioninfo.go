package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		if v == "" {
			log.Println("нет значения")
			continue
		}
		err := dp.Parse(v)
		if err != nil {
			log.Println("не парсится")
			continue
		}
		info, erro := dp.ActionInfo()
		if erro != nil {
			log.Println("не разбирается строка")
			continue
		}
		fmt.Println(info)
	}

}
