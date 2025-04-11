package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ProjectPath represents a project path with its structure
type ProjectPath struct {
	parentFolder string
	childFolder  string
	config       ProjectConfig
}

func (p ProjectPath) FileSystemPath() string {
	return filepath.Join(p.parentFolder, p.childFolder)
}

func (p ProjectPath) TextPath() string {
	return p.parentFolder + "/" + p.childFolder + "/"
}

func (p ProjectPath) Namespace() string {
	return p.config.empresa + "." + p.childFolder
}

// ProjectConfig holds all configuration for project creation
type ProjectConfig struct {
	empresa      string
	solution     string
	solutionName string
	framework    string
}

// TestingFramework represents a supported testing framework
type TestingFramework string

const (
	XUnit TestingFramework = "xunit"
	NUnit TestingFramework = "nunit"
)

var validFrameworks = map[TestingFramework]bool{
	XUnit: true,
	NUnit: true,
}

// IsValid checks if the testing framework is supported
func (t TestingFramework) IsValid() bool {
	return validFrameworks[t]
}

// String returns the framework name as a string
func (t TestingFramework) String() string {
	return string(t)
}

// ParseTestingFramework creates a TestingFramework from string
func ParseTestingFramework(framework string) (TestingFramework, error) {
	tf := TestingFramework(strings.ToLower(framework))
	if !tf.IsValid() {
		return "", fmt.Errorf("unsupported testing framework: %s", framework)
	}
	return tf, nil
}

// CommandExecutor handles command execution
type CommandExecutor struct {
	dryRun bool
}

// ProjectCreator handles project creation
type ProjectCreator struct {
	executor CommandExecutor
	config   ProjectConfig
}

func NewProjectCreator(executor CommandExecutor, config ProjectConfig) ProjectCreator {
	return ProjectCreator{
		executor: executor,
		config:   config,
	}
}

func (pc ProjectCreator) CreateProject(path ProjectPath, projectType string) error {
	return pc.executor.Execute("dotnet", "new", projectType,
		"--output", path.FileSystemPath(),
		"--name", path.Namespace(),
		"--framework", pc.config.framework)
}

func (pc ProjectCreator) CreateTestProject(path ProjectPath, framework TestingFramework) error {
	return pc.CreateProject(path, framework.String())
}

func (pc ProjectCreator) CreateSolution() error {
	return pc.executor.Execute("dotnet", "new", "sln", "--name", pc.config.solution)
}

func (pc ProjectCreator) AddProjectToSolution(path ProjectPath) error {
	return pc.executor.Execute("dotnet", "sln", pc.config.solutionName, "add", path.TextPath())
}

func (pc ProjectCreator) AddProjectReference(from, to ProjectPath) error {
	return pc.executor.Execute("dotnet", "add", from.TextPath(), "reference", to.TextPath())
}
