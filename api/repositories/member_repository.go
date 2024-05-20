package repositories

import (
	"errors"
	"flowerpot/models"

	"gorm.io/gorm"
)

type MemberRepository interface {
	List() ([]models.Member, error)
	ActiveList() ([]models.Member, error)
	Get(id uint) (models.Member, error)
	Insert(models.Member) (models.Member, error)
	Update(id uint, member models.Member) error
	Delete(id uint) error
}

type memberRepository struct {
	db *gorm.DB
}

func InitMemberRepositoryDB(db *gorm.DB) MemberRepository {
	return &memberRepository{
		db,
	}
}

func (mr *memberRepository) List() ([]models.Member, error) {
	var members []models.Member

	result := mr.db.Order("members.created_at ASC").Find(&members)
	return members, result.Error
}

func (mr *memberRepository) ActiveList() ([]models.Member, error) {
	var members []models.Member
	activeStatus := "active"

	result := mr.db.Where("status = ?", activeStatus).Order("members.orcreated_atder ASC").Find(&members)
	return members, result.Error
}

func (mr *memberRepository) Get(id uint) (models.Member, error) {
	var member models.Member
	result := mr.db.First(&member, id)
	return member, result.Error
}

func (mr *memberRepository) Insert(member models.Member) (models.Member, error) {
	result := mr.db.Model(models.Member{}).Create(&member)
	if result.Error != nil {
		return member, result.Error
	} else if result.RowsAffected == 0 {
		return member, errors.New("no rows were Inserted")
	}
	return member, nil
}

func (mr *memberRepository) Update(id uint, member models.Member) error {
	result := mr.db.Model(models.Member{}).Where("id = ?", id).Updates(&member)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no rows were Updated")
	}
	return nil
}

func (mr *memberRepository) Delete(id uint) error {
	result := mr.db.Delete(&models.Member{}, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no rows were Deleted")
	}
	return nil
}
