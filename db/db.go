package db

import (
	"errors"
	"eth_approve_parse/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Sqlite *gorm.DB
)

func InitDB() error {
	var err error
	Sqlite, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return err
	}

	// 迁移模型
	if err := Sqlite.AutoMigrate(&model.Approval{}); err != nil {
		return err
	}

	return nil
}

func UpdateApproval(approval *model.Approval) error {

	var existingApproval *model.Approval
	// 查询数据库，检查是否存在符合条件的记录
	result := Sqlite.Where("contract = ? AND owner = ? AND spender = ?", approval.Contract, approval.Owner, approval.Spender).First(&existingApproval)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	// 如果不存在符合条件的记录，则插入新的记录
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := Sqlite.Create(approval).Error; err != nil {
			return err
		}
	} else {
		// 一个transaction 内可能会包含多个 approve 的 调用， 根据 LogIndex 大小决定顺序， 值越大数据越新
		// 且approval的 BlockNumber 大于等于存在的数据
		// 且approval的 LogIndex 大于存在的数据
		// 说明是更新的数据，则更新该记录的其他数据
		if existingApproval.BlockNumber <= approval.BlockNumber && existingApproval.LogIndex < approval.LogIndex {
			existingApproval.Value = approval.Value
			existingApproval.TxHash = approval.TxHash
			if err := Sqlite.Save(&existingApproval).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
