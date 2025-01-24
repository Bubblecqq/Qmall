package repository

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type IActivityRepository interface {
	IActivityInfoRepo
	// GetActivityById 查询活动
	GetActivityById(id int64) (*model.Activity, error)
	DeleteActivityById(id, userId int64) error

	GetActivityProductById(id int64) (*model.ActivityProduct, error)
	GetActivityProductByProductsId(productId int64) (*model.ActivityProduct, error)

	GetActivityProductSkuById(id int64) (*model.ActivityProductSku, error)

	// AddActivity 添加活动
	AddActivity(activityName string, isOnline int32, startTime, endTime time.Time) (*model.Activity, error)

	AddActivityProduct(activityProduct *model.ActivityProduct) (*model.ActivityProduct, error)

	AddActivityProductSku(activityProductId, productId, productSkuId, userId int64, price, originalPrice float64, number, stock int64) (*model.ActivityProductSku, error)

	AddActivityTime(activityId int64, TimeName string, isEnable int32) (*model.ActivityTime, error)

	GetActivityInfo(productId int64) ([]model.ActivityInfo, error)

	GetActivityInfoByProductsNum(productId int64, productsNum string) (*model.ActivityInfo, error)

	GetActivityTimeById(in *pb.GetActivityTimeByIdReq) (*model.ActivityTime, error)
}

type ActivityRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (a *ActivityRepository) GetActivityTimeById(in *pb.GetActivityTimeByIdReq) (*model.ActivityTime, error) {
	activityTime := &model.ActivityTime{}

	tx := a.mysqlClient.Model(&model.ActivityTime{}).Find(&activityTime, in.Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if activityTime.ID <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	fmt.Printf("当前提取到的活动表结果：%v\n", activityTime)
	return activityTime, nil
}

func (a *ActivityRepository) GetActivityInfoByProductsNum(productId int64, productsNum string) (*model.ActivityInfo, error) {
	fmt.Printf("正在获取活动详细信息>>>>>>当前请求的商品Id：%v", productId)

	activityInfo := &model.ActivityInfo{}

	tx := a.mysqlClient.
		Table("activity ac").
		Select(`ac.name as activity_name,ap.product_name as product_name,
       ap.product_starting_price as product_starting_price,
       ap.category_id as category_id,ap.product_id as product_id,
       at.id as activity_time_id,at.start_time ,at.end_time ,
       ac.is_online,aps.product_sku_id as product_sku_id`).
		Joins("left join mall_marketing.activity_time at on ac.id = at.activity_id").
		Joins("left join mall_marketing.activity_product ap on at.id = ap.activity_time_id").
		Joins("left join mall_marketing.activity_product_sku aps on ap.id = aps.activity_product_id").
		Where("ap.product_id =? and ap.products_num=?", productId, productsNum).
		Find(&activityInfo)
	fmt.Println("activityInfo>>>>>", activityInfo)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return activityInfo, nil
}

func (a *ActivityRepository) GetActivityInfo(productId int64) ([]model.ActivityInfo, error) {
	fmt.Printf("正在获取活动详细信息>>>>>>当前请求的商品Id：%v", productId)

	//activityInfo := &model.ActivityInfo{}

	var activityInfo []model.ActivityInfo

	tx := a.mysqlClient.
		Table("activity ac").
		Select(`ac.name as activity_name,ap.product_name as product_name,
       ap.product_starting_price as product_starting_price,
       ap.category_id as category_id,ap.product_id as product_id,
       at.id as activity_time_id,at.start_time ,at.end_time ,
       ac.is_online`).
		Joins("left join mall_marketing.activity_time at on ac.id = at.activity_id").
		Joins("left join mall_marketing.activity_product ap on at.id = ap.activity_time_id").
		Where("ap.product_id =?", productId).
		Find(&activityInfo)
	fmt.Println("activityInfo>>>>>", activityInfo)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return activityInfo, nil
}

func (a *ActivityRepository) GetActivityProductByProductsId(productId int64) (*model.ActivityProduct, error) {
	activityProduct := &model.ActivityProduct{}
	tx := a.mysqlClient.Model(&model.ActivityProduct{}).Where("product_id=?", productId).Find(&activityProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if activityProduct.ID <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return activityProduct, nil
}

func (a *ActivityRepository) GetActivityProductById(id int64) (*model.ActivityProduct, error) {
	activityProduct := &model.ActivityProduct{}
	tx := a.mysqlClient.Model(&model.ActivityProduct{}).Find(&activityProduct, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return activityProduct, nil
}

func (a *ActivityRepository) GetActivityProductSkuById(id int64) (*model.ActivityProductSku, error) {
	activityProductSku := &model.ActivityProductSku{}
	tx := a.mysqlClient.Model(&model.ActivityProduct{}).Find(&activityProductSku, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return activityProductSku, nil
}

func (a *ActivityRepository) AddActivity(activityName string, isOnline int32, startTime, endTime time.Time) (*model.Activity, error) {
	activity := &model.Activity{
		CreateTime:        time.Now(),
		ActivityStartTime: startTime,
		ActivityEndTime:   endTime,
		Name:              activityName,
		IsDeleted:         0,
		IsOnline:          isOnline,
	}

	tx := a.mysqlClient.Model(&model.Activity{}).Create(&activity)
	if tx.Error != nil {
		fmt.Printf("当前活动名：%v不在活动期间！原因见：%v\n", activityName, tx.Error)
		return nil, tx.Error
	}
	return activity, nil
}

func (a *ActivityRepository) AddActivityProduct(activityProduct *model.ActivityProduct) (*model.ActivityProduct, error) {

	tx := a.mysqlClient.Model(&model.ActivityProduct{}).Create(&activityProduct)
	if tx.Error != nil {
		fmt.Printf("当前商品Id：%v及活动时间Id：%v不在活动期间！原因见：%v\n", activityProduct.ProductID, activityProduct.ActivityTimeID, tx.Error)
		return nil, tx.Error
	}
	return activityProduct, nil
}

func (a *ActivityRepository) AddActivityProductSku(activityProductId, productId, productSkuId, userId int64, price, originalPrice float64, number, stock int64) (*model.ActivityProductSku, error) {

	activityProductSku := &model.ActivityProductSku{
		ActivityProductID: activityProductId,
		ProductID:         productId,
		ProductSkuID:      productSkuId,
		CreateUser:        userId,
		OriginalPrice:     originalPrice,
		Stock:             stock,
		Price:             price,
		Number:            number,
		CreateTime:        time.Now(),
	}

	tx := a.mysqlClient.Model(&model.ActivityProductSku{}).Create(&activityProductSku)
	if tx.Error != nil {
		fmt.Printf("当前活动商品Id：%v，库存Id：%v创建失败！原因见：%v\n", activityProductId, productSkuId, tx.Error)
		return nil, tx.Error
	}
	return activityProductSku, nil
}

func (a *ActivityRepository) AddActivityTime(activityId int64, TimeName string, isEnable int32) (*model.ActivityTime, error) {
	activityTime := &model.ActivityTime{}
	activity := &model.Activity{}
	tx := a.mysqlClient.Model(&model.Activity{}).Find(&activity, activityId)
	if tx.Error != nil {
		fmt.Printf("当前活动Id：%v未开放！原因见：%v\n", activityId, tx.Error)
		return nil, tx.Error
	}
	activityTime.ActivityID = activity.ID
	activityTime.TimeName = TimeName
	activityTime.CreateTime = time.Now()
	activityTime.StartTime = time.Date(activity.ActivityStartTime.Year(), activity.ActivityStartTime.Month(), activity.ActivityStartTime.Day(), activity.ActivityStartTime.Hour(), activity.ActivityStartTime.Minute(), activity.ActivityEndTime.Second(), 0, time.UTC)
	activityTime.EndTime = time.Date(activity.ActivityEndTime.Year(), activity.ActivityEndTime.Month(), activity.ActivityEndTime.Day(), activity.ActivityEndTime.Hour(), activity.ActivityEndTime.Minute(), activity.ActivityEndTime.Second(), 0, time.UTC)

	activityTime.IsEnable = isEnable
	//activityTime.UpdateTime = &time.Time{}
	activityTime.IsDeleted = 0
	create := a.mysqlClient.Model(&model.ActivityTime{}).Create(&activityTime)
	if create.Error != nil {
		fmt.Printf("当前活动：%v时间未开放！原因见：%v\n", TimeName, create.Error)
		return nil, tx.Error
	}
	return activityTime, nil
}

func (a *ActivityRepository) GetActivityById(id int64) (*model.Activity, error) {
	activity := &model.Activity{}
	fmt.Printf("正在获取活动表信息！\n")
	tx := a.mysqlClient.Model(&model.Activity{}).Find(&activity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Printf("活动信息获取成功！，活动内容如下：%v\n", activity)
	return activity, nil
}

func (a *ActivityRepository) DeleteActivityById(id, userId int64) error {
	activityById, err := a.GetActivityById(id)

	if err != nil {
		fmt.Printf("当前id：%v的活动不存在！\n", id)
		return err
	}
	activityById.IsDeleted = 1
	updateTime := time.Now()
	activityById.UpdateTime = &updateTime
	activityById.UpdateUser = userId
	tx := a.mysqlClient.Model(&model.Activity{}).Updates(activityById)
	if tx.Error != nil {
		fmt.Printf("当前id：%v的活动更新失败！原因见：%v\n", id, tx.Error)
		return tx.Error
	}
	return nil
}

func NewActivityRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IActivityRepository {
	return &ActivityRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
