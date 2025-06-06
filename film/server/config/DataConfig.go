package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"time"
)

/*
 定义一些数据库存放的key值, 以及程序运行时的相关参数配置
*/

// -------------------------System Config-----------------------------------
const (

	// ListenerPort web服务监听的端口
	ListenerPort = "3601"

	// MAXGoroutine max goroutine, 执行spider中对协程的数量限制
	MAXGoroutine = 10

	FilmPictureUploadDir = "./static/upload/gallery"
	FilmPictureUrlPath   = "/upload/pic/poster/"
	FilmPictureAccess    = "/api/upload/pic/poster/"
)

// -------------------------redis key-----------------------------------
const (
	// CategoryTreeKey 分类树 key
	CategoryTreeKey = "CategoryTree"
	FilmExpired     = time.Hour * 24 * 365 * 10
	// MovieListInfoKey movies分类列表 key
	MovieListInfoKey = "MovieList:Cid%d"

	// MovieDetailKey movie detail影视详情信息 可以
	MovieDetailKey = "MovieDetail:Cid%d:Id%d"
	// MovieBasicInfoKey 影片基本信息, 简略版本
	MovieBasicInfoKey = "MovieBasicInfo:Cid%d:Id%d"

	// MultipleSiteDetail 多站点影片信息存储key
	MultipleSiteDetail = "MultipleSource:%s"

	// SearchInfoTemp redis暂存检索数据信息
	SearchInfoTemp = "Search:SearchInfoTemp"

	// SearchTitle 影片分类标题key
	SearchTitle = "Search:Pid%d:Title"
	// SearchTag 影片剧情标签key
	SearchTag = "Search:Pid%d:%s"

	// VirtualPictureKey 待同步图片临时存储 key
	VirtualPictureKey = "VirtualPicture"
	// MaxScanCount redis Scan 操作每次扫描的数据量, 每次最多扫描300条数据
	MaxScanCount = 300
)

const (
	AuthUserClaims = "UserClaims"
)

// -------------------------manage 管理后台相关key----------------------------------
const (
	// FilmSourceListKey 采集 API 信息列表key
	FilmSourceListKey = "Config:Collect:FilmSource"
	// ManageConfigExpired 管理配置key 长期有效, 暂定10年
	ManageConfigExpired = time.Hour * 24 * 365 * 10
	// SiteConfigBasic 网站参数配置
	SiteConfigBasic = "SystemConfig:SiteConfig:Basic"
	// BannersKey 轮播组件key 你
	BannersKey = "SystemConfig:Banners"

	// FilmCrontabKey 定时任务列表信息
	FilmCrontabKey = "Cron:Task:Film"
	// DefaultUpdateSpec 每20分钟执行一次
	DefaultUpdateSpec = "0 */20 * * * ?"
	// EveryWeekSpec 每周日凌晨4点更新一次
	EveryWeekSpec = "0 0 4 * * 0"
	// DefaultUpdateTime 每次采集最近 3 小时内更新的影片
	DefaultUpdateTime = 3
)

// -------------------------Web API相关redis key-----------------------------------
const (
	// IndexCacheKey , 首页数据缓存
	IndexCacheKey = "IndexCache"
)

// -------------------------Database Connection Params-----------------------------------

const (
	// SearchTableName 存放检索信息的数据表名
	SearchTableName        = "search"
	UserTableName          = "users"
	UserIdInitialVal       = 10000
	FileTableName          = "files"
	FailureRecordTableName = "failure_records"

	//mysql服务配置信息 root:root 设置mysql账户的用户名和密码

	//MysqlDsn = "root:root@(localhost:3306)/FilmSite?charset=utf8mb4&parseTime=True&loc=Local"
	//MysqlDsn = "root:MuBai0916$@(113.44.5.201:3610)/FilmSite?charset=utf8mb4&parseTime=True&loc=Local"

	// MysqlDsn docker compose 环境下的链接信息 mysql:3306 为 docker compose 中 mysql服务对应的网络名称和端口
	//MysqlDsn = "root:root@(mysql:3306)/FilmSite?charset=utf8mb4&parseTime=True&loc=Local"

	/*
		redis 配置信息
		RedisAddr host:port
		RedisPassword redis访问密码
		RedisDBNo 使用第几号库
	*/
	//RedisAddr     = `113.44.5.201:3620`
	//RedisPassword = `MuBai0916$`
	//RedisDBNo     = 0

	// RedisAddr docker compose 环境下运行使用如下配置信息
	//RedisAddr     = `redis:6379`
	//RedisPassword = `root`
	//RedisDBNo     = 0
)

var (
	MysqlDBName    = os.Getenv("MYSQL_DBNAME")
	MysqlHost      = os.Getenv("MYSQL_HOST")
	MysqlPort      = os.Getenv("MYSQL_PORT")
	MysqlUser      = os.Getenv("MYSQL_USER")
	MysqlPassword  = os.Getenv("MYSQL_PASSWORD")
	RedisHost      = os.Getenv("REDIS_HOST")
	RedisPort      = os.Getenv("REDIS_PORT")
	RedisNo        = os.Getenv("REDIS_NO")
	RedisPassword  = os.Getenv("REDIS_PASSWORD")
	RedisAddr      = Default(RedisHost, "localhost") + `:` + Default(RedisPort, "6379")
	RedisDBNo, err = strconv.Atoi(RedisNo)

	MysqlDsn = Default(MysqlUser, "root") + ":" +
		Default(MysqlPassword, "root") + "@(" +
		Default(MysqlHost, "localhost") + ":" +
		Default(MysqlPort, "3306") + ")/" +
		Default(MysqlDBName, "FilmSite") + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func Default(in string, defaultValue string) string {
	if len(in) == 0 {
		return defaultValue
	} else {
		return in
	}
}
