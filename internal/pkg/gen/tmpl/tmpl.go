package tmpl

import (
	"html/template"
	"os"
)

func tmpl(filePath, textTmpl string, tmplData interface{}) error {

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(textTmpl)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, tmplData)
	if err != nil {
		return err
	}

	return nil
}
