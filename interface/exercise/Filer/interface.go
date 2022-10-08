package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Filer interface {
	GetFatherId() uint32
	GetID() uint32
	GetProjectID() uint32
}

//type Document interface {
//	CREATE(tx *gorm.DB) (uint32, error)
//}

func Create(db *gorm.DB, f Filer) (uint32, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(f).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	isFatherProject := false
	fatherId := f.GetFatherId()

	if fatherId == 0 {
		isFatherProject = true
		fatherId = f.GetProjectID()
	}

	if err := AddChildren(f, tx, isFatherProject, fatherId); err != nil {
		tx.Rollback()
		return 0, err
	}
	ID := f.GetID()
	return ID, tx.Commit().Error
}

func (doc *DocModel) GetFatherId() uint32 {
	return doc.FatherId
}

func (doc *DocModel) GetID() uint32 {
	return doc.ID
}

func (doc *DocModel) GetProjectID() uint32 {
	return doc.ProjectID
}

func (file *FileModel) GetFatherId() uint32 {
	return file.FatherId
}

func (file *FileModel) GetID() uint32 {
	return file.ID
}

func (file *FileModel) GetProjectID() uint32 {
	return file.ProjectID
}

func AddChildren(f Filer, tx *gorm.DB, isFatherProject bool, fatherId uint32) error {
	id := f.GetID()

	item := GetFolder(fatherId)

	newChildren, err := addChildren(item.Children, id)
	if err != nil {
		return err
	}

	item.Children = newChildren

	return tx.Save(item).Error
}

func main() {
	var Type int
	fmt.Println("请输入要创建的类型,1---doc,2---file")
	fmt.Scanf("%d", &Type)
	var doc *DocModel
	var file *FileModel
	var document Filer
	switch Type {
	case 1:
		document = doc
		break
	case 2:
		document = file
		break
	}
	var DB *gorm.DB
	Create(DB, document)

}
