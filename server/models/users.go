/*
 * @Author: F1nley
 * @Date: 2021-10-01 11:17:55
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 11:31:41
 * @Description: 用户
 */
package models

type User struct {
	Id    int64  `xorm:pk autoincr`
	EMail string `xorm:unique notnull`
	Pwd   string `xorm: notnull`
}

type Message struct {
	Id         int64  `xorm:pk autoincr`
	Uid        int64  `xorm: notnull`
	Content    string `xorm: notnull`
	CreateTime int64  `xorm: created notnull`
}
