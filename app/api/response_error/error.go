package response_error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	TCPConnectionFailed = Error{Code: "tcp_connection_failed", Message: "tcp connection failed"}
	ReadChallengeError  = Error{Code: "read_challenge_error", Message: "can't read challenge"}
	ParseChallengeError = Error{Code: "parse_challenge_error", Message: "can't parse challenge"}
	WriteSolutionError  = Error{Code: "write_solution_error", Message: "can't write solution"}
	ReadQuoteError      = Error{Code: "read_quote_error", Message: "can't read quote"}
)

func ReturnError(c *gin.Context, res Error) {
	c.JSON(http.StatusInternalServerError, map[string]Error{"error": res})
}
