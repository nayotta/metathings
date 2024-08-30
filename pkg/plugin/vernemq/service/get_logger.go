package metathings_plugin_vernemq_service

import "github.com/sirupsen/logrus"

func (s *VernemqPluginService) getLogger() logrus.FieldLogger {
	return s.logger
}
