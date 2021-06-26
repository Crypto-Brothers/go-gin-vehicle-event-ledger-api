package model

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type VehicleToken struct {
	Name           string            `json:"name"`
	Symbol         string            `json:"symbol"`
	DecimalPlaces  int               `json:"decimalplaces"`
	InitialSupply  int               `json:"initialsupply"`
	TreasuryId     hedera.AccountID  `json:"treasuryid"`
	TreasuryKey    hedera.PrivateKey `json:"treasurykey"`
	AdminKey       hedera.PrivateKey `json:"adminkey"`
	KycKey         hedera.Key        `json:"kyckey"`
	FreezeKey      hedera.Key        `json:"freezekey"`
	WipeKey        hedera.Key        `json:"wipekey"`
	SupplyKey      hedera.Key        `json:"supplykey"`
	Freeze         bool              `json:"freeze"`
	ExpirationTime time.Time         `json:"expirationtime"`
	Memo           string            `json:"memo"`
}
