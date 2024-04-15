package main

import (
	"gorm.io/gen"
	a "test-gorm-gen/a/models"
	b "test-gorm-gen/b/models"
)

func main() {
	c := gen.Config{
		OutPath:      "query",
		OutFile:      "",
		ModelPkgPath: "models",
		Mode:         gen.WithQueryInterface,
	}
	c.WithImportPkgPath(
		`a "test-gorm-gen/a/models"`,
		`b "test-gorm-gen/b/models"`,
	)
	g := gen.NewGenerator(c)
	g.ApplyBasic(a.A{}, b.B{})
	g.Execute()
}
