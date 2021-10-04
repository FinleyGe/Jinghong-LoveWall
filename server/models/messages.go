/*
 * @Author: F1nley
 * @Date: 2021-10-01 21:04:19
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 15:22:39
 * @Description:
 */
package models

import "time"

type Message struct {
	Id             int64  `xorm:pk autoincr`
	Uid            int64  `xorm: notnull`
	Content        string `xorm: notnull`
	Anonymous      bool   `xorm: notnull`
	Permit_comment bool   `xorm: notnull`
	CommentsID     []int64
	CreateTime     time.Time `xorm: created`
	UpdateTime     time.Time `xorm: updated`
}

type Comment struct {
	Id         int64     `xorm:pk autoincr`
	Uid        int64     `xorm: notnull`
	Content    string    `xorm: notnull`
	Anonymous  bool      `xorm: notnull`
	CreateTime time.Time `xorm: created`
}
