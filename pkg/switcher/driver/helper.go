package metathings_switcher_driver

import (
	pb "github.com/bigdatagz/metathings/pkg/proto/switcher"
)

func (s SwitcherState) ToString() string {
	return _switcher_st_psr.ToString(pb.SwitcherState(s))
}
