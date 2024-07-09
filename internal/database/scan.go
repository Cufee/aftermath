package database

import (
	"database/sql"

	"reflect"

	"github.com/pkg/errors"
)

type scanAssignable interface {
	ScanValues(columns []string) ([]any, error)
	AssignValues(columns []string, values []any) error
}

func rowsToRecords[R scanAssignable](rs *sql.Rows, columns []string) ([]R, error) {
	var records []R

	recordType := reflect.TypeOf((*R)(nil)).Elem()
	if recordType.Kind() != reflect.Ptr {
		return nil, errors.New("type parameter must be a pointer to a struct")
	}

	for rs.Next() {
		recordPtr := reflect.New(recordType.Elem())
		record := recordPtr.Interface().(R)

		values, err := record.ScanValues(columns)
		if err != nil {
			return nil, err
		}
		if err := rs.Scan(values...); err != nil {
			return nil, err
		}

		err = record.AssignValues(columns, values)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
