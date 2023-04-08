package order

import (
	"github.com/geraldofigueiredo/goodtaste/services/web/front/pkg/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type OrderHandlerTestSuite struct {
	suite.Suite
	require *require.Assertions

	orderSvc *mocks.OrderService
	logger   *zap.SugaredLogger

	handler orderHandler
}

var (
	newOrderJSON = `{ "infos": ["1","5","3"], "owner_name":"testing" }`
	okJSON       = `Ok`
)

func TestOrderHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(OrderHandlerTestSuite))
}

func (h *OrderHandlerTestSuite) SetupTest() {
	h.require = require.New(h.T())

	h.logger = zap.NewNop().Sugar()
	h.orderSvc = mocks.NewOrderService(h.T())

	h.handler = orderHandler{
		orderSvc: h.orderSvc,
		logger:   h.logger,
	}
}

func (h *OrderHandlerTestSuite) TestNewOrder() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(newOrderJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h.orderSvc.
		On("NewOrder").
		Return(nil)

	if assert.NoError(h.T(), h.handler.NewOrder(c)) {
		assert.Equal(h.T(), http.StatusCreated, rec.Code)
		assert.Equal(h.T(), okJSON, rec.Body.String())
	}
}
