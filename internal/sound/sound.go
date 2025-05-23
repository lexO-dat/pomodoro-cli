package sound

import (
	"fmt"
	"os/exec"
	"runtime"
)

// plays a notification sound
func Play() {
	switch runtime.GOOS {
	case "linux":
		if _, err := exec.LookPath("paplay"); err == nil {
			exec.Command("paplay", "/usr/share/sounds/alsa/Front_Left.wav").Start()
		} else if _, err := exec.LookPath("aplay"); err == nil {
			exec.Command("aplay", "/usr/share/sounds/alsa/Front_Left.wav").Start()
		} else if _, err := exec.LookPath("speaker-test"); err == nil {
			exec.Command("speaker-test", "-t", "sine", "-f", "800", "-l", "1").Start()
		} else {
			fmt.Print("\a")
		}
	default:
		fmt.Print("\a")
	}
}
