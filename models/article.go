package models

import "time"

type Article struct {
	//字段名一般为首字母大写
	Id    int    `orm:"pk;auto"`
	Aname string `orm:"size(20)"`
	// auto_now 每次model保存时都会对时间自动更新
	Atime    time.Time    `orm:"auto_now"`
	Acount   int          `orm:"default(0);null"`
	Acontent string       // 论文内容
	Aimg     string       // 论文存放路径
	Atype    *ArticleType `orm:"rel(fk)"` //设置一对多关系，一篇论文对应一个类型，论文类型为外键
	User     *User        `orm:"rel(fk)"` //设置一对多的关系，一篇论文对应一个创建者
}

func (m *Article) Tablename() string {
	return TableName("article")
}
