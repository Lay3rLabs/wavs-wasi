// Code generated by wit-bindgen-go. DO NOT EDIT.

package host

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wavs:worker@0.4.0-alpha.2".

//go:wasmimport host get-eth-chain-config
//go:noescape
func wasmimport_GetEthChainConfig(chainName0 *uint8, chainName1 uint32, result *cm.Option[EthChainConfig])

//go:wasmimport host get-cosmos-chain-config
//go:noescape
func wasmimport_GetCosmosChainConfig(chainName0 *uint8, chainName1 uint32, result *cm.Option[CosmosChainConfig])

//go:wasmimport host log
//go:noescape
func wasmimport_Log(level0 uint32, message0 *uint8, message1 uint32)
