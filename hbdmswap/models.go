package hbdmswap

import (
	sjson "encoding/json"
	"fmt"
	"strconv"
)

type Heartbeat struct {
	Heartbeat                 int         `json:"heartbeat"`
	EstimatedRecoveryTime     interface{} `json:"estimated_recovery_time"`
	SwapHeartbeat             int         `json:"swap_heartbeat"`
	SwapEstimatedRecoveryTime interface{} `json:"swap_estimated_recovery_time"`
}

type HeartbeatResult struct {
	Status string    `json:"status"`
	Data   Heartbeat `json:"data"`
	Ts     int64     `json:"ts"`
}

type AccountInfo struct {
	Symbol            string  `json:"symbol"`
	MarginBalance     float64 `json:"margin_balance"`
	MarginPosition    float64 `json:"margin_position"`
	MarginFrozen      float64 `json:"margin_frozen"`
	MarginAvailable   float64 `json:"margin_available"`
	ProfitReal        float64 `json:"profit_real"`
	ProfitUnreal      float64 `json:"profit_unreal"`
	RiskRate          float64 `json:"risk_rate"` //*interface{}
	WithdrawAvailable float64 `json:"withdraw_available"`
	LiquidationPrice  float64 `json:"liquidation_price"` //*interface{}
	LeverRate         float64 `json:"lever_rate"`
	AdjustFactor      float64 `json:"adjust_factor"`
	MarginStatic      float64 `json:"margin_static"`
	ContractCode      string  `json:"contract_code"`
}

type AccountInfoResult struct {
	Status  string        `json:"status"` // "ok" , "error"
	ErrCode int           `json:"err_code"`
	ErrMsg  string        `json:"err_msg"`
	Data    []AccountInfo `json:"data"`
	Ts      int64         `json:"ts"`
}

type Position struct {
	Symbol         string  `json:"symbol"`
	ContractCode   string  `json:"contract_code"`
	Volume         float64 `json:"volume"`
	Available      float64 `json:"available"`
	Frozen         float64 `json:"frozen"`
	CostOpen       float64 `json:"cost_open"`
	CostHold       float64 `json:"cost_hold"`
	ProfitUnreal   float64 `json:"profit_unreal"`
	ProfitRate     float64 `json:"profit_rate"`
	Profit         float64 `json:"profit"`
	PositionMargin float64 `json:"position_margin"`
	LeverRate      int     `json:"lever_rate"`
	Direction      string  `json:"direction"` // "buy":买 "sell":卖
	LastPrice      float64 `json:"last_price"`
}

type PositionInfoResult struct {
	Status  string     `json:"status"`
	ErrCode int        `json:"err_code"`
	ErrMsg  string     `json:"err_msg"`
	Data    []Position `json:"data"`
	Ts      int64      `json:"ts"`
}

type Tick struct {
	Asks    [][]float64 `json:"asks"`
	Bids    [][]float64 `json:"bids"`
	Ch      string      `json:"ch"`
	ID      int64       `json:"id"`
	MrID    int64       `json:"mrid"`
	Ts      int64       `json:"ts"`
	Version int64       `json:"version"`
}

type MarketDepthResult struct {
	Ch      string `json:"ch"`
	Status  string `json:"status"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Tick    Tick   `json:"tick"`
	Ts      int64  `json:"ts"`
}

type KLine struct {
	Amount float64 `json:"amount"`
	Close  float64 `json:"close"`
	Count  int     `json:"count"`
	High   float64 `json:"high"`
	ID     int     `json:"id"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Vol    int     `json:"vol"`
}

type KLineResult struct {
	Ch      string  `json:"ch"`
	Data    []KLine `json:"data"`
	Status  string  `json:"status"`
	ErrCode int     `json:"err_code"`
	ErrMsg  string  `json:"err_msg"`
	Ts      int64   `json:"ts"`
}

type OrderData struct {
	OrderID       int64  `json:"order_id"`
	OrderIDStr    string `json:"order_id_str"`
	ClientOrderID int64  `json:"client_order_id"`
}

type OrderResult struct {
	Status  string    `json:"status"`
	ErrCode int       `json:"err_code"`
	ErrMsg  string    `json:"err_msg"`
	Data    OrderData `json:"data"`
	Ts      int64     `json:"ts"`
}

type CancelError struct {
	OrderID string `json:"order_id"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type CancelData struct {
	Errors    []CancelError `json:"errors"`
	Successes string        `json:"successes"`
}

type CancelResult struct {
	Status  string     `json:"status"`
	ErrCode int        `json:"err_code"`
	ErrMsg  string     `json:"err_msg"`
	Data    CancelData `json:"data"`
	Ts      int64      `json:"ts"`
}

type Order struct {
	Symbol            string           `json:"symbol"`
	ContractCode      string           `json:"contract_code"`
	Volume            float64          `json:"volume"`
	Price             float64          `json:"price"`
	OrderPriceTypeRaw sjson.RawMessage `json:"order_price_type"` // 1限价单，3对手价，4闪电平仓，5计划委托，6post_only
	OrderType         int              `json:"order_type"`
	Direction         string           `json:"direction"`
	Offset            string           `json:"offset"`
	LeverRate         int              `json:"lever_rate"`
	OrderID           int64            `json:"order_id"`
	ClientOrderID     string           `json:"client_order_id"`
	CreatedAt         int64            `json:"created_at"`
	TradeVolume       float64          `json:"trade_volume"`
	TradeTurnover     float64          `json:"trade_turnover"`
	Fee               float64          `json:"fee"`
	TradeAvgPrice     float64          `json:"trade_avg_price"`
	MarginFrozen      float64          `json:"margin_frozen"`
	Profit            float64          `json:"profit"`
	Status            int              `json:"status"`
	OrderSource       string           `json:"order_source"`
	OrderIDStr        string           `json:"order_id_str"`
	FeeAsset          string           `json:"fee_asset"`
}

func (o *Order) OrderPriceType() string {
	// 1：限价单，3：对手价，4：闪电平仓，5：计划委托，6：post_only
	d, err := o.OrderPriceTypeRaw.MarshalJSON()
	if err != nil {
		return ""
	}
	s := string(d)
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		switch i {
		case 1:
			return "limit"
		case 3:
			return "opponent"
		case 4:
			return "lightning"
		case 5:
			return "trigger_order"
		case 6:
			return "post_only"
		default:
			return fmt.Sprint(i)
		}
	}
	return s
}

type OrderInfoResult struct {
	Status  string  `json:"status"`
	ErrCode int     `json:"err_code"`
	ErrMsg  string  `json:"err_msg"`
	Data    []Order `json:"data"`
	Ts      int64   `json:"ts"`
}

type OpenOrdersData struct {
	Orders      []Order `json:"orders"`
	TotalPage   int     `json:"total_page"`
	CurrentPage int     `json:"current_page"`
	TotalSize   int     `json:"total_size"`
}

type OpenOrdersResult struct {
	Status  string         `json:"status"`
	ErrCode int            `json:"err_code"`
	ErrMsg  string         `json:"err_msg"`
	Data    OpenOrdersData `json:"data"`
	Ts      int64          `json:"ts"`
}

type HisOrdersData struct {
	Orders      []Order `json:"orders"`
	TotalPage   int     `json:"total_page"`
	CurrentPage int     `json:"current_page"`
	TotalSize   int     `json:"total_size"`
}

type HisOrdersResult struct {
	Status  string        `json:"status"`
	ErrCode int           `json:"err_code"`
	ErrMsg  string        `json:"err_msg"`
	Data    HisOrdersData `json:"data"`
	Ts      int64         `json:"ts"`
}

type WSTickerTick struct {
	ID     int64   `json:"id"`
	MrID   int64   `json:"mrid"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Amount float64 `json:"amount"`
	Vol    float64 `json:"vol"`
	Count  int64   `json:"count"`
}

type WSTicker struct {
	Ch   string       `json:"ch"`
	Ts   int64        `json:"ts"`
	Tick WSTickerTick `json:"tick"`
}

type WSTick struct {
	MrID    int64       `json:"mrid"`
	ID      int         `json:"id"`
	Bids    [][]float64 `json:"bids"`
	Asks    [][]float64 `json:"asks"`
	Ts      int64       `json:"ts"`
	Version int         `json:"version"`
	Ch      string      `json:"ch"`
}

type WSDepth struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick Tick   `json:"tick"`
}

type WSTickHF struct {
	Asks    [][]float64 `json:"asks"`
	Bids    [][]float64 `json:"bids"`
	Ch      string      `json:"ch"`
	Event   string      `json:"event"`
	ID      int64       `json:"id"`
	Mrid    int64       `json:"mrid"`
	Ts      int64       `json:"ts"`
	Version int         `json:"version"`
}

type WSDepthHF struct {
	Ch   string   `json:"ch"`
	Tick WSTickHF `json:"tick"`
	Ts   int64    `json:"ts"`
}

type WSTradeItem struct {
	Amount    int     `json:"amount"`
	Ts        int64   `json:"ts"`
	ID        int64   `json:"id"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}

type WSTradeTick struct {
	ID   int64         `json:"id"`
	Ts   int64         `json:"ts"`
	Data []WSTradeItem `json:"data"`
}

type WSTrade struct {
	Ch   string      `json:"ch"`
	Ts   int64       `json:"ts"`
	Tick WSTradeTick `json:"tick"`
}

type WSMyTrade struct {
	TradeID       int64   `json:"trade_id"`
	ID            string  `json:"id"`
	TradeVolume   float64 `json:"trade_volume"`
	TradePrice    float64 `json:"trade_price"`
	TradeFee      float64 `json:"trade_fee"`
	FeeAsset      string  `json:"fee_asset"`
	TradeTurnover float64 `json:"trade_turnover"`
	CreatedAt     int64   `json:"created_at"`
	Role          string  `json:"role"`
}

type WSOrder struct {
	Op             string      `json:"op"`
	Topic          string      `json:"topic"`
	Ts             int64       `json:"ts"`
	Symbol         string      `json:"symbol"`
	ContractCode   string      `json:"contract_code"`
	Volume         float64     `json:"volume"`
	Price          float64     `json:"price"`
	OrderPriceType string      `json:"order_price_type"`
	Direction      string      `json:"direction"`
	Offset         string      `json:"offset"`
	Status         int         `json:"status"`
	LeverRate      float64     `json:"lever_rate"`
	OrderID        int64       `json:"order_id"`
	OrderIDStr     string      `json:"order_id_str"`
	ClientOrderID  int64       `json:"client_order_id"`
	OrderSource    string      `json:"order_source"`
	OrderType      int         `json:"order_type"`
	CreatedAt      int64       `json:"created_at"`
	TradeVolume    float64     `json:"trade_volume"`
	TradeTurnover  float64     `json:"trade_turnover"`
	Fee            float64     `json:"fee"`
	TradeAvgPrice  float64     `json:"trade_avg_price"`
	MarginFrozen   float64     `json:"margin_frozen"`
	Profit         float64     `json:"profit"`
	Trade          []WSMyTrade `json:"trade"`
}

type WSAccountData struct {
	Symbol            string  `json:"symbol"`
	ContractCode      string  `json:"contract_code"`
	MarginBalance     float64 `json:"margin_balance"`
	MarginStatic      float64 `json:"margin_static"`
	MarginPosition    float64 `json:"margin_position"`
	MarginFrozen      float64 `json:"margin_frozen"`
	MarginAvailable   float64 `json:"margin_available"`
	ProfitReal        float64 `json:"profit_real"`
	ProfitUnreal      float64 `json:"profit_unreal"`
	WithdrawAvailable float64 `json:"withdraw_available"`
	RiskRate          float64 `json:"risk_rate"`
	LiquidationPrice  float64 `json:"liquidation_price"`
	LeverRate         float64 `json:"lever_rate"`
	AdjustFactor      float64 `json:"adjust_factor"`
}

type WSAccounts struct {
	Op    string          `json:"op"`
	Topic string          `json:"topic"`
	Ts    int64           `json:"ts"`
	Event string          `json:"event"`
	Data  []WSAccountData `json:"data"`
}

type WSPositionData struct {
	Symbol         string  `json:"symbol"`
	ContractCode   string  `json:"contract_code"`
	Volume         float64 `json:"volume"`
	Available      float64 `json:"available"`
	Frozen         float64 `json:"frozen"`
	CostOpen       float64 `json:"cost_open"`
	CostHold       float64 `json:"cost_hold"`
	ProfitUnreal   float64 `json:"profit_unreal"`
	ProfitRate     float64 `json:"profit_rate"`
	Profit         float64 `json:"profit"`
	PositionMargin float64 `json:"position_margin"`
	LeverRate      float64 `json:"lever_rate"`
	Direction      string  `json:"direction"`
	LastPrice      float64 `json:"last_price"`
}

type WSPositions struct {
	Op    string           `json:"op"`
	Topic string           `json:"topic"`
	Ts    int64            `json:"ts"`
	Event string           `json:"event"`
	Data  []WSPositionData `json:"data"`
}

type WSLiquidationOrderData struct {
	Symbol       string  `json:"symbol"`
	ContractCode string  `json:"contract_code"`
	Direction    string  `json:"direction"`
	Offset       string  `json:"offset"`
	Volume       float64 `json:"volume"`
	Price        float64 `json:"price"`
	CreatedAt    int64   `json:"created_at"`
}

type WSLiquidationOrders struct {
	Op    string                   `json:"op"`
	Topic string                   `json:"topic"`
	Ts    int64                    `json:"ts"`
	Data  []WSLiquidationOrderData `json:"data"`
}
