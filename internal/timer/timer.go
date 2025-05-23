package timer

import "time"

// represents the type of timer session
type SessionType int

const (
	Work SessionType = iota
	Break
)

// returns the string representation of the session type
func (s SessionType) String() string {
	switch s {
	case Work:
		return "Work"
	case Break:
		return "Break"
	default:
		return "Unknown"
	}
}

// holds the timer configuration
type Config struct {
	WorkDuration  time.Duration
	BreakDuration time.Duration
}

// represents a pomodoro timer
type Timer struct {
	config      Config
	remaining   time.Duration
	total       time.Duration
	running     bool
	sessionType SessionType
}

// creates a new timer with the given configuration
func New(config Config) *Timer {
	return &Timer{
		config:      config,
		remaining:   config.WorkDuration,
		total:       config.WorkDuration,
		running:     false,
		sessionType: Work,
	}
}

// starts the timer
func (t *Timer) Start() {
	t.running = true
}

// pauses the timer
func (t *Timer) Pause() {
	t.running = false
}

// toggles the timer between running and paused
func (t *Timer) Toggle() {
	t.running = !t.running
}

// decrements the timer by one second if running
// Returns true if the session completed
func (t *Timer) Tick() bool {
	if !t.running || t.remaining <= 0 {
		return false
	}

	t.remaining -= time.Second
	if t.remaining <= 0 {
		t.running = false
		t.switchSession()
		return true
	}
	return false
}

// resets the timer to the initial work session
func (t *Timer) Reset() {
	t.sessionType = Work
	t.remaining = t.config.WorkDuration
	t.total = t.config.WorkDuration
	t.running = false
}

// switches between work and break sessions
func (t *Timer) switchSession() {
	if t.sessionType == Work {
		t.sessionType = Break
		t.remaining = t.config.BreakDuration
		t.total = t.config.BreakDuration
	} else {
		t.sessionType = Work
		t.remaining = t.config.WorkDuration
		t.total = t.config.WorkDuration
	}
}

// returns the remaining time
func (t *Timer) Remaining() time.Duration {
	return t.remaining
}

// returns the total duration for the current session
func (t *Timer) Total() time.Duration {
	return t.total
}

// returns whether the timer is running
func (t *Timer) IsRunning() bool {
	return t.running
}

// returns the current session type
func (t *Timer) SessionType() SessionType {
	return t.sessionType
}

// returns the progress as a percentage (0-100)
func (t *Timer) Progress() float64 {
	if t.total == 0 {
		return 0
	}
	elapsed := t.total - t.remaining
	return float64(elapsed) / float64(t.total) * 100
}

// returns the remaining time as a percentage (0-100)
func (t *Timer) RemainingPercent() float64 {
	if t.total == 0 {
		return 0
	}
	return float64(t.remaining) / float64(t.total) * 100
}
