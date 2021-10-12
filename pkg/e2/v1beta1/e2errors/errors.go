// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2errors

import (
	"fmt"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"google.golang.org/grpc/status"
)

// E2APType is an E2AP error type
type E2APType int

// Error type constants
const (
	Unknown E2APType = iota

	RICUnspecified

	RICRANFunctionIDInvalid

	RICActionNotSupported

	RICExcessiveActions

	RICDuplicateAction

	RICDuplicateEvent

	RICFunctionResourceLimit

	RICRequestIDUnknown

	RICInconsistentActionSubsequentActionSequence

	RICControlMessageInvalid

	RICCallProcessIDInvalid

	RICServiceUnspecified

	RICServiceFunctionNotRequired

	RICServiceExcessiveFunctions

	RICServiceRICResourceLimit

	ProtocolUnspecified

	ProtocolTransferSyntaxError

	ProtocolAbstractSyntaxErrorReject

	ProtocolAbstractSyntaxErrorIgnoreAndNotify

	ProtocolMessageNotCompatibleWithReceiverState

	ProtocolSemanticError

	ProtocolAbstractSyntaxErrorFalselyConstructedMessage

	MiscUnspecified

	MiscControlProcessingOverload

	MiscHardwareFailure

	MiscOMIntervention
)

// TypedError is a typed error
type TypedError struct {
	// E2APType is the E2AP error type
	E2APType E2APType
	// Message is the error message
	Message string
}

func (e *TypedError) Error() string {
	return e.Message
}

var _ error = &TypedError{}

// FromGRPC creates a typed error from a gRPC error
func FromGRPC(err error) error {
	if err == nil {
		return nil
	}

	stat, ok := status.FromError(err)
	if !ok {
		return New(Unknown, err.Error())
	}
	details := stat.Details()

	if len(details) == 0 {
		return New(Unknown, err.Error())
	}

	switch t := details[0].(type) {
	case *e2api.Error:
		cause := t.GetCause().GetCause()
		switch c := cause.(type) {
		case *e2api.Error_Cause_Ric_:
			switch c.Ric.Type {
			case e2api.Error_Cause_Ric_UNSPECIFIED:
				return New(RICUnspecified, stat.Message())
			case e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID:
				return New(RICRANFunctionIDInvalid, stat.Message())
			case e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED:
				return New(RICActionNotSupported, stat.Message())
			case e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS:
				return New(RICExcessiveActions, stat.Message())
			case e2api.Error_Cause_Ric_DUPLICATE_ACTION:
				return New(RICDuplicateAction, stat.Message())
			case e2api.Error_Cause_Ric_DUPLICATE_EVENT:
				return New(RICDuplicateEvent, stat.Message())
			case e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT:
				return New(RICFunctionResourceLimit, stat.Message())
			case e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN:
				return New(RICRequestIDUnknown, stat.Message())
			case e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
				return New(RICInconsistentActionSubsequentActionSequence, stat.Message())
			case e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID:
				return New(RICControlMessageInvalid, stat.Message())
			case e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID:
				return New(RICCallProcessIDInvalid, stat.Message())
			default:
				return New(RICUnspecified, stat.Message())

			}
		case *e2api.Error_Cause_Protocol_:
			switch c.Protocol.Type {
			case e2api.Error_Cause_Protocol_UNSPECIFIED:
				return New(ProtocolUnspecified, stat.Message())
			case e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR:
				return New(ProtocolTransferSyntaxError, stat.Message())
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT:
				return New(ProtocolAbstractSyntaxErrorReject, stat.Message())
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
				return New(ProtocolAbstractSyntaxErrorIgnoreAndNotify, stat.Message())
			case e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
				return New(ProtocolMessageNotCompatibleWithReceiverState, stat.Message())
			case e2api.Error_Cause_Protocol_SEMANTIC_ERROR:
				return New(ProtocolSemanticError, stat.Message())
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
				return New(ProtocolAbstractSyntaxErrorFalselyConstructedMessage, stat.Message())

			}
		case *e2api.Error_Cause_Misc_:
			switch c.Misc.Type {
			case e2api.Error_Cause_Misc_UNSPECIFIED:
				return New(MiscUnspecified, stat.Message())
			case e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD:
				return New(MiscControlProcessingOverload, stat.Message())
			case e2api.Error_Cause_Misc_OM_INTERVENTION:
				return New(MiscOMIntervention, stat.Message())
			case e2api.Error_Cause_Misc_HARDWARE_FAILURE:
				return New(MiscHardwareFailure, stat.Message())

			}
		case *e2api.Error_Cause_RicService_:
			switch c.RicService.Type {
			case e2api.Error_Cause_RicService_UNSPECIFIED:
				return New(RICServiceUnspecified, stat.Message())
			case e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED:
				return New(RICServiceFunctionNotRequired, stat.Message())
			case e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS:
				return New(RICServiceExcessiveFunctions, stat.Message())
			case e2api.Error_Cause_RicService_RIC_RESOURCE_LIMIT:
				return New(RICServiceRICResourceLimit, stat.Message())

			}

		default:
			return New(Unknown, stat.Message())
		}
	default:
		return New(Unknown, stat.Message())
	}
	return New(Unknown, stat.Message())
}

// New creates a new typed error
func New(t E2APType, msg string, args ...interface{}) error {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return &TypedError{
		E2APType: t,
		Message:  msg,
	}
}

// NewUnknown returns a new Unknown error
func NewUnknown(msg string, args ...interface{}) error {
	return New(Unknown, msg, args...)
}

// NewRICUnspecified returns a new RIC Unspecified error
func NewRICUnspecified(msg string, args ...interface{}) error {
	return New(RICUnspecified, msg, args...)
}

// NewRICRANFunctionIDInvalid returns a new RICRANFunctionIDInvalid error
func NewRICRANFunctionIDInvalid(msg string, args ...interface{}) error {
	return New(RICRANFunctionIDInvalid, msg, args...)
}

// NewRICActionNotSupported returns a new RICActionNotSupported error
func NewRICActionNotSupported(msg string, args ...interface{}) error {
	return New(RICActionNotSupported, msg, args...)
}

// NewRICExcessiveActions returns a new RICExcessiveActions error
func NewRICExcessiveActions(msg string, args ...interface{}) error {
	return New(RICExcessiveActions, msg, args...)
}

// NewRICDuplicateAction returns a new RICDuplicateAction  error
func NewRICDuplicateAction(msg string, args ...interface{}) error {
	return New(RICDuplicateAction, msg, args...)
}

// NewRICDuplicateEvent returns a new RICDuplicateEvent error
func NewRICDuplicateEvent(msg string, args ...interface{}) error {
	return New(RICDuplicateEvent, msg, args...)
}

// NewRICFunctionResourceLimit returns a new RICFunctionResourceLimit error
func NewRICFunctionResourceLimit(msg string, args ...interface{}) error {
	return New(RICFunctionResourceLimit, msg, args...)
}

// NewRICRequestIDUnknown returns a new RICRequestIDUnknown error
func NewRICRequestIDUnknown(msg string, args ...interface{}) error {
	return New(RICRequestIDUnknown, msg, args...)
}

// NewRICInconsistentActionSubsequentActionSequence returns a new RICInconsistentActionSubsequentActionSequence error
func NewRICInconsistentActionSubsequentActionSequence(msg string, args ...interface{}) error {
	return New(RICInconsistentActionSubsequentActionSequence, msg, args...)
}

// NewRICControlMessageInvalid returns a new RICControlMessageInvalid error
func NewRICControlMessageInvalid(msg string, args ...interface{}) error {
	return New(RICControlMessageInvalid, msg, args...)
}

// NewRICCallProcessIDInvalid returns a new RICCallProcessIDInvalid error
func NewRICCallProcessIDInvalid(msg string, args ...interface{}) error {
	return New(RICCallProcessIDInvalid, msg, args...)
}

// NewRICServiceUnspecified returns a new RICServiceUnspecified error
func NewRICServiceUnspecified(msg string, args ...interface{}) error {
	return New(RICServiceUnspecified, msg, args...)
}

// NewRICServiceFunctionNotRequired returns a new 	RICServiceFunctionNotRequired error
func NewRICServiceFunctionNotRequired(msg string, args ...interface{}) error {
	return New(RICServiceFunctionNotRequired, msg, args...)
}

// NewRICServiceExcessiveFunctions returns a new RICServiceExcessiveFunctions error
func NewRICServiceExcessiveFunctions(msg string, args ...interface{}) error {
	return New(RICServiceExcessiveFunctions, msg, args...)
}

// NewRICServiceRICResourceLimit returns a new 	RICServiceRICResourceLimit error
func NewRICServiceRICResourceLimit(msg string, args ...interface{}) error {
	return New(RICServiceRICResourceLimit, msg, args...)
}

// NewProtocolUnspecified returns a new ProtocolUnspecified error
func NewProtocolUnspecified(msg string, args ...interface{}) error {
	return New(ProtocolUnspecified, msg, args...)
}

// NewProtocolTransferSyntaxError returns a new ProtocolTransferSyntaxError error
func NewProtocolTransferSyntaxError(msg string, args ...interface{}) error {
	return New(ProtocolTransferSyntaxError, msg, args...)
}

// NewProtocolAbstractSyntaxErrorReject returns a new ProtocolAbstractSyntaxErrorReject error
func NewProtocolAbstractSyntaxErrorReject(msg string, args ...interface{}) error {
	return New(ProtocolAbstractSyntaxErrorReject, msg, args...)
}

// NewProtocolAbstractSyntaxErrorIgnoreAndNotify returns a new 	ProtocolAbstractSyntaxErrorIgnoreAndNotify error
func NewProtocolAbstractSyntaxErrorIgnoreAndNotify(msg string, args ...interface{}) error {
	return New(ProtocolAbstractSyntaxErrorIgnoreAndNotify, msg, args...)
}

// NewProtocolMessageNotCompatibleWithReceiverState returns a new ProtocolMessageNotCompatibleWithReceiverState error
func NewProtocolMessageNotCompatibleWithReceiverState(msg string, args ...interface{}) error {
	return New(ProtocolMessageNotCompatibleWithReceiverState, msg, args...)
}

// NewProtocolSemanticError returns a new ProtocolSemanticError error
func NewProtocolSemanticError(msg string, args ...interface{}) error {
	return New(ProtocolSemanticError, msg, args...)
}

// NewProtocolAbstractSyntaxErrorFalselyConstructedMessage returns a new ProtocolAbstractSyntaxErrorFalselyConstructedMessage error
func NewProtocolAbstractSyntaxErrorFalselyConstructedMessage(msg string, args ...interface{}) error {
	return New(ProtocolAbstractSyntaxErrorFalselyConstructedMessage, msg, args...)
}

// NewMiscUnspecified returns a new MiscUnspecified error
func NewMiscUnspecified(msg string, args ...interface{}) error {
	return New(MiscUnspecified, msg, args...)
}

// NewMiscControlProcessingOverload returns a new MiscControlProcessingOverload error
func NewMiscControlProcessingOverload(msg string, args ...interface{}) error {
	return New(MiscControlProcessingOverload, msg, args...)
}

// NewMiscHardwareFailure returns a new MiscHardwareFailure error
func NewMiscHardwareFailure(msg string, args ...interface{}) error {
	return New(MiscHardwareFailure, msg, args...)
}

// NewMiscOMIntervention returns a new MiscOMIntervention error
func NewMiscOMIntervention(msg string, args ...interface{}) error {
	return New(MiscOMIntervention, msg, args...)
}

// TypeOf returns the type of the given error
func TypeOf(err error) E2APType {
	if typed, ok := err.(*TypedError); ok {
		return typed.E2APType
	}
	return Unknown
}

// IsType checks whether the given error is of the given type
func IsType(err error, t E2APType) bool {
	if typed, ok := err.(*TypedError); ok {
		return typed.E2APType == t
	}
	return false
}

// IsRICUnspecified checks whether the given error is a RIC Unspecified error
func IsRICUnspecified(err error) bool {
	return IsType(err, RICUnspecified)
}

// IsRICRANFunctionIDInvalid checks whether the given error is a RICRANFunctionIDInvalid error
func IsRICRANFunctionIDInvalid(err error) bool {
	return IsType(err, RICRANFunctionIDInvalid)
}

// IsRICActionNotSupported checks whether the given error is a RICActionNotSupported error
func IsRICActionNotSupported(err error) bool {
	return IsType(err, RICActionNotSupported)
}

// IsRICExcessiveActions checks whether the given error is a RICExcessiveActions error
func IsRICExcessiveActions(err error) bool {
	return IsType(err, RICExcessiveActions)
}

// IsRICDuplicateAction checks whether the given error is a RICDuplicateAction error
func IsRICDuplicateAction(err error) bool {
	return IsType(err, RICDuplicateAction)
}

// IsRICDuplicateEvent checks whether the given error is a RICDuplicateEvent error
func IsRICDuplicateEvent(err error) bool {
	return IsType(err, RICDuplicateEvent)
}

// IsRICFunctionResourceLimit checks whether the given error is a 	RICFunctionResourceLimit error
func IsRICFunctionResourceLimit(err error) bool {
	return IsType(err, RICFunctionResourceLimit)
}

// IsRICRequestIDUnknown checks whether the given error is a RICRequestIDUnknown error
func IsRICRequestIDUnknown(err error) bool {
	return IsType(err, RICRequestIDUnknown)
}

// IsRICInconsistentActionSubsequentActionSequence checks whether the given error is a 	RICInconsistentActionSubsequentActionSequence error
func IsRICInconsistentActionSubsequentActionSequence(err error) bool {
	return IsType(err, RICInconsistentActionSubsequentActionSequence)
}

// IsRICControlMessageInvalid checks whether the given error is a RICControlMessageInvalid error
func IsRICControlMessageInvalid(err error) bool {
	return IsType(err, RICControlMessageInvalid)
}

// IsRICCallProcessIDInvalid checks whether the given error is a RICCallProcessIDInvalid error
func IsRICCallProcessIDInvalid(err error) bool {
	return IsType(err, RICCallProcessIDInvalid)
}

// IsRICServiceUnspecified checks whether the given error is a RICServiceUnspecified error
func IsRICServiceUnspecified(err error) bool {
	return IsType(err, RICServiceUnspecified)
}

// IsRICServiceFunctionNotRequired checks whether the given error is a RICServiceFunctionNotRequired error
func IsRICServiceFunctionNotRequired(err error) bool {
	return IsType(err, RICServiceFunctionNotRequired)
}

// IsRICServiceExcessiveFunctions checks whether the given error is a RICServiceExcessiveFunctions error
func IsRICServiceExcessiveFunctions(err error) bool {
	return IsType(err, RICServiceExcessiveFunctions)
}

// IsRICServiceRICResourceLimit checks whether the given error is a RICServiceRICResourceLimit error
func IsRICServiceRICResourceLimit(err error) bool {
	return IsType(err, RICServiceRICResourceLimit)
}

// IsProtocolUnspecified checks whether the given error is a ProtocolUnspecified error
func IsProtocolUnspecified(err error) bool {
	return IsType(err, ProtocolUnspecified)
}

// IsProtocolTransferSyntaxError checks whether the given error is a ProtocolTransferSyntaxError error
func IsProtocolTransferSyntaxError(err error) bool {
	return IsType(err, ProtocolTransferSyntaxError)
}

// IsProtocolAbstractSyntaxErrorReject checks whether the given error is a ProtocolAbstractSyntaxErrorReject error
func IsProtocolAbstractSyntaxErrorReject(err error) bool {
	return IsType(err, ProtocolAbstractSyntaxErrorReject)
}

// IsProtocolAbstractSyntaxErrorIgnoreAndNotify checks whether the given error is a ProtocolAbstractSyntaxErrorIgnoreAndNotify error
func IsProtocolAbstractSyntaxErrorIgnoreAndNotify(err error) bool {
	return IsType(err, ProtocolAbstractSyntaxErrorIgnoreAndNotify)
}

// IsProtocolMessageNotCompatibleWithReceiverState checks whether the given error is a ProtocolMessageNotCompatibleWithReceiverState error
func IsProtocolMessageNotCompatibleWithReceiverState(err error) bool {
	return IsType(err, ProtocolMessageNotCompatibleWithReceiverState)
}

// IsProtocolSemanticError checks whether the given error is a 	ProtocolSemanticError error
func IsProtocolSemanticError(err error) bool {
	return IsType(err, ProtocolSemanticError)
}

// IsProtocolAbstractSyntaxErrorFalselyConstructedMessage checks whether the given error is a ProtocolAbstractSyntaxErrorFalselyConstructedMessage error
func IsProtocolAbstractSyntaxErrorFalselyConstructedMessage(err error) bool {
	return IsType(err, ProtocolAbstractSyntaxErrorFalselyConstructedMessage)
}

// IsMiscUnspecified checks whether the given error is a MiscUnspecified error
func IsMiscUnspecified(err error) bool {
	return IsType(err, MiscUnspecified)
}

// IsMiscControlProcessingOverload checks whether the given error is a 	MiscControlProcessingOverload error
func IsMiscControlProcessingOverload(err error) bool {
	return IsType(err, MiscControlProcessingOverload)
}

// IsMiscHardwareFailure checks whether the given error is a MiscHardwareFailure error
func IsMiscHardwareFailure(err error) bool {
	return IsType(err, MiscHardwareFailure)
}

// IsMiscOMIntervention checks whether the given error is a MiscOMIntervention error
func IsMiscOMIntervention(err error) bool {
	return IsType(err, MiscOMIntervention)
}

// IsE2APError checks if a given error is an E2AP error
func IsE2APError(err error) bool {
	if err == nil {
		return false
	}

	stat, ok := status.FromError(err)
	if !ok {
		return false
	}
	details := stat.Details()

	if len(details) == 0 {
		return false
	}

	switch details[0].(type) {
	case *e2api.Error:
		return true

	default:
		return false
	}
}
