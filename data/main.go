package main

import (
	"log/slog"
	"time"

	"github.com/go-drift/drift/pkg/core"
	"github.com/go-drift/drift/pkg/drift"
	"github.com/go-drift/drift/pkg/layout"
	"github.com/go-drift/drift/pkg/overlay"
	"github.com/go-drift/drift/pkg/theme"
	"github.com/go-drift/drift/pkg/widgets"
)

func main() {
	drift.NewApp(App()).Run()
}

func App() core.Widget {
	// https://driftframework.dev/docs/guides/theming#providing-theme
	return theme.Theme{
		Data: theme.DefaultDarkTheme(), // or DefaultLightTheme()
		// We add an overlay to use later with Toast
		Child: overlay.Overlay{
			Child: app{},
		},
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
	status string
	err    error
	data   *Data
}

func (s *appState) InitState() {
	s.status = "Loading"
	s.data = &Data{}

	go func() {
		err := s.data.Load()
		slog.Info("Data: Load", "err", err)
		drift.Dispatch(func() {
			s.SetState(func() {
				s.err = err
				s.UpdateStatus()
				slog.Info("Data: Dispatch", "err", err, "status", s.status)
			})
		})
	}()
}

func (s *appState) UpdateStatus() {
	if s.err != nil {
		s.status = s.err.Error()
	}
	if s.data == nil {
		s.status = "No data"
	}
	s.status = "OK"
}

func (s *appState) Build(ctx core.BuildContext) core.Widget {
	slog.Info("Build()")
	_, colors, textTheme := theme.UseTheme(ctx)

	return widgets.Container{
		Padding: layout.EdgeInsetsAll(8),
		Color:   colors.Background,
		Child: widgets.Column{
			Children: []core.Widget{
				theme.TextOf(ctx, "State: "+s.status, textTheme.DisplaySmall),
				widgets.Expanded{
					Child: widgets.ListViewBuilder{
						ItemCount: len(s.data.Items),
						ItemBuilder: func(ctx core.BuildContext, index int) core.Widget {
							return theme.TextOf(ctx, s.data.Items[index], textTheme.BodyLarge)
						},
					},
				},
				widgets.Row{
					MainAxisAlignment: widgets.MainAxisAlignmentSpaceBetween,
					Children: []core.Widget{
						theme.ButtonOf(ctx, "➕\nAdd item", func() {
							s.SetState(func() {
								s.data.Items = append(s.data.Items, time.Now().Format(time.RFC822))
							})
						}),
						theme.ButtonOf(ctx, "❌\nClear all", func() {
							s.SetState(func() {
								s.data.Items = make([]string, 0)
							})
						}),
						theme.ButtonOf(ctx, "💾\nSave", func() {
							err := s.data.Save()
							slog.Info("Data: Save", "err", err)
							drift.Dispatch(func() {
								s.SetState(func() {
									s.err = err
									s.UpdateStatus()
									slog.Info("Data: Dispatch", "err", err, "status", s.status)
								})
							})
						}),
					},
				},
			},
		},
	}
}
