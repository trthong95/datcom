// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testItems(t *testing.T) {
	t.Parallel()

	query := Items()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Items().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Item exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ItemExists to return true, but got false.")
	}
}

func testItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	itemFound, err := FindItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if itemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Items().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Items().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	itemOne := &Item{}
	itemTwo := &Item{}
	if err = randomize.Struct(seed, itemOne, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}
	if err = randomize.Struct(seed, itemTwo, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = itemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = itemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Items().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	itemOne := &Item{}
	itemTwo := &Item{}
	if err = randomize.Struct(seed, itemOne, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}
	if err = randomize.Struct(seed, itemTwo, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = itemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = itemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func itemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func itemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Item) error {
	*o = Item{}
	return nil
}

func testItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Item{}
	o := &Item{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, itemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Item object: %s", err)
	}

	AddItemHook(boil.BeforeInsertHook, itemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	itemBeforeInsertHooks = []ItemHook{}

	AddItemHook(boil.AfterInsertHook, itemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	itemAfterInsertHooks = []ItemHook{}

	AddItemHook(boil.AfterSelectHook, itemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	itemAfterSelectHooks = []ItemHook{}

	AddItemHook(boil.BeforeUpdateHook, itemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	itemBeforeUpdateHooks = []ItemHook{}

	AddItemHook(boil.AfterUpdateHook, itemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	itemAfterUpdateHooks = []ItemHook{}

	AddItemHook(boil.BeforeDeleteHook, itemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	itemBeforeDeleteHooks = []ItemHook{}

	AddItemHook(boil.AfterDeleteHook, itemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	itemAfterDeleteHooks = []ItemHook{}

	AddItemHook(boil.BeforeUpsertHook, itemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	itemBeforeUpsertHooks = []ItemHook{}

	AddItemHook(boil.AfterUpsertHook, itemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	itemAfterUpsertHooks = []ItemHook{}
}

func testItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(itemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testItemToManyOrders(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b, c Order

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ItemID, a.ID)
	queries.Assign(&c.ItemID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Orders().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ItemID, b.ItemID) {
			bFound = true
		}
		if queries.Equal(v.ItemID, c.ItemID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ItemSlice{&a}
	if err = a.L.LoadOrders(ctx, tx, false, (*[]*Item)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Orders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Orders = nil
	if err = a.L.LoadOrders(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Orders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testItemToManyAddOpOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b, c, d, e Order

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Order{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Order{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrders(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ItemID) {
			t.Error("foreign key was wrong value", a.ID, first.ItemID)
		}
		if !queries.Equal(a.ID, second.ItemID) {
			t.Error("foreign key was wrong value", a.ID, second.ItemID)
		}

		if first.R.Item != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Item != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Orders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Orders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Orders().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testItemToManySetOpOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b, c, d, e Order

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Order{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetOrders(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Orders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetOrders(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Orders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ItemID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ItemID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ItemID) {
		t.Error("foreign key was wrong value", a.ID, d.ItemID)
	}
	if !queries.Equal(a.ID, e.ItemID) {
		t.Error("foreign key was wrong value", a.ID, e.ItemID)
	}

	if b.R.Item != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Item != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Item != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Item != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Orders[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Orders[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testItemToManyRemoveOpOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b, c, d, e Order

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Order{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddOrders(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Orders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveOrders(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Orders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ItemID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ItemID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Item != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Item != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Item != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Item != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Orders) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Orders[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Orders[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testItemToOneMenuUsingMenu(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Item
	var foreign Menu

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, menuDBTypes, false, menuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Menu struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.MenuID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Menu().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ItemSlice{&local}
	if err = local.L.LoadMenu(ctx, tx, false, (*[]*Item)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Menu = nil
	if err = local.L.LoadMenu(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testItemToOneSetOpMenuUsingMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b, c Menu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Menu{&b, &c} {
		err = a.SetMenu(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Menu != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Items[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.MenuID, x.ID) {
			t.Error("foreign key was wrong value", a.MenuID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MenuID))
		reflect.Indirect(reflect.ValueOf(&a.MenuID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.MenuID, x.ID) {
			t.Error("foreign key was wrong value", a.MenuID, x.ID)
		}
	}
}

func testItemToOneRemoveOpMenuUsingMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Item
	var b Menu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetMenu(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveMenu(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Menu().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Menu != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.MenuID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Items) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Items().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	itemDBTypes = map[string]string{`ID`: `integer`, `ItemName`: `character varying`, `MenuID`: `integer`}
	_           = bytes.MinRead
)

func testItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(itemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(itemColumns) == len(itemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, itemDBTypes, true, itemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(itemColumns) == len(itemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Item{}
	if err = randomize.Struct(seed, o, itemDBTypes, true, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, itemDBTypes, true, itemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(itemColumns, itemPrimaryKeyColumns) {
		fields = itemColumns
	} else {
		fields = strmangle.SetComplement(
			itemColumns,
			itemPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(itemColumns) == len(itemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Item{}
	if err = randomize.Struct(seed, &o, itemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Item: %s", err)
	}

	count, err := Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, itemDBTypes, false, itemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Item: %s", err)
	}

	count, err = Items().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}