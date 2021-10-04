/*
 * @Author: F1nley
 * @Date: 2021-10-01 21:04:19
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:43:42
 * @Description:
 */
package models

type Message struct {
	Id         int64  `xorm:pk autoincr`
	Uid        int64  `xorm: notnull`
	Content    string `xorm: notnull`
	Anonymous  bool   `xorm: notnull`
	CreateTime int64  `xorm: created notnull`
}
