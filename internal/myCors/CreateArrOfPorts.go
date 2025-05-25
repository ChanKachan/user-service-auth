package myCors

import (
	"errors"
	"strconv"
)

func CreateArrOfPorts(ports ...string) ([]string, error) {
	var result []string

	for _, port := range ports {
		result = append(result, "http://localhost:"+port+"/")
		for i := 1; i < 10; i++ {
			result = append(result, "http://localhost:"+port+"/")
			if len(port) == 4 && port[len(port)-1] == '0' {
				result = append(result, port[:len(port)-1]+strconv.Itoa(i))
			} else {
				return nil, errors.New("Порт должен состоять из 4-х чисел с '0' в конце")
			}
		}
	}
	return result, nil
}
