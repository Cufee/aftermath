//go:build exclude

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-jet/jet/v2/generator/metadata"
	generator "github.com/go-jet/jet/v2/generator/sqlite"
	"github.com/go-jet/jet/v2/generator/template"
	"github.com/go-jet/jet/v2/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	dsn := filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME"))
	out := filepath.Join(basepath, "gen")

	err := generator.GenerateDSN(
		dsn,
		out,
		template.Default(sqlite.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									return defaultTableModelField.UseTags(
										fmt.Sprintf(`db:"%s"`, columnMetaData.Name),
									)
								})
						}),
					)
			}),
	)
	if err != nil {
		panic(err)
	}
}
