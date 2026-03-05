package main

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/go-drift/drift/pkg/core"
	"github.com/go-drift/drift/pkg/drift"
	"github.com/go-drift/drift/pkg/layout"
	"github.com/go-drift/drift/pkg/platform"
	"github.com/go-drift/drift/pkg/theme"
	"github.com/go-drift/drift/pkg/widgets"
)

func main() {
	drift.NewApp(App()).Run()
}

func App() core.Widget {
	// https://driftframework.dev/docs/guides/theming#providing-theme
	return theme.Theme{
		Data:  theme.DefaultDarkTheme(), // or DefaultLightTheme()
		Child: app{},
	}
}

type app struct {
	core.StatefulBase
}

func (app) CreateState() core.State {
	return &appState{}
}

// https://driftframework.dev/docs/guides/state-management#lifecycle-order
type appState struct {
	core.StateBase
	filterText string
	items      []string
}

func (s *appState) InitState() {
	// Populate items with dummy data
	for i := range 30 {
		s.items = append(s.items, fmt.Sprintf("Item %d", i))
	}
}

func (s *appState) Build(ctx core.BuildContext) core.Widget {
	_, colors, textTheme := theme.UseTheme(ctx)

	filterCtrl := platform.NewTextEditingController(s.filterText)

	// Create a new list of items filtered by filterText
	var filteredItems []string
	if s.filterText == "" {
		filteredItems = s.items
	} else {
		filteredItems = slices.Collect(
			func(yield func(string) bool) {
				for _, v := range s.items {
					if strings.Contains(v, s.filterText) {
						if !yield(v) {
							return
						}
					}
				}
			})
	}

	return widgets.Container{
		Color: colors.Background,
		Child: widgets.Padded(
			layout.EdgeInsetsAll(8),
			widgets.Column{
				MainAxisAlignment:  widgets.MainAxisAlignmentStart,
				CrossAxisAlignment: widgets.CrossAxisAlignmentStart,
				MainAxisSize:       widgets.MainAxisSizeMax,
				Children: []core.Widget{
					theme.TextOf(ctx, "List filter demo", textTheme.HeadlineMedium),
					widgets.VSpace(16),
					widgets.Row{
						Children: []core.Widget{
							widgets.Expanded{
								Child: theme.TextFieldOf(ctx, filterCtrl).
									WithPlaceholder("Filter").
									WithInputAction(platform.TextInputActionSearch).
									WithOnSubmitted(func(v string) {
										slog.Info("OnSubmitted", "value", v)
										s.SetState(func() {
											s.filterText = v
										})
									}),
							},
							widgets.HSpace(16),
							theme.ButtonOf(ctx, "⌫", func() {
								s.SetState(func() {
									s.filterText = ""
								})
							}),
						},
					},
					widgets.VSpace(16),
					// widgets.Expanded is need to stop the list going past the bottom of the page
					widgets.Expanded{
						Child: widgets.ListViewBuilder{
							ItemCount: len(filteredItems),
							ItemBuilder: func(ctx core.BuildContext, index int) core.Widget {
								item := widgets.PaddingOnly(
									0, 8, 0, 8,
									widgets.Row{
										MainAxisAlignment: widgets.MainAxisAlignmentSpaceBetween,
										Children: []core.Widget{
											theme.TextOf(ctx, filteredItems[index], textTheme.BodyLarge),
											theme.IconOf(ctx, ">"),
										},
									},
								)
								return widgets.Tap(func() {
									slog.Info("TAPPED", "index", index)
								}, item)
							},
						},
					},
				},
			},
		),
	}
}
