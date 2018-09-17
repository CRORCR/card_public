package modes

import (
	"fmt"
	"time"
	"errors"
	"../db"
	"../../lib"
)

/*
 * 描述：用户管理员表
 *
 *  type_id     : 用户创建类型
 *  phone       ：只可以用手机号
 *  role        ：角色：
 *  status      ：状态：0 正常 1 禁用 2 违法操作
 *
 ********************************************************************************/
type Users struct {
	Id             uint32 `xorm:"not null 'id'"`
	TypeId         uint32 `xorm:"not null 'type_id'"`
	OpenId         string `xorm:"not null 'open_id'"`
	SharId         string `xorm:"not null 'share_id'"`
	IconUrl        string `xorm:"not null 'iconurl'"`
	Phone          string `xorm:"not null 'phone'"`
	Name           string `xorm:"not null 'name'"`
	NumberId       string `xorm:"not null 'number_id'"`
	JwtToken       string `xorm:"not null 'jwt_token'"`
	Token          string `xorm:"not null 'token'"`
	LoginPass      string `xorm:"not null 'loginpass'"`
	PayPass        string `xorm:"not null 'paypass'"`
	CreateAt       int64  `xorm:"not null 'create_at'"`
	UpdateAt       int64  `xorm:"not null 'update_at'"`
	Email          string `xorm:"not null 'email'"`
	Sex            uint8  `xorm:"not null 'sex'"`
	Role           uint8  `xorm:"not null 'role'"`
	Attesta        uint8  `xorm:"not null 'attesta'"`
	Status         uint8  `xorm:"not null 'status'"`
	UnionidAndroid string `xorm:"not null 'unionid_android'"`
	UnionidIos     string `xorm:"not null 'unionid_ios'"`
}

/*
 * 描述：根据用户分享ID获取用户信息
 *
 *******************************************************************************/
func (this *Users) GetShareUser( usershareid *string, user *Users ) error {
	fmt.Println("用户分享ID", usershareid)
	fage, err := db.GetDBHand(0).Table("users").Where( "share_id = ?", usershareid).Get( user )
	if !fage || nil != err {
		return errors.New( fmt.Sprintf("用户分享ID: %s 不存在，或数据库操作失败", usershareid ) )
	}
	fmt.Println("PUBLIC", user )
	return nil
}

/*
 * 描述：根据手机号获取用户信息
 *
 *******************************************************************************/
func (this *Users) GetPhoneUser( phone *string, user *Users ) error {
	fmt.Println("GetPhoneUser Phone:", phone)
	fage, err := db.GetDBHand(0).Table("users").Where( "phone = ?", phone ).Get( user )
	if !fage || nil != err {
		return errors.New( fmt.Sprintf("手机用户: %s 不存在，或数据库操作失败", phone ) )
	}
	return nil
}


/*
 * 描述：判断手机账号存不存在
 *
 *******************************************************************************/
func (this *Users) IsPhoneUser( phone *string, user *Users )error {
	fage, err := db.GetDBHand(0).Table("users").
			       Where("phone = ?", phone ).
			       Get( user )
	if !fage || nil != err {
		return errors.New("用户不存在")
	}
	return nil
}

/*
 * 描述： UnionId 绑定手机号
 *
 *	前置条件:  UnionId.Phone 不可以为空
 *
 *******************************************************************************/
 type UnionId struct{
	Union	string
	Type	string
	Phone	string
}
func ( this *Users)UnionBing( unid *UnionId, user *Users ) error{
	user.Phone = unid.Phone
	if "unionid_android" == unid.Type {
		user.UnionidAndroid = unid.Union
	}else if "unionid_ios"  == unid.Type {
		user.UnionidIos = unid.Union
	}
	if nil == user.IsPhoneUser( &unid.Phone, user ){ // 如果存在手机号
		_, err := db.GetDBHand(0).Table("users").
				Where("phone = ? ", unid.Phone ).
				Cols( unid.Type ).
				Update(user)
		return err
	}
	return user.PhoneSave( user, nil )
}


/*
 * 描述： 设置头像
 *
 *	前置条件:  UnionId.Phone 不可以为空
 *
 *******************************************************************************/
func ( this *Users)SetIcon( user *Users, nfage *bool ) error{
	count, err := db.GetDBHand(0).Table("users").
				Where("share_id = ? ", user.SharId ).
				Cols( "iconurl" ).
				Update(user)
	if 1 == count && nil == err{
		return nil
	}
	return err
}


/*
 * 描述： 功能创建一个新用户
 *
 *	前置条件:  1: this.Phone 不可以为空。
 *		   2: users 表中不存在此手机的唯一性。
 *
 *******************************************************************************/
func (this *Users) PhoneSave( user *Users, strNull *string )error {
	user.CreateAt = time.Now().Unix()
	user.SharId   = lib.StrMd5Str(fmt.Sprintf("%s%d", user.Phone, user.CreateAt))
	if user.IconUrl == "" {
		user.IconUrl  = lib.GetUserLib().HeadIcon
	}
	_, err := db.GetDBHand(0).Table("users").Insert( user )
	return err
}

/*
 * 描述：查询用户是否已经存在
 *
 *******************************************************************************/
func (this *Users) GetUserByUnionidAndroid( union *string, user *Users )error{
	bil, err := db.GetDBHand(0).Table("users").
				    Where("unionid_android = ? ", union ).
				    Get( user )
	if !bil {
		return errors.New("用户不存在")
	}
	return err
}

/*
 * 描述：查询用户是否已经存在
 *
 *******************************************************************************/
func (this *Users) GetUserByUnionidIos( union *string, user *Users )error{
	bil, err := db.GetDBHand(0).Table("users").
					Where("unionid_ios = ? ", union).
					Get( user )
	if !bil {
		return errors.New("用户不存在")
	}
	return err
}




