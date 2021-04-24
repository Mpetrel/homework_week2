package main

import (
	"database/sql"
	"github.com/pkg/errors"
)
/**
 * 处理一类属于正常范围内的一些错误
 */

// IsEmptyError 检查根因错误是否是记录不存在的错误
func IsEmptyError(err error) bool {
	return errors.Cause(err) == sql.ErrNoRows
}