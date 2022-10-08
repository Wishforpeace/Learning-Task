package main

import "github.com/jinzhu/gorm"

func CreateDoc(db *gorm.DB, doc *DocModel) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(doc).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	isFatherProject := false
	fatherId := doc.FatherId
	if doc.FatherId == 0 {
		isFatherProject = true
		fatherId = doc.ProjectID
	}

	if err := AddDocChildren(tx, isFatherProject, fatherId, doc); err != nil {
		tx.Rollback()
		return 0, err
	}

	return doc.ID, tx.Commit().Error
}

func CreateFile(db *gorm.DB, file *FileModel) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(file).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	isFatherProject := false
	fatherId := file.FatherId
	if file.FatherId == 0 {
		isFatherProject = true
		fatherId = file.ProjectID
	}

	if err := AddFileChildren(tx, isFatherProject, fatherId, file); err != nil {
		tx.Rollback()
		return 0, err
	}

	return file.ID, tx.Commit().Error
}

// AddDocChildren ... 新增 doc 文件树
func AddDocChildren(tx *gorm.DB, isFatherProject bool, fatherId uint32, doc *DocModel) error {
	id := doc.ID

	item := GetFolder(fatherId)

	newChildren, err := addChildren(item.Children, id)
	if err != nil {
		return err
	}

	item.Children = newChildren

	return tx.Save(item).Error
}

// AddFileChildren ... 新增 file 文件树
func AddFileChildren(tx *gorm.DB, isFatherProject bool, fatherId uint32, file *FileModel) error {
	id := file.ID

	item := GetFolder(fatherId)

	newChildren, err := addChildren(item.Children, id)
	if err != nil {
		return err
	}

	item.Children = newChildren

	return tx.Save(item).Error
}

// FileModel ... 文件物理模型
type FileModel struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	URL       string `json:"url" gorm:"column:url;" binding:"required"`
	Re        bool   `json:"re" gorm:"column:re;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	FatherId  uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

// DocModel ... 文档物理模型
type DocModel struct {
	ID        uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Content   string `json:"content" gorm:"column:content;" binding:"required"`
	Re        bool   `json:"re" gorm:"column:re;" binding:"required"`
	ProjectID uint32 `json:"projectId" gorm:"column:project_id;" binding:"required"`
	FatherId  uint32 `json:"father_id" gorm:"column:father_id;" binding:"required"`
}

func addChildren(children string, id uint32) (string, error) {
	return "", nil
}

type Item struct {
	Children string
}

func GetFolder(fatherId uint32) Item {
	return Item{}
}
