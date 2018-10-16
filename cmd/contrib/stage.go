package cmd_contrib

type StageOptioner interface {
	GetStageP() *string
	GetStage() string
	SetStage(string)
}

type StageOption struct {
	Stage string
}

func (self *StageOption) GetStageP() *string {
	return &self.Stage
}

func (self *StageOption) GetStage() string {
	return self.Stage
}

func (self *StageOption) SetStage(stage string) {
	self.Stage = stage
}
