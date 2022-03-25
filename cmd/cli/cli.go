/*
CLI tooling

- generate ENV-docs
*/
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/cmd/cli/internal/tool"
)

func init() {
	viper.SetDefault("SERVICE_NAME", "cli")

	rootCmd := &cobra.Command{
		Use:   "shortctl",
		Short: "Shortlink it's sandbox for experiments",
		Long:  "Demo microservice architecture and best practices",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().String("o", "./docs/env.md", "Output file path")
	if err := viper.BindPFlag("o", rootCmd.Flags().Lookup("o")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().String("include-dir", "cmd,internal,pkg", "Include directories")
	if err := viper.BindPFlag("include-dir", rootCmd.Flags().Lookup("include-dir")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().String("exclude-dir", "vendor,node_modules,dist,ui", "Exclude directories")
	if err := viper.BindPFlag("exclude-dir", rootCmd.Flags().Lookup("exclude-dir")); err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Generate docs
	if err := doc.GenMarkdownTree(rootCmd, "./docs"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Short", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("Link", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
	if err != nil {
		panic(err)
	}

	dirs := []string{}
	config := Config{}
	filePath := viper.GetString("o")

	includeDirs := viper.GetString("include-dir")
	findDirs := strings.Split(includeDirs, ",")

	excludeDirs := viper.GetString("exclude-dir")
	skipDirs := strings.Split(excludeDirs, ",")

	for _, dir := range findDirs {
		resp, errGetDirectories := tool.GetDirectories(dir, skipDirs)
		if errGetDirectories != nil {
			fmt.Println(errGetDirectories)
			return
		}

		dirs = append(dirs, resp...)
	}

	for _, dir := range dirs {
		pterm.DefaultSection.Printf(`Search in directory %s`, dir)
		config.setConfigDocs(dir, &config)
	}

	payload := config.renderMDTable(config)

	if err := tool.SaveToFile(filePath, payload); err != nil {
		fmt.Println(err)
		return
	}
}

//gocyclo:ignore
func (*Config) setConfigDocs(path string, config *Config) { // nolint:gocognit
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			pterm.Success.Printf("working on file %v\n", fileName)
			ast.Inspect(file, func(n ast.Node) bool {
				if stmt, ok := n.(*ast.ExprStmt); ok { // nolint:nestif
					ast.Inspect(stmt.X, func(n ast.Node) bool {
						if x, ok := n.(*ast.CallExpr); ok {
							if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
								if ident, ok := fun.X.(*ast.Ident); ok {
									if ident.Name == "viper" && fun.Sel.Name == "SetDefault" {
										env := ENV{
											pos:      x.Args[0].(*ast.BasicLit).Pos(),
											fileName: fileName, // nolint:scopelint
											key:      x.Args[0].(*ast.BasicLit).Value,
										}

										switch arg := x.Args[1].(type) {
										case *ast.BasicLit:
											env.value = tool.TrimQuotes(arg.Value)
											env.kind = arg.Kind.String()
										case *ast.Ident:
											if arg.Obj != nil {
												switch variable := arg.Obj.Decl.(type) {
												case *ast.AssignStmt:
													c := variable.Rhs[0].(*ast.CallExpr)

													str := []interface{}{}
													for i := range c.Args {
														str = append(str, tool.TrimQuotes(c.Args[i].(*ast.BasicLit).Value))
													}

													env.value = fmt.Sprintf(str[0].(string), str[1:]...)
												}
											} else {
												env.value = tool.TrimQuotes(arg.Name)
											}
										case *ast.SelectorExpr:
											env.value = tool.TrimQuotes(fmt.Sprintf("%s.%s", arg.X.(*ast.Ident).Name, arg.Sel.Name))
										}

										config.envs = append(config.envs, env)
									}
								}
							}
						}

						return true
					})

					return true
				}

				return true
			})

			for _, comment := range file.Comments {
				for _, item := range comment.List {
					line := fset.Position(item.Pos()).Line

					for index, conf := range config.envs {
						currentLine := fset.Position(conf.pos).Line
						if line == currentLine && fileName == conf.fileName {
							config.envs[index].describe = item.Text[3:] // skip comments symbols
						}
					}
				}
			}
		}
	}
}

func (*Config) renderMDTable(conf Config) string {
	str := `<!---
File generated by cli. DO NOT EDIT.
-->

# ENVIRONMENT

|Name | Default Value | Description |
|---|---|---|
`

	for _, env := range conf.envs {
		str += fmt.Sprintf("| %s | %s | %s |\n", env.key, env.value, env.describe)
	}

	return str
}
