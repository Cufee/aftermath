// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package api

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"errors"
	"net/http"

	"github.com/cufee/aftermath/cmd/frontend/components/connections"
	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
)

var RemoveConnection handler.Endpoint = func(ctx *handler.Context) error {
	user, err := ctx.SessionUser()
	if err != nil {
		return ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	connID := ctx.Path("connectionId")
	if connID == "" {
		return ctx.Error(errors.New("connection id path parameter is empty"))
	}

	err = ctx.Database().DeleteUserConnection(ctx, user.ID, connID)
	if err != nil {
		return ctx.Error(err, "failed to delete a connection")
	}

	return nil
}

var SetDefaultConnection handler.Partial = func(ctx *handler.Context) (templ.Component, error) {
	user, err := ctx.SessionUser(database.WithConnections())
	if err != nil {
		return nil, ctx.Redirect("/login", http.StatusTemporaryRedirect)
	}

	connID := ctx.Path("connectionId")
	if connID == "" {
		return nil, ctx.Error(errors.New("connection id path parameter is empty"))
	}

	var previous, updated models.UserConnection
	for _, conn := range user.Connections {
		if conn.Metadata["default"] == true {
			previous = conn
		}

		conn.Metadata["default"] = conn.ID == connID
		u, err := ctx.Database().UpdateConnection(ctx, conn)
		if err != nil {
			return nil, ctx.Error(err, "failed to update a connection")
		}

		if conn.ID == connID {
			updated = u
		}
	}
	var ids []string
	if previous.ReferenceID != "" {
		ids = append(ids, previous.ReferenceID)
	}
	if updated.ReferenceID != "" {
		ids = append(ids, updated.ReferenceID)
	}
	if len(ids) < 1 {
		ctx.SetStatus(http.StatusNotFound)
		return nil, ctx.Error(errors.New("connection not found"), "connection not found")
	}

	accounts, err := ctx.Database().GetAccounts(ctx.Context, ids)
	if err != nil {
		return nil, ctx.Error(err, "failed to find connected accounts")
	}

	var withProps []connections.ConnectionWithAccount
	for _, a := range accounts {
		println(previous.ID, updated.ID, a.ID)
		if previous.ReferenceID == a.ID {
			withProps = append(withProps, connections.ConnectionWithAccount{
				UserConnection: previous,
				Account:        a,
			})
		}
		if updated.ReferenceID == a.ID {
			withProps = append(withProps, connections.ConnectionWithAccount{
				UserConnection: updated,
				Account:        a,
			})
		}
	}
	return updatedConnections(withProps, updated.ID), nil
}

func updatedConnections(conns []connections.ConnectionWithAccount, defaultID string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, c := range conns {
			templ_7745c5c3_Err = connections.ConnectionCard(c, defaultID).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}
