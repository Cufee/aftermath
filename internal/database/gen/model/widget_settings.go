//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type WidgetSettings struct {
	ID                 string  `sql:"primary_key" db:"id"`
	CreatedAt          string  `db:"created_at"`
	UpdatedAt          string  `db:"updated_at"`
	ReferenceID        string  `db:"reference_id"`
	Title              *string `db:"title"`
	SessionFrom        *string `db:"session_from"`
	Metadata           []byte  `db:"metadata"`
	Styles             []byte  `db:"styles"`
	UserID             string  `db:"user_id"`
	SessionReferenceID *string `db:"session_reference_id"`
}
