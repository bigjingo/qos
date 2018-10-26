package app

import (
	"github.com/QOSGroup/qbase/account"
	"github.com/QOSGroup/qbase/baseabci"
	"github.com/QOSGroup/qbase/context"
	qosacc "github.com/QOSGroup/qos/account"
	"github.com/QOSGroup/qos/mapper"
	"github.com/QOSGroup/qos/txs/approve"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"io"
)

const (
	appName = "QOS"
)

type QOSApp struct {
	*baseabci.BaseApp
}

func NewApp(logger log.Logger, db dbm.DB, traceStore io.Writer) *QOSApp {

	baseApp := baseabci.NewBaseApp(appName, logger, db, RegisterCodec)
	baseApp.SetCommitMultiStoreTracer(traceStore)

	app := &QOSApp{
		BaseApp: baseApp,
	}

	// 设置 InitChainer
	app.SetInitChainer(app.initChainer)

	// 账户mapper
	app.RegisterAccountProto(qosacc.ProtoQOSAccount)

	// 基础信息操作mapper
	app.RegisterMapper(mapper.NewMainMapper())

	// QCP mapper
	// qbase 默认已注入

	// 预授权mapper
	app.RegisterMapper(approve.NewApproveMapper())

	// Mount stores and load the latest state.
	err := app.LoadLatestVersion()
	if err != nil {
		cmn.Exit(err.Error())
	}
	return app
}

// 初始配置
func (app *QOSApp) initChainer(ctx context.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// 上下文中获取mapper
	mainMapper := ctx.Mapper(mapper.GetMainStoreKey()).(*mapper.MainMapper)
	accountMapper := ctx.Mapper(account.GetAccountKVStoreName()).(*account.AccountMapper)

	// 反序列化app_state
	stateJSON := req.AppStateBytes
	genesisState := &qosacc.GenesisState{}
	err := accountMapper.GetCodec().UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		panic(err)
	}

	// 保存CA
	mainMapper.SetRootCA(genesisState.CAPubKey)

	// 保存初始账户
	for _, acc := range genesisState.Accounts {
		accountMapper.SetAccount(acc)
	}

	return abci.ResponseInitChain{}
}