package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/sync/errgroup"
)

// IssueCommentEventHandleFunc represents a callback function triggered on github.IssueCommentEvent.
// deliveryID (type: string) is the unique webhook delivery ID.
// eventName (type: string) is the name of the event.
// event (type: *github.IssueCommentEvent) is the webhook payload.
type IssueCommentEventHandleFunc func(deliveryID string, eventName string, event *github.IssueCommentEvent) error

// OnIssueCommentCreated registers callbacks listening to events of type github.IssueCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentCreated must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssueCommentCreated(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = append(g.onIssueCommentEvent[action], callbacks...)
}

// SetOnIssueCommentCreated registers callbacks listening to events of type github.IssueCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentCreatedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssueCommentCreated(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "created"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = callbacks
}

func (g *EventHandler) handleIssueCommentCreated(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "created"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentCreated() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssueCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssueCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssueCommentEvent[action] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnIssueCommentEdited registers callbacks listening to events of type github.IssueCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentEdited must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssueCommentEdited(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = append(g.onIssueCommentEvent[action], callbacks...)
}

// SetOnIssueCommentEdited registers callbacks listening to events of type github.IssueCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentEditedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssueCommentEdited(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "edited"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = callbacks
}

func (g *EventHandler) handleIssueCommentEdited(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "edited"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentEdited() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssueCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssueCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssueCommentEvent[action] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnIssueCommentDeleted registers callbacks listening to events of type github.IssueCommentEvent.
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentDeleted must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssueCommentDeleted(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = append(g.onIssueCommentEvent[action], callbacks...)
}

// SetOnIssueCommentDeleted registers callbacks listening to events of type github.IssueCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentDeletedAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssueCommentDeleted(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const action = "deleted"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[action] = callbacks
}

func (g *EventHandler) handleIssueCommentDeleted(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}

	const action = "deleted"
	if action != *event.Action {
		return fmt.Errorf(
			"handleIssueCommentDeleted() called with wrong action, want %s, got %s",
			action,
			*event.Action,
		)
	}

	err := g.handleIssueCommentEventAny(deliveryID, eventName, event)
	if err != nil {
		return err
	}
	if _, ok := g.onIssueCommentEvent[action]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssueCommentEvent[action] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// OnIssueCommentEventAny registers callbacks listening to events of type github.IssueCommentEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnIssueCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) OnIssueCommentEventAny(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[any] = append(g.onIssueCommentEvent[any], callbacks...)
}

// SetOnIssueCommentEventAny registers callbacks listening to events of type github.IssueCommentEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnIssueCommentEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
func (g *EventHandler) SetOnIssueCommentEventAny(callbacks ...IssueCommentEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// "action" is used to register handleFuncs on action types.
	// "*" - triggers on all action types or when the event does not have actions
	const any = "*"

	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onIssueCommentEvent == nil {
		g.onIssueCommentEvent = make(map[string][]IssueCommentEventHandleFunc)
	}
	g.onIssueCommentEvent[any] = callbacks
}

func (g *EventHandler) handleIssueCommentEventAny(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
	if event == nil {
		return fmt.Errorf("event was empty or nil")
	}
	const any = "*"
	if _, ok := g.onIssueCommentEvent[any]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onIssueCommentEvent[any] {
		handle := h
		eg.Go(func() error {
			err := handle(deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// IssueCommentEvent handles github.IssueCommentEvent
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered by OnBeforeAny are executed in parallel.
// 2) All callbacks registered by OnIssueCommentEventAny are executed in parallel.
// 3) Optional: All callbacks registered via OnIssueCommentEvent... are executed in parallel in case the Event has actions.
// 4) All callbacks registered by OnAfterAny are executed in parallel.
// on any error all callbacks registered by OnError are executed in parallel.
//
func (g *EventHandler) IssueCommentEvent(deliveryID string, eventName string, event *github.IssueCommentEvent) error {

	if event == nil || event.Action == nil || *event.Action == "" {
		return fmt.Errorf("event action was empty or nil")
	}
	action := *event.Action

	err := g.handleBeforeAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}

	switch action {

	case "created":
		err := g.handleIssueCommentCreated(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "edited":
		err := g.handleIssueCommentEdited(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	case "deleted":
		err := g.handleIssueCommentDeleted(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}

	default:
		err := g.handleIssueCommentEventAny(deliveryID, eventName, event)
		if err != nil {
			return g.handleError(deliveryID, eventName, event, err)
		}
	}

	err = g.handleAfterAny(deliveryID, eventName, event)
	if err != nil {
		return g.handleError(deliveryID, eventName, event, err)
	}
	return nil
}
