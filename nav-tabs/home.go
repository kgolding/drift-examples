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

type HomePage struct {
	core.StatefulBase
}

func (HomePage) CreateState() core.State {
	return &HomePageState{}
}

type HomePageState struct {
	core.StateBase
	settings navigation.RouteSettings
	items    []string
}

func (s *HomePageState) InitState() {
	slog.Info("HomePage.InitState()")
	// Populate items with dummy data
	for i := range 30 {
		s.items = append(s.items, fmt.Sprintf("Item %d", i))
	}
}

func (s *HomePageState) Build(ctx core.BuildContext) core.Widget {
	slog.Info("HomePage.Build()", "settings", s.settings)

	textTheme := theme.TextThemeOf(ctx)

	header := theme.TextOf(ctx, "Home", textTheme.TitleLarge.WithColor(graphics.ColorGreen))

	content := widgets.Expanded{
		Child: widgets.ListViewBuilder{
			ItemCount: len(s.items),
			ItemBuilder: func(ctx core.BuildContext, index int) core.Widget {
				item := widgets.PaddingOnly(
					0, 8, 0, 8,
					widgets.Row{
						MainAxisAlignment: widgets.MainAxisAlignmentSpaceBetween,
						Children: []core.Widget{
							theme.TextOf(ctx, s.items[index], textTheme.BodyLarge),
							theme.IconOf(ctx, ">"),
						},
					},
				)
				return widgets.Tap(func() {
					slog.Info("TAPPED", "index", index)
					nav := navigation.NavigatorOf(ctx)
					if nav == nil {
						return
					}
					nav.PushNamed("/alert", index)
				}, item)
			},
		},
	}

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
