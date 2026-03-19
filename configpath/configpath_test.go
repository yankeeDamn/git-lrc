package configpath

import "testing"

func TestNormalizeWindowsHomePathFromMSYSDrivePath(t *testing.T) {
	got := normalizeHomeForOS("windows", "/c/Users/Alice")
	want := `C:\Users\Alice`
	if got != want {
		t.Fatalf("normalize windows msys path: got %q want %q", got, want)
	}
}

func TestNormalizeWindowsHomePathFromNativePath(t *testing.T) {
	got := normalizeHomeForOS("windows", `C:\Users\Alice`)
	want := `C:\Users\Alice`
	if got != want {
		t.Fatalf("normalize windows native path: got %q want %q", got, want)
	}
}

func TestNormalizeHomeForLinuxKeepsPOSIXPath(t *testing.T) {
	got := normalizeHomeForOS("linux", "/home/alice")
	want := "/home/alice"
	if got != want {
		t.Fatalf("normalize linux path: got %q want %q", got, want)
	}
}
