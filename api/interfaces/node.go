package interfaces

import (
	"errors"
	database "gomboc/api/database"
	models "gomboc/api/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/clause"
)

func GetNodeById(id int) *models.NodeModel {
	var node *models.NodeModel
	db := database.New().Database
	result := map[string]interface{}{}

	db.Model(models.NodeModel{ID: id}).First(&result, id)

	return node
}

func GetNodeByPublicId(public_id string) (models.NodeModel, error) {
	node := models.NodeModel{}
	db := database.New().Database
	result := db.Model(&models.NodeModel{}).Where("public_id", public_id).First(&node)

	return node, result.Error
}

func GetAllNodes(page int, limit int, isActive bool, isPending bool) ([]models.NodeModel, error) {
	nodes := []models.NodeModel{}

	var err error
	db := database.New().Database

	offset := (page - 1) * limit
	query := db.Model(&models.NodeModel{}).Select("*").Offset(offset).Limit(limit)
	query = query.Where("is_active = ?", isActive)
	query = query.Where("is_pending = ?", isPending)

	result := query.Find(&nodes)

	if result.Error != nil {
		err = result.Error
		log.Err(result.Error).Msgf("Fail to execute 'CreateNode' operation: %s", err.Error())
	}

	return nodes, err
}

func CreateNode(node models.NodeModel) (models.NodeModel, error) {
	db := database.New().Database.Clauses(clause.Returning{})

	var err error
	result := db.Create(&node)

	if result.Error != nil {
		err = result.Error
		log.Err(result.Error).Msgf("Fail to execute 'CreateNode' operation: %s", err.Error())
	}

	return node, err
}

func CreateOrUpdateNode(node *models.NodeModel) error {
	db := database.New().Database.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "ID"}, {Name: "PublicId"}},
	})

	if err := db.Create(&node).Error; err != nil {
		log.Err(err).Msgf("Fail to execute 'CreateOrUpdateNode' operation: %s", err.Error())
		return err
	}

	return nil
}

// func GetNodeByPublicId(public_id string) (models.NodeModel, error) {
// 	var node models.NodeModel = models.NodeModel{}

// 	db := database.New().Database
// 	result := db.Model(&models.NodeModel{}).Where("public_id", public_id).First(&node)

// 	return node, result.Error
// }

func UpdateNode(iNode models.NodeModel) (models.NodeModel, error) {
	oNode := models.NodeModel{}

	db := database.New().Database
	result := db.Model(&oNode).Select("*").Omit("ID").Where("ID = ?", iNode.ID).Updates(iNode)

	return oNode, result.Error
}

func DeleteNode() (node models.NodeModel, e error) {
	return node, errors.New("function 'DeleteUser' not implemented yet")
}
