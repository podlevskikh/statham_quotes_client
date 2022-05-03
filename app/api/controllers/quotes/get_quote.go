package quotes

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/podlevskikh/statham_quotes_client/app/api/response_error"
	"github.com/podlevskikh/statham_quotes_client/app/api/response_success"
	"github.com/podlevskikh/statham_quotes_server/models/hashcash"
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"strings"
)

type GetQuote struct {
	logger *zerolog.Logger
}

func NewGetQuote(logger *zerolog.Logger) *GetQuote {
	return &GetQuote{logger: logger}
}

func (s *GetQuote) HTTPHandler(c *gin.Context) {
	conn, err := net.Dial("tcp", specs.TCPHost+":"+specs.TCPPort)
	if err != nil {
		response_error.ReturnError(c, response_error.TCPConnectionFailed)
		return
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Err(err).Msg("close connection")
		}
	}()

	reader := bufio.NewReader(conn)
	challenge, err := reader.ReadString('\n')
	if err != nil {
		response_error.ReturnError(c, response_error.ReadChallengeError)
		return
	}

	hash, err := hashcash.Parse(strings.Trim(challenge, "\n"))
	if err != nil {
		response_error.ReturnError(c, response_error.ParseChallengeError)
		return
	}

	hash.Solute()
	_, err = conn.Write([]byte(hash.ToString() + "\n"))
	if err != nil {
		response_error.ReturnError(c, response_error.WriteSolutionError)
		return
	}

	quote, err := reader.ReadString('\n')
	if err != nil {
		response_error.ReturnError(c, response_error.ReadQuoteError)
		return
	}

	c.JSON(http.StatusOK, response_success.FromTextResponse(strings.Trim(quote, "\n")))
}
