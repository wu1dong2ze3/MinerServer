package cgminer

type SendCmdContainer struct {
	Command   string `json:"command"`
	Parameter string `json:"parameter"`
}

func Create(command, parameter string) *SendCmdContainer {
	return &SendCmdContainer{command, parameter}
}
