package main

import (
	"github.com/go-drift/drift/pkg/graphics"
	"github.com/go-drift/drift/pkg/theme"
)

func MyTheme() *theme.ThemeData {
	// Start with a default theme
	t := theme.DefaultDarkTheme()

	// And modify
	t.ColorScheme.Primary = graphics.ColorBlue
	t.ColorScheme.OnPrimary = graphics.ColorWhite
	t.ColorScheme.Background = graphics.RGB(0x33, 0x00, 0x00)

	buttonTheme := t.ButtonThemeOf()
	buttonTheme.FontSize = 32
	t.ButtonTheme = &buttonTheme

	return t
}
