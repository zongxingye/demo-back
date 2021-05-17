package db

import (
	"errors"
	"fmt"
	"time"
	"xingyeblog/model"
)

type ArticleInfo struct {
	Id       int    `json:"id" xorm:"id"`
	UserId   int    `json:"user_id" xorm:"user_id"`
	Content  string `json:"content" xorm:"contents"`
	Title    string `json:"title" xorm:"title"`
	Spot     int    `json:"spot" xorm:"spot"`
	CreateAt int    `json:"create_at" xorm:"create_at"`
	UpdateAt int    `json:"update_at" xorm:"update_at"`
}

// user column
var articleCols = []string{"id", "user_id", "contents", "title", "spot", "create_at", "update_at"}
var c_articleCols = []string{"user_id", "contents", "title", "spot", "create_at", "update_at"}

type ArticleOp struct {
}

func (ArticleOp) findAllArticlesBykey(key string, val interface{}) ([]map[string]string, error) {
	sql := "select * from articles where (" + key + "= ?)"
	results, err := db.SQL(sql, val).QueryString()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(results)
	return results, err
}

// 获取具体key的文章(通用)
func (ArticleOp) findOneByKey(key string, val interface{}) ([]map[string]string, error) {
	sql := "select * from articles where (" + key + "= ?)"
	result, err := db.SQL(sql, val).QueryString()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(result)
	return result, err
}

//// 获取具体id的文章
//func (ArticleOp) findOneArticleById(id int) (*model.Article, error) {
//	articleTemp := new(model.Article)
//	articleTemp.Id = id
//	has, err := db.Table("articles").Cols(articleCols...).Get(articleTemp)
//	if nil != err {
//		fmt.Println(err)
//		return nil, err
//	}
//	if !has {
//		fmt.Println(has)
//		return nil, nil
//	}
//	return articleTemp, nil
//}

// 更新文章
//func (ArticleOp) UpdateArticle(id int, article *model.Article) error {
//	_, err := db.Table("articles").ID(id).Cols(c_articleCols...).Update(article)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	return nil
//}

// 更新文章
func (ArticleOp) UpdateArticle(id int, article *model.Article) error {
	sql := "update articles set title = ?,spot= ?, contents=?,update_at = ? where id = ?"
	res, err := db.Exec(sql, article.Title, article.Spot, article.Contents, article.UpdateAt, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return nil
}

// 插入文章by userid
func (ArticleOp) InsertArticle(article *model.Article) error {
	_, err := db.Table("articles").Cols(c_articleCols...).Insert(article)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 删除文章
func (ArticleOp) DeleteArticle(id int) error {
	sql := "delete from articles where id = ?"
	res, err := db.Exec(sql, id)
	fmt.Println("111111", res, err)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("该文章已被删除")
	}
	return nil
}

// 通过uid获取所有博客文章
func GetAllArticlesByUid(uid int) ([]map[string]string, error) {
	a := ArticleOp{}
	res, err := a.findAllArticlesBykey("user_id", uid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

// 获取指定id的文章
func GetOneArticleById(id int) ([]map[string]string, error) {
	a := ArticleOp{}
	res, err := a.findOneByKey("id", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

// 创建文章（根据用户id）
func CreateArticleWithUid(uid int, article ArticleInfo) error {
	articleTmp := new(model.Article)
	articleTmp.UserId = uid
	articleTmp.Contents = article.Content
	articleTmp.Title = article.Title
	articleTmp.Spot = article.Spot
	articleTmp.CreateAt = int(time.Now().Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println(articleTmp.Contents)
	a := ArticleOp{}
	err := a.InsertArticle(articleTmp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// update
func UpdateArticle(id int, article ArticleInfo) error {
	articleTmp := new(model.Article)
	articleTmp.Id = id
	articleTmp.UserId = article.UserId
	articleTmp.Contents = article.Content
	articleTmp.Title = article.Title
	articleTmp.Spot = article.Spot
	articleTmp.UpdateAt = article.UpdateAt
	a := ArticleOp{}
	err := a.UpdateArticle(id, articleTmp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// delete
func DeleteArticle(id int) error {
	a := ArticleOp{}
	err := a.DeleteArticle(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
