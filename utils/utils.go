package utils

import (
	"crypto/md5"
	"fmt"
)

// 获取汉字首字母
func ChineseCharacterInitials(hans string) string {

	args := pinyin.NewArgs()

	rows := pinyin.Pinyin(hans, args)

	strResult := ""
	for i := 0; i < len(rows); i++ {
		if len(rows[i]) != 0 {
			str := rows[i][0]
			pi := str[0:1]
			strResult += pi
		}
	}
	return strResult
}

func MD5(s string) string {
	data := []byte(s)
	hashed := md5.Sum(data)
	return fmt.Sprintf("%x", hashed)
}
