package ui

import (
	"fmt"
	"strings"
	"time"

	"pomodoro/internal/sound"
	"pomodoro/internal/timer"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	greenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))  // Bright Green
	redStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
	/* titleStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Align(lipgloss.Center)
	statusStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Align(lipgloss.Center)
	timeStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")).Align(lipgloss.Center) */
	helpStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Align(lipgloss.Center)
	containerStyle = lipgloss.NewStyle().Align(lipgloss.Center).Padding(1, 2)
)

// UI state
type Model struct {
	timer    *timer.Timer
	showHelp bool
	width    int
	height   int
}

// new UI model
func NewModel(config timer.Config) *Model {
	return &Model{
		timer:    timer.New(config),
		showHelp: false,
	}
}

// initializes the model
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
		tea.EnterAltScreen,
		func() tea.Msg {
			fmt.Printf("\033]0;%s\007", "🍅 Pomodoro Timer")
			return nil
		},
	)
}

// handles messages
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Sequence(tea.ExitAltScreen, tea.Quit)
		case " ":
			if !m.showHelp {
				m.timer.Toggle()
			}
		case "h", "?":
			m.showHelp = !m.showHelp
		case "r":
			if !m.showHelp {
				m.timer.Reset()
			}
		}
	case time.Time:
		if !m.showHelp {
			completed := m.timer.Tick()
			if completed {
				sound.Play()
			}
		}
		return m, tickCmd()
	}
	return m, nil
}

// renders the UI
func (m *Model) View() string {
	if m.showHelp {
		return m.helpView()
	}

	var title string
	var emoji string
	sessionType := m.timer.SessionType()
	if sessionType == timer.Work {
		title = "Work Session"
		emoji = "🍅"
	} else {
		title = "Break Time"
		emoji = "☕"
	}

	// Get timer state
	remainingPercent := m.timer.RemainingPercent()
	timeDisplay := formatDuration(m.timer.Remaining())

	// Status
	var status string
	if m.timer.Remaining() <= 0 {
		if sessionType == timer.Work {
			status = "Work session complete! Starting break..."
		} else {
			status = "Break complete! Ready for work..."
		}
	} else if m.timer.IsRunning() {
		status = "Running... (SPACE to pause)"
	} else {
		status = "Paused (SPACE to start)"
	}

	// Creation of the descending bar: starts green, becomes more red as time runs out
	barWidth := min(60, m.width-10) // Responsive bar width
	if barWidth < 20 {
		barWidth = 20
	}
	greenBars := int(remainingPercent / 100 * float64(barWidth))
	redBars := barWidth - greenBars

	bar := greenStyle.Render(strings.Repeat("█", greenBars)) +
		redStyle.Render(strings.Repeat("█", redBars))

	// Create the main content
	content := fmt.Sprintf(`%s %s

%s

%s
%.0f%% remaining

%s

Press H for help`,
		emoji,
		title,
		timeDisplay,
		bar,
		remainingPercent,
		status,
	)

	// Center the content vertically and horizontally
	styledContent := lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		containerStyle.Render(content),
	)

	return styledContent
}

// render of the help screen
func (m *Model) helpView() string {
	helpContent := `🍅 POMODORO TIMER - COMMANDS

					SPACE     Start/Pause timer
					R         Reset to work session
					H or ?    Toggle help screen
					Q/ESC     Quit application

					Press H to return to timer`

	// Center the help content
	styledHelp := lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		helpStyle.Width(m.width-4).Render(helpContent),
	)

	return styledHelp
}

// formats a duration into MM:SS format
func formatDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

// returns a command that sends a tick message every second
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

// well, i think that i dont have to explain this xd
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
