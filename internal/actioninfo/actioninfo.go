package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

// DataParser описывает интерфейс для структур,
// которые разбирают входные данные
// и формируют строку с информацией об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор строк с данными об активности.
//
// Принимает:
//   - []string — слайс строк с данными;
//   - DataParser — структуру, реализующую интерфейс.
//
// Ошибки логируются.
func Info(dataset []string, dp DataParser) {
	for _, line := range dataset {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		err := dp.Parse(line)
		if err != nil {
			log.Printf("ошибка разбора: %v (строка: %s)", err, line)
			continue
		}
		info, ierr := dp.ActionInfo()
		if ierr != nil {
			log.Printf("ошибка получения информации о действии: %v (строка: %s)", ierr, line)
			continue
		}
		fmt.Println(info)
	}
}
