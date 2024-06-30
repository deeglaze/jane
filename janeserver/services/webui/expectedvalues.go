package webui

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"a10/operations"
	"a10/structures"
)

type evstruct struct {
	EV structures.ExpectedValue
	E  structures.Element
	I  structures.Intent
}

func showExpectedValues(c echo.Context) error {
	fmt.Println("here")
	es, _ := operations.GetExpectedValuesAll()

	evs := []evstruct{}

	for _, j := range es {
		e, _ := operations.GetElementByItemID(j.ElementID)
		i, _ := operations.GetIntentByItemID(j.IntentID)
		evs = append(evs, evstruct{j, e, i})
	}

	return c.Render(http.StatusOK, "evs.html", evs)
}

func showExpectedValue(c echo.Context) error {
	ev, _ := operations.GetExpectedValueByItemID(c.Param("itemid"))

	e, _ := operations.GetElementByItemID(ev.ElementID)
	p, _ := operations.GetIntentByItemID(ev.IntentID)

	evstr := evstruct{ev, e, p}
	return c.Render(http.StatusOK, "ev.html", evstr)
}

func newExpectedValue(c echo.Context) error {
	return c.Render(http.StatusOK, "editexpectedvalue.html", nil)
}
