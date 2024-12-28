package config

import "path"

// Base app configuration key value
type app struct {
	// App name
	Name string `json:"name"`

	// App description
	Desc string `json:"desc"`

	// Port number that app will running
	Port int `json:"port"`

	// Environtment of the app
	//
	// "development" or "production"
	Env string `json:"env"`

	// Prefork mode on REST API server
	//
	// If true, app will running on prefork mode (multi child proccess)
	Prefork bool `json:"prefork"`

	// Is abbreviation of Auto Update Watcher Interval
	//
	// Value must be int, and interval should be in second(s)
	AUWI int `json:"auwi"`

	// Full path where program executable places
	ProgramFile string `json:"programFile"`

	// Full path of working directory location
	WorkingDir string `json:"workingDir"`

	TimeZone string `json:"timeZone"`
}

// Debug global debug flag based on app env.
// Return debug true/false
func (a *app) Debug() bool {
	return a.Env != "production"
}

// ResolveFilePathInWorkDir return full path of given file name in working directory
func (a *app) ResolveFilePathInWorkDir(f string) string {
	return path.Join(a.WorkingDir, "/", f)
}
