package initialize

import (
	"ecommerce-backend-golang/global"
	"ecommerce-backend-golang/internal/po"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: true, // Disable implicit transaction
	})
	// Check for error
	checkErrorPanic(err, "Failed to connect to MySQL database")
	// Logger connection status
	global.Logger.Info("Connected to MySQL database")
	// Set for global.Mdb
	global.Mdb = db
	// Set for pool
	SetPool()
	// Migrate tables
	migrateTables()
	// genTableDAO()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Failed to set pool %s", err)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenCons)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxTime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Printf("Failed to migrate tables %s", err)
	}
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",                                                 // output path
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	// g.GenerateAllTable() // generate all tables
	// g.GenerateModel("go_crm_user") // generate table `go_crm_user`
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
