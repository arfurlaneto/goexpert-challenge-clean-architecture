package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenExistingOrders_WhenList_ThenShouldListExistingOrders() {
	_, err := suite.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (1, 100, 10, 110), (2, 200, 20, 220)")
	suite.NoError(err)

	repo := NewOrderRepository(suite.Db)
	orders, err := repo.List()

	suite.NoError(err)
	suite.Len(orders, 2)

	suite.Equal(orders[0].ID, "1")
	suite.Equal(orders[0].Price, 100.0)
	suite.Equal(orders[0].Tax, 10.0)
	suite.Equal(orders[0].FinalPrice, 110.0)

	suite.Equal(orders[1].ID, "2")
	suite.Equal(orders[1].Price, 200.0)
	suite.Equal(orders[1].Tax, 20.0)
	suite.Equal(orders[1].FinalPrice, 220.0)
}
