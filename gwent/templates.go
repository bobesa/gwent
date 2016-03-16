package gwent

import (
	"encoding/json"
	"io/ioutil"
)

var AllTemplates []CardTemplate

func init() {
	AllTemplates = make([]CardTemplate, 0)

	templateId := int64(0)
	directory := "./cards"

	files, _ := ioutil.ReadDir(directory)
	for _, file := range files {
		data, err := ioutil.ReadFile(directory + "/" + file.Name())
		if err == nil {
			var templates []CardTemplate
			err = json.Unmarshal(data, &templates)
			//Go trough all templates and assign IDs
			if err == nil {
				for _, template := range templates {
					templateId++
					template.Id = templateId
					AllTemplates = append(AllTemplates, template)
				}
			}
		}
	}

}
