package main

import (
	"net/http"

	"fmt"

	"github.com/teerkevin23/rundoo/internal/version"
)

func (app *application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"Version": version.Get(),
	}

	return data
}

func (app *application) backgroundTask(fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.reportError(fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			app.reportError(err)
		}
	}()
}
