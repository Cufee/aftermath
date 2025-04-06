// package database
//go:build exclude

package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-jet/jet/v2/generator/metadata"
	generator "github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	"github.com/go-jet/jet/v2/postgres"

	_ "github.com/lib/pq"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	out := filepath.Join(basepath, "gen")

	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = generator.GenerateDB(
		db,
		"public",
		out,
		template.Default(postgres.Dialect).
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
