package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func KillProcessOnPort(port uint16) error {
	switch runtime.GOOS {
	case "windows":
		return killOnWindows(port)
	case "linux", "darwin": // macOS and Linux
		return killOnUnix(port)
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func killOnUnix(port uint16) error {
	// Find the PID using lsof
	cmd := exec.Command("lsof", "-t", fmt.Sprintf("-i:%d", port))
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	pid := strings.TrimSpace(string(output))
	if pid == "" {
		return fmt.Errorf("no process found on port %d", port)
	}

	// Kill the process using kill
	killCmd := exec.Command("kill", "-9", pid)
	if err := killCmd.Run(); err != nil {
		return err
	}

	return nil
}

func killOnWindows(port uint16) error {
	// Run netstat to find the PID using the port
	cmd := exec.Command("cmd", "/C", fmt.Sprintf(`netstat -ano | findstr ":%d"`, port))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("netstat command failed: %v", err)
	}

	// Parse the output to extract the PID
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf(":%d", port)) {
			fields := strings.Fields(line)
			if len(fields) > 4 {
				pid := fields[len(fields)-1]

				// Kill the process using the PID
				killCmd := exec.Command("taskkill", "/F", "/PID", pid)
				var killOut bytes.Buffer
				killCmd.Stdout = &killOut
				killCmd.Stderr = &killOut

				if err := killCmd.Run(); err != nil {
					return fmt.Errorf("taskkill command failed: %v", err)
				}
				fmt.Printf("Killed process with PID: %s\n", pid)
				break
			}
		}
	}

	return nil
}
