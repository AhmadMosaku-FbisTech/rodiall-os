// Package rodialutils provides shared CLI helpers and utilities
// for all Rodial services (agent, proxy, updater, ddns, etc.)
//
// Usage example:
//     import "rodial-utils"
//     rodialutils.PrintBanner("rodial-agent", "v1.0")
//     opts := rodialutils.ParseArgs()
//     rodialutils.LogInfo("Starting agent...")
//
package rodialutils

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lrodialutils
#include "cli_helper.h"
*/
import "C"
import "unsafe"

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// CLIOptions holds parsed command line arguments
type CLIOptions struct {
	ConfigPath string
	LogLevel   string
	Verbose    bool
	ShowHelp   bool
}

// PrintBanner displays a consistent banner for all Rodial services
func PrintBanner(service, version string) {
	fmt.Printf("\n🛰️  %s (%s)\n", strings.ToUpper(service), version)
	fmt.Println("───────────────────────────────────────────────")
	fmt.Printf("Build Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Rodial OS © 2025 — All rights reserved.\n")
}

// ParseArgs parses standard Rodial CLI arguments
func ParseArgs() *CLIOptions {
	opts := &CLIOptions{}
	flag.StringVar(&opts.ConfigPath, "config", "/etc/config/rodial", "Path to configuration file")
	flag.StringVar(&opts.LogLevel, "log", "info", "Log level (debug, info, warn, error)")
	flag.BoolVar(&opts.Verbose, "v", false, "Enable verbose logging")
	flag.BoolVar(&opts.ShowHelp, "h", false, "Show help message")
	flag.Parse()

	if opts.ShowHelp {
		flag.Usage()
		os.Exit(0)
	}
	return opts
}

// LogInfo prints an informational message
func LogInfo(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

// LogDebug prints a debug message if verbose is enabled
func LogDebug(enabled bool, msg string) {
	if enabled {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

// LogError prints an error message and exits if fatal is true
func LogError(msg string, fatal bool) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s\n", msg)
	if fatal {
		os.Exit(1)
	}
}

// CheckFileExists validates file existence
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// LoadEnv safely reads environment variables
func LoadEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
//go:build cgo
// +build cgo

package rodialutils


// Banner prints a consistent Rodial banner via C backend
func Banner(service, version string) {
	cs := C.CString(service)
	cv := C.CString(version)
	C.rodial_print_banner(cs, cv)
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(cv))
}

func LogInfo(msg string) {
	cs := C.CString(msg)
	C.rodial_log_info(cs)
	C.free(unsafe.Pointer(cs))
}

func LogError(msg string, fatal bool) {
	cs := C.CString(msg)
	if fatal {
		C.rodial_log_error(cs, 1)
	} else {
		C.rodial_log_error(cs, 0)
	}
	C.free(unsafe.Pointer(cs))
}

func FileExists(path string) bool {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	return C.rodial_file_exists(cp) != 0
}

// rodial-utils/src/cli_helper.h
#ifndef RODIAL_CLI_HELPER_H
#define RODIAL_CLI_HELPER_H

#ifdef __cplusplus
extern "C" {
#endif

// Prints a consistent Rodial banner
void rodial_print_banner(const char* service, const char* version);

// Logs an informational message
void rodial_log_info(const char* message);

// Logs an error message; if fatal is non-zero, exits the program
void rodial_log_error(const char* message, int fatal);

// Checks if a file exists at the given path; returns 1 if exists, 0 otherwise
int rodial_file_exists(const char* path);

#ifdef __cplusplus
}
#endif

#endif // RODIAL_CLI_HELPER_H
