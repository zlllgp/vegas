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

	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
)

func newRuleMetum(db *gorm.DB, opts ...gen.DOOption) ruleMetum {
	_ruleMetum := ruleMetum{}

	_ruleMetum.ruleMetumDo.UseDB(db, opts...)
	_ruleMetum.ruleMetumDo.UseModel(&model.RuleMetum{})

	tableName := _ruleMetum.ruleMetumDo.TableName()
	_ruleMetum.ALL = field.NewAsterisk(tableName)
	_ruleMetum.ID = field.NewInt64(tableName, "id")
	_ruleMetum.CreatedAt = field.NewTime(tableName, "created_at")
	_ruleMetum.UpdatedAt = field.NewTime(tableName, "updated_at")
	_ruleMetum.DeletedAt = field.NewField(tableName, "deleted_at")
	_ruleMetum.CreatorID = field.NewInt64(tableName, "creator_id")
	_ruleMetum.CreatorName = field.NewString(tableName, "creator_name")
	_ruleMetum.Name = field.NewString(tableName, "name")
	_ruleMetum.Description = field.NewString(tableName, "description")

	_ruleMetum.fillFieldMap()

	return _ruleMetum
}

type ruleMetum struct {
	ruleMetumDo

	ALL         field.Asterisk
	ID          field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	CreatorID   field.Int64
	CreatorName field.String
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (r ruleMetum) Table(newTableName string) *ruleMetum {
	r.ruleMetumDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r ruleMetum) As(alias string) *ruleMetum {
	r.ruleMetumDo.DO = *(r.ruleMetumDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *ruleMetum) updateTableName(table string) *ruleMetum {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.CreatorID = field.NewInt64(table, "creator_id")
	r.CreatorName = field.NewString(table, "creator_name")
	r.Name = field.NewString(table, "name")
	r.Description = field.NewString(table, "description")

	r.fillFieldMap()

	return r
}

func (r *ruleMetum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *ruleMetum) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["creator_id"] = r.CreatorID
	r.fieldMap["creator_name"] = r.CreatorName
	r.fieldMap["name"] = r.Name
	r.fieldMap["description"] = r.Description
}

func (r ruleMetum) clone(db *gorm.DB) ruleMetum {
	r.ruleMetumDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r ruleMetum) replaceDB(db *gorm.DB) ruleMetum {
	r.ruleMetumDo.ReplaceDB(db)
	return r
}

type ruleMetumDo struct{ gen.DO }

type IRuleMetumDo interface {
	gen.SubQuery
	Debug() IRuleMetumDo
	WithContext(ctx context.Context) IRuleMetumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRuleMetumDo
	WriteDB() IRuleMetumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRuleMetumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRuleMetumDo
	Not(conds ...gen.Condition) IRuleMetumDo
	Or(conds ...gen.Condition) IRuleMetumDo
	Select(conds ...field.Expr) IRuleMetumDo
	Where(conds ...gen.Condition) IRuleMetumDo
	Order(conds ...field.Expr) IRuleMetumDo
	Distinct(cols ...field.Expr) IRuleMetumDo
	Omit(cols ...field.Expr) IRuleMetumDo
	Join(table schema.Tabler, on ...field.Expr) IRuleMetumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRuleMetumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRuleMetumDo
	Group(cols ...field.Expr) IRuleMetumDo
	Having(conds ...gen.Condition) IRuleMetumDo
	Limit(limit int) IRuleMetumDo
	Offset(offset int) IRuleMetumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRuleMetumDo
	Unscoped() IRuleMetumDo
	Create(values ...*model.RuleMetum) error
	CreateInBatches(values []*model.RuleMetum, batchSize int) error
	Save(values ...*model.RuleMetum) error
	First() (*model.RuleMetum, error)
	Take() (*model.RuleMetum, error)
	Last() (*model.RuleMetum, error)
	Find() ([]*model.RuleMetum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RuleMetum, err error)
	FindInBatches(result *[]*model.RuleMetum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RuleMetum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRuleMetumDo
	Assign(attrs ...field.AssignExpr) IRuleMetumDo
	Joins(fields ...field.RelationField) IRuleMetumDo
	Preload(fields ...field.RelationField) IRuleMetumDo
	FirstOrInit() (*model.RuleMetum, error)
	FirstOrCreate() (*model.RuleMetum, error)
	FindByPage(offset int, limit int) (result []*model.RuleMetum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRuleMetumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r ruleMetumDo) Debug() IRuleMetumDo {
	return r.withDO(r.DO.Debug())
}

func (r ruleMetumDo) WithContext(ctx context.Context) IRuleMetumDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r ruleMetumDo) ReadDB() IRuleMetumDo {
	return r.Clauses(dbresolver.Read)
}

func (r ruleMetumDo) WriteDB() IRuleMetumDo {
	return r.Clauses(dbresolver.Write)
}

func (r ruleMetumDo) Session(config *gorm.Session) IRuleMetumDo {
	return r.withDO(r.DO.Session(config))
}

func (r ruleMetumDo) Clauses(conds ...clause.Expression) IRuleMetumDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r ruleMetumDo) Returning(value interface{}, columns ...string) IRuleMetumDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r ruleMetumDo) Not(conds ...gen.Condition) IRuleMetumDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r ruleMetumDo) Or(conds ...gen.Condition) IRuleMetumDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r ruleMetumDo) Select(conds ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r ruleMetumDo) Where(conds ...gen.Condition) IRuleMetumDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r ruleMetumDo) Order(conds ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r ruleMetumDo) Distinct(cols ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r ruleMetumDo) Omit(cols ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r ruleMetumDo) Join(table schema.Tabler, on ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r ruleMetumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r ruleMetumDo) RightJoin(table schema.Tabler, on ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r ruleMetumDo) Group(cols ...field.Expr) IRuleMetumDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r ruleMetumDo) Having(conds ...gen.Condition) IRuleMetumDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r ruleMetumDo) Limit(limit int) IRuleMetumDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r ruleMetumDo) Offset(offset int) IRuleMetumDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r ruleMetumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRuleMetumDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r ruleMetumDo) Unscoped() IRuleMetumDo {
	return r.withDO(r.DO.Unscoped())
}

func (r ruleMetumDo) Create(values ...*model.RuleMetum) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r ruleMetumDo) CreateInBatches(values []*model.RuleMetum, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r ruleMetumDo) Save(values ...*model.RuleMetum) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r ruleMetumDo) First() (*model.RuleMetum, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RuleMetum), nil
	}
}

func (r ruleMetumDo) Take() (*model.RuleMetum, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RuleMetum), nil
	}
}

func (r ruleMetumDo) Last() (*model.RuleMetum, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RuleMetum), nil
	}
}

func (r ruleMetumDo) Find() ([]*model.RuleMetum, error) {
	result, err := r.DO.Find()
	return result.([]*model.RuleMetum), err
}

func (r ruleMetumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RuleMetum, err error) {
	buf := make([]*model.RuleMetum, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r ruleMetumDo) FindInBatches(result *[]*model.RuleMetum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r ruleMetumDo) Attrs(attrs ...field.AssignExpr) IRuleMetumDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r ruleMetumDo) Assign(attrs ...field.AssignExpr) IRuleMetumDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r ruleMetumDo) Joins(fields ...field.RelationField) IRuleMetumDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r ruleMetumDo) Preload(fields ...field.RelationField) IRuleMetumDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r ruleMetumDo) FirstOrInit() (*model.RuleMetum, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RuleMetum), nil
	}
}

func (r ruleMetumDo) FirstOrCreate() (*model.RuleMetum, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RuleMetum), nil
	}
}

func (r ruleMetumDo) FindByPage(offset int, limit int) (result []*model.RuleMetum, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r ruleMetumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r ruleMetumDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r ruleMetumDo) Delete(models ...*model.RuleMetum) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *ruleMetumDo) withDO(do gen.Dao) *ruleMetumDo {
	r.DO = *do.(*gen.DO)
	return r
}
