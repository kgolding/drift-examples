package main

import (
	"github.com/go-drift/drift/pkg/graphics"
	"github.com/go-drift/drift/pkg/widgets"
)

func HomeIcon() widgets.Icon {
	return widgets.Icon{
		Glyph: "🏠",
		Size:  24,
		Color: graphics.RGB(20, 20, 20),
	}
}
func GearIcon() widgets.Icon {
	return widgets.Icon{
		Glyph: "⚙",
		Size:  24,
		Color: graphics.RGB(20, 20, 20),
	}
}
