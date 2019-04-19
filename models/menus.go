// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Menu is an object representing the database table.
type Menu struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID         int       `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	MenuName        string    `boil:"menu_name" json:"menu_name" toml:"menu_name" yaml:"menu_name"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Deadline        time.Time `boil:"deadline" json:"deadline" toml:"deadline" yaml:"deadline"`
	PaymentReminder time.Time `boil:"payment_reminder" json:"payment_reminder" toml:"payment_reminder" yaml:"payment_reminder"`
	Status          int       `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *menuR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L menuL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MenuColumns = struct {
	ID              string
	OwnerID         string
	MenuName        string
	CreatedAt       string
	Deadline        string
	PaymentReminder string
	Status          string
}{
	ID:              "id",
	OwnerID:         "owner_id",
	MenuName:        "menu_name",
	CreatedAt:       "created_at",
	Deadline:        "deadline",
	PaymentReminder: "payment_reminder",
	Status:          "status",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MenuWhere = struct {
	ID              whereHelperint
	OwnerID         whereHelperint
	MenuName        whereHelperstring
	CreatedAt       whereHelpertime_Time
	Deadline        whereHelpertime_Time
	PaymentReminder whereHelpertime_Time
	Status          whereHelperint
}{
	ID:              whereHelperint{field: `id`},
	OwnerID:         whereHelperint{field: `owner_id`},
	MenuName:        whereHelperstring{field: `menu_name`},
	CreatedAt:       whereHelpertime_Time{field: `created_at`},
	Deadline:        whereHelpertime_Time{field: `deadline`},
	PaymentReminder: whereHelpertime_Time{field: `payment_reminder`},
	Status:          whereHelperint{field: `status`},
}

// MenuRels is where relationship names are stored.
var MenuRels = struct {
	Owner           string
	Items           string
	PeopleInCharges string
}{
	Owner:           "Owner",
	Items:           "Items",
	PeopleInCharges: "PeopleInCharges",
}

// menuR is where relationships are stored.
type menuR struct {
	Owner           *User
	Items           ItemSlice
	PeopleInCharges PeopleInChargeSlice
}

// NewStruct creates a new relationship struct
func (*menuR) NewStruct() *menuR {
	return &menuR{}
}

// menuL is where Load methods for each relationship are stored.
type menuL struct{}

var (
	menuColumns               = []string{"id", "owner_id", "menu_name", "created_at", "deadline", "payment_reminder", "status"}
	menuColumnsWithoutDefault = []string{"owner_id", "menu_name", "created_at", "deadline", "payment_reminder", "status"}
	menuColumnsWithDefault    = []string{"id"}
	menuPrimaryKeyColumns     = []string{"id"}
)

type (
	// MenuSlice is an alias for a slice of pointers to Menu.
	// This should generally be used opposed to []Menu.
	MenuSlice []*Menu
	// MenuHook is the signature for custom Menu hook methods
	MenuHook func(context.Context, boil.ContextExecutor, *Menu) error

	menuQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	menuType                 = reflect.TypeOf(&Menu{})
	menuMapping              = queries.MakeStructMapping(menuType)
	menuPrimaryKeyMapping, _ = queries.BindMapping(menuType, menuMapping, menuPrimaryKeyColumns)
	menuInsertCacheMut       sync.RWMutex
	menuInsertCache          = make(map[string]insertCache)
	menuUpdateCacheMut       sync.RWMutex
	menuUpdateCache          = make(map[string]updateCache)
	menuUpsertCacheMut       sync.RWMutex
	menuUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var menuBeforeInsertHooks []MenuHook
var menuBeforeUpdateHooks []MenuHook
var menuBeforeDeleteHooks []MenuHook
var menuBeforeUpsertHooks []MenuHook

var menuAfterInsertHooks []MenuHook
var menuAfterSelectHooks []MenuHook
var menuAfterUpdateHooks []MenuHook
var menuAfterDeleteHooks []MenuHook
var menuAfterUpsertHooks []MenuHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Menu) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Menu) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Menu) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Menu) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Menu) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Menu) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Menu) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Menu) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Menu) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range menuAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMenuHook registers your hook function for all future operations.
func AddMenuHook(hookPoint boil.HookPoint, menuHook MenuHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		menuBeforeInsertHooks = append(menuBeforeInsertHooks, menuHook)
	case boil.BeforeUpdateHook:
		menuBeforeUpdateHooks = append(menuBeforeUpdateHooks, menuHook)
	case boil.BeforeDeleteHook:
		menuBeforeDeleteHooks = append(menuBeforeDeleteHooks, menuHook)
	case boil.BeforeUpsertHook:
		menuBeforeUpsertHooks = append(menuBeforeUpsertHooks, menuHook)
	case boil.AfterInsertHook:
		menuAfterInsertHooks = append(menuAfterInsertHooks, menuHook)
	case boil.AfterSelectHook:
		menuAfterSelectHooks = append(menuAfterSelectHooks, menuHook)
	case boil.AfterUpdateHook:
		menuAfterUpdateHooks = append(menuAfterUpdateHooks, menuHook)
	case boil.AfterDeleteHook:
		menuAfterDeleteHooks = append(menuAfterDeleteHooks, menuHook)
	case boil.AfterUpsertHook:
		menuAfterUpsertHooks = append(menuAfterUpsertHooks, menuHook)
	}
}

// One returns a single menu record from the query.
func (q menuQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Menu, error) {
	o := &Menu{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for menus")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Menu records from the query.
func (q menuQuery) All(ctx context.Context, exec boil.ContextExecutor) (MenuSlice, error) {
	var o []*Menu

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Menu slice")
	}

	if len(menuAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Menu records in the query.
func (q menuQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count menus rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q menuQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if menus exists")
	}

	return count > 0, nil
}

// Owner pointed to by the foreign key.
func (o *Menu) Owner(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.OwnerID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Items retrieves all the item's Items with an executor.
func (o *Menu) Items(mods ...qm.QueryMod) itemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"items\".\"menu_id\"=?", o.ID),
	)

	query := Items(queryMods...)
	queries.SetFrom(query.Query, "\"items\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"items\".*"})
	}

	return query
}

// PeopleInCharges retrieves all the people_in_charge's PeopleInCharges with an executor.
func (o *Menu) PeopleInCharges(mods ...qm.QueryMod) peopleInChargeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"people_in_charge\".\"menu_id\"=?", o.ID),
	)

	query := PeopleInCharges(queryMods...)
	queries.SetFrom(query.Query, "\"people_in_charge\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"people_in_charge\".*"})
	}

	return query
}

// LoadOwner allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (menuL) LoadOwner(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		object = maybeMenu.(*Menu)
	} else {
		slice = *maybeMenu.(*[]*Menu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.OwnerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.OwnerID {
					continue Outer
				}
			}

			args = append(args, obj.OwnerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(menuAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Owner = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.OwnerMenus = append(foreign.R.OwnerMenus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OwnerID == foreign.ID {
				local.R.Owner = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OwnerMenus = append(foreign.R.OwnerMenus, local)
				break
			}
		}
	}

	return nil
}

// LoadItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (menuL) LoadItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		object = maybeMenu.(*Menu)
	} else {
		slice = *maybeMenu.(*[]*Menu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`items`), qm.WhereIn(`menu_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load items")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if len(itemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Items = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &itemR{}
			}
			foreign.R.Menu = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MenuID {
				local.R.Items = append(local.R.Items, foreign)
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.Menu = local
				break
			}
		}
	}

	return nil
}

// LoadPeopleInCharges allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (menuL) LoadPeopleInCharges(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		object = maybeMenu.(*Menu)
	} else {
		slice = *maybeMenu.(*[]*Menu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`people_in_charge`), qm.WhereIn(`menu_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load people_in_charge")
	}

	var resultSlice []*PeopleInCharge
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice people_in_charge")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on people_in_charge")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for people_in_charge")
	}

	if len(peopleInChargeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PeopleInCharges = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &peopleInChargeR{}
			}
			foreign.R.Menu = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MenuID {
				local.R.PeopleInCharges = append(local.R.PeopleInCharges, foreign)
				if foreign.R == nil {
					foreign.R = &peopleInChargeR{}
				}
				foreign.R.Menu = local
				break
			}
		}
	}

	return nil
}

// SetOwner of the menu to the related item.
// Sets o.R.Owner to related.
// Adds o to related.R.OwnerMenus.
func (o *Menu) SetOwner(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"menus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"owner_id"}),
		strmangle.WhereClause("\"", "\"", 2, menuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OwnerID = related.ID
	if o.R == nil {
		o.R = &menuR{
			Owner: related,
		}
	} else {
		o.R.Owner = related
	}

	if related.R == nil {
		related.R = &userR{
			OwnerMenus: MenuSlice{o},
		}
	} else {
		related.R.OwnerMenus = append(related.R.OwnerMenus, o)
	}

	return nil
}

// AddItems adds the given related objects to the existing relationships
// of the menu, optionally inserting them as new records.
// Appends related to o.R.Items.
// Sets related.R.Menu appropriately.
func (o *Menu) AddItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Item) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MenuID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
				strmangle.WhereClause("\"", "\"", 2, itemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MenuID = o.ID
		}
	}

	if o.R == nil {
		o.R = &menuR{
			Items: related,
		}
	} else {
		o.R.Items = append(o.R.Items, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &itemR{
				Menu: o,
			}
		} else {
			rel.R.Menu = o
		}
	}
	return nil
}

// AddPeopleInCharges adds the given related objects to the existing relationships
// of the menu, optionally inserting them as new records.
// Appends related to o.R.PeopleInCharges.
// Sets related.R.Menu appropriately.
func (o *Menu) AddPeopleInCharges(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PeopleInCharge) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MenuID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"people_in_charge\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
				strmangle.WhereClause("\"", "\"", 2, peopleInChargePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MenuID = o.ID
		}
	}

	if o.R == nil {
		o.R = &menuR{
			PeopleInCharges: related,
		}
	} else {
		o.R.PeopleInCharges = append(o.R.PeopleInCharges, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &peopleInChargeR{
				Menu: o,
			}
		} else {
			rel.R.Menu = o
		}
	}
	return nil
}

// Menus retrieves all the records using an executor.
func Menus(mods ...qm.QueryMod) menuQuery {
	mods = append(mods, qm.From("\"menus\""))
	return menuQuery{NewQuery(mods...)}
}

// FindMenu retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMenu(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Menu, error) {
	menuObj := &Menu{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"menus\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, menuObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from menus")
	}

	return menuObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Menu) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menus provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	menuInsertCacheMut.RLock()
	cache, cached := menuInsertCache[key]
	menuInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			menuColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(menuType, menuMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"menus\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"menus\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into menus")
	}

	if !cached {
		menuInsertCacheMut.Lock()
		menuInsertCache[key] = cache
		menuInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Menu.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Menu) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	menuUpdateCacheMut.RLock()
	cache, cached := menuUpdateCache[key]
	menuUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			menuColumns,
			menuPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update menus, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"menus\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, menuPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, append(wl, menuPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update menus row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for menus")
	}

	if !cached {
		menuUpdateCacheMut.Lock()
		menuUpdateCache[key] = cache
		menuUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q menuQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for menus")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MenuSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"menus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, menuPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all menu")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Menu) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menus provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	menuUpsertCacheMut.RLock()
	cache, cached := menuUpsertCache[key]
	menuUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			menuColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			menuColumns,
			menuPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert menus, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(menuPrimaryKeyColumns))
			copy(conflict, menuPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"menus\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(menuType, menuMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert menus")
	}

	if !cached {
		menuUpsertCacheMut.Lock()
		menuUpsertCache[key] = cache
		menuUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Menu record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Menu) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Menu provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), menuPrimaryKeyMapping)
	sql := "DELETE FROM \"menus\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for menus")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q menuQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no menuQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MenuSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Menu slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(menuBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"menus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menus")
	}

	if len(menuAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Menu) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMenu(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MenuSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MenuSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"menus\".* FROM \"menus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MenuSlice")
	}

	*o = slice

	return nil
}

// MenuExists checks if the Menu row exists.
func MenuExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"menus\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if menus exists")
	}

	return exists, nil
}
