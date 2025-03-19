package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// NewProjectConfig creates a default configuration
func NewProjectConfig(empresa string, solutionNameFlag string) ProjectConfig {
	return ProjectConfig{
		empresa:      empresa,
		solution:     solutionNameFlag,
		solutionName: solutionNameFlag + ".sln",
		framework:    "net9.0",
	}
}

// NewProjectPath creates a new project path
func NewProjectPath(config ProjectConfig, parentFolder, childFolder string) ProjectPath {
	return ProjectPath{
		parentFolder: parentFolder,
		childFolder:  childFolder,
		config:       config,
	}
}

func (e CommandExecutor) Execute(cmd string, args ...string) error {
	fmt.Printf("Ejecutando: %s %s\n", cmd, strings.Join(args, " "))
	if e.dryRun {
		return nil
	}
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func CreateSimpleSolution(config ProjectConfig, testingFramework TestingFramework) error {
	executor := CommandExecutor{dryRun: false}
	creator := NewProjectCreator(executor, config)

	// Define project paths
	paths := struct {
		api         ProjectPath
		domain      ProjectPath
		application ProjectPath
		infra       ProjectPath
		testDomain  ProjectPath
		testApp     ProjectPath
	}{
		api:         NewProjectPath(config, "Presentation", "Api"),
		domain:      NewProjectPath(config, "Core", "Domain"),
		application: NewProjectPath(config, "Core", "Application"),
		infra:       NewProjectPath(config, "Infrastructure", "Services"),
		testDomain:  NewProjectPath(config, "Test", "DomainTest"),
		testApp:     NewProjectPath(config, "Test", "ApplicationTest"),
	}

	// Create projects
	if err := creator.CreateProject(paths.api, "webapi"); err != nil {
		return fmt.Errorf("error creating API: %v", err)
	}

	for _, p := range []struct {
		path ProjectPath
		name string
	}{
		{paths.domain, "Domain"},
		{paths.application, "Application"},
		{paths.infra, "Infrastructure"},
	} {
		if err := creator.CreateProject(p.path, "classlib"); err != nil {
			return fmt.Errorf("error creating %s: %v", p.name, err)
		}
	}

	// Create test projects
	for _, p := range []struct {
		path ProjectPath
		name string
	}{
		{paths.testDomain, "Domain Test"},
		{paths.testApp, "Application Test"},
	} {
		if err := creator.CreateTestProject(p.path, testingFramework); err != nil {
			return fmt.Errorf("error creating %s: %v", p.name, err)
		}
	}

	// Create and configure solution
	if err := creator.CreateSolution(); err != nil {
		return fmt.Errorf("error creating solution: %v", err)
	}

	// Add projects to solution
	for _, p := range []ProjectPath{
		paths.api, paths.domain, paths.application, paths.infra,
		paths.testDomain, paths.testApp,
	} {
		if err := creator.AddProjectToSolution(p); err != nil {
			return fmt.Errorf("error adding project to solution: %v", err)
		}
	}

	// Add project references
	references := []struct {
		from, to ProjectPath
	}{
		{paths.testDomain, paths.domain},
		{paths.testApp, paths.application},
		{paths.application, paths.domain},
		{paths.application, paths.infra},
		{paths.api, paths.application},
	}

	for _, ref := range references {
		if err := creator.AddProjectReference(ref.from, ref.to); err != nil {
			return fmt.Errorf("error adding reference: %v", err)
		}
	}

	return creator.executor.Execute("dotnet", "build")
}

// go run creator.go -empresa Cynergy -solution DemoVps
func main() {
	empresaFlag := flag.String("empresa", "MyCompany", "Nombre de la empresa")
	solutionNameFlag := flag.String("solution", "KataSolution", "Nombre de la solución")
	testFrameworkFlag := flag.String("test", "nunit", "Framework de testing (xunit/nunit)")
	flag.Parse()

	config := NewProjectConfig(*empresaFlag, *solutionNameFlag)

	testFramework, err := ParseTestingFramework(*testFrameworkFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Creando estructura de solución .NET...")
	if err := CreateSimpleSolution(config, testFramework); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("¡Solución creada exitosamente!")
}
