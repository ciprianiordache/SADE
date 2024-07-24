package middleware

import "sade-backend/pkg/logger"

var log *logger.Logger

func init() {
	log = logger.New("../log/log_request.log")
}
