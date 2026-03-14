package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// Tool represents a tool to be tested in the container
type Tool struct {
	Name    string
	Cmd     string
	Args    []string
	Timeout time.Duration
}

// TestResult represents the result of a tool test
type TestResult struct {
	Tool    Tool
	Success bool
	Output  string
	Error   error
}

func testContainer(containerName string) {
	ctx := context.Background()
	ctx = log.Logger.WithContext(ctx)

	log.Info().Str("container", containerName).Msg("Testing container environment")

	// Check if container is running
	if !isContainerRunning(containerName) {
		log.Error().Str("container", containerName).Msg("Container is not running")
		os.Exit(1)
	}

	// Define tools to test
	tools := []Tool{
		// Programming Languages
		{Name: "Go", Cmd: "go", Args: []string{"version"}, Timeout: 10 * time.Second},
		{Name: "Rust/Cargo", Cmd: "cargo", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Python", Cmd: "python", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Python3", Cmd: "python3", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Node.js", Cmd: "node", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "npm", Cmd: "npm", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "yarn", Cmd: "yarn", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "pnpm", Cmd: "pnpm", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Java", Cmd: "java", Args: []string{"-version"}, Timeout: 10 * time.Second},
		{Name: "Maven", Cmd: "mvn", Args: []string{"--version"}, Timeout: 30 * time.Second},
		{Name: "Gradle", Cmd: "gradle", Args: []string{"--version"}, Timeout: 30 * time.Second},
		{Name: ".NET", Cmd: "dotnet", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "GCC", Cmd: "gcc", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Clang", Cmd: "clang", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "CMake", Cmd: "cmake", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// DevOps & Container Tools
		{Name: "Docker", Cmd: "docker", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Docker Compose", Cmd: "docker-compose", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Podman", Cmd: "podman", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "kubectl", Cmd: "kubectl", Args: []string{"version", "--client"}, Timeout: 10 * time.Second},
		{Name: "Helm", Cmd: "helm", Args: []string{"version"}, Timeout: 10 * time.Second},

		// Version Control
		{Name: "Git", Cmd: "git", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "GitHub CLI", Cmd: "gh", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Build Tools
		{Name: "Task", Cmd: "task", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Make", Cmd: "make", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Python Tools
		{Name: "uv", Cmd: "uv", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Typst
		{Name: "Typst", Cmd: "typst", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Go Tools
		{Name: "gopls", Cmd: "gopls", Args: []string{"version"}, Timeout: 10 * time.Second},
		{Name: "golangci-lint", Cmd: "golangci-lint", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Shell & Terminal
		{Name: "Zsh", Cmd: "zsh", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Fish", Cmd: "fish", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Editor
		{Name: "Vim", Cmd: "vim", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Nano", Cmd: "nano", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "Micro", Cmd: "micro", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Utilities
		{Name: "jq", Cmd: "jq", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "wget", Cmd: "wget", Args: []string{"--version"}, Timeout: 10 * time.Second},
		{Name: "unzip", Cmd: "unzip", Args: []string{"-v"}, Timeout: 10 * time.Second},
		{Name: "rsync", Cmd: "rsync", Args: []string{"--version"}, Timeout: 10 * time.Second},

		// Claude Code
		{Name: "Claude Code", Cmd: "claude", Args: []string{"--version"}, Timeout: 10 * time.Second},
	}

	// Run tests
	results := make([]TestResult, 0, len(tools))
	passed := 0
	failed := 0

	for _, tool := range tools {
		result := testTool(ctx, containerName, tool)
		results = append(results, result)
		if result.Success {
			passed++
		} else {
			failed++
		}
	}

	// Print results
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("       Container Environment Test       ")
	fmt.Println("========================================")
	fmt.Println()

	for _, result := range results {
		status := "✓ PASS"
		if !result.Success {
			status = "✗ FAIL"
		}
		fmt.Printf("%-20s %s\n", result.Tool.Name+":", status)
		if result.Success && result.Output != "" {
			// Print first line of output
			lines := strings.Split(result.Output, "\n")
			if len(lines) > 0 && lines[0] != "" {
				fmt.Printf("    %s\n", strings.TrimSpace(lines[0]))
			}
		}
		if !result.Success && result.Error != nil {
			fmt.Printf("    Error: %v\n", result.Error)
		}
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("Total: %d passed, %d failed\n", passed, failed)
	fmt.Println("========================================")

	if failed > 0 {
		log.Error().Int("failed", failed).Msg("Some tests failed")
		os.Exit(1)
	}

	log.Info().Msg("All tests passed!")
}

func isContainerRunning(containerName string) bool {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerName)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func testTool(ctx context.Context, containerName string, tool Tool) TestResult {
	result := TestResult{
		Tool: tool,
	}

	// Build docker exec command
	args := []string{"exec", containerName, tool.Cmd}
	args = append(args, tool.Args...)

	cmd := exec.Command("docker", args...)

	// Set timeout
	ctx, cancel := context.WithTimeout(ctx, tool.Timeout)
	defer cancel()

	cmd = exec.CommandContext(ctx, "docker", args...)

	output, err := cmd.CombinedOutput()
	result.Output = string(output)
	result.Error = err

	if err != nil {
		result.Success = false
	} else {
		result.Success = true
	}

	return result
}