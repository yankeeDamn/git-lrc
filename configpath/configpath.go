package configpath

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const configFileName = ".lrc.toml"

// ResolveHomeDir returns a stable absolute home directory path for this process.
func ResolveHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to determine home directory: %w", err)
	}

	normalized := normalizeHomeForOS(runtime.GOOS, homeDir)
	if strings.TrimSpace(normalized) == "" {
		return "", fmt.Errorf("failed to determine home directory: empty path")
	}

	return normalized, nil
}

// ResolveConfigPath returns the canonical ~/.lrc.toml path.
func ResolveConfigPath() (string, error) {
	homeDir, err := ResolveHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func normalizeHomeForOS(goos, homeDir string) string {
	normalized := strings.TrimSpace(homeDir)
	if normalized == "" {
		return ""
	}

	if goos == "windows" {
		normalized = normalizeWindowsHomePath(normalized)
	}

	return filepath.Clean(normalized)
}

// normalizeWindowsHomePath converts common MSYS-style paths to native form.
func normalizeWindowsHomePath(homeDir string) string {
	normalized := strings.TrimSpace(homeDir)
	if len(normalized) >= 3 && normalized[0] == '/' && isASCIILetter(normalized[1]) && normalized[2] == '/' {
		drive := strings.ToUpper(string(normalized[1]))
		normalized = drive + ":" + normalized[2:]
	}
	if len(normalized) >= 2 && isASCIILetter(normalized[0]) && normalized[1] == ':' {
		normalized = strings.ToUpper(string(normalized[0])) + normalized[1:]
	}
	return strings.ReplaceAll(normalized, "/", `\`)
}

func isASCIILetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
