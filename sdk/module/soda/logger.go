package metathings_module_soda_sdk

import "github.com/sirupsen/logrus"

func (c *sodaClient) GetLogger() *logrus.Entry {
	return c.logger.WithFields(logrus.Fields{
		"#instance": "sodaClient",
	})
}
