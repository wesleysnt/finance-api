package pkg

import (
	"cmp"
	"slices"

	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
)

func ListRouters(e *echo.Echo) {
	routes := e.Routes()

	slices.SortFunc(routes, func(a, b *echo.Route) int {
		return cmp.Or(
			cmp.Compare(a.Method, b.Method),
			cmp.Compare(a.Path, b.Path),
		)
	})

	for _, r := range routes {
		color.Cyanln(r.Method, r.Path)
	}
}
