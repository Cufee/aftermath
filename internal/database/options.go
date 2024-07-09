package database

import "time"

type baseQueryOptions struct {
	referenceIDIn    map[string]struct{}
	referenceIDNotIn map[string]struct{}

	createdAfter  *time.Time
	createdBefore *time.Time

	fields []string
}

type Query func(*baseQueryOptions)

// --- snapshot query options ---

/*
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id in (ids) will be added to the query
*/
func WithReferenceIDIn(ids ...string) Query {
	return func(q *baseQueryOptions) {
		if q.referenceIDIn == nil {
			q.referenceIDIn = make(map[string]struct{})
		}
		for _, id := range ids {
			q.referenceIDIn[id] = struct{}{}
		}
	}
}

/*
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id not in (ids) will be added to the query
*/
func WithReferenceIDNotIn(ids ...string) Query {
	return func(q *baseQueryOptions) {
		if q.referenceIDNotIn == nil {
			q.referenceIDNotIn = make(map[string]struct{})
		}
		for _, id := range ids {
			q.referenceIDNotIn[id] = struct{}{}
		}
	}
}

/*
Adds a created_at lt constraint
  - if this constraint is set, records will be sorted by created_at ASC
*/
func WithCreatedAfter(after time.Time) Query {
	return func(q *baseQueryOptions) {
		q.createdAfter = &after
	}
}

/*
Adds a created_at gt constraint
  - if this constraint is set, records will be sorted by created_at DESC
*/
func WithCreatedBefore(before time.Time) Query {
	return func(q *baseQueryOptions) {
		q.createdBefore = &before
	}
}

/*
Set fields that will be selected.
  - Some fields like id, created_at, updated_at will always be selected
  - Passing 0 length fields will result in select *
*/
func WithSelect(fields ...string) Query {
	return func(q *baseQueryOptions) {
		// make sure fields are unique
		fieldsSet := make(map[string]struct{})
		for _, field := range append(q.fields, fields...) {
			fieldsSet[field] = struct{}{}
		}

		// passing 0 length fields will result in select *
		q.fields = nil
		for field := range fieldsSet {
			q.fields = append(q.fields, field)
		}
	}
}

/*
Returns a slice of fields for a select statement with all duplicates removed and required fields added
  - if query.fields is nil, returns a nil slice
*/
func (q *baseQueryOptions) selectFields(required ...string) []string {
	if q.fields == nil {
		return nil
	}

	// maker sure all required fields are part of the slice
	var requiredFields = make(map[string]struct{})
	for _, field := range required {
		requiredFields[field] = struct{}{}
	}

	for _, field := range q.fields {
		for required := range requiredFields {
			if field == required {
				delete(requiredFields, required)
			}
		}
	}
	for field := range requiredFields {
		q.fields = append(q.fields, field)
	}

	return q.fields
}

func (q *baseQueryOptions) refIDIn() []any {
	if len(q.referenceIDIn) == 0 {
		return nil
	}
	var ids []any
	for id := range q.referenceIDIn {
		ids = append(ids, id)
	}
	return ids
}

func (q *baseQueryOptions) refIDNotIn() []any {
	if len(q.referenceIDNotIn) == 0 {
		return nil
	}
	var ids []any
	for id := range q.referenceIDNotIn {
		ids = append(ids, id)
	}
	return ids
}
