// Code generated by wit-bindgen-go. DO NOT EDIT.

package layertypes

import (
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// TriggerSourceCronShape is used for storage in variant or result types.
type TriggerSourceCronShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(TriggerSourceCron{})]byte
}

// TriggerDataCosmosContractEventShape is used for storage in variant or result types.
type TriggerDataCosmosContractEventShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(TriggerDataCosmosContractEvent{})]byte
}
