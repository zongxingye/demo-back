package db

import (
	"fmt"
	"time"
	"xingyeblog/model"
)

// user column
var userCols = []string{"id", "username", "password", "nickname", "create_at", "update_at"}

type UserOp struct {
}

// Register方法插入用户（先检查再插入）
func Register(userName string, pwd string) error {
	exist, err := CheckUser(userName)
	fmt.Println(err)

	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("用户已经存在")
	}
	// 数据库插入
	userTmp := model.User{UserName: userName, Pwd: pwd, CreateAt: int(time.Now().Unix())}
	err = InsertRegister(userTmp)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("插入用户失败", err)
		return err
	}
	return nil
}

// 检查user是否存在
func CheckUser(userName string) (bool, error) {
	var op = UserOp{}
	userTpm, err := op.GetUserByUserName(userName)
	fmt.Println("usertpm is :", userTpm)
	if err != nil {
		return false, err
	}
	if userTpm == nil {
		return false, nil
	}
	return true, nil
}

func InsertRegister(user model.User) error {
	var op = UserOp{}
	err := op.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Get query one record by id
func (UserOp) GetUserByUserName(userName string) (*model.User, error) {
	userTmp := new(model.User)
	userTmp.UserName = userName
	has, err := db.Table("users").Cols(userCols...).Get(userTmp)
	if nil != err {
		return nil, err
	}
	if !has {
		fmt.Println(has)
		return nil, nil
	}
	return userTmp, nil
}

// insert操作
func (UserOp) InsertUser(user model.User) error {
	fmt.Println("11111", user)
	_, err := db.Table("users").Cols(userCols...).Insert(&user)
	//_, err := db.Insert(&file)
	return err
}

// 登陆情况检查（用户名密码是否匹配）
func LoginCheck(userName, password string) (bool, model.User, error) {
	var loginReq model.User
	loginReq.UserName = userName
	loginReq.Pwd = password
	var op = UserOp{}
	userTpm, err := op.GetUserByUserName(loginReq.UserName)
	if err != nil {
		return false, model.User{}, err
	}
	if userTpm == nil {
		return false, model.User{}, nil
	}
	if !(loginReq.UserName == userTpm.UserName && loginReq.Pwd == userTpm.Pwd) {
		return false, *userTpm, fmt.Errorf("用户信息错误")
	}

	return true, *userTpm, nil

}
