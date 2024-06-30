package db

import (
	"fmt"

	"github.com/majezanu/grpc-course-microservices/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	OrderId    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (adapter Adapter) Get(id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := adapter.db.First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		OrderID:    paymentEntity.OrderId,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}
	return payment, res.Error
}

func (adapter Adapter) Save(payment *domain.Payment) error {
	paymentModel := Payment{
		CustomerID: payment.CustomerID,
		TotalPrice: payment.TotalPrice,
		OrderId:    payment.OrderID,
	}
	res := adapter.db.Create(&paymentModel)

	if res.Error == nil {
		payment.ID = int64(paymentModel.ID)
		payment.CreatedAt = paymentModel.CreatedAt.UnixNano()
	}
	return res.Error
}
