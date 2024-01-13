package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose // данные для сериализации и десериализации
	FileInput     string                // имя файла, который нужно перекодировать
	FileOutput    string                // имя файла с результатом перекодирования
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose // данные для сериализации и десериализации
	FileInput     string                // имя файла, который нужно перекодировать
	FileOutput    string                // имя файла с результатом перекодирования
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// 1. читаем файл j.FileInput с информацией в формате json в переменную jsonData []byte
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("Error in reading jsonFIle: %s\n", err.Error())
		return err
	}
	// fmt.Println(string(jsonData))

	// создаем переменную dockerCompose, чтобы конвертировать эти данные в YAML
	// var dockerCompose models.DockerCompose - можно было бы создать, но у нас есть поле типа DockerCompose у входных даннх j

	// 2. Десериализуем переменную jsonData []byte, то есть преобразуем в структуру models.DockerCompose
	err = json.Unmarshal(jsonData, &j.DockerCompose)
	if err != nil {
		fmt.Printf("Error in deserialisation (unmarshalling json data): %s\n", err.Error())
		return err
	}

	// 3. Имея структуру models.DockerCompose сериализуем данные из переменной dockerCompose в слайс байт
	// & означает, что мы передаём указатель на dockerCompose
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("Error in yaml serialising data: %s\n", err.Error())
		return err
	}

	// 4. Создаем файл
	file, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("Error in creating file: %s\n", err.Error())
		return err
	}
	defer file.Close()

	// 5. записываем слайс байт в файл
	_, err = file.Write(yamlData)
	if err != nil {
		fmt.Printf("Error in writing file: %s\n", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод

	// 1. читаем файл y.FileInput с информацией в формате yaml в переменную yamlData []byte
	var yamlData []byte
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("Error in reading y.FileInput: %s\n", err.Error())
		return err
	}

	// создаем переменную dockerCompose, чтобы конвертировать эти данные в JSON
	// var dockerCompose models.DockerCompose - можно было бы создать, но у нас есть поле типа DockerCompose у входных даннх j

	// 2. Десериализуем переменную yamlData []byte, то есть преобразуем в структуру models.DockerCompose?
	err = yaml.Unmarshal(yamlData, &y.DockerCompose)
	if err != nil {
		fmt.Printf("Error in deserialising (unmarshalling yaml data: %s\n)", err.Error())
		return err
	}

	// 3. Имея структуру models.DockerCompose сериализуем данные из переменной dockerCompose в слайс байт
	// & означает, что мы передаём указатель на dockerCompose
	out, err := json.MarshalIndent(y.DockerCompose, "", "    ")
	if err != nil {
		fmt.Printf("Error in serialising dockerCompose: %s\n", err.Error())
		return err
	}

	// 4. Создаем файл
	file, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("Error in opening file: %s\n", err.Error())
		return err
	}
	defer file.Close()

	// 5. записываем слайс байт в файл
	_, err = file.Write(out)
	if err != nil {
		fmt.Printf("Error in writing file %s\n", err.Error())
		return err
	}

	return nil
}
