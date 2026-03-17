//go:build windows

package interactive

import (
	"errors"
	"syscall"
)

// ERROR_INVALID_HANDLE from Win32 API (winerror.h).
const windowsErrorInvalidHandle syscall.Errno = 6

func isIgnorableSyncError(err error) bool {
	// Keep this list narrow and Windows-specific; unexpected sync errors should be surfaced.
	// ERROR_INVALID_HANDLE is common when stdout/stderr are attached to non-syncable console handles.
	return errors.Is(err, syscall.EINVAL) || errors.Is(err, syscall.ENOTSUP) || errors.Is(err, windowsErrorInvalidHandle)
}
