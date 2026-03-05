package main

import (
	"fmt"
	"log/slog"

	"github.com/go-drift/drift/pkg/core"
	"github.com/go-drift/drift/pkg/graphics"
	"github.com/go-drift/drift/pkg/navigation"
	"github.com/go-drift/drift/pkg/theme"
	"github.com/go-drift/drift/pkg/widgets"
)

type SettingsPage struct {
	core.StatefulBase
}

func (SettingsPage) CreateState() core.State {
	return &SettingsPageState{}
}

type SettingsPageState struct {
	core.StateBase
	settings   navigation.RouteSettings
	filterText string
	items      []string
}

func (s *SettingsPageState) InitState() {
	slog.Info("SettingsPage.InitState()")
	// Populate items with dummy data
	for i := range 100 {
		s.items = append(s.items, fmt.Sprintf("Item %d", i))
	}
}

func (s *SettingsPageState) Dispose() {
	slog.Info("SettingsPage.Dispose")
}

func (s *SettingsPageState) Build(ctx core.BuildContext) core.Widget {
	slog.Info("SettingsPage.Build()", "settings", s.settings)
	textTheme := theme.TextThemeOf(ctx)

	header := widgets.Row{
		MainAxisAlignment:  widgets.MainAxisAlignmentSpaceBetween,
		CrossAxisAlignment: widgets.CrossAxisAlignmentCenter,
		Children: []core.Widget{
			theme.TextOf(ctx, "Settings", textTheme.TitleLarge.WithColor(graphics.ColorGreen)),
		},
	}

	content := theme.TextOf(ctx, "Settings content", textTheme.BodyLarge)

	return widgets.SafeArea{
		Child: widgets.Column{
			Children: []core.Widget{
				header,
				widgets.VSpace(16),
				content,
			},
		},
	}
}
