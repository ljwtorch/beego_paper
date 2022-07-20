package models

type ArticleType struct {
	Id       int        `json:"id"`
	Typename string     `orm:"size(20)",json:"typename"`
	Articles []*Article `orm:"reverse(many)"` // 设置一对多的反向关系，一个类型可以包含多篇论文
}

func (m *ArticleType) TableName() string {
	return TableName("article_type")
}
