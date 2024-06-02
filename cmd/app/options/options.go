package options

import (
	"context"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	defaultConfigFilePath = "etc/config.yaml"
)

type Options struct {
	VIPER      *viper.Viper
	LOG        *zap.Logger
	DB         *gorm.DB
	GinEngine  *gin.Engine
	ConfigFile string
	Config     *config.Config
	Factory    database.ShareFactory
}

func NewOptions() *Options {
	return &Options{
		ConfigFile: defaultConfigFilePath,
	}
}

func (o *Options) viper(path ...string) *viper.Viper {
	var configFile string
	if len(path) == 0 {
		flag.StringVar(&configFile, "c", "", "chose server config file")
		flag.Parse()
		if "" == configFile {
			//获取配置文件路径的顺序： cmd > env > default
			if env := os.Getenv("SERVER_CONFIG_FILE"); env == "" {
				configFile = defaultConfigFilePath
				fmt.Println("use default server config file: ", configFile)
			} else {
				configFile = env
				fmt.Println("server config file get from SERVER_CONFIG_FILE: ", configFile)
			}
		} else {
			fmt.Println("use cmd input path: ", configFile)
		}
	} else {
		configFile = path[0]
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config file failed"))
	}
	v.WatchConfig()
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("config file changed:", event.Name)
		if err = v.Unmarshal(&o.ConfigFile); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&o.Config); err != nil {
		fmt.Println(err)
	}
	return v
}

func (o *Options) logger() error {
	logType := strings.ToLower(o.Config.Log.Format)
	if logType == "file" {
		if err := utils.EnsureDirectoryExists(o.Config.Log.DirectorPath); err != nil {
			return err
		}
	}
	if err := log.SetupLoggers(logType, o.Config.Log.DirectorPath, o.Config.Log.Level); err != nil {
		return err
	}
	return nil
}

func (o *Options) SetupServices() error {
	if err := o.logger(); err != nil {
		return err
	}

	if err := o.GormMysql(); err != nil {
		return err
	}
	return nil
}

func (o *Options) Initialize() error {
	o.GinEngine = gin.Default()
	o.viper()

	if err := o.SetupServices(); err != nil {
		return err
	}
	return nil
}

// 根据配置决定是够需要开启mysql的日志
func (o *Options) gormConfig(mod string) *gorm.Config {
	setLogger := func(level string) logger.Interface {
		switch level {
		case "Silent":
			return logger.Default.LogMode(logger.Silent)
		case "Error":
			return logger.Default.LogMode(logger.Error)
		case "Warn":
			return logger.Default.LogMode(logger.Warn)
		case "Info":
			return logger.Default.LogMode(logger.Info)
		default:
			return logger.Default.LogMode(logger.Silent)
		}
	}

	logMode := setLogger(o.Config.Mysql.LogZap)
	return &gorm.Config{
		Logger: logMode,
		//在数据库迁移期间不自动创建外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func (o *Options) GormMysql() (err error) {
	config := o.Config.Mysql
	//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Addr + ")/" + config.DBName + "?" + config.Config
	sqlConfig := mysql.Config{
		DSN:                       dsn,   //数据源名称
		DefaultStringSize:         256,   //string 字段默认长度
		DisableDatetimePrecision:  true,  //禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, //根据版本自动配置
	}

	o.DB, err = gorm.Open(mysql.New(sqlConfig), o.gormConfig(config.LogMode))
	if err != nil {
		log.Logger.LogError("mysql connection failed", zap.Any("err: ", err))
		return
	}
	sqlDB, err := o.DB.DB()
	if err != nil {
		log.Logger.LogError("connect db failed", zap.Any("err:", err))
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	if err = sqlDB.Ping(); err != nil {
		o.LOG.Error("ping db failed", zap.Any("err:", err))
		return
	}

	o.Factory = database.NewFactory(o.DB)
	//初始化数据库表
	if o.Config.System.AutoMigrateDb {
		if err = database.MySqlTables(o.DB); err != nil {
			return
		}
	}
	return
}

func (o *Options) BindConfigurationFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ConfigFile, "config", "", "please specify the configuration file path of server")
}

func (o *Options) RunHttpServer() error {
	server := &http.Server{
		ReadHeaderTimeout: 50 * time.Second,
		WriteTimeout:      50 * time.Second,
		Handler:           o.GinEngine,
		Addr:              fmt.Sprintf("%s:%d", o.Config.Http.Addr, o.Config.Http.Port),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			err = fmt.Errorf("listen: %w", err)
			return
		}
	}()
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		err = fmt.Errorf("server forced to shutdown: %v", err)
		return err
	}

	return nil
}
