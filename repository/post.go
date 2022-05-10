package repository

import (
	"encoding/json"
	"os"
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
	postDao  *PostDao     //指向Dao结构体的指针变量
	postOnce sync.Once    //单例模式，可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的
	rwMutex  sync.RWMutex //支持单写多读，而sync.Mutex：互斥锁
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

//向文件/数据库中插入新发布的帖子
func (*PostDao) InsertNewPost(post *Post) error {
	//0600 0rw-------
	open, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer open.Close()
	postStr, _ := json.Marshal(post) //序列号为json格式
	//写入文件时需要string
	//注意windows中回车换行符
	if _, err = open.WriteString(string(postStr) + "\n"); err != nil {
		return err
	}
	//更新Map时考虑并发安全性问题
	rwMutex.Lock()
	postList, ok := postIndexMap[post.ParentId]
	//若字典中不存在该主题ID，则新建
	if !ok {
		postIndexMap[post.ParentId] = []*Post{post}
	} else {
		postList = append(postList, post)
		postIndexMap[post.ParentId] = postList
	}
	rwMutex.Unlock()
	return nil
}
