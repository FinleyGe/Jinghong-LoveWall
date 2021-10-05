/*
 * @Author: F1nley
 * @Date: 2021-10-03 10:22:11
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-05 10:25:25
 * @Description:tokens
 */
package models

type Token struct {
	Uid   int64  `xorm: "pk unique notnull"`
	Token string `xorm: "unique notnull"`
}
