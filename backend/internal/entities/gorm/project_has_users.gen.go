// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
)

func newProjectHasUser(db *gorm.DB, opts ...gen.DOOption) projectHasUser {
	_projectHasUser := projectHasUser{}

	_projectHasUser.projectHasUserDo.UseDB(db, opts...)
	_projectHasUser.projectHasUserDo.UseModel(&model.ProjectHasUser{})

	tableName := _projectHasUser.projectHasUserDo.TableName()
	_projectHasUser.ALL = field.NewAsterisk(tableName)
	_projectHasUser.ProjectID = field.NewString(tableName, "project_id")
	_projectHasUser.UserID = field.NewString(tableName, "user_id")
	_projectHasUser.Role = field.NewString(tableName, "role")

	_projectHasUser.fillFieldMap()

	return _projectHasUser
}

type projectHasUser struct {
	projectHasUserDo projectHasUserDo

	ALL       field.Asterisk
	ProjectID field.String
	UserID    field.String
	Role      field.String

	fieldMap map[string]field.Expr
}

func (p projectHasUser) Table(newTableName string) *projectHasUser {
	p.projectHasUserDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectHasUser) As(alias string) *projectHasUser {
	p.projectHasUserDo.DO = *(p.projectHasUserDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectHasUser) updateTableName(table string) *projectHasUser {
	p.ALL = field.NewAsterisk(table)
	p.ProjectID = field.NewString(table, "project_id")
	p.UserID = field.NewString(table, "user_id")
	p.Role = field.NewString(table, "role")

	p.fillFieldMap()

	return p
}

func (p *projectHasUser) WithContext(ctx context.Context) *projectHasUserDo {
	return p.projectHasUserDo.WithContext(ctx)
}

func (p projectHasUser) TableName() string { return p.projectHasUserDo.TableName() }

func (p projectHasUser) Alias() string { return p.projectHasUserDo.Alias() }

func (p projectHasUser) Columns(cols ...field.Expr) gen.Columns {
	return p.projectHasUserDo.Columns(cols...)
}

func (p *projectHasUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectHasUser) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["project_id"] = p.ProjectID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["role"] = p.Role
}

func (p projectHasUser) clone(db *gorm.DB) projectHasUser {
	p.projectHasUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectHasUser) replaceDB(db *gorm.DB) projectHasUser {
	p.projectHasUserDo.ReplaceDB(db)
	return p
}

type projectHasUserDo struct{ gen.DO }

func (p projectHasUserDo) Debug() *projectHasUserDo {
	return p.withDO(p.DO.Debug())
}

func (p projectHasUserDo) WithContext(ctx context.Context) *projectHasUserDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectHasUserDo) ReadDB() *projectHasUserDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectHasUserDo) WriteDB() *projectHasUserDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectHasUserDo) Session(config *gorm.Session) *projectHasUserDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectHasUserDo) Clauses(conds ...clause.Expression) *projectHasUserDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectHasUserDo) Returning(value interface{}, columns ...string) *projectHasUserDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectHasUserDo) Not(conds ...gen.Condition) *projectHasUserDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectHasUserDo) Or(conds ...gen.Condition) *projectHasUserDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectHasUserDo) Select(conds ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectHasUserDo) Where(conds ...gen.Condition) *projectHasUserDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectHasUserDo) Order(conds ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectHasUserDo) Distinct(cols ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectHasUserDo) Omit(cols ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectHasUserDo) Join(table schema.Tabler, on ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectHasUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectHasUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectHasUserDo) Group(cols ...field.Expr) *projectHasUserDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectHasUserDo) Having(conds ...gen.Condition) *projectHasUserDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectHasUserDo) Limit(limit int) *projectHasUserDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectHasUserDo) Offset(offset int) *projectHasUserDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectHasUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *projectHasUserDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectHasUserDo) Unscoped() *projectHasUserDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectHasUserDo) Create(values ...*model.ProjectHasUser) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectHasUserDo) CreateInBatches(values []*model.ProjectHasUser, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectHasUserDo) Save(values ...*model.ProjectHasUser) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectHasUserDo) First() (*model.ProjectHasUser, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectHasUser), nil
	}
}

func (p projectHasUserDo) Take() (*model.ProjectHasUser, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectHasUser), nil
	}
}

func (p projectHasUserDo) Last() (*model.ProjectHasUser, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectHasUser), nil
	}
}

func (p projectHasUserDo) Find() ([]*model.ProjectHasUser, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectHasUser), err
}

func (p projectHasUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectHasUser, err error) {
	buf := make([]*model.ProjectHasUser, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectHasUserDo) FindInBatches(result *[]*model.ProjectHasUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectHasUserDo) Attrs(attrs ...field.AssignExpr) *projectHasUserDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectHasUserDo) Assign(attrs ...field.AssignExpr) *projectHasUserDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectHasUserDo) Joins(fields ...field.RelationField) *projectHasUserDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectHasUserDo) Preload(fields ...field.RelationField) *projectHasUserDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectHasUserDo) FirstOrInit() (*model.ProjectHasUser, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectHasUser), nil
	}
}

func (p projectHasUserDo) FirstOrCreate() (*model.ProjectHasUser, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectHasUser), nil
	}
}

func (p projectHasUserDo) FindByPage(offset int, limit int) (result []*model.ProjectHasUser, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p projectHasUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectHasUserDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectHasUserDo) Delete(models ...*model.ProjectHasUser) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectHasUserDo) withDO(do gen.Dao) *projectHasUserDo {
	p.DO = *do.(*gen.DO)
	return p
}