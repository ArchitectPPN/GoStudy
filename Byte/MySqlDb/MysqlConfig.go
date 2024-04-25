package MySqlDb

type MySqlConfig struct {
	env          string
	userName     string
	password     string
	port         string
	hostName     string
	databaseName string
}

func NewMysqlConfig() (mysqlConfig *MySqlConfig) {
	return &MySqlConfig{}
}

func (mysqlConfig *MySqlConfig) SetEnv(env string) {
	mysqlConfig.env = env
}

func (mysqlConfig *MySqlConfig) GetEnv() string {
	return mysqlConfig.env
}

func (mysqlConfig *MySqlConfig) SetUserName(userName string) {
	mysqlConfig.userName = userName
}

func (mysqlConfig *MySqlConfig) SetPassword(password string) {
	mysqlConfig.password = password
}

func (mysqlConfig *MySqlConfig) SetPort(port string) {
	mysqlConfig.port = port
}

func (mysqlConfig *MySqlConfig) SetHostName(hostName string) {
	mysqlConfig.hostName = hostName
}

func (mysqlConfig *MySqlConfig) SetDatabaseName(databaseName string) {
	mysqlConfig.databaseName = databaseName
}

func (mysqlConfig *MySqlConfig) GetMysqlDbConfig() string {
	//"crm_dev:YKdYXQfhjjzDr281@tcp(rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com:3306)/crm_dev"
	dsn := mysqlConfig.userName + ":" + mysqlConfig.password + "@" + "tcp" + "(" + mysqlConfig.hostName + ":" + mysqlConfig.port + ")/" + mysqlConfig.databaseName
	return dsn
}

func (mysqlConfig *MySqlConfig) SetDevConfig() {
	mysqlConfig.SetEnv("dev")
	mysqlConfig.SetDatabaseName("crm_dev")
	mysqlConfig.SetHostName("rm-uf67h6texeyf9f325.mysql.rds.aliyuncs.com")
	mysqlConfig.SetPort("3306")
	mysqlConfig.SetUserName("crm_dev")
	mysqlConfig.SetPassword("YKdYXQfhjjzDr281")
}

func (mysqlConfig *MySqlConfig) SetSitConfig() {
	mysqlConfig.SetEnv("sit")
	mysqlConfig.SetDatabaseName("crm_sit")
	mysqlConfig.SetHostName("rm-uf6v4cfq4a000sr3w.mysql.rds.aliyuncs.com")
	mysqlConfig.SetPort("3306")
	mysqlConfig.SetUserName("crm_sit")
	mysqlConfig.SetPassword("UOGRxk4SVvt0Zncl")
}

func (mysqlConfig *MySqlConfig) SetProdConfig() {
	mysqlConfig.SetEnv("prod")
	mysqlConfig.SetDatabaseName("crm")
	mysqlConfig.SetHostName("192.168.35.118")
	mysqlConfig.SetPort("3306")
	mysqlConfig.SetUserName("crm_pro_read")
	mysqlConfig.SetPassword("zXkRr9Jo7JpVCCkT")
}
