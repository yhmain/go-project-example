package repository

import (
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao  //指向Dao结构体的指针变量
	postOnce sync.Once //单例模式，可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的
)

//初始化单例
func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

//通过主题ID查询其对应的所有的帖子
//func (*PostDao)表示为 *PostDao的成员方法
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
