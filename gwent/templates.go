package gwent

import (
	"encoding/json"
	"io/ioutil"
)

var allTemplates []CardTemplate

func init() {
	allTemplates = make([]CardTemplate, 0)

	templateID := int64(0)
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
					templateID++
					template.ID = templateID
					allTemplates = append(allTemplates, template)
				}
			}
		}
	}

}
