// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnDiscussionEventAny(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventAny(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventAnyAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventAny(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventAny(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventAny(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",

				event: &github.DiscussionEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",

				event: &github.DiscussionEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventAny(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleDiscussionEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventCreated(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventCreated(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventCreatedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventCreated(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventCreated(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventCreated(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventCreated(t *testing.T) {
	action := DiscussionEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventCreated(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventEdited(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventEdited(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventEditedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventEdited(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventEdited(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventEdited(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventEdited(t *testing.T) {
	action := DiscussionEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventEdited(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventDeleted(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventDeleted(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventDeletedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventDeleted(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventDeleted(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventDeleted(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventDeleted(t *testing.T) {
	action := DiscussionEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventDeleted(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventPinned(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventPinned(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventPinnedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventPinnedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventPinned(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventPinned(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventPinned(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventPinnedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventPinnedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventPinned(t *testing.T) {
	action := DiscussionEventPinnedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventPinned(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventPinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventPinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnpinned(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnpinnedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventUnpinnedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventUnpinned(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventUnpinned(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventUnpinned(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnpinnedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventUnpinnedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventUnpinned(t *testing.T) {
	action := DiscussionEventUnpinnedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnpinned(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventUnpinned(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventUnpinned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventLocked(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventLocked(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventLockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventLockedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventLocked(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventLocked(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventLocked(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventLockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventLockedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventLocked(t *testing.T) {
	action := DiscussionEventLockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventLocked(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventLocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventLocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnlocked(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnlockedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventUnlockedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventUnlocked(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventUnlocked(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventUnlocked(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnlockedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventUnlockedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventUnlocked(t *testing.T) {
	action := DiscussionEventUnlockedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnlocked(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventUnlocked(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventUnlocked() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventTransferred(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventTransferred(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventTransferredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventTransferredAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventTransferred(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventTransferred(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventTransferred(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventTransferredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventTransferredAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventTransferred(t *testing.T) {
	action := DiscussionEventTransferredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventTransferred(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventTransferred(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventTransferred() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventCategoryChanged(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventCategoryChanged(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventCategoryChangedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventCategoryChangedAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventCategoryChanged(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventCategoryChanged(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventCategoryChanged(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventCategoryChangedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventCategoryChangedAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventCategoryChanged(t *testing.T) {
	action := DiscussionEventCategoryChangedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventCategoryChanged(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventCategoryChanged(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventCategoryChanged() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventAnswered(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventAnswered(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventAnsweredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventAnsweredAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventAnswered(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventAnswered(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventAnswered(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventAnsweredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventAnsweredAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventAnswered(t *testing.T) {
	action := DiscussionEventAnsweredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventAnswered(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventAnswered(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventAnswered() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventUnanswered(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnanswered(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnansweredAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventUnansweredAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventUnanswered(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventUnanswered(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventUnanswered(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnansweredAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventUnansweredAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventUnanswered(t *testing.T) {
	action := DiscussionEventUnansweredAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnanswered(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventUnanswered(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventUnanswered() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventLabeled(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventLabeled(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventLabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventLabeledAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventLabeled(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventLabeled(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventLabeled(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventLabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventLabeledAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventLabeled(t *testing.T) {
	action := DiscussionEventLabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventLabeled(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventLabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventLabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnDiscussionEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple DiscussionEventHandleFunc",
			args: args{
				callbacks: []DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnlabeled(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnlabeledAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onDiscussionEvent[DiscussionEventUnlabeledAction]))
			}
		})
	}
}

func TestSetOnDiscussionEventUnlabeled(t *testing.T) {
	type args struct {
		callbacks []DiscussionEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single DiscussionEventHandleFunc",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple DiscussionEventHandleFuncs",
			args: args{
				[]DiscussionEventHandleFunc{
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnDiscussionEventUnlabeled(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				return nil
			})
			g.SetOnDiscussionEventUnlabeled(tt.args.callbacks...)
			if len(g.onDiscussionEvent[DiscussionEventUnlabeledAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onDiscussionEvent[DiscussionEventUnlabeledAction]), tt.want)
			}
		})
	}
}

func TestHandleDiscussionEventUnlabeled(t *testing.T) {
	action := DiscussionEventUnlabeledAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.DiscussionEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "discussion",
				event:      &github.DiscussionEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnDiscussionEventUnlabeled(func(deliveryID string, eventName string, event *github.DiscussionEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleDiscussionEventUnlabeled(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleDiscussionEventUnlabeled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
