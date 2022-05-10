package service

import (
	"errors"
	"time"
	"unicode/utf16"

	idworker "github.com/gitstliu/go-id-worker"
	"github.com/yhmain/go-project-example/repository"
)

var idGen *idworker.IdWorker

//ID生成器的初始化
func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1, 1) //WORKERID位数 (用于对工作进程进行编码), 数据中心ID位数 (用于对数据中心进行编码)
}

//供controller调用，写入新帖子
func PublishPost(topicId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, content).Do()
}

//impl?
func NewPublishPostFlow(topicId int64, content string) *PublishNewPost {
	return &PublishNewPost{
		content: content,
		topicId: topicId,
	}
}

//定义提交新帖的结构体流
type PublishNewPost struct {
	content string
	topicId int64
	postId  int64
}

func (f *PublishNewPost) Do() (int64, error) {
	//;写法为在if语句之前添加一个执行语句  ;后面的才为if的判断条件
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

//检查内容长度是否符号要求
func (f *PublishNewPost) checkParam() error {
	//rune: int32的别名，用它来区分字符值和整数值
	//fmt.Println(len("Go语言编程"))  				// 输出：14
	//eg. fmt.Println(len([]rune("Go语言编程")))	// 输出：6
	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

//调用Dao层，将新帖存入数据库/文件
func (f *PublishNewPost) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: time.Now().Unix(), //获取时间戳
	}
	id, err := idGen.NextId() //生成新ID
	if err != nil {
		return err
	}
	post.Id = id
	//调用Dao层
	if err := repository.NewPostDaoInstance().InsertNewPost(post); err != nil {
		return err
	}
	//更新该帖子在内存中的ID
	f.postId = post.Id
	return nil
}
