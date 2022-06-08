package model

import (
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/fbbyqsyea/gin-framework/global"
)

type SqlParse struct {
	InsertColumns string                 // insert
	InsertValues  []interface{}          // insert
	UpdateColumns map[string]interface{} // update
	SelectColumns string                 // select
	Where         []interface{}          // where=key,condition1,condition2
	Page, Limit   uint64                 // page,limit
	OrderBy       string                 // order by
}

func (s *SqlParse) Parse(i interface{}) {
	keys := reflect.TypeOf(i)
	values := reflect.ValueOf(i)
	s.Recursion(keys, values)
}

func (s *SqlParse) Recursion(key reflect.Type, val reflect.Value) {
	if key.Kind() == reflect.Ptr {
		key = key.Elem()
		val = val.Elem()
	}
	for i := 0; i < key.NumField(); i++ {
		k := key.Field(i).Type
		v := val.Field(i)
		if k.Kind() == reflect.Struct {
			s.Recursion(k, v)
		} else {
			s.Condition(key.Field(i), v)
		}
	}
}

func (s *SqlParse) Condition(key reflect.StructField, value reflect.Value) {
	if (!value.IsZero() && value.CanInterface()) || key.Tag.Get("order") != "" {
		v := value.Interface()
		// insert
		tag := key.Tag.Get("insert")
		if tag != "" {
			if s.InsertColumns != "" {
				s.InsertColumns = s.InsertColumns + "," + tag
			} else {
				s.InsertColumns = tag
			}
			s.InsertValues = append(s.InsertValues, v)
		}

		// update
		tag = key.Tag.Get("update")
		if tag != "" {
			s.UpdateColumns[tag] = v
		}
		// select
		tag = key.Tag.Get("select")
		if tag != "" {
			if s.SelectColumns != "" {
				s.SelectColumns = s.SelectColumns + "," + tag
			} else {
				s.SelectColumns = tag
			}
		}
		// where
		tag = key.Tag.Get("where")
		if tag != "" {
			column, assign := "", ""
			where := strings.Split(tag, ",")
			if len(where) > 1 {
				column, assign = where[0], strings.ToLower(where[1])
			} else {
				column, assign = where[0], "eq"
			}
			switch assign {
			case "eq":
				s.Where = append(s.Where, sq.Eq{column: v})
			case "not_eq":
				s.Where = append(s.Where, sq.NotEq{column: v})
			case "gt":
				s.Where = append(s.Where, sq.Gt{column: v})
			case "egt":
				s.Where = append(s.Where, sq.GtOrEq{column: v})
			case "lt":
				s.Where = append(s.Where, sq.Lt{column: v})
			case "elt":
				s.Where = append(s.Where, sq.LtOrEq{column: v})
			case "like":
				s.Where = append(s.Where, sq.Like{column: v})
			case "not_like":
				s.Where = append(s.Where, sq.NotILike{column: v})
			case "ilike":
				s.Where = append(s.Where, sq.ILike{column: v})
			case "not_ilke":
				s.Where = append(s.Where, sq.NotILike{column: v})
			}
		}
		// page
		tag = key.Tag.Get("page")
		if page, ok := v.(uint64); tag != "" && ok {
			s.Page = page
		}
		// limit
		tag = key.Tag.Get("limit")
		if limit, ok := v.(uint64); tag != "" && ok {
			s.Limit = limit
		}
		// order by
		tag = key.Tag.Get("order")
		if tag != "" {
			if s.OrderBy != "" {
				s.OrderBy = s.OrderBy + "," + tag
			} else {
				s.OrderBy = tag
			}
		}
	}
}

type SqlExec struct {
	TableName string
	Parse     *SqlParse
	Result    interface{}
}

func NewSqlExec(req, resp interface{}, tb string) *SqlExec {
	s := &SqlExec{
		TableName: tb,
		Parse: &SqlParse{
			SelectColumns: "*",
			UpdateColumns: make(map[string]interface{})},
		Result: resp,
	}
	s.Parse.Parse(req)
	return s
}

func (s *SqlExec) Find(isMaster bool) error {
	sb := sq.Select(s.Parse.SelectColumns).From(s.TableName)
	for _, where := range s.Parse.Where {
		sb = sb.Where(where)
	}
	sql, data, err := sb.ToSql()
	if err != nil {
		return err
	}
	if isMaster {
		return global.DB.Master.Get(s.Result, sql, data...)
	}
	return global.DB.Replica.Get(s.Result, sql, data...)
}

func (s *SqlExec) Select(isMaster bool) error {
	sb := sq.Select(s.Parse.SelectColumns).From(s.TableName)
	for _, where := range s.Parse.Where {
		sb = sb.Where(where)
	}
	if s.Parse.Limit != 0 {
		sb = sb.Limit(s.Parse.Limit)
	}
	offset := (s.Parse.Page - 1) * s.Parse.Limit
	if offset != 0 {
		sb = sb.Offset(offset)
	}
	if s.Parse.OrderBy != "" {
		sb = sb.OrderBy(s.Parse.OrderBy)
	}
	sql, data, err := sb.ToSql()
	if err != nil {
		return err
	}
	if isMaster {
		return global.DB.Master.Select(s.Result, sql, data...)
	}
	return global.DB.Replica.Select(s.Result, sql, data...)
}

func (s *SqlExec) Count(isMaster bool) error {
	sb := sq.Select("count(*) as count").From(s.TableName)
	for _, where := range s.Parse.Where {
		sb = sb.Where(where)
	}
	sql, data, err := sb.ToSql()
	if err != nil {
		return err
	}
	if isMaster {
		return global.DB.Master.Get(s.Result, sql, data...)
	}
	return global.DB.Replica.Get(s.Result, sql, data...)
}

func (s *SqlExec) Insert() (lastInsertId int64, err error) {
	sql, data, err := sq.Insert(s.TableName).Columns(s.Parse.InsertColumns).Values(s.Parse.InsertValues...).ToSql()
	if err != nil {
		return
	}
	result, err := global.DB.Master.Exec(sql, data...)
	if err != nil {
		return
	}
	lastInsertId, err = result.LastInsertId()
	return
}

func (s *SqlExec) Update() (affectedRow int64, err error) {
	ub := sq.Update(s.TableName)
	for k, v := range s.Parse.UpdateColumns {
		ub = ub.Set(k, v)
	}
	for _, where := range s.Parse.Where {
		ub = ub.Where(where)
	}
	sql, data, err := ub.ToSql()
	if err != nil {
		return
	}
	result, err := global.DB.Master.Exec(sql, data...)
	if err != nil {
		return
	}
	affectedRow, err = result.RowsAffected()
	return
}
