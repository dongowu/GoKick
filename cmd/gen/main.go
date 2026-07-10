package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dongowu/gokick/internal/gen"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: gokick gen <entity> [--module=<name>]")
	}

	if args[0] != "gen" && args[0] != "generate" {
		return fmt.Errorf("unknown command: %s\nSupported: gen|generate", args[0])
	}

	if len(args) < 2 {
		return fmt.Errorf("usage: gokick gen <entity> [--module=<name>]")
	}

	entity := args[1]
	module := deriveModule(entity)

	// 解析额外参数
	for i := 2; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--module=") {
			module = strings.TrimPrefix(args[i], "--module=")
		}
	}

	data := gen.TemplateData{
		Entity:      strings.Title(entity),
		Module:      module,
		EntityLower: strings.ToLower(entity),
	}

	// 生成 handler
	handlerContent, err := gen.GenerateHandler(data)
	if err != nil {
		return fmt.Errorf("generate handler: %w", err)
	}
	if err := gen.WriteFile(fmt.Sprintf("internal/handler/%s.go", strings.ToLower(entity)), handlerContent); err != nil {
		return fmt.Errorf("write handler: %w", err)
	}

	// 生成 service
	serviceContent, err := gen.GenerateService(data)
	if err != nil {
		return fmt.Errorf("generate service: %w", err)
	}
	if err := gen.WriteFile(fmt.Sprintf("internal/service/%s.go", strings.ToLower(entity)), serviceContent); err != nil {
		return fmt.Errorf("write service: %w", err)
	}

	// 生成 repository
	repoContent, err := gen.GenerateRepository(data)
	if err != nil {
		return fmt.Errorf("generate repository: %w", err)
	}
	if err := gen.WriteFile(fmt.Sprintf("internal/repository/%s/%s.go", module, strings.ToLower(entity)), repoContent); err != nil {
		return fmt.Errorf("write repository: %w", err)
	}

	// 生成 model
	modelContent, err := gen.GenerateModel(data)
	if err != nil {
		return fmt.Errorf("generate model: %w", err)
	}
	if err := gen.WriteFile(fmt.Sprintf("internal/model/%s.go", strings.ToLower(entity)), modelContent); err != nil {
		return fmt.Errorf("write model: %w", err)
	}

	fmt.Printf("✅ Generated %s module:\n", module)
	fmt.Printf("  - internal/handler/%s.go\n", strings.ToLower(entity))
	fmt.Printf("  - internal/service/%s.go\n", strings.ToLower(entity))
	fmt.Printf("  - internal/repository/%s/%s.go\n", module, strings.ToLower(entity))
	fmt.Printf("  - internal/model/%s.go\n", strings.ToLower(entity))

	return nil
}

func deriveModule(entity string) string {
	entity = strings.ToLower(entity)
	var result strings.Builder
	for i, r := range entity {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(r - 'A' + 'a')
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
