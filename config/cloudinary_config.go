package config

type CloudinaryConfig struct {
	CloudName  string
	ApiKey     string
	ApiSecret  string
	FolderName string
}

var CloudConfig = CloudinaryConfig{
	CloudName:  getENV("CLOUD_NAME", ""),
	ApiKey:     getENV("API_KEY", ""),
	ApiSecret:  getENV("API_SECRET", ""),
	FolderName: getENV("FOLDER_NAME", ""),
}
