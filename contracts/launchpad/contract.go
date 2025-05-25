// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package launchpad

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InitParams is an auto generated low-level Go binding around an user-defined struct.
type InitParams struct {
	Admin           common.Address
	Token           common.Address
	PubStartTime    *big.Int
	PubEndTime      *big.Int
	MinInvestment   *big.Int
	MaxInvestment   *big.Int
	TotalPublicSale *big.Int
	CapRatio        *big.Int
	PubPrice        *big.Int
	RaiseRate       *big.Int
	LiquidityRate   *big.Int
	PlatformRate    *big.Int
	LiquidityWallet common.Address
	RaiseWallet     common.Address
}

// Investment is an auto generated low-level Go binding around an user-defined struct.
type Investment struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Claimed          bool
}

// LaunchpadDetails is an auto generated low-level Go binding around an user-defined struct.
type LaunchpadDetails struct {
	Token                 common.Address
	StartTime             *big.Int
	EndTime               *big.Int
	MinInvestment         *big.Int
	MaxInvestment         *big.Int
	TotalPublicSubscribed *big.Int
	TokenPrice            *big.Int
	LiquidityRate         *big.Int
	RaiseRate             *big.Int
	PlatformRate          *big.Int
	TotalPublicSale       *big.Int
	CapRatio              *big.Int
	SoftCap               *big.Int
	HardCap               *big.Int
	TotalInvestors        *big.Int
	LiquidityWallet       common.Address
	RaiseWallet           common.Address
}

// WhitelistDetails is an auto generated low-level Go binding around an user-defined struct.
type WhitelistDetails struct {
	WhitelistStartTime      *big.Int
	WhitelistEndTime        *big.Int
	WhitelistPrice          *big.Int
	TotalWhitelistSale      *big.Int
	TotalWhitelistSubscribe *big.Int
}

// LaunchpadMetaData contains all meta data concerning the Launchpad contract.
var LaunchpadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capRatio\",\"type\":\"uint256\"}],\"name\":\"CapRatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"claimable\",\"type\":\"bool\"}],\"name\":\"ClaimableSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teamWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"platformWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformAmount\",\"type\":\"uint256\"}],\"name\":\"FundsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"}],\"name\":\"Invested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minInvestment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxInvestment\",\"type\":\"uint256\"}],\"name\":\"InvestmentLimitsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityWallet\",\"type\":\"address\"}],\"name\":\"LiquidityWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicPrice\",\"type\":\"uint256\"}],\"name\":\"PublicPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicEndTime\",\"type\":\"uint256\"}],\"name\":\"PublicSaleTimesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"raiseWallet\",\"type\":\"address\"}],\"name\":\"RaiseWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformRate\",\"type\":\"uint256\"}],\"name\":\"RateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEth\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemainingTokensTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"investorAddresses\",\"type\":\"address[]\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTotalPublicSale\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalPublicSale\",\"type\":\"uint256\"}],\"name\":\"TotalPublicSaleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"}],\"name\":\"WhitelistInvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"whitelistStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"whitelistEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"whitelistPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWhitelistSale\",\"type\":\"uint256\"}],\"name\":\"WhitelistParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxInvestment\",\"type\":\"uint256\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSubscribed\",\"type\":\"uint256\"}],\"name\":\"estimateAmountBePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBePaid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"estimateAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"estimateWhitelistAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHardCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getInvestment\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structInvestment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLaunchpadDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPublicSubscribed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPublicSale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"softCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hardCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalInvestors\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"liquidityWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raiseWallet\",\"type\":\"address\"}],\"internalType\":\"structLaunchpadDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSoftCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInvestors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWhitelistSold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"whitelistStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"whitelistEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"whitelistPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWhitelistSale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWhitelistSubscribe\",\"type\":\"uint256\"}],\"internalType\":\"structWhitelistDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pubStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPublicSale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"liquidityWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raiseWallet\",\"type\":\"address\"}],\"internalType\":\"structInitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"investable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"investments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isClaimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInWhitelistPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxInvestment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minInvestment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raiseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raiseWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_whitelists\",\"type\":\"address[]\"}],\"name\":\"removeWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_capRatio\",\"type\":\"uint256\"}],\"name\":\"setCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setClaimable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxInvestment\",\"type\":\"uint256\"}],\"name\":\"setInvestmentThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityWallet\",\"type\":\"address\"}],\"name\":\"setLiquidityWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pubPrice\",\"type\":\"uint256\"}],\"name\":\"setPubPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pubStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pubEndTime\",\"type\":\"uint256\"}],\"name\":\"setPublicSaleTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_raiseWallet\",\"type\":\"address\"}],\"name\":\"setRaiseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformRate\",\"type\":\"uint256\"}],\"name\":\"setRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalPublicSale\",\"type\":\"uint256\"}],\"name\":\"setTotalPublicSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_whitelists\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxInvestments\",\"type\":\"uint256[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_whitelistStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_whitelistEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_whitelistPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalWhitelistSale\",\"type\":\"uint256\"}],\"name\":\"setWhitelistParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInvestors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPublicSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPublicSubscribed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWhitelistSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWhitelistSubscribed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWhitelisted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistInvest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelists\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"investorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvestment\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LaunchpadABI is the input ABI used to generate the binding from.
// Deprecated: Use LaunchpadMetaData.ABI instead.
var LaunchpadABI = LaunchpadMetaData.ABI

// Launchpad is an auto generated Go binding around an Ethereum contract.
type Launchpad struct {
	LaunchpadCaller     // Read-only binding to the contract
	LaunchpadTransactor // Write-only binding to the contract
	LaunchpadFilterer   // Log filterer for contract events
}

// LaunchpadCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaunchpadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchpadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaunchpadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchpadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaunchpadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchpadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaunchpadSession struct {
	Contract     *Launchpad        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaunchpadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaunchpadCallerSession struct {
	Contract *LaunchpadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LaunchpadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaunchpadTransactorSession struct {
	Contract     *LaunchpadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LaunchpadRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaunchpadRaw struct {
	Contract *Launchpad // Generic contract binding to access the raw methods on
}

// LaunchpadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaunchpadCallerRaw struct {
	Contract *LaunchpadCaller // Generic read-only contract binding to access the raw methods on
}

// LaunchpadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaunchpadTransactorRaw struct {
	Contract *LaunchpadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaunchpad creates a new instance of Launchpad, bound to a specific deployed contract.
func NewLaunchpad(address common.Address, backend bind.ContractBackend) (*Launchpad, error) {
	contract, err := bindLaunchpad(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Launchpad{LaunchpadCaller: LaunchpadCaller{contract: contract}, LaunchpadTransactor: LaunchpadTransactor{contract: contract}, LaunchpadFilterer: LaunchpadFilterer{contract: contract}}, nil
}

// NewLaunchpadCaller creates a new read-only instance of Launchpad, bound to a specific deployed contract.
func NewLaunchpadCaller(address common.Address, caller bind.ContractCaller) (*LaunchpadCaller, error) {
	contract, err := bindLaunchpad(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaunchpadCaller{contract: contract}, nil
}

// NewLaunchpadTransactor creates a new write-only instance of Launchpad, bound to a specific deployed contract.
func NewLaunchpadTransactor(address common.Address, transactor bind.ContractTransactor) (*LaunchpadTransactor, error) {
	contract, err := bindLaunchpad(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaunchpadTransactor{contract: contract}, nil
}

// NewLaunchpadFilterer creates a new log filterer instance of Launchpad, bound to a specific deployed contract.
func NewLaunchpadFilterer(address common.Address, filterer bind.ContractFilterer) (*LaunchpadFilterer, error) {
	contract, err := bindLaunchpad(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaunchpadFilterer{contract: contract}, nil
}

// bindLaunchpad binds a generic wrapper to an already deployed contract.
func bindLaunchpad(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LaunchpadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Launchpad *LaunchpadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Launchpad.Contract.LaunchpadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Launchpad *LaunchpadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.Contract.LaunchpadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Launchpad *LaunchpadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Launchpad.Contract.LaunchpadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Launchpad *LaunchpadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Launchpad.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Launchpad *LaunchpadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Launchpad *LaunchpadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Launchpad.Contract.contract.Transact(opts, method, params...)
}

// CONFIGROLE is a free data retrieval call binding the contract method 0xa4d19feb.
//
// Solidity: function CONFIG_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadCaller) CONFIGROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "CONFIG_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGROLE is a free data retrieval call binding the contract method 0xa4d19feb.
//
// Solidity: function CONFIG_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadSession) CONFIGROLE() ([32]byte, error) {
	return _Launchpad.Contract.CONFIGROLE(&_Launchpad.CallOpts)
}

// CONFIGROLE is a free data retrieval call binding the contract method 0xa4d19feb.
//
// Solidity: function CONFIG_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadCallerSession) CONFIGROLE() ([32]byte, error) {
	return _Launchpad.Contract.CONFIGROLE(&_Launchpad.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Launchpad.Contract.DEFAULTADMINROLE(&_Launchpad.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Launchpad *LaunchpadCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Launchpad.Contract.DEFAULTADMINROLE(&_Launchpad.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Launchpad *LaunchpadCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Launchpad *LaunchpadSession) Admin() (common.Address, error) {
	return _Launchpad.Contract.Admin(&_Launchpad.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Launchpad *LaunchpadCallerSession) Admin() (common.Address, error) {
	return _Launchpad.Contract.Admin(&_Launchpad.CallOpts)
}

// CapRatio is a free data retrieval call binding the contract method 0x569a747d.
//
// Solidity: function capRatio() view returns(uint256)
func (_Launchpad *LaunchpadCaller) CapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "capRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CapRatio is a free data retrieval call binding the contract method 0x569a747d.
//
// Solidity: function capRatio() view returns(uint256)
func (_Launchpad *LaunchpadSession) CapRatio() (*big.Int, error) {
	return _Launchpad.Contract.CapRatio(&_Launchpad.CallOpts)
}

// CapRatio is a free data retrieval call binding the contract method 0x569a747d.
//
// Solidity: function capRatio() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) CapRatio() (*big.Int, error) {
	return _Launchpad.Contract.CapRatio(&_Launchpad.CallOpts)
}

// EstimateAmountBePaid is a free data retrieval call binding the contract method 0xa565db92.
//
// Solidity: function estimateAmountBePaid(uint256 amountSubscribed) view returns(uint256 amountBePaid)
func (_Launchpad *LaunchpadCaller) EstimateAmountBePaid(opts *bind.CallOpts, amountSubscribed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "estimateAmountBePaid", amountSubscribed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateAmountBePaid is a free data retrieval call binding the contract method 0xa565db92.
//
// Solidity: function estimateAmountBePaid(uint256 amountSubscribed) view returns(uint256 amountBePaid)
func (_Launchpad *LaunchpadSession) EstimateAmountBePaid(amountSubscribed *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateAmountBePaid(&_Launchpad.CallOpts, amountSubscribed)
}

// EstimateAmountBePaid is a free data retrieval call binding the contract method 0xa565db92.
//
// Solidity: function estimateAmountBePaid(uint256 amountSubscribed) view returns(uint256 amountBePaid)
func (_Launchpad *LaunchpadCallerSession) EstimateAmountBePaid(amountSubscribed *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateAmountBePaid(&_Launchpad.CallOpts, amountSubscribed)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xfb5a4535.
//
// Solidity: function estimateAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadCaller) EstimateAmountOut(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "estimateAmountOut", amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xfb5a4535.
//
// Solidity: function estimateAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadSession) EstimateAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateAmountOut(&_Launchpad.CallOpts, amountIn)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xfb5a4535.
//
// Solidity: function estimateAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadCallerSession) EstimateAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateAmountOut(&_Launchpad.CallOpts, amountIn)
}

// EstimateWhitelistAmountOut is a free data retrieval call binding the contract method 0x63344e18.
//
// Solidity: function estimateWhitelistAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadCaller) EstimateWhitelistAmountOut(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "estimateWhitelistAmountOut", amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateWhitelistAmountOut is a free data retrieval call binding the contract method 0x63344e18.
//
// Solidity: function estimateWhitelistAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadSession) EstimateWhitelistAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateWhitelistAmountOut(&_Launchpad.CallOpts, amountIn)
}

// EstimateWhitelistAmountOut is a free data retrieval call binding the contract method 0x63344e18.
//
// Solidity: function estimateWhitelistAmountOut(uint256 amountIn) view returns(uint256 amountOut)
func (_Launchpad *LaunchpadCallerSession) EstimateWhitelistAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _Launchpad.Contract.EstimateWhitelistAmountOut(&_Launchpad.CallOpts, amountIn)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Launchpad *LaunchpadCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Launchpad *LaunchpadSession) Factory() (common.Address, error) {
	return _Launchpad.Contract.Factory(&_Launchpad.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Launchpad *LaunchpadCallerSession) Factory() (common.Address, error) {
	return _Launchpad.Contract.Factory(&_Launchpad.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Launchpad *LaunchpadCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Launchpad *LaunchpadSession) GetAdmin() (common.Address, error) {
	return _Launchpad.Contract.GetAdmin(&_Launchpad.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Launchpad *LaunchpadCallerSession) GetAdmin() (common.Address, error) {
	return _Launchpad.Contract.GetAdmin(&_Launchpad.CallOpts)
}

// GetHardCap is a free data retrieval call binding the contract method 0xbf0a07bd.
//
// Solidity: function getHardCap() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getHardCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHardCap is a free data retrieval call binding the contract method 0xbf0a07bd.
//
// Solidity: function getHardCap() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetHardCap() (*big.Int, error) {
	return _Launchpad.Contract.GetHardCap(&_Launchpad.CallOpts)
}

// GetHardCap is a free data retrieval call binding the contract method 0xbf0a07bd.
//
// Solidity: function getHardCap() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetHardCap() (*big.Int, error) {
	return _Launchpad.Contract.GetHardCap(&_Launchpad.CallOpts)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(address user) view returns((address,uint256,bool))
func (_Launchpad *LaunchpadCaller) GetInvestment(opts *bind.CallOpts, user common.Address) (Investment, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getInvestment", user)

	if err != nil {
		return *new(Investment), err
	}

	out0 := *abi.ConvertType(out[0], new(Investment)).(*Investment)

	return out0, err

}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(address user) view returns((address,uint256,bool))
func (_Launchpad *LaunchpadSession) GetInvestment(user common.Address) (Investment, error) {
	return _Launchpad.Contract.GetInvestment(&_Launchpad.CallOpts, user)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(address user) view returns((address,uint256,bool))
func (_Launchpad *LaunchpadCallerSession) GetInvestment(user common.Address) (Investment, error) {
	return _Launchpad.Contract.GetInvestment(&_Launchpad.CallOpts, user)
}

// GetLaunchpadDetails is a free data retrieval call binding the contract method 0xe26127e8.
//
// Solidity: function getLaunchpadDetails() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address))
func (_Launchpad *LaunchpadCaller) GetLaunchpadDetails(opts *bind.CallOpts) (LaunchpadDetails, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getLaunchpadDetails")

	if err != nil {
		return *new(LaunchpadDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(LaunchpadDetails)).(*LaunchpadDetails)

	return out0, err

}

// GetLaunchpadDetails is a free data retrieval call binding the contract method 0xe26127e8.
//
// Solidity: function getLaunchpadDetails() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address))
func (_Launchpad *LaunchpadSession) GetLaunchpadDetails() (LaunchpadDetails, error) {
	return _Launchpad.Contract.GetLaunchpadDetails(&_Launchpad.CallOpts)
}

// GetLaunchpadDetails is a free data retrieval call binding the contract method 0xe26127e8.
//
// Solidity: function getLaunchpadDetails() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address))
func (_Launchpad *LaunchpadCallerSession) GetLaunchpadDetails() (LaunchpadDetails, error) {
	return _Launchpad.Contract.GetLaunchpadDetails(&_Launchpad.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetPrice() (*big.Int, error) {
	return _Launchpad.Contract.GetPrice(&_Launchpad.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetPrice() (*big.Int, error) {
	return _Launchpad.Contract.GetPrice(&_Launchpad.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Launchpad *LaunchpadCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Launchpad *LaunchpadSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Launchpad.Contract.GetRoleAdmin(&_Launchpad.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Launchpad *LaunchpadCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Launchpad.Contract.GetRoleAdmin(&_Launchpad.CallOpts, role)
}

// GetSoftCap is a free data retrieval call binding the contract method 0x4d52a512.
//
// Solidity: function getSoftCap() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetSoftCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getSoftCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSoftCap is a free data retrieval call binding the contract method 0x4d52a512.
//
// Solidity: function getSoftCap() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetSoftCap() (*big.Int, error) {
	return _Launchpad.Contract.GetSoftCap(&_Launchpad.CallOpts)
}

// GetSoftCap is a free data retrieval call binding the contract method 0x4d52a512.
//
// Solidity: function getSoftCap() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetSoftCap() (*big.Int, error) {
	return _Launchpad.Contract.GetSoftCap(&_Launchpad.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Launchpad *LaunchpadCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Launchpad *LaunchpadSession) GetToken() (common.Address, error) {
	return _Launchpad.Contract.GetToken(&_Launchpad.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Launchpad *LaunchpadCallerSession) GetToken() (common.Address, error) {
	return _Launchpad.Contract.GetToken(&_Launchpad.CallOpts)
}

// GetTotalInvested is a free data retrieval call binding the contract method 0x5f945733.
//
// Solidity: function getTotalInvested() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetTotalInvested(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getTotalInvested")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInvested is a free data retrieval call binding the contract method 0x5f945733.
//
// Solidity: function getTotalInvested() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetTotalInvested() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalInvested(&_Launchpad.CallOpts)
}

// GetTotalInvested is a free data retrieval call binding the contract method 0x5f945733.
//
// Solidity: function getTotalInvested() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetTotalInvested() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalInvested(&_Launchpad.CallOpts)
}

// GetTotalInvestors is a free data retrieval call binding the contract method 0xd7766d79.
//
// Solidity: function getTotalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetTotalInvestors(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getTotalInvestors")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInvestors is a free data retrieval call binding the contract method 0xd7766d79.
//
// Solidity: function getTotalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetTotalInvestors() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalInvestors(&_Launchpad.CallOpts)
}

// GetTotalInvestors is a free data retrieval call binding the contract method 0xd7766d79.
//
// Solidity: function getTotalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetTotalInvestors() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalInvestors(&_Launchpad.CallOpts)
}

// GetTotalWhitelistSold is a free data retrieval call binding the contract method 0xa52ea34f.
//
// Solidity: function getTotalWhitelistSold() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetTotalWhitelistSold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getTotalWhitelistSold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWhitelistSold is a free data retrieval call binding the contract method 0xa52ea34f.
//
// Solidity: function getTotalWhitelistSold() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetTotalWhitelistSold() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalWhitelistSold(&_Launchpad.CallOpts)
}

// GetTotalWhitelistSold is a free data retrieval call binding the contract method 0xa52ea34f.
//
// Solidity: function getTotalWhitelistSold() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetTotalWhitelistSold() (*big.Int, error) {
	return _Launchpad.Contract.GetTotalWhitelistSold(&_Launchpad.CallOpts)
}

// GetWhitelistDetails is a free data retrieval call binding the contract method 0x64374d10.
//
// Solidity: function getWhitelistDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Launchpad *LaunchpadCaller) GetWhitelistDetails(opts *bind.CallOpts) (WhitelistDetails, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getWhitelistDetails")

	if err != nil {
		return *new(WhitelistDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(WhitelistDetails)).(*WhitelistDetails)

	return out0, err

}

// GetWhitelistDetails is a free data retrieval call binding the contract method 0x64374d10.
//
// Solidity: function getWhitelistDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Launchpad *LaunchpadSession) GetWhitelistDetails() (WhitelistDetails, error) {
	return _Launchpad.Contract.GetWhitelistDetails(&_Launchpad.CallOpts)
}

// GetWhitelistDetails is a free data retrieval call binding the contract method 0x64374d10.
//
// Solidity: function getWhitelistDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Launchpad *LaunchpadCallerSession) GetWhitelistDetails() (WhitelistDetails, error) {
	return _Launchpad.Contract.GetWhitelistDetails(&_Launchpad.CallOpts)
}

// GetWhitelistRemaining is a free data retrieval call binding the contract method 0x173e5777.
//
// Solidity: function getWhitelistRemaining() view returns(uint256)
func (_Launchpad *LaunchpadCaller) GetWhitelistRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "getWhitelistRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistRemaining is a free data retrieval call binding the contract method 0x173e5777.
//
// Solidity: function getWhitelistRemaining() view returns(uint256)
func (_Launchpad *LaunchpadSession) GetWhitelistRemaining() (*big.Int, error) {
	return _Launchpad.Contract.GetWhitelistRemaining(&_Launchpad.CallOpts)
}

// GetWhitelistRemaining is a free data retrieval call binding the contract method 0x173e5777.
//
// Solidity: function getWhitelistRemaining() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) GetWhitelistRemaining() (*big.Int, error) {
	return _Launchpad.Contract.GetWhitelistRemaining(&_Launchpad.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Launchpad *LaunchpadCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Launchpad *LaunchpadSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Launchpad.Contract.HasRole(&_Launchpad.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Launchpad *LaunchpadCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Launchpad.Contract.HasRole(&_Launchpad.CallOpts, role, account)
}

// Investable is a free data retrieval call binding the contract method 0x276d7459.
//
// Solidity: function investable(address user) view returns(bool)
func (_Launchpad *LaunchpadCaller) Investable(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "investable", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Investable is a free data retrieval call binding the contract method 0x276d7459.
//
// Solidity: function investable(address user) view returns(bool)
func (_Launchpad *LaunchpadSession) Investable(user common.Address) (bool, error) {
	return _Launchpad.Contract.Investable(&_Launchpad.CallOpts, user)
}

// Investable is a free data retrieval call binding the contract method 0x276d7459.
//
// Solidity: function investable(address user) view returns(bool)
func (_Launchpad *LaunchpadCallerSession) Investable(user common.Address) (bool, error) {
	return _Launchpad.Contract.Investable(&_Launchpad.CallOpts, user)
}

// Investments is a free data retrieval call binding the contract method 0x96b98862.
//
// Solidity: function investments(address ) view returns(address investorAddress, uint256 investmentAmount, bool claimed)
func (_Launchpad *LaunchpadCaller) Investments(opts *bind.CallOpts, arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Claimed          bool
}, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "investments", arg0)

	outstruct := new(struct {
		InvestorAddress  common.Address
		InvestmentAmount *big.Int
		Claimed          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InvestorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.InvestmentAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Investments is a free data retrieval call binding the contract method 0x96b98862.
//
// Solidity: function investments(address ) view returns(address investorAddress, uint256 investmentAmount, bool claimed)
func (_Launchpad *LaunchpadSession) Investments(arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Claimed          bool
}, error) {
	return _Launchpad.Contract.Investments(&_Launchpad.CallOpts, arg0)
}

// Investments is a free data retrieval call binding the contract method 0x96b98862.
//
// Solidity: function investments(address ) view returns(address investorAddress, uint256 investmentAmount, bool claimed)
func (_Launchpad *LaunchpadCallerSession) Investments(arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Claimed          bool
}, error) {
	return _Launchpad.Contract.Investments(&_Launchpad.CallOpts, arg0)
}

// IsClaimable is a free data retrieval call binding the contract method 0x74478bb3.
//
// Solidity: function isClaimable() view returns(bool)
func (_Launchpad *LaunchpadCaller) IsClaimable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "isClaimable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimable is a free data retrieval call binding the contract method 0x74478bb3.
//
// Solidity: function isClaimable() view returns(bool)
func (_Launchpad *LaunchpadSession) IsClaimable() (bool, error) {
	return _Launchpad.Contract.IsClaimable(&_Launchpad.CallOpts)
}

// IsClaimable is a free data retrieval call binding the contract method 0x74478bb3.
//
// Solidity: function isClaimable() view returns(bool)
func (_Launchpad *LaunchpadCallerSession) IsClaimable() (bool, error) {
	return _Launchpad.Contract.IsClaimable(&_Launchpad.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_Launchpad *LaunchpadCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "isFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_Launchpad *LaunchpadSession) IsFinished() (bool, error) {
	return _Launchpad.Contract.IsFinished(&_Launchpad.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_Launchpad *LaunchpadCallerSession) IsFinished() (bool, error) {
	return _Launchpad.Contract.IsFinished(&_Launchpad.CallOpts)
}

// IsInWhitelistPeriod is a free data retrieval call binding the contract method 0xe33b631d.
//
// Solidity: function isInWhitelistPeriod() view returns(bool)
func (_Launchpad *LaunchpadCaller) IsInWhitelistPeriod(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "isInWhitelistPeriod")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWhitelistPeriod is a free data retrieval call binding the contract method 0xe33b631d.
//
// Solidity: function isInWhitelistPeriod() view returns(bool)
func (_Launchpad *LaunchpadSession) IsInWhitelistPeriod() (bool, error) {
	return _Launchpad.Contract.IsInWhitelistPeriod(&_Launchpad.CallOpts)
}

// IsInWhitelistPeriod is a free data retrieval call binding the contract method 0xe33b631d.
//
// Solidity: function isInWhitelistPeriod() view returns(bool)
func (_Launchpad *LaunchpadCallerSession) IsInWhitelistPeriod() (bool, error) {
	return _Launchpad.Contract.IsInWhitelistPeriod(&_Launchpad.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address user) view returns(bool)
func (_Launchpad *LaunchpadCaller) IsWhitelisted(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "isWhitelisted", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address user) view returns(bool)
func (_Launchpad *LaunchpadSession) IsWhitelisted(user common.Address) (bool, error) {
	return _Launchpad.Contract.IsWhitelisted(&_Launchpad.CallOpts, user)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address user) view returns(bool)
func (_Launchpad *LaunchpadCallerSession) IsWhitelisted(user common.Address) (bool, error) {
	return _Launchpad.Contract.IsWhitelisted(&_Launchpad.CallOpts, user)
}

// LiquidityRate is a free data retrieval call binding the contract method 0x548cbeb0.
//
// Solidity: function liquidityRate() view returns(uint256)
func (_Launchpad *LaunchpadCaller) LiquidityRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "liquidityRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityRate is a free data retrieval call binding the contract method 0x548cbeb0.
//
// Solidity: function liquidityRate() view returns(uint256)
func (_Launchpad *LaunchpadSession) LiquidityRate() (*big.Int, error) {
	return _Launchpad.Contract.LiquidityRate(&_Launchpad.CallOpts)
}

// LiquidityRate is a free data retrieval call binding the contract method 0x548cbeb0.
//
// Solidity: function liquidityRate() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) LiquidityRate() (*big.Int, error) {
	return _Launchpad.Contract.LiquidityRate(&_Launchpad.CallOpts)
}

// LiquidityWallet is a free data retrieval call binding the contract method 0xd4698016.
//
// Solidity: function liquidityWallet() view returns(address)
func (_Launchpad *LaunchpadCaller) LiquidityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "liquidityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityWallet is a free data retrieval call binding the contract method 0xd4698016.
//
// Solidity: function liquidityWallet() view returns(address)
func (_Launchpad *LaunchpadSession) LiquidityWallet() (common.Address, error) {
	return _Launchpad.Contract.LiquidityWallet(&_Launchpad.CallOpts)
}

// LiquidityWallet is a free data retrieval call binding the contract method 0xd4698016.
//
// Solidity: function liquidityWallet() view returns(address)
func (_Launchpad *LaunchpadCallerSession) LiquidityWallet() (common.Address, error) {
	return _Launchpad.Contract.LiquidityWallet(&_Launchpad.CallOpts)
}

// MaxInvestment is a free data retrieval call binding the contract method 0x002e1316.
//
// Solidity: function maxInvestment() view returns(uint256)
func (_Launchpad *LaunchpadCaller) MaxInvestment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "maxInvestment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxInvestment is a free data retrieval call binding the contract method 0x002e1316.
//
// Solidity: function maxInvestment() view returns(uint256)
func (_Launchpad *LaunchpadSession) MaxInvestment() (*big.Int, error) {
	return _Launchpad.Contract.MaxInvestment(&_Launchpad.CallOpts)
}

// MaxInvestment is a free data retrieval call binding the contract method 0x002e1316.
//
// Solidity: function maxInvestment() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) MaxInvestment() (*big.Int, error) {
	return _Launchpad.Contract.MaxInvestment(&_Launchpad.CallOpts)
}

// MinInvestment is a free data retrieval call binding the contract method 0x8ac2c680.
//
// Solidity: function minInvestment() view returns(uint256)
func (_Launchpad *LaunchpadCaller) MinInvestment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "minInvestment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinInvestment is a free data retrieval call binding the contract method 0x8ac2c680.
//
// Solidity: function minInvestment() view returns(uint256)
func (_Launchpad *LaunchpadSession) MinInvestment() (*big.Int, error) {
	return _Launchpad.Contract.MinInvestment(&_Launchpad.CallOpts)
}

// MinInvestment is a free data retrieval call binding the contract method 0x8ac2c680.
//
// Solidity: function minInvestment() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) MinInvestment() (*big.Int, error) {
	return _Launchpad.Contract.MinInvestment(&_Launchpad.CallOpts)
}

// PlatformRate is a free data retrieval call binding the contract method 0x36e0cac6.
//
// Solidity: function platformRate() view returns(uint256)
func (_Launchpad *LaunchpadCaller) PlatformRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "platformRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformRate is a free data retrieval call binding the contract method 0x36e0cac6.
//
// Solidity: function platformRate() view returns(uint256)
func (_Launchpad *LaunchpadSession) PlatformRate() (*big.Int, error) {
	return _Launchpad.Contract.PlatformRate(&_Launchpad.CallOpts)
}

// PlatformRate is a free data retrieval call binding the contract method 0x36e0cac6.
//
// Solidity: function platformRate() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) PlatformRate() (*big.Int, error) {
	return _Launchpad.Contract.PlatformRate(&_Launchpad.CallOpts)
}

// PubEndTime is a free data retrieval call binding the contract method 0x4222eb93.
//
// Solidity: function pubEndTime() view returns(uint256)
func (_Launchpad *LaunchpadCaller) PubEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "pubEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PubEndTime is a free data retrieval call binding the contract method 0x4222eb93.
//
// Solidity: function pubEndTime() view returns(uint256)
func (_Launchpad *LaunchpadSession) PubEndTime() (*big.Int, error) {
	return _Launchpad.Contract.PubEndTime(&_Launchpad.CallOpts)
}

// PubEndTime is a free data retrieval call binding the contract method 0x4222eb93.
//
// Solidity: function pubEndTime() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) PubEndTime() (*big.Int, error) {
	return _Launchpad.Contract.PubEndTime(&_Launchpad.CallOpts)
}

// PubPrice is a free data retrieval call binding the contract method 0xbddb7be7.
//
// Solidity: function pubPrice() view returns(uint256)
func (_Launchpad *LaunchpadCaller) PubPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "pubPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PubPrice is a free data retrieval call binding the contract method 0xbddb7be7.
//
// Solidity: function pubPrice() view returns(uint256)
func (_Launchpad *LaunchpadSession) PubPrice() (*big.Int, error) {
	return _Launchpad.Contract.PubPrice(&_Launchpad.CallOpts)
}

// PubPrice is a free data retrieval call binding the contract method 0xbddb7be7.
//
// Solidity: function pubPrice() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) PubPrice() (*big.Int, error) {
	return _Launchpad.Contract.PubPrice(&_Launchpad.CallOpts)
}

// PubStartTime is a free data retrieval call binding the contract method 0xf5ebb47d.
//
// Solidity: function pubStartTime() view returns(uint256)
func (_Launchpad *LaunchpadCaller) PubStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "pubStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PubStartTime is a free data retrieval call binding the contract method 0xf5ebb47d.
//
// Solidity: function pubStartTime() view returns(uint256)
func (_Launchpad *LaunchpadSession) PubStartTime() (*big.Int, error) {
	return _Launchpad.Contract.PubStartTime(&_Launchpad.CallOpts)
}

// PubStartTime is a free data retrieval call binding the contract method 0xf5ebb47d.
//
// Solidity: function pubStartTime() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) PubStartTime() (*big.Int, error) {
	return _Launchpad.Contract.PubStartTime(&_Launchpad.CallOpts)
}

// RaiseRate is a free data retrieval call binding the contract method 0x1a2385ee.
//
// Solidity: function raiseRate() view returns(uint256)
func (_Launchpad *LaunchpadCaller) RaiseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "raiseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiseRate is a free data retrieval call binding the contract method 0x1a2385ee.
//
// Solidity: function raiseRate() view returns(uint256)
func (_Launchpad *LaunchpadSession) RaiseRate() (*big.Int, error) {
	return _Launchpad.Contract.RaiseRate(&_Launchpad.CallOpts)
}

// RaiseRate is a free data retrieval call binding the contract method 0x1a2385ee.
//
// Solidity: function raiseRate() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) RaiseRate() (*big.Int, error) {
	return _Launchpad.Contract.RaiseRate(&_Launchpad.CallOpts)
}

// RaiseWallet is a free data retrieval call binding the contract method 0x7d76465c.
//
// Solidity: function raiseWallet() view returns(address)
func (_Launchpad *LaunchpadCaller) RaiseWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "raiseWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaiseWallet is a free data retrieval call binding the contract method 0x7d76465c.
//
// Solidity: function raiseWallet() view returns(address)
func (_Launchpad *LaunchpadSession) RaiseWallet() (common.Address, error) {
	return _Launchpad.Contract.RaiseWallet(&_Launchpad.CallOpts)
}

// RaiseWallet is a free data retrieval call binding the contract method 0x7d76465c.
//
// Solidity: function raiseWallet() view returns(address)
func (_Launchpad *LaunchpadCallerSession) RaiseWallet() (common.Address, error) {
	return _Launchpad.Contract.RaiseWallet(&_Launchpad.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Launchpad *LaunchpadCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Launchpad *LaunchpadSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Launchpad.Contract.SupportsInterface(&_Launchpad.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Launchpad *LaunchpadCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Launchpad.Contract.SupportsInterface(&_Launchpad.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Launchpad *LaunchpadCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Launchpad *LaunchpadSession) Token() (common.Address, error) {
	return _Launchpad.Contract.Token(&_Launchpad.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Launchpad *LaunchpadCallerSession) Token() (common.Address, error) {
	return _Launchpad.Contract.Token(&_Launchpad.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Launchpad *LaunchpadCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Launchpad *LaunchpadSession) TokenDecimals() (uint8, error) {
	return _Launchpad.Contract.TokenDecimals(&_Launchpad.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Launchpad *LaunchpadCallerSession) TokenDecimals() (uint8, error) {
	return _Launchpad.Contract.TokenDecimals(&_Launchpad.CallOpts)
}

// TotalInvestors is a free data retrieval call binding the contract method 0x29b8caff.
//
// Solidity: function totalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalInvestors(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalInvestors")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInvestors is a free data retrieval call binding the contract method 0x29b8caff.
//
// Solidity: function totalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalInvestors() (*big.Int, error) {
	return _Launchpad.Contract.TotalInvestors(&_Launchpad.CallOpts)
}

// TotalInvestors is a free data retrieval call binding the contract method 0x29b8caff.
//
// Solidity: function totalInvestors() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalInvestors() (*big.Int, error) {
	return _Launchpad.Contract.TotalInvestors(&_Launchpad.CallOpts)
}

// TotalPublicSale is a free data retrieval call binding the contract method 0xe66e995d.
//
// Solidity: function totalPublicSale() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalPublicSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalPublicSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPublicSale is a free data retrieval call binding the contract method 0xe66e995d.
//
// Solidity: function totalPublicSale() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalPublicSale() (*big.Int, error) {
	return _Launchpad.Contract.TotalPublicSale(&_Launchpad.CallOpts)
}

// TotalPublicSale is a free data retrieval call binding the contract method 0xe66e995d.
//
// Solidity: function totalPublicSale() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalPublicSale() (*big.Int, error) {
	return _Launchpad.Contract.TotalPublicSale(&_Launchpad.CallOpts)
}

// TotalPublicSubscribed is a free data retrieval call binding the contract method 0xec468dbf.
//
// Solidity: function totalPublicSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalPublicSubscribed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalPublicSubscribed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPublicSubscribed is a free data retrieval call binding the contract method 0xec468dbf.
//
// Solidity: function totalPublicSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalPublicSubscribed() (*big.Int, error) {
	return _Launchpad.Contract.TotalPublicSubscribed(&_Launchpad.CallOpts)
}

// TotalPublicSubscribed is a free data retrieval call binding the contract method 0xec468dbf.
//
// Solidity: function totalPublicSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalPublicSubscribed() (*big.Int, error) {
	return _Launchpad.Contract.TotalPublicSubscribed(&_Launchpad.CallOpts)
}

// TotalWhitelistSale is a free data retrieval call binding the contract method 0x968ee09e.
//
// Solidity: function totalWhitelistSale() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalWhitelistSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalWhitelistSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWhitelistSale is a free data retrieval call binding the contract method 0x968ee09e.
//
// Solidity: function totalWhitelistSale() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalWhitelistSale() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelistSale(&_Launchpad.CallOpts)
}

// TotalWhitelistSale is a free data retrieval call binding the contract method 0x968ee09e.
//
// Solidity: function totalWhitelistSale() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalWhitelistSale() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelistSale(&_Launchpad.CallOpts)
}

// TotalWhitelistSubscribed is a free data retrieval call binding the contract method 0x491caa85.
//
// Solidity: function totalWhitelistSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalWhitelistSubscribed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalWhitelistSubscribed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWhitelistSubscribed is a free data retrieval call binding the contract method 0x491caa85.
//
// Solidity: function totalWhitelistSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalWhitelistSubscribed() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelistSubscribed(&_Launchpad.CallOpts)
}

// TotalWhitelistSubscribed is a free data retrieval call binding the contract method 0x491caa85.
//
// Solidity: function totalWhitelistSubscribed() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalWhitelistSubscribed() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelistSubscribed(&_Launchpad.CallOpts)
}

// TotalWhitelisted is a free data retrieval call binding the contract method 0x04072b20.
//
// Solidity: function totalWhitelisted() view returns(uint256)
func (_Launchpad *LaunchpadCaller) TotalWhitelisted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "totalWhitelisted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWhitelisted is a free data retrieval call binding the contract method 0x04072b20.
//
// Solidity: function totalWhitelisted() view returns(uint256)
func (_Launchpad *LaunchpadSession) TotalWhitelisted() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelisted(&_Launchpad.CallOpts)
}

// TotalWhitelisted is a free data retrieval call binding the contract method 0x04072b20.
//
// Solidity: function totalWhitelisted() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) TotalWhitelisted() (*big.Int, error) {
	return _Launchpad.Contract.TotalWhitelisted(&_Launchpad.CallOpts)
}

// WhitelistEndTime is a free data retrieval call binding the contract method 0xebdfd722.
//
// Solidity: function whitelistEndTime() view returns(uint256)
func (_Launchpad *LaunchpadCaller) WhitelistEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "whitelistEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistEndTime is a free data retrieval call binding the contract method 0xebdfd722.
//
// Solidity: function whitelistEndTime() view returns(uint256)
func (_Launchpad *LaunchpadSession) WhitelistEndTime() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistEndTime(&_Launchpad.CallOpts)
}

// WhitelistEndTime is a free data retrieval call binding the contract method 0xebdfd722.
//
// Solidity: function whitelistEndTime() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) WhitelistEndTime() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistEndTime(&_Launchpad.CallOpts)
}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_Launchpad *LaunchpadCaller) WhitelistPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "whitelistPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_Launchpad *LaunchpadSession) WhitelistPrice() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistPrice(&_Launchpad.CallOpts)
}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) WhitelistPrice() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistPrice(&_Launchpad.CallOpts)
}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Launchpad *LaunchpadCaller) WhitelistStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "whitelistStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Launchpad *LaunchpadSession) WhitelistStartTime() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistStartTime(&_Launchpad.CallOpts)
}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Launchpad *LaunchpadCallerSession) WhitelistStartTime() (*big.Int, error) {
	return _Launchpad.Contract.WhitelistStartTime(&_Launchpad.CallOpts)
}

// Whitelists is a free data retrieval call binding the contract method 0x1e7be210.
//
// Solidity: function whitelists(address ) view returns(address investorAddress, uint256 investmentAmount, uint256 maxInvestment, bool claimed)
func (_Launchpad *LaunchpadCaller) Whitelists(opts *bind.CallOpts, arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	MaxInvestment    *big.Int
	Claimed          bool
}, error) {
	var out []interface{}
	err := _Launchpad.contract.Call(opts, &out, "whitelists", arg0)

	outstruct := new(struct {
		InvestorAddress  common.Address
		InvestmentAmount *big.Int
		MaxInvestment    *big.Int
		Claimed          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InvestorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.InvestmentAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxInvestment = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Whitelists is a free data retrieval call binding the contract method 0x1e7be210.
//
// Solidity: function whitelists(address ) view returns(address investorAddress, uint256 investmentAmount, uint256 maxInvestment, bool claimed)
func (_Launchpad *LaunchpadSession) Whitelists(arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	MaxInvestment    *big.Int
	Claimed          bool
}, error) {
	return _Launchpad.Contract.Whitelists(&_Launchpad.CallOpts, arg0)
}

// Whitelists is a free data retrieval call binding the contract method 0x1e7be210.
//
// Solidity: function whitelists(address ) view returns(address investorAddress, uint256 investmentAmount, uint256 maxInvestment, bool claimed)
func (_Launchpad *LaunchpadCallerSession) Whitelists(arg0 common.Address) (struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	MaxInvestment    *big.Int
	Claimed          bool
}, error) {
	return _Launchpad.Contract.Whitelists(&_Launchpad.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Launchpad *LaunchpadTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Launchpad *LaunchpadSession) Claim() (*types.Transaction, error) {
	return _Launchpad.Contract.Claim(&_Launchpad.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Launchpad *LaunchpadTransactorSession) Claim() (*types.Transaction, error) {
	return _Launchpad.Contract.Claim(&_Launchpad.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_Launchpad *LaunchpadTransactor) EmergencyWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "emergencyWithdraw", token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_Launchpad *LaunchpadSession) EmergencyWithdraw(token common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.EmergencyWithdraw(&_Launchpad.TransactOpts, token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_Launchpad *LaunchpadTransactorSession) EmergencyWithdraw(token common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.EmergencyWithdraw(&_Launchpad.TransactOpts, token)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_Launchpad *LaunchpadTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_Launchpad *LaunchpadSession) Finish() (*types.Transaction, error) {
	return _Launchpad.Contract.Finish(&_Launchpad.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_Launchpad *LaunchpadTransactorSession) Finish() (*types.Transaction, error) {
	return _Launchpad.Contract.Finish(&_Launchpad.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.GrantRole(&_Launchpad.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.GrantRole(&_Launchpad.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xdeba076b.
//
// Solidity: function initialize((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) returns()
func (_Launchpad *LaunchpadTransactor) Initialize(opts *bind.TransactOpts, params InitParams) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "initialize", params)
}

// Initialize is a paid mutator transaction binding the contract method 0xdeba076b.
//
// Solidity: function initialize((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) returns()
func (_Launchpad *LaunchpadSession) Initialize(params InitParams) (*types.Transaction, error) {
	return _Launchpad.Contract.Initialize(&_Launchpad.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xdeba076b.
//
// Solidity: function initialize((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) returns()
func (_Launchpad *LaunchpadTransactorSession) Initialize(params InitParams) (*types.Transaction, error) {
	return _Launchpad.Contract.Initialize(&_Launchpad.TransactOpts, params)
}

// Invest is a paid mutator transaction binding the contract method 0xe8b5e51f.
//
// Solidity: function invest() payable returns()
func (_Launchpad *LaunchpadTransactor) Invest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "invest")
}

// Invest is a paid mutator transaction binding the contract method 0xe8b5e51f.
//
// Solidity: function invest() payable returns()
func (_Launchpad *LaunchpadSession) Invest() (*types.Transaction, error) {
	return _Launchpad.Contract.Invest(&_Launchpad.TransactOpts)
}

// Invest is a paid mutator transaction binding the contract method 0xe8b5e51f.
//
// Solidity: function invest() payable returns()
func (_Launchpad *LaunchpadTransactorSession) Invest() (*types.Transaction, error) {
	return _Launchpad.Contract.Invest(&_Launchpad.TransactOpts)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] _whitelists) returns()
func (_Launchpad *LaunchpadTransactor) RemoveWhitelist(opts *bind.TransactOpts, _whitelists []common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "removeWhitelist", _whitelists)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] _whitelists) returns()
func (_Launchpad *LaunchpadSession) RemoveWhitelist(_whitelists []common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RemoveWhitelist(&_Launchpad.TransactOpts, _whitelists)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] _whitelists) returns()
func (_Launchpad *LaunchpadTransactorSession) RemoveWhitelist(_whitelists []common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RemoveWhitelist(&_Launchpad.TransactOpts, _whitelists)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Launchpad *LaunchpadTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Launchpad *LaunchpadSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RenounceRole(&_Launchpad.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Launchpad *LaunchpadTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RenounceRole(&_Launchpad.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RevokeRole(&_Launchpad.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Launchpad *LaunchpadTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.RevokeRole(&_Launchpad.TransactOpts, role, account)
}

// SetCapRatio is a paid mutator transaction binding the contract method 0xc06fe4ab.
//
// Solidity: function setCapRatio(uint256 _capRatio) returns()
func (_Launchpad *LaunchpadTransactor) SetCapRatio(opts *bind.TransactOpts, _capRatio *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setCapRatio", _capRatio)
}

// SetCapRatio is a paid mutator transaction binding the contract method 0xc06fe4ab.
//
// Solidity: function setCapRatio(uint256 _capRatio) returns()
func (_Launchpad *LaunchpadSession) SetCapRatio(_capRatio *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetCapRatio(&_Launchpad.TransactOpts, _capRatio)
}

// SetCapRatio is a paid mutator transaction binding the contract method 0xc06fe4ab.
//
// Solidity: function setCapRatio(uint256 _capRatio) returns()
func (_Launchpad *LaunchpadTransactorSession) SetCapRatio(_capRatio *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetCapRatio(&_Launchpad.TransactOpts, _capRatio)
}

// SetClaimable is a paid mutator transaction binding the contract method 0x268f993c.
//
// Solidity: function setClaimable() returns()
func (_Launchpad *LaunchpadTransactor) SetClaimable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setClaimable")
}

// SetClaimable is a paid mutator transaction binding the contract method 0x268f993c.
//
// Solidity: function setClaimable() returns()
func (_Launchpad *LaunchpadSession) SetClaimable() (*types.Transaction, error) {
	return _Launchpad.Contract.SetClaimable(&_Launchpad.TransactOpts)
}

// SetClaimable is a paid mutator transaction binding the contract method 0x268f993c.
//
// Solidity: function setClaimable() returns()
func (_Launchpad *LaunchpadTransactorSession) SetClaimable() (*types.Transaction, error) {
	return _Launchpad.Contract.SetClaimable(&_Launchpad.TransactOpts)
}

// SetInvestmentThreshold is a paid mutator transaction binding the contract method 0x7788a98a.
//
// Solidity: function setInvestmentThreshold(uint256 _minInvestment, uint256 _maxInvestment) returns()
func (_Launchpad *LaunchpadTransactor) SetInvestmentThreshold(opts *bind.TransactOpts, _minInvestment *big.Int, _maxInvestment *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setInvestmentThreshold", _minInvestment, _maxInvestment)
}

// SetInvestmentThreshold is a paid mutator transaction binding the contract method 0x7788a98a.
//
// Solidity: function setInvestmentThreshold(uint256 _minInvestment, uint256 _maxInvestment) returns()
func (_Launchpad *LaunchpadSession) SetInvestmentThreshold(_minInvestment *big.Int, _maxInvestment *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetInvestmentThreshold(&_Launchpad.TransactOpts, _minInvestment, _maxInvestment)
}

// SetInvestmentThreshold is a paid mutator transaction binding the contract method 0x7788a98a.
//
// Solidity: function setInvestmentThreshold(uint256 _minInvestment, uint256 _maxInvestment) returns()
func (_Launchpad *LaunchpadTransactorSession) SetInvestmentThreshold(_minInvestment *big.Int, _maxInvestment *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetInvestmentThreshold(&_Launchpad.TransactOpts, _minInvestment, _maxInvestment)
}

// SetLiquidityWallet is a paid mutator transaction binding the contract method 0x296f0a0c.
//
// Solidity: function setLiquidityWallet(address _liquidityWallet) returns()
func (_Launchpad *LaunchpadTransactor) SetLiquidityWallet(opts *bind.TransactOpts, _liquidityWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setLiquidityWallet", _liquidityWallet)
}

// SetLiquidityWallet is a paid mutator transaction binding the contract method 0x296f0a0c.
//
// Solidity: function setLiquidityWallet(address _liquidityWallet) returns()
func (_Launchpad *LaunchpadSession) SetLiquidityWallet(_liquidityWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.SetLiquidityWallet(&_Launchpad.TransactOpts, _liquidityWallet)
}

// SetLiquidityWallet is a paid mutator transaction binding the contract method 0x296f0a0c.
//
// Solidity: function setLiquidityWallet(address _liquidityWallet) returns()
func (_Launchpad *LaunchpadTransactorSession) SetLiquidityWallet(_liquidityWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.SetLiquidityWallet(&_Launchpad.TransactOpts, _liquidityWallet)
}

// SetPubPrice is a paid mutator transaction binding the contract method 0xdcd5994e.
//
// Solidity: function setPubPrice(uint256 _pubPrice) returns()
func (_Launchpad *LaunchpadTransactor) SetPubPrice(opts *bind.TransactOpts, _pubPrice *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setPubPrice", _pubPrice)
}

// SetPubPrice is a paid mutator transaction binding the contract method 0xdcd5994e.
//
// Solidity: function setPubPrice(uint256 _pubPrice) returns()
func (_Launchpad *LaunchpadSession) SetPubPrice(_pubPrice *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetPubPrice(&_Launchpad.TransactOpts, _pubPrice)
}

// SetPubPrice is a paid mutator transaction binding the contract method 0xdcd5994e.
//
// Solidity: function setPubPrice(uint256 _pubPrice) returns()
func (_Launchpad *LaunchpadTransactorSession) SetPubPrice(_pubPrice *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetPubPrice(&_Launchpad.TransactOpts, _pubPrice)
}

// SetPublicSaleTime is a paid mutator transaction binding the contract method 0x54645d6a.
//
// Solidity: function setPublicSaleTime(uint256 _pubStartTime, uint256 _pubEndTime) returns()
func (_Launchpad *LaunchpadTransactor) SetPublicSaleTime(opts *bind.TransactOpts, _pubStartTime *big.Int, _pubEndTime *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setPublicSaleTime", _pubStartTime, _pubEndTime)
}

// SetPublicSaleTime is a paid mutator transaction binding the contract method 0x54645d6a.
//
// Solidity: function setPublicSaleTime(uint256 _pubStartTime, uint256 _pubEndTime) returns()
func (_Launchpad *LaunchpadSession) SetPublicSaleTime(_pubStartTime *big.Int, _pubEndTime *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetPublicSaleTime(&_Launchpad.TransactOpts, _pubStartTime, _pubEndTime)
}

// SetPublicSaleTime is a paid mutator transaction binding the contract method 0x54645d6a.
//
// Solidity: function setPublicSaleTime(uint256 _pubStartTime, uint256 _pubEndTime) returns()
func (_Launchpad *LaunchpadTransactorSession) SetPublicSaleTime(_pubStartTime *big.Int, _pubEndTime *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetPublicSaleTime(&_Launchpad.TransactOpts, _pubStartTime, _pubEndTime)
}

// SetRaiseWallet is a paid mutator transaction binding the contract method 0x9623a563.
//
// Solidity: function setRaiseWallet(address _raiseWallet) returns()
func (_Launchpad *LaunchpadTransactor) SetRaiseWallet(opts *bind.TransactOpts, _raiseWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setRaiseWallet", _raiseWallet)
}

// SetRaiseWallet is a paid mutator transaction binding the contract method 0x9623a563.
//
// Solidity: function setRaiseWallet(address _raiseWallet) returns()
func (_Launchpad *LaunchpadSession) SetRaiseWallet(_raiseWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.SetRaiseWallet(&_Launchpad.TransactOpts, _raiseWallet)
}

// SetRaiseWallet is a paid mutator transaction binding the contract method 0x9623a563.
//
// Solidity: function setRaiseWallet(address _raiseWallet) returns()
func (_Launchpad *LaunchpadTransactorSession) SetRaiseWallet(_raiseWallet common.Address) (*types.Transaction, error) {
	return _Launchpad.Contract.SetRaiseWallet(&_Launchpad.TransactOpts, _raiseWallet)
}

// SetRates is a paid mutator transaction binding the contract method 0xd004f38b.
//
// Solidity: function setRates(uint256 _raiseRate, uint256 _liquidityRate, uint256 _platformRate) returns()
func (_Launchpad *LaunchpadTransactor) SetRates(opts *bind.TransactOpts, _raiseRate *big.Int, _liquidityRate *big.Int, _platformRate *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setRates", _raiseRate, _liquidityRate, _platformRate)
}

// SetRates is a paid mutator transaction binding the contract method 0xd004f38b.
//
// Solidity: function setRates(uint256 _raiseRate, uint256 _liquidityRate, uint256 _platformRate) returns()
func (_Launchpad *LaunchpadSession) SetRates(_raiseRate *big.Int, _liquidityRate *big.Int, _platformRate *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetRates(&_Launchpad.TransactOpts, _raiseRate, _liquidityRate, _platformRate)
}

// SetRates is a paid mutator transaction binding the contract method 0xd004f38b.
//
// Solidity: function setRates(uint256 _raiseRate, uint256 _liquidityRate, uint256 _platformRate) returns()
func (_Launchpad *LaunchpadTransactorSession) SetRates(_raiseRate *big.Int, _liquidityRate *big.Int, _platformRate *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetRates(&_Launchpad.TransactOpts, _raiseRate, _liquidityRate, _platformRate)
}

// SetTotalPublicSale is a paid mutator transaction binding the contract method 0x849e7a99.
//
// Solidity: function setTotalPublicSale(uint256 _totalPublicSale) returns()
func (_Launchpad *LaunchpadTransactor) SetTotalPublicSale(opts *bind.TransactOpts, _totalPublicSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setTotalPublicSale", _totalPublicSale)
}

// SetTotalPublicSale is a paid mutator transaction binding the contract method 0x849e7a99.
//
// Solidity: function setTotalPublicSale(uint256 _totalPublicSale) returns()
func (_Launchpad *LaunchpadSession) SetTotalPublicSale(_totalPublicSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetTotalPublicSale(&_Launchpad.TransactOpts, _totalPublicSale)
}

// SetTotalPublicSale is a paid mutator transaction binding the contract method 0x849e7a99.
//
// Solidity: function setTotalPublicSale(uint256 _totalPublicSale) returns()
func (_Launchpad *LaunchpadTransactorSession) SetTotalPublicSale(_totalPublicSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetTotalPublicSale(&_Launchpad.TransactOpts, _totalPublicSale)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf013e0e1.
//
// Solidity: function setWhitelist(address[] _whitelists, uint256[] _maxInvestments) returns()
func (_Launchpad *LaunchpadTransactor) SetWhitelist(opts *bind.TransactOpts, _whitelists []common.Address, _maxInvestments []*big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setWhitelist", _whitelists, _maxInvestments)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf013e0e1.
//
// Solidity: function setWhitelist(address[] _whitelists, uint256[] _maxInvestments) returns()
func (_Launchpad *LaunchpadSession) SetWhitelist(_whitelists []common.Address, _maxInvestments []*big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetWhitelist(&_Launchpad.TransactOpts, _whitelists, _maxInvestments)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0xf013e0e1.
//
// Solidity: function setWhitelist(address[] _whitelists, uint256[] _maxInvestments) returns()
func (_Launchpad *LaunchpadTransactorSession) SetWhitelist(_whitelists []common.Address, _maxInvestments []*big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetWhitelist(&_Launchpad.TransactOpts, _whitelists, _maxInvestments)
}

// SetWhitelistParams is a paid mutator transaction binding the contract method 0x95f8f13a.
//
// Solidity: function setWhitelistParams(uint256 _whitelistStartTime, uint256 _whitelistEndTime, uint256 _whitelistPrice, uint256 _totalWhitelistSale) returns()
func (_Launchpad *LaunchpadTransactor) SetWhitelistParams(opts *bind.TransactOpts, _whitelistStartTime *big.Int, _whitelistEndTime *big.Int, _whitelistPrice *big.Int, _totalWhitelistSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "setWhitelistParams", _whitelistStartTime, _whitelistEndTime, _whitelistPrice, _totalWhitelistSale)
}

// SetWhitelistParams is a paid mutator transaction binding the contract method 0x95f8f13a.
//
// Solidity: function setWhitelistParams(uint256 _whitelistStartTime, uint256 _whitelistEndTime, uint256 _whitelistPrice, uint256 _totalWhitelistSale) returns()
func (_Launchpad *LaunchpadSession) SetWhitelistParams(_whitelistStartTime *big.Int, _whitelistEndTime *big.Int, _whitelistPrice *big.Int, _totalWhitelistSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetWhitelistParams(&_Launchpad.TransactOpts, _whitelistStartTime, _whitelistEndTime, _whitelistPrice, _totalWhitelistSale)
}

// SetWhitelistParams is a paid mutator transaction binding the contract method 0x95f8f13a.
//
// Solidity: function setWhitelistParams(uint256 _whitelistStartTime, uint256 _whitelistEndTime, uint256 _whitelistPrice, uint256 _totalWhitelistSale) returns()
func (_Launchpad *LaunchpadTransactorSession) SetWhitelistParams(_whitelistStartTime *big.Int, _whitelistEndTime *big.Int, _whitelistPrice *big.Int, _totalWhitelistSale *big.Int) (*types.Transaction, error) {
	return _Launchpad.Contract.SetWhitelistParams(&_Launchpad.TransactOpts, _whitelistStartTime, _whitelistEndTime, _whitelistPrice, _totalWhitelistSale)
}

// WhitelistClaim is a paid mutator transaction binding the contract method 0x1f131fb4.
//
// Solidity: function whitelistClaim() returns()
func (_Launchpad *LaunchpadTransactor) WhitelistClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "whitelistClaim")
}

// WhitelistClaim is a paid mutator transaction binding the contract method 0x1f131fb4.
//
// Solidity: function whitelistClaim() returns()
func (_Launchpad *LaunchpadSession) WhitelistClaim() (*types.Transaction, error) {
	return _Launchpad.Contract.WhitelistClaim(&_Launchpad.TransactOpts)
}

// WhitelistClaim is a paid mutator transaction binding the contract method 0x1f131fb4.
//
// Solidity: function whitelistClaim() returns()
func (_Launchpad *LaunchpadTransactorSession) WhitelistClaim() (*types.Transaction, error) {
	return _Launchpad.Contract.WhitelistClaim(&_Launchpad.TransactOpts)
}

// WhitelistInvest is a paid mutator transaction binding the contract method 0xdcfe6cb6.
//
// Solidity: function whitelistInvest() payable returns()
func (_Launchpad *LaunchpadTransactor) WhitelistInvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Launchpad.contract.Transact(opts, "whitelistInvest")
}

// WhitelistInvest is a paid mutator transaction binding the contract method 0xdcfe6cb6.
//
// Solidity: function whitelistInvest() payable returns()
func (_Launchpad *LaunchpadSession) WhitelistInvest() (*types.Transaction, error) {
	return _Launchpad.Contract.WhitelistInvest(&_Launchpad.TransactOpts)
}

// WhitelistInvest is a paid mutator transaction binding the contract method 0xdcfe6cb6.
//
// Solidity: function whitelistInvest() payable returns()
func (_Launchpad *LaunchpadTransactorSession) WhitelistInvest() (*types.Transaction, error) {
	return _Launchpad.Contract.WhitelistInvest(&_Launchpad.TransactOpts)
}

// LaunchpadCapRatioUpdatedIterator is returned from FilterCapRatioUpdated and is used to iterate over the raw logs and unpacked data for CapRatioUpdated events raised by the Launchpad contract.
type LaunchpadCapRatioUpdatedIterator struct {
	Event *LaunchpadCapRatioUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadCapRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadCapRatioUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadCapRatioUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadCapRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadCapRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadCapRatioUpdated represents a CapRatioUpdated event raised by the Launchpad contract.
type LaunchpadCapRatioUpdated struct {
	CapRatio *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCapRatioUpdated is a free log retrieval operation binding the contract event 0x45b44375c75c6ed136375d5182c1ad54f7a391afb5f8e2d6b5bc9007e3ae5a47.
//
// Solidity: event CapRatioUpdated(uint256 capRatio)
func (_Launchpad *LaunchpadFilterer) FilterCapRatioUpdated(opts *bind.FilterOpts) (*LaunchpadCapRatioUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "CapRatioUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadCapRatioUpdatedIterator{contract: _Launchpad.contract, event: "CapRatioUpdated", logs: logs, sub: sub}, nil
}

// WatchCapRatioUpdated is a free log subscription operation binding the contract event 0x45b44375c75c6ed136375d5182c1ad54f7a391afb5f8e2d6b5bc9007e3ae5a47.
//
// Solidity: event CapRatioUpdated(uint256 capRatio)
func (_Launchpad *LaunchpadFilterer) WatchCapRatioUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadCapRatioUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "CapRatioUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadCapRatioUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "CapRatioUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCapRatioUpdated is a log parse operation binding the contract event 0x45b44375c75c6ed136375d5182c1ad54f7a391afb5f8e2d6b5bc9007e3ae5a47.
//
// Solidity: event CapRatioUpdated(uint256 capRatio)
func (_Launchpad *LaunchpadFilterer) ParseCapRatioUpdated(log types.Log) (*LaunchpadCapRatioUpdated, error) {
	event := new(LaunchpadCapRatioUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "CapRatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadClaimableSetIterator is returned from FilterClaimableSet and is used to iterate over the raw logs and unpacked data for ClaimableSet events raised by the Launchpad contract.
type LaunchpadClaimableSetIterator struct {
	Event *LaunchpadClaimableSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadClaimableSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadClaimableSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadClaimableSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadClaimableSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadClaimableSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadClaimableSet represents a ClaimableSet event raised by the Launchpad contract.
type LaunchpadClaimableSet struct {
	Claimable bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimableSet is a free log retrieval operation binding the contract event 0x92b7925d83a5ab241e205db556ec406c8f63c33f6025d40e34a4e213ef7456e9.
//
// Solidity: event ClaimableSet(bool claimable)
func (_Launchpad *LaunchpadFilterer) FilterClaimableSet(opts *bind.FilterOpts) (*LaunchpadClaimableSetIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "ClaimableSet")
	if err != nil {
		return nil, err
	}
	return &LaunchpadClaimableSetIterator{contract: _Launchpad.contract, event: "ClaimableSet", logs: logs, sub: sub}, nil
}

// WatchClaimableSet is a free log subscription operation binding the contract event 0x92b7925d83a5ab241e205db556ec406c8f63c33f6025d40e34a4e213ef7456e9.
//
// Solidity: event ClaimableSet(bool claimable)
func (_Launchpad *LaunchpadFilterer) WatchClaimableSet(opts *bind.WatchOpts, sink chan<- *LaunchpadClaimableSet) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "ClaimableSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadClaimableSet)
				if err := _Launchpad.contract.UnpackLog(event, "ClaimableSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimableSet is a log parse operation binding the contract event 0x92b7925d83a5ab241e205db556ec406c8f63c33f6025d40e34a4e213ef7456e9.
//
// Solidity: event ClaimableSet(bool claimable)
func (_Launchpad *LaunchpadFilterer) ParseClaimableSet(log types.Log) (*LaunchpadClaimableSet, error) {
	event := new(LaunchpadClaimableSet)
	if err := _Launchpad.contract.UnpackLog(event, "ClaimableSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Launchpad contract.
type LaunchpadClaimedIterator struct {
	Event *LaunchpadClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadClaimed represents a Claimed event raised by the Launchpad contract.
type LaunchpadClaimed struct {
	InvestorAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed investorAddress, uint256 amount)
func (_Launchpad *LaunchpadFilterer) FilterClaimed(opts *bind.FilterOpts, investorAddress []common.Address) (*LaunchpadClaimedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "Claimed", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadClaimedIterator{contract: _Launchpad.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed investorAddress, uint256 amount)
func (_Launchpad *LaunchpadFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *LaunchpadClaimed, investorAddress []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "Claimed", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadClaimed)
				if err := _Launchpad.contract.UnpackLog(event, "Claimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed investorAddress, uint256 amount)
func (_Launchpad *LaunchpadFilterer) ParseClaimed(log types.Log) (*LaunchpadClaimed, error) {
	event := new(LaunchpadClaimed)
	if err := _Launchpad.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadFundsDistributedIterator is returned from FilterFundsDistributed and is used to iterate over the raw logs and unpacked data for FundsDistributed events raised by the Launchpad contract.
type LaunchpadFundsDistributedIterator struct {
	Event *LaunchpadFundsDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadFundsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadFundsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadFundsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadFundsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadFundsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadFundsDistributed represents a FundsDistributed event raised by the Launchpad contract.
type LaunchpadFundsDistributed struct {
	TeamWallet      common.Address
	LiquidityWallet common.Address
	PlatformWallet  common.Address
	TeamAmount      *big.Int
	LiquidityAmount *big.Int
	PlatformAmount  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundsDistributed is a free log retrieval operation binding the contract event 0x4a88bd679da4aaee6664a8e505eb8eaa53f983a1c99740536e734e6e31dc08d8.
//
// Solidity: event FundsDistributed(address indexed teamWallet, address indexed liquidityWallet, address indexed platformWallet, uint256 teamAmount, uint256 liquidityAmount, uint256 platformAmount)
func (_Launchpad *LaunchpadFilterer) FilterFundsDistributed(opts *bind.FilterOpts, teamWallet []common.Address, liquidityWallet []common.Address, platformWallet []common.Address) (*LaunchpadFundsDistributedIterator, error) {

	var teamWalletRule []interface{}
	for _, teamWalletItem := range teamWallet {
		teamWalletRule = append(teamWalletRule, teamWalletItem)
	}
	var liquidityWalletRule []interface{}
	for _, liquidityWalletItem := range liquidityWallet {
		liquidityWalletRule = append(liquidityWalletRule, liquidityWalletItem)
	}
	var platformWalletRule []interface{}
	for _, platformWalletItem := range platformWallet {
		platformWalletRule = append(platformWalletRule, platformWalletItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "FundsDistributed", teamWalletRule, liquidityWalletRule, platformWalletRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadFundsDistributedIterator{contract: _Launchpad.contract, event: "FundsDistributed", logs: logs, sub: sub}, nil
}

// WatchFundsDistributed is a free log subscription operation binding the contract event 0x4a88bd679da4aaee6664a8e505eb8eaa53f983a1c99740536e734e6e31dc08d8.
//
// Solidity: event FundsDistributed(address indexed teamWallet, address indexed liquidityWallet, address indexed platformWallet, uint256 teamAmount, uint256 liquidityAmount, uint256 platformAmount)
func (_Launchpad *LaunchpadFilterer) WatchFundsDistributed(opts *bind.WatchOpts, sink chan<- *LaunchpadFundsDistributed, teamWallet []common.Address, liquidityWallet []common.Address, platformWallet []common.Address) (event.Subscription, error) {

	var teamWalletRule []interface{}
	for _, teamWalletItem := range teamWallet {
		teamWalletRule = append(teamWalletRule, teamWalletItem)
	}
	var liquidityWalletRule []interface{}
	for _, liquidityWalletItem := range liquidityWallet {
		liquidityWalletRule = append(liquidityWalletRule, liquidityWalletItem)
	}
	var platformWalletRule []interface{}
	for _, platformWalletItem := range platformWallet {
		platformWalletRule = append(platformWalletRule, platformWalletItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "FundsDistributed", teamWalletRule, liquidityWalletRule, platformWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadFundsDistributed)
				if err := _Launchpad.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDistributed is a log parse operation binding the contract event 0x4a88bd679da4aaee6664a8e505eb8eaa53f983a1c99740536e734e6e31dc08d8.
//
// Solidity: event FundsDistributed(address indexed teamWallet, address indexed liquidityWallet, address indexed platformWallet, uint256 teamAmount, uint256 liquidityAmount, uint256 platformAmount)
func (_Launchpad *LaunchpadFilterer) ParseFundsDistributed(log types.Log) (*LaunchpadFundsDistributed, error) {
	event := new(LaunchpadFundsDistributed)
	if err := _Launchpad.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadInvestedIterator is returned from FilterInvested and is used to iterate over the raw logs and unpacked data for Invested events raised by the Launchpad contract.
type LaunchpadInvestedIterator struct {
	Event *LaunchpadInvested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadInvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadInvested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadInvested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadInvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadInvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadInvested represents a Invested event raised by the Launchpad contract.
type LaunchpadInvested struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInvested is a free log retrieval operation binding the contract event 0xc3f75dfc78f6efac88ad5abb5e606276b903647d97b2a62a1ef89840a658bbc3.
//
// Solidity: event Invested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) FilterInvested(opts *bind.FilterOpts, investorAddress []common.Address) (*LaunchpadInvestedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "Invested", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadInvestedIterator{contract: _Launchpad.contract, event: "Invested", logs: logs, sub: sub}, nil
}

// WatchInvested is a free log subscription operation binding the contract event 0xc3f75dfc78f6efac88ad5abb5e606276b903647d97b2a62a1ef89840a658bbc3.
//
// Solidity: event Invested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) WatchInvested(opts *bind.WatchOpts, sink chan<- *LaunchpadInvested, investorAddress []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "Invested", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadInvested)
				if err := _Launchpad.contract.UnpackLog(event, "Invested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvested is a log parse operation binding the contract event 0xc3f75dfc78f6efac88ad5abb5e606276b903647d97b2a62a1ef89840a658bbc3.
//
// Solidity: event Invested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) ParseInvested(log types.Log) (*LaunchpadInvested, error) {
	event := new(LaunchpadInvested)
	if err := _Launchpad.contract.UnpackLog(event, "Invested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadInvestmentLimitsUpdatedIterator is returned from FilterInvestmentLimitsUpdated and is used to iterate over the raw logs and unpacked data for InvestmentLimitsUpdated events raised by the Launchpad contract.
type LaunchpadInvestmentLimitsUpdatedIterator struct {
	Event *LaunchpadInvestmentLimitsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadInvestmentLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadInvestmentLimitsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadInvestmentLimitsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadInvestmentLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadInvestmentLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadInvestmentLimitsUpdated represents a InvestmentLimitsUpdated event raised by the Launchpad contract.
type LaunchpadInvestmentLimitsUpdated struct {
	MinInvestment *big.Int
	MaxInvestment *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvestmentLimitsUpdated is a free log retrieval operation binding the contract event 0x4c14c3ff8e0ac5cb918f60f7ce707bc2c0eae10e69c8df94e19c904400309894.
//
// Solidity: event InvestmentLimitsUpdated(uint256 minInvestment, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) FilterInvestmentLimitsUpdated(opts *bind.FilterOpts) (*LaunchpadInvestmentLimitsUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "InvestmentLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadInvestmentLimitsUpdatedIterator{contract: _Launchpad.contract, event: "InvestmentLimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchInvestmentLimitsUpdated is a free log subscription operation binding the contract event 0x4c14c3ff8e0ac5cb918f60f7ce707bc2c0eae10e69c8df94e19c904400309894.
//
// Solidity: event InvestmentLimitsUpdated(uint256 minInvestment, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) WatchInvestmentLimitsUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadInvestmentLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "InvestmentLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadInvestmentLimitsUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "InvestmentLimitsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvestmentLimitsUpdated is a log parse operation binding the contract event 0x4c14c3ff8e0ac5cb918f60f7ce707bc2c0eae10e69c8df94e19c904400309894.
//
// Solidity: event InvestmentLimitsUpdated(uint256 minInvestment, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) ParseInvestmentLimitsUpdated(log types.Log) (*LaunchpadInvestmentLimitsUpdated, error) {
	event := new(LaunchpadInvestmentLimitsUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "InvestmentLimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadLiquidityWalletUpdatedIterator is returned from FilterLiquidityWalletUpdated and is used to iterate over the raw logs and unpacked data for LiquidityWalletUpdated events raised by the Launchpad contract.
type LaunchpadLiquidityWalletUpdatedIterator struct {
	Event *LaunchpadLiquidityWalletUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadLiquidityWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadLiquidityWalletUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadLiquidityWalletUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadLiquidityWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadLiquidityWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadLiquidityWalletUpdated represents a LiquidityWalletUpdated event raised by the Launchpad contract.
type LaunchpadLiquidityWalletUpdated struct {
	LiquidityWallet common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWalletUpdated is a free log retrieval operation binding the contract event 0xc0f8ad2683cf240266d4f868f5b59a16c13e4a0ff63dfcfb0f0cd3843a121784.
//
// Solidity: event LiquidityWalletUpdated(address indexed liquidityWallet)
func (_Launchpad *LaunchpadFilterer) FilterLiquidityWalletUpdated(opts *bind.FilterOpts, liquidityWallet []common.Address) (*LaunchpadLiquidityWalletUpdatedIterator, error) {

	var liquidityWalletRule []interface{}
	for _, liquidityWalletItem := range liquidityWallet {
		liquidityWalletRule = append(liquidityWalletRule, liquidityWalletItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "LiquidityWalletUpdated", liquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadLiquidityWalletUpdatedIterator{contract: _Launchpad.contract, event: "LiquidityWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityWalletUpdated is a free log subscription operation binding the contract event 0xc0f8ad2683cf240266d4f868f5b59a16c13e4a0ff63dfcfb0f0cd3843a121784.
//
// Solidity: event LiquidityWalletUpdated(address indexed liquidityWallet)
func (_Launchpad *LaunchpadFilterer) WatchLiquidityWalletUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadLiquidityWalletUpdated, liquidityWallet []common.Address) (event.Subscription, error) {

	var liquidityWalletRule []interface{}
	for _, liquidityWalletItem := range liquidityWallet {
		liquidityWalletRule = append(liquidityWalletRule, liquidityWalletItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "LiquidityWalletUpdated", liquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadLiquidityWalletUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityWalletUpdated is a log parse operation binding the contract event 0xc0f8ad2683cf240266d4f868f5b59a16c13e4a0ff63dfcfb0f0cd3843a121784.
//
// Solidity: event LiquidityWalletUpdated(address indexed liquidityWallet)
func (_Launchpad *LaunchpadFilterer) ParseLiquidityWalletUpdated(log types.Log) (*LaunchpadLiquidityWalletUpdated, error) {
	event := new(LaunchpadLiquidityWalletUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadPublicPriceUpdatedIterator is returned from FilterPublicPriceUpdated and is used to iterate over the raw logs and unpacked data for PublicPriceUpdated events raised by the Launchpad contract.
type LaunchpadPublicPriceUpdatedIterator struct {
	Event *LaunchpadPublicPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadPublicPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadPublicPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadPublicPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadPublicPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadPublicPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadPublicPriceUpdated represents a PublicPriceUpdated event raised by the Launchpad contract.
type LaunchpadPublicPriceUpdated struct {
	PublicPrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicPriceUpdated is a free log retrieval operation binding the contract event 0xb6929b692b182f5174d872e8742af072aa2786771d49d1d2c4bf1f19921b9b16.
//
// Solidity: event PublicPriceUpdated(uint256 publicPrice)
func (_Launchpad *LaunchpadFilterer) FilterPublicPriceUpdated(opts *bind.FilterOpts) (*LaunchpadPublicPriceUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "PublicPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadPublicPriceUpdatedIterator{contract: _Launchpad.contract, event: "PublicPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicPriceUpdated is a free log subscription operation binding the contract event 0xb6929b692b182f5174d872e8742af072aa2786771d49d1d2c4bf1f19921b9b16.
//
// Solidity: event PublicPriceUpdated(uint256 publicPrice)
func (_Launchpad *LaunchpadFilterer) WatchPublicPriceUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadPublicPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "PublicPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadPublicPriceUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "PublicPriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicPriceUpdated is a log parse operation binding the contract event 0xb6929b692b182f5174d872e8742af072aa2786771d49d1d2c4bf1f19921b9b16.
//
// Solidity: event PublicPriceUpdated(uint256 publicPrice)
func (_Launchpad *LaunchpadFilterer) ParsePublicPriceUpdated(log types.Log) (*LaunchpadPublicPriceUpdated, error) {
	event := new(LaunchpadPublicPriceUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "PublicPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadPublicSaleTimesUpdatedIterator is returned from FilterPublicSaleTimesUpdated and is used to iterate over the raw logs and unpacked data for PublicSaleTimesUpdated events raised by the Launchpad contract.
type LaunchpadPublicSaleTimesUpdatedIterator struct {
	Event *LaunchpadPublicSaleTimesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadPublicSaleTimesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadPublicSaleTimesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadPublicSaleTimesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadPublicSaleTimesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadPublicSaleTimesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadPublicSaleTimesUpdated represents a PublicSaleTimesUpdated event raised by the Launchpad contract.
type LaunchpadPublicSaleTimesUpdated struct {
	PublicStartTime *big.Int
	PublicEndTime   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicSaleTimesUpdated is a free log retrieval operation binding the contract event 0x3b98dc731102f7200b8622807fde6953cbcc4b216aa165a7ae6383b594913bba.
//
// Solidity: event PublicSaleTimesUpdated(uint256 publicStartTime, uint256 publicEndTime)
func (_Launchpad *LaunchpadFilterer) FilterPublicSaleTimesUpdated(opts *bind.FilterOpts) (*LaunchpadPublicSaleTimesUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "PublicSaleTimesUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadPublicSaleTimesUpdatedIterator{contract: _Launchpad.contract, event: "PublicSaleTimesUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicSaleTimesUpdated is a free log subscription operation binding the contract event 0x3b98dc731102f7200b8622807fde6953cbcc4b216aa165a7ae6383b594913bba.
//
// Solidity: event PublicSaleTimesUpdated(uint256 publicStartTime, uint256 publicEndTime)
func (_Launchpad *LaunchpadFilterer) WatchPublicSaleTimesUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadPublicSaleTimesUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "PublicSaleTimesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadPublicSaleTimesUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "PublicSaleTimesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicSaleTimesUpdated is a log parse operation binding the contract event 0x3b98dc731102f7200b8622807fde6953cbcc4b216aa165a7ae6383b594913bba.
//
// Solidity: event PublicSaleTimesUpdated(uint256 publicStartTime, uint256 publicEndTime)
func (_Launchpad *LaunchpadFilterer) ParsePublicSaleTimesUpdated(log types.Log) (*LaunchpadPublicSaleTimesUpdated, error) {
	event := new(LaunchpadPublicSaleTimesUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "PublicSaleTimesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRaiseWalletUpdatedIterator is returned from FilterRaiseWalletUpdated and is used to iterate over the raw logs and unpacked data for RaiseWalletUpdated events raised by the Launchpad contract.
type LaunchpadRaiseWalletUpdatedIterator struct {
	Event *LaunchpadRaiseWalletUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRaiseWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRaiseWalletUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRaiseWalletUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRaiseWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRaiseWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRaiseWalletUpdated represents a RaiseWalletUpdated event raised by the Launchpad contract.
type LaunchpadRaiseWalletUpdated struct {
	RaiseWallet common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaiseWalletUpdated is a free log retrieval operation binding the contract event 0xce924bc6fca022e9db5ae717a726b0a3324aed37582cc9642efb37da499b8257.
//
// Solidity: event RaiseWalletUpdated(address indexed raiseWallet)
func (_Launchpad *LaunchpadFilterer) FilterRaiseWalletUpdated(opts *bind.FilterOpts, raiseWallet []common.Address) (*LaunchpadRaiseWalletUpdatedIterator, error) {

	var raiseWalletRule []interface{}
	for _, raiseWalletItem := range raiseWallet {
		raiseWalletRule = append(raiseWalletRule, raiseWalletItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RaiseWalletUpdated", raiseWalletRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRaiseWalletUpdatedIterator{contract: _Launchpad.contract, event: "RaiseWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchRaiseWalletUpdated is a free log subscription operation binding the contract event 0xce924bc6fca022e9db5ae717a726b0a3324aed37582cc9642efb37da499b8257.
//
// Solidity: event RaiseWalletUpdated(address indexed raiseWallet)
func (_Launchpad *LaunchpadFilterer) WatchRaiseWalletUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadRaiseWalletUpdated, raiseWallet []common.Address) (event.Subscription, error) {

	var raiseWalletRule []interface{}
	for _, raiseWalletItem := range raiseWallet {
		raiseWalletRule = append(raiseWalletRule, raiseWalletItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RaiseWalletUpdated", raiseWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRaiseWalletUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "RaiseWalletUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRaiseWalletUpdated is a log parse operation binding the contract event 0xce924bc6fca022e9db5ae717a726b0a3324aed37582cc9642efb37da499b8257.
//
// Solidity: event RaiseWalletUpdated(address indexed raiseWallet)
func (_Launchpad *LaunchpadFilterer) ParseRaiseWalletUpdated(log types.Log) (*LaunchpadRaiseWalletUpdated, error) {
	event := new(LaunchpadRaiseWalletUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "RaiseWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRateUpdatedIterator is returned from FilterRateUpdated and is used to iterate over the raw logs and unpacked data for RateUpdated events raised by the Launchpad contract.
type LaunchpadRateUpdatedIterator struct {
	Event *LaunchpadRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRateUpdated represents a RateUpdated event raised by the Launchpad contract.
type LaunchpadRateUpdated struct {
	RaiseRate     *big.Int
	LiquidityRate *big.Int
	PlatformRate  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateUpdated is a free log retrieval operation binding the contract event 0xa387c3601f61e88a17e341cfebf7ab826a8cd4544cd72b2dc0cf12988e77754b.
//
// Solidity: event RateUpdated(uint256 raiseRate, uint256 liquidityRate, uint256 platformRate)
func (_Launchpad *LaunchpadFilterer) FilterRateUpdated(opts *bind.FilterOpts) (*LaunchpadRateUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadRateUpdatedIterator{contract: _Launchpad.contract, event: "RateUpdated", logs: logs, sub: sub}, nil
}

// WatchRateUpdated is a free log subscription operation binding the contract event 0xa387c3601f61e88a17e341cfebf7ab826a8cd4544cd72b2dc0cf12988e77754b.
//
// Solidity: event RateUpdated(uint256 raiseRate, uint256 liquidityRate, uint256 platformRate)
func (_Launchpad *LaunchpadFilterer) WatchRateUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRateUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "RateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRateUpdated is a log parse operation binding the contract event 0xa387c3601f61e88a17e341cfebf7ab826a8cd4544cd72b2dc0cf12988e77754b.
//
// Solidity: event RateUpdated(uint256 raiseRate, uint256 liquidityRate, uint256 platformRate)
func (_Launchpad *LaunchpadFilterer) ParseRateUpdated(log types.Log) (*LaunchpadRateUpdated, error) {
	event := new(LaunchpadRateUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "RateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Launchpad contract.
type LaunchpadRefundedIterator struct {
	Event *LaunchpadRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRefunded represents a Refunded event raised by the Launchpad contract.
type LaunchpadRefunded struct {
	InvestorAddress common.Address
	AmountEth       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed investorAddress, uint256 amountEth)
func (_Launchpad *LaunchpadFilterer) FilterRefunded(opts *bind.FilterOpts, investorAddress []common.Address) (*LaunchpadRefundedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "Refunded", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRefundedIterator{contract: _Launchpad.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed investorAddress, uint256 amountEth)
func (_Launchpad *LaunchpadFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *LaunchpadRefunded, investorAddress []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "Refunded", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRefunded)
				if err := _Launchpad.contract.UnpackLog(event, "Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefunded is a log parse operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed investorAddress, uint256 amountEth)
func (_Launchpad *LaunchpadFilterer) ParseRefunded(log types.Log) (*LaunchpadRefunded, error) {
	event := new(LaunchpadRefunded)
	if err := _Launchpad.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRemainingTokensTransferredIterator is returned from FilterRemainingTokensTransferred and is used to iterate over the raw logs and unpacked data for RemainingTokensTransferred events raised by the Launchpad contract.
type LaunchpadRemainingTokensTransferredIterator struct {
	Event *LaunchpadRemainingTokensTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRemainingTokensTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRemainingTokensTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRemainingTokensTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRemainingTokensTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRemainingTokensTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRemainingTokensTransferred represents a RemainingTokensTransferred event raised by the Launchpad contract.
type LaunchpadRemainingTokensTransferred struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemainingTokensTransferred is a free log retrieval operation binding the contract event 0xc3c490623455bc2969f780b6ae0cab230dd6fa5139df78381363816a5266842c.
//
// Solidity: event RemainingTokensTransferred(address indexed to, uint256 amount)
func (_Launchpad *LaunchpadFilterer) FilterRemainingTokensTransferred(opts *bind.FilterOpts, to []common.Address) (*LaunchpadRemainingTokensTransferredIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RemainingTokensTransferred", toRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRemainingTokensTransferredIterator{contract: _Launchpad.contract, event: "RemainingTokensTransferred", logs: logs, sub: sub}, nil
}

// WatchRemainingTokensTransferred is a free log subscription operation binding the contract event 0xc3c490623455bc2969f780b6ae0cab230dd6fa5139df78381363816a5266842c.
//
// Solidity: event RemainingTokensTransferred(address indexed to, uint256 amount)
func (_Launchpad *LaunchpadFilterer) WatchRemainingTokensTransferred(opts *bind.WatchOpts, sink chan<- *LaunchpadRemainingTokensTransferred, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RemainingTokensTransferred", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRemainingTokensTransferred)
				if err := _Launchpad.contract.UnpackLog(event, "RemainingTokensTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemainingTokensTransferred is a log parse operation binding the contract event 0xc3c490623455bc2969f780b6ae0cab230dd6fa5139df78381363816a5266842c.
//
// Solidity: event RemainingTokensTransferred(address indexed to, uint256 amount)
func (_Launchpad *LaunchpadFilterer) ParseRemainingTokensTransferred(log types.Log) (*LaunchpadRemainingTokensTransferred, error) {
	event := new(LaunchpadRemainingTokensTransferred)
	if err := _Launchpad.contract.UnpackLog(event, "RemainingTokensTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the Launchpad contract.
type LaunchpadRemovedFromWhitelistIterator struct {
	Event *LaunchpadRemovedFromWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRemovedFromWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRemovedFromWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the Launchpad contract.
type LaunchpadRemovedFromWhitelist struct {
	InvestorAddresses []common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x17dbdef66bfe28946fe1bf49cc6946533e58c6e0ecbf907152e406b630eb4f27.
//
// Solidity: event RemovedFromWhitelist(address[] indexed investorAddresses)
func (_Launchpad *LaunchpadFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts, investorAddresses [][]common.Address) (*LaunchpadRemovedFromWhitelistIterator, error) {

	var investorAddressesRule []interface{}
	for _, investorAddressesItem := range investorAddresses {
		investorAddressesRule = append(investorAddressesRule, investorAddressesItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RemovedFromWhitelist", investorAddressesRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRemovedFromWhitelistIterator{contract: _Launchpad.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0x17dbdef66bfe28946fe1bf49cc6946533e58c6e0ecbf907152e406b630eb4f27.
//
// Solidity: event RemovedFromWhitelist(address[] indexed investorAddresses)
func (_Launchpad *LaunchpadFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *LaunchpadRemovedFromWhitelist, investorAddresses [][]common.Address) (event.Subscription, error) {

	var investorAddressesRule []interface{}
	for _, investorAddressesItem := range investorAddresses {
		investorAddressesRule = append(investorAddressesRule, investorAddressesItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RemovedFromWhitelist", investorAddressesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRemovedFromWhitelist)
				if err := _Launchpad.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedFromWhitelist is a log parse operation binding the contract event 0x17dbdef66bfe28946fe1bf49cc6946533e58c6e0ecbf907152e406b630eb4f27.
//
// Solidity: event RemovedFromWhitelist(address[] indexed investorAddresses)
func (_Launchpad *LaunchpadFilterer) ParseRemovedFromWhitelist(log types.Log) (*LaunchpadRemovedFromWhitelist, error) {
	event := new(LaunchpadRemovedFromWhitelist)
	if err := _Launchpad.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Launchpad contract.
type LaunchpadRoleAdminChangedIterator struct {
	Event *LaunchpadRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRoleAdminChanged represents a RoleAdminChanged event raised by the Launchpad contract.
type LaunchpadRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Launchpad *LaunchpadFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LaunchpadRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRoleAdminChangedIterator{contract: _Launchpad.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Launchpad *LaunchpadFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LaunchpadRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRoleAdminChanged)
				if err := _Launchpad.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Launchpad *LaunchpadFilterer) ParseRoleAdminChanged(log types.Log) (*LaunchpadRoleAdminChanged, error) {
	event := new(LaunchpadRoleAdminChanged)
	if err := _Launchpad.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Launchpad contract.
type LaunchpadRoleGrantedIterator struct {
	Event *LaunchpadRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRoleGranted represents a RoleGranted event raised by the Launchpad contract.
type LaunchpadRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LaunchpadRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRoleGrantedIterator{contract: _Launchpad.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LaunchpadRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRoleGranted)
				if err := _Launchpad.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) ParseRoleGranted(log types.Log) (*LaunchpadRoleGranted, error) {
	event := new(LaunchpadRoleGranted)
	if err := _Launchpad.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Launchpad contract.
type LaunchpadRoleRevokedIterator struct {
	Event *LaunchpadRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadRoleRevoked represents a RoleRevoked event raised by the Launchpad contract.
type LaunchpadRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LaunchpadRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadRoleRevokedIterator{contract: _Launchpad.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LaunchpadRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadRoleRevoked)
				if err := _Launchpad.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Launchpad *LaunchpadFilterer) ParseRoleRevoked(log types.Log) (*LaunchpadRoleRevoked, error) {
	event := new(LaunchpadRoleRevoked)
	if err := _Launchpad.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadTotalPublicSaleUpdatedIterator is returned from FilterTotalPublicSaleUpdated and is used to iterate over the raw logs and unpacked data for TotalPublicSaleUpdated events raised by the Launchpad contract.
type LaunchpadTotalPublicSaleUpdatedIterator struct {
	Event *LaunchpadTotalPublicSaleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadTotalPublicSaleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadTotalPublicSaleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadTotalPublicSaleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadTotalPublicSaleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadTotalPublicSaleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadTotalPublicSaleUpdated represents a TotalPublicSaleUpdated event raised by the Launchpad contract.
type LaunchpadTotalPublicSaleUpdated struct {
	OldTotalPublicSale *big.Int
	NewTotalPublicSale *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTotalPublicSaleUpdated is a free log retrieval operation binding the contract event 0xbec5698ba60c77d292152f4abe5a9401ac081a3f216eb3aa73f1f8ec508c6ece.
//
// Solidity: event TotalPublicSaleUpdated(uint256 oldTotalPublicSale, uint256 newTotalPublicSale)
func (_Launchpad *LaunchpadFilterer) FilterTotalPublicSaleUpdated(opts *bind.FilterOpts) (*LaunchpadTotalPublicSaleUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "TotalPublicSaleUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadTotalPublicSaleUpdatedIterator{contract: _Launchpad.contract, event: "TotalPublicSaleUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalPublicSaleUpdated is a free log subscription operation binding the contract event 0xbec5698ba60c77d292152f4abe5a9401ac081a3f216eb3aa73f1f8ec508c6ece.
//
// Solidity: event TotalPublicSaleUpdated(uint256 oldTotalPublicSale, uint256 newTotalPublicSale)
func (_Launchpad *LaunchpadFilterer) WatchTotalPublicSaleUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadTotalPublicSaleUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "TotalPublicSaleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadTotalPublicSaleUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "TotalPublicSaleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTotalPublicSaleUpdated is a log parse operation binding the contract event 0xbec5698ba60c77d292152f4abe5a9401ac081a3f216eb3aa73f1f8ec508c6ece.
//
// Solidity: event TotalPublicSaleUpdated(uint256 oldTotalPublicSale, uint256 newTotalPublicSale)
func (_Launchpad *LaunchpadFilterer) ParseTotalPublicSaleUpdated(log types.Log) (*LaunchpadTotalPublicSaleUpdated, error) {
	event := new(LaunchpadTotalPublicSaleUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "TotalPublicSaleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadWhitelistInvestedIterator is returned from FilterWhitelistInvested and is used to iterate over the raw logs and unpacked data for WhitelistInvested events raised by the Launchpad contract.
type LaunchpadWhitelistInvestedIterator struct {
	Event *LaunchpadWhitelistInvested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadWhitelistInvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadWhitelistInvested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadWhitelistInvested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadWhitelistInvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadWhitelistInvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadWhitelistInvested represents a WhitelistInvested event raised by the Launchpad contract.
type LaunchpadWhitelistInvested struct {
	InvestorAddress  common.Address
	InvestmentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWhitelistInvested is a free log retrieval operation binding the contract event 0x2ca2384c617adf7f1394ce760ec338019cdaf29234245327ceeaf51e27125bb6.
//
// Solidity: event WhitelistInvested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) FilterWhitelistInvested(opts *bind.FilterOpts, investorAddress []common.Address) (*LaunchpadWhitelistInvestedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "WhitelistInvested", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadWhitelistInvestedIterator{contract: _Launchpad.contract, event: "WhitelistInvested", logs: logs, sub: sub}, nil
}

// WatchWhitelistInvested is a free log subscription operation binding the contract event 0x2ca2384c617adf7f1394ce760ec338019cdaf29234245327ceeaf51e27125bb6.
//
// Solidity: event WhitelistInvested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) WatchWhitelistInvested(opts *bind.WatchOpts, sink chan<- *LaunchpadWhitelistInvested, investorAddress []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "WhitelistInvested", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadWhitelistInvested)
				if err := _Launchpad.contract.UnpackLog(event, "WhitelistInvested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistInvested is a log parse operation binding the contract event 0x2ca2384c617adf7f1394ce760ec338019cdaf29234245327ceeaf51e27125bb6.
//
// Solidity: event WhitelistInvested(address indexed investorAddress, uint256 investmentAmount)
func (_Launchpad *LaunchpadFilterer) ParseWhitelistInvested(log types.Log) (*LaunchpadWhitelistInvested, error) {
	event := new(LaunchpadWhitelistInvested)
	if err := _Launchpad.contract.UnpackLog(event, "WhitelistInvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadWhitelistParametersUpdatedIterator is returned from FilterWhitelistParametersUpdated and is used to iterate over the raw logs and unpacked data for WhitelistParametersUpdated events raised by the Launchpad contract.
type LaunchpadWhitelistParametersUpdatedIterator struct {
	Event *LaunchpadWhitelistParametersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadWhitelistParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadWhitelistParametersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadWhitelistParametersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadWhitelistParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadWhitelistParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadWhitelistParametersUpdated represents a WhitelistParametersUpdated event raised by the Launchpad contract.
type LaunchpadWhitelistParametersUpdated struct {
	WhitelistStartTime *big.Int
	WhitelistEndTime   *big.Int
	WhitelistPrice     *big.Int
	TotalWhitelistSale *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWhitelistParametersUpdated is a free log retrieval operation binding the contract event 0x57ad337a69a1b1b4921c921b25cab99991e0ebaba5915b0143e2df05d17e6596.
//
// Solidity: event WhitelistParametersUpdated(uint256 whitelistStartTime, uint256 whitelistEndTime, uint256 whitelistPrice, uint256 totalWhitelistSale)
func (_Launchpad *LaunchpadFilterer) FilterWhitelistParametersUpdated(opts *bind.FilterOpts) (*LaunchpadWhitelistParametersUpdatedIterator, error) {

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "WhitelistParametersUpdated")
	if err != nil {
		return nil, err
	}
	return &LaunchpadWhitelistParametersUpdatedIterator{contract: _Launchpad.contract, event: "WhitelistParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistParametersUpdated is a free log subscription operation binding the contract event 0x57ad337a69a1b1b4921c921b25cab99991e0ebaba5915b0143e2df05d17e6596.
//
// Solidity: event WhitelistParametersUpdated(uint256 whitelistStartTime, uint256 whitelistEndTime, uint256 whitelistPrice, uint256 totalWhitelistSale)
func (_Launchpad *LaunchpadFilterer) WatchWhitelistParametersUpdated(opts *bind.WatchOpts, sink chan<- *LaunchpadWhitelistParametersUpdated) (event.Subscription, error) {

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "WhitelistParametersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadWhitelistParametersUpdated)
				if err := _Launchpad.contract.UnpackLog(event, "WhitelistParametersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistParametersUpdated is a log parse operation binding the contract event 0x57ad337a69a1b1b4921c921b25cab99991e0ebaba5915b0143e2df05d17e6596.
//
// Solidity: event WhitelistParametersUpdated(uint256 whitelistStartTime, uint256 whitelistEndTime, uint256 whitelistPrice, uint256 totalWhitelistSale)
func (_Launchpad *LaunchpadFilterer) ParseWhitelistParametersUpdated(log types.Log) (*LaunchpadWhitelistParametersUpdated, error) {
	event := new(LaunchpadWhitelistParametersUpdated)
	if err := _Launchpad.contract.UnpackLog(event, "WhitelistParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchpadWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the Launchpad contract.
type LaunchpadWhitelistedIterator struct {
	Event *LaunchpadWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LaunchpadWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchpadWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LaunchpadWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LaunchpadWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchpadWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchpadWhitelisted represents a Whitelisted event raised by the Launchpad contract.
type LaunchpadWhitelisted struct {
	InvestorAddress common.Address
	MaxInvestment   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0x6ea640312e182de387819fbeb13be00db3171a445412852248559054871c4199.
//
// Solidity: event Whitelisted(address indexed investorAddress, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) FilterWhitelisted(opts *bind.FilterOpts, investorAddress []common.Address) (*LaunchpadWhitelistedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.FilterLogs(opts, "Whitelisted", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return &LaunchpadWhitelistedIterator{contract: _Launchpad.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0x6ea640312e182de387819fbeb13be00db3171a445412852248559054871c4199.
//
// Solidity: event Whitelisted(address indexed investorAddress, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *LaunchpadWhitelisted, investorAddress []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}

	logs, sub, err := _Launchpad.contract.WatchLogs(opts, "Whitelisted", investorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchpadWhitelisted)
				if err := _Launchpad.contract.UnpackLog(event, "Whitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelisted is a log parse operation binding the contract event 0x6ea640312e182de387819fbeb13be00db3171a445412852248559054871c4199.
//
// Solidity: event Whitelisted(address indexed investorAddress, uint256 maxInvestment)
func (_Launchpad *LaunchpadFilterer) ParseWhitelisted(log types.Log) (*LaunchpadWhitelisted, error) {
	event := new(LaunchpadWhitelisted)
	if err := _Launchpad.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
