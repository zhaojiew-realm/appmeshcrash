package model

type RemoteHost struct {
	Host string
	Port string
	Path string
}

func (r *RemoteHost) GetUrl() string {
	return "http://" + r.Host + ":" + r.Port + r.Path
}
