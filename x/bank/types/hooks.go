package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _BankHooks=&MultiBankHooks{}

type MultiBankHooks []BankHooks

func NewMultiBankHooks(hooks ...BankHooks) MultiBankHooks{
	return hooks
}

func (h MultiBankHooks) BeforeSendCoins(ctx sdk.Context,fromAddr sdk.AccAddress,toAddr sdk.AccAddress,amt sdk.Coins)error{
	for i:=range h{
		if err:=h[i].BeforeSendCoins(ctx,fromAddr,toAddr,amt);err!=nil{
			return err
		}
	}
	return nil
}

func (h MultiBankHooks) AfterSendCoins(ctx sdk.Context,fromAddr sdk.AccAddress,toAddr sdk.AccAddress,amt sdk.Coins)error{
	for i:=range h{
		if err:=h[i].AfterSendCoins(ctx,fromAddr,toAddr,amt);err!=nil{
			return err
		}
	}
	return nil
}