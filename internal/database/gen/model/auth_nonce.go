//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type AuthNonce struct {
	ID         string `sql:"primary_key" db:"id"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	Active     bool   `db:"active"`
	ExpiresAt  string `db:"expires_at"`
	Identifier string `db:"identifier"`
	PublicID   string `db:"public_id"`
	Metadata   []byte `db:"metadata"`
}
