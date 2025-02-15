package locations

import (
	"github.com/grow-assistant/swoop-menu-creator/swoop/pkg/log"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = log.NewLogger()
}
