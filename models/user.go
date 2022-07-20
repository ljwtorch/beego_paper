package models

type User struct {
	Name     string     `orm:"pk",orm:"description(用户名)"`
	Pwd      string     `orm:"description(密码)"`
	Articles []*Article `orm:"reverse(many)"` //设置一对多关系
}

func (m *User) tableName() string {
	return TableName("user")
}
