// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"test_demo/db/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Product{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Product{}) fail: %s", err)
	}
}

func Test_productQuery(t *testing.T) {
	product := newProduct(db)
	product = *product.As(product.TableName())
	_do := product.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(product.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <product> fail:", err)
		return
	}

	_, ok := product.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from product success")
	}

	err = _do.Create(&model.Product{})
	if err != nil {
		t.Error("create item in table <product> fail:", err)
	}

	err = _do.Save(&model.Product{})
	if err != nil {
		t.Error("create item in table <product> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Product{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <product> fail:", err)
	}

	_, err = _do.Select(product.ALL).Take()
	if err != nil {
		t.Error("Take() on table <product> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <product> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <product> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <product> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Product{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <product> fail:", err)
	}

	_, err = _do.Select(product.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <product> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <product> fail:", err)
	}

	_, err = _do.Select(product.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <product> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <product> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <product> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <product> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Product{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <product> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <product> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <product> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <product> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <product> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <product> fail:", err)
	}
}
