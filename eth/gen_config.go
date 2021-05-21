// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"

	"github.com/yunhailanuxgk/go-jinbao/common"
	"github.com/yunhailanuxgk/go-jinbao/common/hexutil"
	"github.com/yunhailanuxgk/go-jinbao/consensus/ethash"
	"github.com/yunhailanuxgk/go-jinbao/core"
	"github.com/yunhailanuxgk/go-jinbao/eth/downloader"
	"github.com/yunhailanuxgk/go-jinbao/eth/gasprice"
)

func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		LightServ               int  `toml:",omitempty"`
		LightPeers              int  `toml:",omitempty"`
		MaxPeers                int  `toml:"-"`
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		Etherbase               common.Address `toml:",omitempty"`
		MinerThreads            int            `toml:",omitempty"`
		ExtraData               hexutil.Bytes  `toml:",omitempty"`
		GasPrice                *big.Int
		EthashCacheDir          string
		EthashCachesInMem       int
		EthashCachesOnDisk      int
		EthashDatasetDir        string
		EthashDatasetsInMem     int
		EthashDatasetsOnDisk    int
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string      `toml:"-"`
		PowMode                 ethash.Mode `toml:"-"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.Etherbase = c.Etherbase
	enc.MinerThreads = c.MinerThreads
	enc.ExtraData = c.ExtraData
	enc.GasPrice = c.GasPrice
	enc.EthashCacheDir = c.Ethash.CacheDir
	enc.EthashCachesInMem = c.Ethash.CachesInMem
	enc.EthashCachesOnDisk = c.Ethash.CachesOnDisk
	enc.EthashDatasetDir = c.Ethash.DatasetDir
	enc.EthashDatasetsInMem = c.Ethash.DatasetsInMem
	enc.EthashDatasetsOnDisk = c.Ethash.DatasetsOnDisk
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.PowMode = c.Ethash.PowMode
	return &enc, nil
}

func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		LightServ               *int  `toml:",omitempty"`
		LightPeers              *int  `toml:",omitempty"`
		MaxPeers                *int  `toml:"-"`
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		Etherbase               *common.Address `toml:",omitempty"`
		MinerThreads            *int            `toml:",omitempty"`
		ExtraData               hexutil.Bytes   `toml:",omitempty"`
		GasPrice                *big.Int
		EthashCacheDir          *string
		EthashCachesInMem       *int
		EthashCachesOnDisk      *int
		EthashDatasetDir        *string
		EthashDatasetsInMem     *int
		EthashDatasetsOnDisk    *int
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string      `toml:"-"`
		PowMode                 *ethash.Mode `toml:"-"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.Etherbase != nil {
		c.Etherbase = *dec.Etherbase
	}
	if dec.MinerThreads != nil {
		c.MinerThreads = *dec.MinerThreads
	}
	if dec.ExtraData != nil {
		c.ExtraData = dec.ExtraData
	}
	if dec.GasPrice != nil {
		c.GasPrice = dec.GasPrice
	}
	if dec.EthashCacheDir != nil {
		c.Ethash.CacheDir = *dec.EthashCacheDir
	}
	if dec.EthashCachesInMem != nil {
		c.Ethash.CachesInMem = *dec.EthashCachesInMem
	}
	if dec.EthashCachesOnDisk != nil {
		c.Ethash.CachesOnDisk = *dec.EthashCachesOnDisk
	}
	if dec.EthashDatasetDir != nil {
		c.Ethash.DatasetDir = *dec.EthashDatasetDir
	}
	if dec.EthashDatasetsInMem != nil {
		c.Ethash.DatasetsInMem = *dec.EthashDatasetsInMem
	}
	if dec.EthashDatasetsOnDisk != nil {
		c.Ethash.DatasetsOnDisk = *dec.EthashDatasetsOnDisk
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.PowMode != nil {
		c.Ethash.PowMode = *dec.PowMode
	}
	return nil
}
