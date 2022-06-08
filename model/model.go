package model

type Model struct {
	TableName string
}

func NewModel(tb string) *Model {
	return &Model{
		TableName: tb,
	}
}

func (m *Model) Get(req interface{}, resp interface{}, isMaster bool) (err error) {
	return NewSqlExec(req, resp, m.TableName).Find(isMaster)
}

func (m *Model) Count(req interface{}, resp interface{}, isMaster bool) (err error) {
	return NewSqlExec(req, resp, m.TableName).Count(isMaster)
}

func (m *Model) Select(req, resp interface{}, isMaster bool) (err error) {
	se := NewSqlExec(req, resp, m.TableName)
	return se.Select(isMaster)
}

func (m *Model) Insert(req interface{}) (lastInsertId int64, err error) {
	return NewSqlExec(req, nil, m.TableName).Insert()
}

func (m *Model) Update(req interface{}) (affextedRow int64, err error) {
	return NewSqlExec(req, nil, m.TableName).Update()
}
