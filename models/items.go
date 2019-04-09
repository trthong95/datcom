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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Item is an object representing the database table.
type Item struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	ItemName null.String `boil:"item_name" json:"item_name,omitempty" toml:"item_name" yaml:"item_name,omitempty"`
	MenuID   null.Int    `boil:"menu_id" json:"menu_id,omitempty" toml:"menu_id" yaml:"menu_id,omitempty"`

	R *itemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L itemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ItemColumns = struct {
	ID       string
	ItemName string
	MenuID   string
}{
	ID:       "id",
	ItemName: "item_name",
	MenuID:   "menu_id",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ItemWhere = struct {
	ID       whereHelperint
	ItemName whereHelpernull_String
	MenuID   whereHelpernull_Int
}{
	ID:       whereHelperint{field: `id`},
	ItemName: whereHelpernull_String{field: `item_name`},
	MenuID:   whereHelpernull_Int{field: `menu_id`},
}

// ItemRels is where relationship names are stored.
var ItemRels = struct {
	Menu   string
	Orders string
}{
	Menu:   "Menu",
	Orders: "Orders",
}

// itemR is where relationships are stored.
type itemR struct {
	Menu   *Menu
	Orders OrderSlice
}

// NewStruct creates a new relationship struct
func (*itemR) NewStruct() *itemR {
	return &itemR{}
}

// itemL is where Load methods for each relationship are stored.
type itemL struct{}

var (
	itemColumns               = []string{"id", "item_name", "menu_id"}
	itemColumnsWithoutDefault = []string{"item_name", "menu_id"}
	itemColumnsWithDefault    = []string{"id"}
	itemPrimaryKeyColumns     = []string{"id"}
)

type (
	// ItemSlice is an alias for a slice of pointers to Item.
	// This should generally be used opposed to []Item.
	ItemSlice []*Item
	// ItemHook is the signature for custom Item hook methods
	ItemHook func(context.Context, boil.ContextExecutor, *Item) error

	itemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	itemType                 = reflect.TypeOf(&Item{})
	itemMapping              = queries.MakeStructMapping(itemType)
	itemPrimaryKeyMapping, _ = queries.BindMapping(itemType, itemMapping, itemPrimaryKeyColumns)
	itemInsertCacheMut       sync.RWMutex
	itemInsertCache          = make(map[string]insertCache)
	itemUpdateCacheMut       sync.RWMutex
	itemUpdateCache          = make(map[string]updateCache)
	itemUpsertCacheMut       sync.RWMutex
	itemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var itemBeforeInsertHooks []ItemHook
var itemBeforeUpdateHooks []ItemHook
var itemBeforeDeleteHooks []ItemHook
var itemBeforeUpsertHooks []ItemHook

var itemAfterInsertHooks []ItemHook
var itemAfterSelectHooks []ItemHook
var itemAfterUpdateHooks []ItemHook
var itemAfterDeleteHooks []ItemHook
var itemAfterUpsertHooks []ItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Item) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Item) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Item) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Item) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Item) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Item) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Item) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Item) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Item) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddItemHook registers your hook function for all future operations.
func AddItemHook(hookPoint boil.HookPoint, itemHook ItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		itemBeforeInsertHooks = append(itemBeforeInsertHooks, itemHook)
	case boil.BeforeUpdateHook:
		itemBeforeUpdateHooks = append(itemBeforeUpdateHooks, itemHook)
	case boil.BeforeDeleteHook:
		itemBeforeDeleteHooks = append(itemBeforeDeleteHooks, itemHook)
	case boil.BeforeUpsertHook:
		itemBeforeUpsertHooks = append(itemBeforeUpsertHooks, itemHook)
	case boil.AfterInsertHook:
		itemAfterInsertHooks = append(itemAfterInsertHooks, itemHook)
	case boil.AfterSelectHook:
		itemAfterSelectHooks = append(itemAfterSelectHooks, itemHook)
	case boil.AfterUpdateHook:
		itemAfterUpdateHooks = append(itemAfterUpdateHooks, itemHook)
	case boil.AfterDeleteHook:
		itemAfterDeleteHooks = append(itemAfterDeleteHooks, itemHook)
	case boil.AfterUpsertHook:
		itemAfterUpsertHooks = append(itemAfterUpsertHooks, itemHook)
	}
}

// One returns a single item record from the query.
func (q itemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Item, error) {
	o := &Item{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Item records from the query.
func (q itemQuery) All(ctx context.Context, exec boil.ContextExecutor) (ItemSlice, error) {
	var o []*Item

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Item slice")
	}

	if len(itemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Item records in the query.
func (q itemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q itemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if items exists")
	}

	return count > 0, nil
}

// Menu pointed to by the foreign key.
func (o *Item) Menu(mods ...qm.QueryMod) menuQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MenuID),
	}

	queryMods = append(queryMods, mods...)

	query := Menus(queryMods...)
	queries.SetFrom(query.Query, "\"menus\"")

	return query
}

// Orders retrieves all the order's Orders with an executor.
func (o *Item) Orders(mods ...qm.QueryMod) orderQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"orders\".\"item_id\"=?", o.ID),
	)

	query := Orders(queryMods...)
	queries.SetFrom(query.Query, "\"orders\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"orders\".*"})
	}

	return query
}

// LoadMenu allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (itemL) LoadMenu(ctx context.Context, e boil.ContextExecutor, singular bool, maybeItem interface{}, mods queries.Applicator) error {
	var slice []*Item
	var object *Item

	if singular {
		object = maybeItem.(*Item)
	} else {
		slice = *maybeItem.(*[]*Item)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &itemR{}
		}
		if !queries.IsNil(object.MenuID) {
			args = append(args, object.MenuID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &itemR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.MenuID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.MenuID) {
				args = append(args, obj.MenuID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`menus`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Menu")
	}

	var resultSlice []*Menu
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Menu")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for menus")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for menus")
	}

	if len(itemAfterSelectHooks) != 0 {
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
		object.R.Menu = foreign
		if foreign.R == nil {
			foreign.R = &menuR{}
		}
		foreign.R.Items = append(foreign.R.Items, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.MenuID, foreign.ID) {
				local.R.Menu = foreign
				if foreign.R == nil {
					foreign.R = &menuR{}
				}
				foreign.R.Items = append(foreign.R.Items, local)
				break
			}
		}
	}

	return nil
}

// LoadOrders allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (itemL) LoadOrders(ctx context.Context, e boil.ContextExecutor, singular bool, maybeItem interface{}, mods queries.Applicator) error {
	var slice []*Item
	var object *Item

	if singular {
		object = maybeItem.(*Item)
	} else {
		slice = *maybeItem.(*[]*Item)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &itemR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &itemR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`orders`), qm.WhereIn(`item_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load orders")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice orders")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Orders = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderR{}
			}
			foreign.R.Item = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ItemID) {
				local.R.Orders = append(local.R.Orders, foreign)
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.Item = local
				break
			}
		}
	}

	return nil
}

// SetMenu of the item to the related item.
// Sets o.R.Menu to related.
// Adds o to related.R.Items.
func (o *Item) SetMenu(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Menu) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
		strmangle.WhereClause("\"", "\"", 2, itemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.MenuID, related.ID)
	if o.R == nil {
		o.R = &itemR{
			Menu: related,
		}
	} else {
		o.R.Menu = related
	}

	if related.R == nil {
		related.R = &menuR{
			Items: ItemSlice{o},
		}
	} else {
		related.R.Items = append(related.R.Items, o)
	}

	return nil
}

// RemoveMenu relationship.
// Sets o.R.Menu to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Item) RemoveMenu(ctx context.Context, exec boil.ContextExecutor, related *Menu) error {
	var err error

	queries.SetScanner(&o.MenuID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("menu_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Menu = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Items {
		if queries.Equal(o.MenuID, ri.MenuID) {
			continue
		}

		ln := len(related.R.Items)
		if ln > 1 && i < ln-1 {
			related.R.Items[i] = related.R.Items[ln-1]
		}
		related.R.Items = related.R.Items[:ln-1]
		break
	}
	return nil
}

// AddOrders adds the given related objects to the existing relationships
// of the item, optionally inserting them as new records.
// Appends related to o.R.Orders.
// Sets related.R.Item appropriately.
func (o *Item) AddOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ItemID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"orders\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"item_id"}),
				strmangle.WhereClause("\"", "\"", 2, orderPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ItemID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &itemR{
			Orders: related,
		}
	} else {
		o.R.Orders = append(o.R.Orders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderR{
				Item: o,
			}
		} else {
			rel.R.Item = o
		}
	}
	return nil
}

// SetOrders removes all previously related items of the
// item replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Item's Orders accordingly.
// Replaces o.R.Orders with related.
// Sets related.R.Item's Orders accordingly.
func (o *Item) SetOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	query := "update \"orders\" set \"item_id\" = null where \"item_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Orders {
			queries.SetScanner(&rel.ItemID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Item = nil
		}

		o.R.Orders = nil
	}
	return o.AddOrders(ctx, exec, insert, related...)
}

// RemoveOrders relationships from objects passed in.
// Removes related items from R.Orders (uses pointer comparison, removal does not keep order)
// Sets related.R.Item.
func (o *Item) RemoveOrders(ctx context.Context, exec boil.ContextExecutor, related ...*Order) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ItemID, nil)
		if rel.R != nil {
			rel.R.Item = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("item_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Orders {
			if rel != ri {
				continue
			}

			ln := len(o.R.Orders)
			if ln > 1 && i < ln-1 {
				o.R.Orders[i] = o.R.Orders[ln-1]
			}
			o.R.Orders = o.R.Orders[:ln-1]
			break
		}
	}

	return nil
}

// Items retrieves all the records using an executor.
func Items(mods ...qm.QueryMod) itemQuery {
	mods = append(mods, qm.From("\"items\""))
	return itemQuery{NewQuery(mods...)}
}

// FindItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Item, error) {
	itemObj := &Item{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, itemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from items")
	}

	return itemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Item) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(itemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	itemInsertCacheMut.RLock()
	cache, cached := itemInsertCache[key]
	itemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			itemColumns,
			itemColumnsWithDefault,
			itemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(itemType, itemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into items")
	}

	if !cached {
		itemInsertCacheMut.Lock()
		itemInsertCache[key] = cache
		itemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Item.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Item) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	itemUpdateCacheMut.RLock()
	cache, cached := itemUpdateCache[key]
	itemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			itemColumns,
			itemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, itemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, append(wl, itemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for items")
	}

	if !cached {
		itemUpdateCacheMut.Lock()
		itemUpdateCache[key] = cache
		itemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q itemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, itemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in item slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all item")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Item) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(itemColumnsWithDefault, o)

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

	itemUpsertCacheMut.RLock()
	cache, cached := itemUpsertCache[key]
	itemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			itemColumns,
			itemColumnsWithDefault,
			itemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			itemColumns,
			itemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(itemPrimaryKeyColumns))
			copy(conflict, itemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(itemType, itemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert items")
	}

	if !cached {
		itemUpsertCacheMut.Lock()
		itemUpsertCache[key] = cache
		itemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Item record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Item) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Item provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), itemPrimaryKeyMapping)
	sql := "DELETE FROM \"items\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q itemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no itemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Item slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(itemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, itemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from item slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for items")
	}

	if len(itemAfterDeleteHooks) != 0 {
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
func (o *Item) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"items\".* FROM \"items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, itemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ItemSlice")
	}

	*o = slice

	return nil
}

// ItemExists checks if the Item row exists.
func ItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"items\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if items exists")
	}

	return exists, nil
}