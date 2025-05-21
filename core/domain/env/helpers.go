// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package env

import "os"

type Env string

// System environment variables
const (
	// LANG represents the locale/language settings of the system.
	// Example: "en_US.UTF-8"
	LANG Env = "LANG"
	// TZ represents the timezone setting of the system.
	// Example: "UTC", "America/New_York"
	TZ Env = "TZ"
	// PWD represents the current working directory.
	// This is typically set by shells and reflects the directory from which a process was started.
	PWD Env = "PWD"
	// OS represents the operating system running the process.
	OS Env = "OS"
	// HOME represents the home directory of the current user.
	HOME Env = "HOME"
	// USER represents the username of the current user.
	USER Env = "USER"
	// SHELL represents the default shell for the current user.
	SHELL Env = "SHELL"
	// PATH represents the search path for executable files.
	PATH Env = "PATH"
	// HOSTNAME represents the hostname of the system.
	HOSTNAME Env = "HOSTNAME"
	// TMPDIR represents the temporary directory for the current user.
	TMPDIR Env = "TMPDIR"
)

// Execution environment variables
const (
	// GOBIN represents the directory where Go binaries are installed.
	GOBIN Env = "GOBIN"
	// GOPROXY represents the proxy used for Go module downloads.
	GOPROXY Env = "GOPROXY"
	// GOPATH represents the directory where Go packages are installed.
	GOPATH Env = "GOPATH"
	// GOROOT represents the directory where the Go runtime is installed.
	GOROOT Env = "GOROOT"
	// GO111MODULE represents the module mode for Go.
	GO111MODULE Env = "GO111MODULE"
	// GOFLAGS represents additional flags for Go commands.
	GOFLAGS Env = "GOFLAGS"
	// GOOS represents the operating system the process is running on.
	GOOS Env = "GOOS"
	// GOMODCACHE Env = "GOMODCACHE"
	GOMODCACHE Env = "GOMODCACHE"
	// GOVERSION  Env = "GOVERSION"
	GOVERSION Env = "GOVERSION"
)

// Terminal environment variables
const (
	//TERM         Env = "TERM"         // TERM is a standard environment variable that indicates the terminal type.
	TERM Env = "TERM"
	// WT_SESSION
	WT_SESSION   Env = "WT_SESSION"   // is specific to Windows Terminal.
	TERM_PROGRAM Env = "TERM_PROGRAM" // is used by macOS Terminal and iTerm2.
	COLORTERM    Env = "COLORTERM"    // indicates support for true color in terminals.
)

const (
	// Editor-related
	EDITOR Env = "EDITOR"
	VISUAL Env = "VISUAL"
)

const (
	// CI/CD & Automation
	CI             Env = "CI"
	GITHUB_ACTIONS Env = "GITHUB_ACTIONS"
	JENKINS_HOME   Env = "JENKINS_HOME"

	// Container-related
	CONTAINER_NAME Env = "CONTAINER_NAME"
	CONTAINER_ID   Env = "CONTAINER_ID"
)

const (

	// Cloud-related
	AWS_REGION  Env = "AWS_REGION"
	AWS_PROFILE Env = "AWS_PROFILE"
	GCP_PROJECT Env = "GCP_PROJECT"
)

const (
	// Debugging & Logs
	DEBUG     Env = "DEBUG"
	LOG_LEVEL Env = "LOG_LEVEL"

	// OS-specific
	DARWIN_VERSION  Env = "DARWIN_VERSION"
	LINUX_VERSION   Env = "LINUX_VERSION"
	WINDOWS_VERSION Env = "WINDOWS_VERSION"
)

// Windows environment variables
const (
	// APPDATA      Env = "APPDATA"      // Windows Application data folder
	APPDATA Env = "APPDATA"
	// LOCALAPPDATA Env = "LOCALAPPDATA" // Windows Local app data folder
	LOCALAPPDATA Env = "LOCALAPPDATA"
	// USERNAME     Env = "USERNAME"     // Windows Current logged-in user
	USERNAME Env = "USERNAME"
	// USERPROFILE  Env = "USERPROFILE"  // Windows User home directory
	USERPROFILE Env = "USERPROFILE"
	// SystemRoot   Env = "SystemRoot"   // Windows system root
	SystemRoot Env = "SystemRoot"
	// ProgramFiles Env = "ProgramFiles" // Windows Program Files directory
	ProgramFiles Env = "ProgramFiles"
	// ComSpec      Env = "ComSpec"      // Windows Path to cmd.exe
	ComSpec Env = "ComSpec"
)

func (e Env) Get() string {
	return os.Getenv(string(e))
}

// Set sets the value of an environment variable.
func (e Env) Set(value string) error {
	return os.Setenv(string(e), value)
}

// Exists checks if an environment variable exists.
func (e Env) Exists() bool {
	_, exists := os.LookupEnv(string(e))
	return exists
}

func (e Env) Equals(value string) bool {
	return os.Getenv(string(e)) == value
}

// Unset removes an environment variable.
func (e Env) Unset() error {
	return os.Unsetenv(string(e))
}

// IsWindows returns true if the OS is Windows.
func IsWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
}

// IsUnix returns true if the OS is Unix-like (Linux/macOS).
func IsUnix() bool {
	return !IsWindows()
}
