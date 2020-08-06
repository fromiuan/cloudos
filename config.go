package cloudos

type CosConfig struct {
	SecretID   string // API密钥ID
	SecretKey  string // API密钥私钥
	Bucket     string // 存储桶名称
	Region     string // 存储桶所属地域
	APIAddress string // API地址(访问域名) 在存储桶列表
	Domain     string // 自定义域名
}

type MinioConfig struct {
	AccessKey string // API密钥ID
	SecretKey string // API密钥私钥
	Bucket    string // 存储桶所属地域
	Endpoint  string // API地址(访问域名) 在存储桶列表
	Domain    string // 自定义域名
	SSL       bool   // 加密链接
}

type OssConfig struct {
	AccessKeyId     string // AccessKey ID
	AccessKeySecret string // Access Key Secret
	Bucket          string // 存储桶名称
	Endpoint        string // API地址(访问域名) 在存储桶列表
}

type ObsConfig struct {
	Bucket    string // 存储桶名称
	AccessKey string // API密钥ID
	SecretKey string // API密钥私钥
	Endpoint  string // API地址(访问域名) 在存储桶列表
}
