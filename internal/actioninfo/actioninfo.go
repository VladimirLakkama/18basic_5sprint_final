package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Parsing error (string %d): %v\n", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Format error (string %d): %v\n", i+1, err)
			continue
		}

		log.Println(info)
	}
}
