// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/zlllgp/vegas/internal/dal/model"
)

func newActivity(db *gorm.DB, opts ...gen.DOOption) activity {
	_activity := activity{}

	_activity.activityDo.UseDB(db, opts...)
	_activity.activityDo.UseModel(&model.Activity{})

	tableName := _activity.activityDo.TableName()
	_activity.ALL = field.NewAsterisk(tableName)
	_activity.ID = field.NewUint(tableName, "id")
	_activity.CreatedAt = field.NewTime(tableName, "created_at")
	_activity.UpdatedAt = field.NewTime(tableName, "updated_at")
	_activity.DeletedAt = field.NewField(tableName, "deleted_at")
	_activity.CreatorId = field.NewInt64(tableName, "creator_id")
	_activity.CreatorName = field.NewString(tableName, "creator_name")
	_activity.Name = field.NewString(tableName, "name")
	_activity.TenantId = field.NewUint8(tableName, "tenant_id not null")

	_activity.fillFieldMap()

	return _activity
}

type activity struct {
	activityDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	CreatorId   field.Int64
	CreatorName field.String
	Name        field.String
	TenantId    field.Uint8

	fieldMap map[string]field.Expr
}

func (a activity) Table(newTableName string) *activity {
	a.activityDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activity) As(alias string) *activity {
	a.activityDo.DO = *(a.activityDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activity) updateTableName(table string) *activity {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.CreatorId = field.NewInt64(table, "creator_id")
	a.CreatorName = field.NewString(table, "creator_name")
	a.Name = field.NewString(table, "name")
	a.TenantId = field.NewUint8(table, "tenant_id not null")

	a.fillFieldMap()

	return a
}

func (a *activity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activity) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["creator_id"] = a.CreatorId
	a.fieldMap["creator_name"] = a.CreatorName
	a.fieldMap["name"] = a.Name
	a.fieldMap["tenant_id not null"] = a.TenantId
}

func (a activity) clone(db *gorm.DB) activity {
	a.activityDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activity) replaceDB(db *gorm.DB) activity {
	a.activityDo.ReplaceDB(db)
	return a
}

type activityDo struct{ gen.DO }

type IActivityDo interface {
	gen.SubQuery
	Debug() IActivityDo
	WithContext(ctx context.Context) IActivityDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActivityDo
	WriteDB() IActivityDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActivityDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActivityDo
	Not(conds ...gen.Condition) IActivityDo
	Or(conds ...gen.Condition) IActivityDo
	Select(conds ...field.Expr) IActivityDo
	Where(conds ...gen.Condition) IActivityDo
	Order(conds ...field.Expr) IActivityDo
	Distinct(cols ...field.Expr) IActivityDo
	Omit(cols ...field.Expr) IActivityDo
	Join(table schema.Tabler, on ...field.Expr) IActivityDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActivityDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActivityDo
	Group(cols ...field.Expr) IActivityDo
	Having(conds ...gen.Condition) IActivityDo
	Limit(limit int) IActivityDo
	Offset(offset int) IActivityDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityDo
	Unscoped() IActivityDo
	Create(values ...*model.Activity) error
	CreateInBatches(values []*model.Activity, batchSize int) error
	Save(values ...*model.Activity) error
	First() (*model.Activity, error)
	Take() (*model.Activity, error)
	Last() (*model.Activity, error)
	Find() ([]*model.Activity, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Activity, err error)
	FindInBatches(result *[]*model.Activity, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Activity) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActivityDo
	Assign(attrs ...field.AssignExpr) IActivityDo
	Joins(fields ...field.RelationField) IActivityDo
	Preload(fields ...field.RelationField) IActivityDo
	FirstOrInit() (*model.Activity, error)
	FirstOrCreate() (*model.Activity, error)
	FindByPage(offset int, limit int) (result []*model.Activity, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActivityDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a activityDo) Debug() IActivityDo {
	return a.withDO(a.DO.Debug())
}

func (a activityDo) WithContext(ctx context.Context) IActivityDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityDo) ReadDB() IActivityDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityDo) WriteDB() IActivityDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityDo) Session(config *gorm.Session) IActivityDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityDo) Clauses(conds ...clause.Expression) IActivityDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityDo) Returning(value interface{}, columns ...string) IActivityDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityDo) Not(conds ...gen.Condition) IActivityDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityDo) Or(conds ...gen.Condition) IActivityDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityDo) Select(conds ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityDo) Where(conds ...gen.Condition) IActivityDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityDo) Order(conds ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityDo) Distinct(cols ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityDo) Omit(cols ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityDo) Join(table schema.Tabler, on ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActivityDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityDo) RightJoin(table schema.Tabler, on ...field.Expr) IActivityDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityDo) Group(cols ...field.Expr) IActivityDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityDo) Having(conds ...gen.Condition) IActivityDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityDo) Limit(limit int) IActivityDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityDo) Offset(offset int) IActivityDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityDo) Unscoped() IActivityDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityDo) Create(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityDo) CreateInBatches(values []*model.Activity, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityDo) Save(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityDo) First() (*model.Activity, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Take() (*model.Activity, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Last() (*model.Activity, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Find() ([]*model.Activity, error) {
	result, err := a.DO.Find()
	return result.([]*model.Activity), err
}

func (a activityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Activity, err error) {
	buf := make([]*model.Activity, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityDo) FindInBatches(result *[]*model.Activity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityDo) Attrs(attrs ...field.AssignExpr) IActivityDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityDo) Assign(attrs ...field.AssignExpr) IActivityDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityDo) Joins(fields ...field.RelationField) IActivityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityDo) Preload(fields ...field.RelationField) IActivityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityDo) FirstOrInit() (*model.Activity, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FirstOrCreate() (*model.Activity, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FindByPage(offset int, limit int) (result []*model.Activity, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a activityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityDo) Delete(models ...*model.Activity) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityDo) withDO(do gen.Dao) *activityDo {
	a.DO = *do.(*gen.DO)
	return a
}
