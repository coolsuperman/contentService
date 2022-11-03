package datamanager

import (
	"contentService/internal/entity"
	"fmt"
)

func (mysql *MysqlHelper) InsertContent(data entity.Content) error {
	result := mysql.db.Table("content").Create(data)
	fmt.Println("result.RowsAffected:", result.RowsAffected, "result.Error:", result.Error)
	return result.Error
}

func (mysql *MysqlHelper) GetContentListByTag(limit, offset, tag, status int) (data []entity.BaseContent, err error) {
	result := mysql.db.Table("content").
		Where("tag & ? = 1 AND status = ", tag, status).
		Limit(limit).Offset(offset).
		Find(&data)
	fmt.Println("result.RowsAffected:", result.RowsAffected, "result.Error:", result.Error)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (mysql *MysqlHelper) GetContentDetail(contentID string) (data *entity.Content, err error) {
	result := mysql.db.Table("content").
		Where("content_id = ?", contentID).
		Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (mysql *MysqlHelper) UpdateContent(contentID string, data entity.Content) error {
	result := mysql.db.Table("content").
		Where("content_id = ?", contentID).
		Updates(data)
	fmt.Println("result.RowsAffected:", result.RowsAffected, "result.Error:", result.Error)
	return result.Error
}
