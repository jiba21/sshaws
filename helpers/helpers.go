package helpers

// Configuration struct
type Configuration struct {
	App    string
	Env    string
	Region string
	Name   string
	User   string
}

//NewConfiguration returns one Configuration object
func NewConfiguration(region, env, app, name, user string) *Configuration {
	config := new(Configuration)
	config.App = app
	config.Env = env
	config.Name = name
	config.Region = region
	config.User = user
	return config
}

//ReturnConfiguration returns a set of strings with current configuration
func ReturnConfiguration(currentConfig Configuration) (string, string, string, string, string) {
	return currentConfig.Env, currentConfig.App, currentConfig.Name, currentConfig.Region, currentConfig.User
}

//Instance struct
type Instance struct {
	Name string
	IP   string
	ID   string
	Size string
}

//NewInstance creator
func NewInstance(name, ip, id, size string) *Instance {
	newInstance := new(Instance)
	newInstance.Name = name
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	return newInstance
}
