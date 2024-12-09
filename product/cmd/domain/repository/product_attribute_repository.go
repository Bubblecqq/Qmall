package repository

import "QMall/product/cmd/domain/model"

type IProductAttributeRepository interface {
	FindProductAttributeGroup(id int64) (*model.ProductAttributeGroup, error)
	GetProductAttributeGroupList() []model.ProductAttributeGroup
	DeleteProductAttributeGroup(id int64) error
}

func (p *ProductRepository) FindProductAttributeGroup(id int64) (*model.ProductAttributeGroup, error) {
	productAttributeGroup := &model.ProductAttributeGroup{}
	err := p.mysqlClient.Model(&model.ProductAttributeGroup{}).Find(&productAttributeGroup, id).Error
	return productAttributeGroup, err
}

func (p *ProductRepository) GetProductAttributeGroupList() []model.ProductAttributeGroup {
	var productAttributeGroupList []model.ProductAttributeGroup
	p.mysqlClient.Model(&model.ProductAttributeGroup{}).Find(&productAttributeGroupList)
	return productAttributeGroupList
}

func (p *ProductRepository) DeleteProductAttributeGroup(id int64) error {
	productAttributeGroup := &model.ProductAttributeGroup{}
	err := p.mysqlClient.Model(&model.ProductAttributeGroup{}).Find(&productAttributeGroup, id).Error
	return err
}
