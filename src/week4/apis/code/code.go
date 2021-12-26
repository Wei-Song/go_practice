package code

import "go_practice/go_practice/src/week4/errors"

var (
	UserNotFound = errors.NotFound("shop.user.UserNotFound", "用户不存在")

	Unknown = errors.Internal("shop.user.Unknown", "未知错误")
)
