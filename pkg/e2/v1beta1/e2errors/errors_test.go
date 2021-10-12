// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2errors

import (
	"errors"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFactories(t *testing.T) {
	assert.Equal(t, RICUnspecified, NewRICUnspecified("").(*TypedError).E2APType)
	assert.Equal(t, "RICUnspecified", NewRICUnspecified("RICUnspecified").Error())

	assert.Equal(t, RICRANFunctionIDInvalid, NewRICRANFunctionIDInvalid("").(*TypedError).E2APType)
	assert.Equal(t, "RICRANFunctionIDInvalid", NewRICRANFunctionIDInvalid("RICRANFunctionIDInvalid").Error())

	assert.Equal(t, RICActionNotSupported, NewRICActionNotSupported("").(*TypedError).E2APType)
	assert.Equal(t, "RICActionNotSupported", NewRICActionNotSupported("RICActionNotSupported").Error())

	assert.Equal(t, RICExcessiveActions, NewRICExcessiveActions("").(*TypedError).E2APType)
	assert.Equal(t, "RICExcessiveActions", NewRICExcessiveActions("RICExcessiveActions").Error())

	assert.Equal(t, RICDuplicateAction, NewRICDuplicateAction("").(*TypedError).E2APType)
	assert.Equal(t, "RICDuplicateAction", NewRICDuplicateAction("RICDuplicateAction").Error())

	assert.Equal(t, RICDuplicateEvent, NewRICDuplicateEvent("").(*TypedError).E2APType)
	assert.Equal(t, "RICDuplicateEvent", NewRICDuplicateEvent("RICDuplicateEvent").Error())

	assert.Equal(t, RICFunctionResourceLimit, NewRICFunctionResourceLimit("").(*TypedError).E2APType)
	assert.Equal(t, "RICFunctionResourceLimit", NewRICFunctionResourceLimit("RICFunctionResourceLimit").Error())

	assert.Equal(t, RICRequestIDUnknown, NewRICRequestIDUnknown("").(*TypedError).E2APType)
	assert.Equal(t, "RICRequestIDUnknown", NewRICRequestIDUnknown("RICRequestIDUnknown").Error())

	assert.Equal(t, RICInconsistentActionSubsequentActionSequence, NewRICInconsistentActionSubsequentActionSequence("").(*TypedError).E2APType)
	assert.Equal(t, "RICInconsistentActionSubsequentActionSequence", NewRICInconsistentActionSubsequentActionSequence("RICInconsistentActionSubsequentActionSequence").Error())

	assert.Equal(t, RICControlMessageInvalid, NewRICControlMessageInvalid("").(*TypedError).E2APType)
	assert.Equal(t, "RICControlMessageInvalid", NewRICControlMessageInvalid("RICControlMessageInvalid").Error())

	assert.Equal(t, RICCallProcessIDInvalid, NewRICCallProcessIDInvalid("").(*TypedError).E2APType)
	assert.Equal(t, "RICCallProcessIDInvalid", NewRICCallProcessIDInvalid("RICCallProcessIDInvalid").Error())

	assert.Equal(t, RICServiceUnspecified, NewRICServiceUnspecified("").(*TypedError).E2APType)
	assert.Equal(t, "RICServiceUnspecified", NewRICServiceUnspecified("RICServiceUnspecified").Error())

	assert.Equal(t, RICServiceFunctionNotRequired, NewRICServiceFunctionNotRequired("").(*TypedError).E2APType)
	assert.Equal(t, "RICServiceFunctionNotRequired", NewRICServiceFunctionNotRequired("RICServiceFunctionNotRequired").Error())

	assert.Equal(t, RICServiceExcessiveFunctions, NewRICServiceExcessiveFunctions("").(*TypedError).E2APType)
	assert.Equal(t, "RICServiceExcessiveFunctions", NewRICServiceExcessiveFunctions("RICServiceExcessiveFunctions").Error())

	assert.Equal(t, RICServiceRICResourceLimit, NewRICServiceRICResourceLimit("").(*TypedError).E2APType)
	assert.Equal(t, "RICServiceRICResourceLimit", NewRICServiceRICResourceLimit("RICServiceRICResourceLimit").Error())

	assert.Equal(t, ProtocolUnspecified, NewProtocolUnspecified("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolUnspecified", NewProtocolUnspecified("ProtocolUnspecified").Error())

	assert.Equal(t, ProtocolTransferSyntaxError, NewProtocolTransferSyntaxError("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolTransferSyntaxError", NewProtocolTransferSyntaxError("ProtocolTransferSyntaxError").Error())

	assert.Equal(t, ProtocolAbstractSyntaxErrorReject, NewProtocolAbstractSyntaxErrorReject("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolAbstractSyntaxErrorReject", NewProtocolAbstractSyntaxErrorReject("ProtocolAbstractSyntaxErrorReject").Error())

	assert.Equal(t, ProtocolAbstractSyntaxErrorIgnoreAndNotify, NewProtocolAbstractSyntaxErrorIgnoreAndNotify("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolAbstractSyntaxErrorIgnoreAndNotify", NewProtocolAbstractSyntaxErrorIgnoreAndNotify("ProtocolAbstractSyntaxErrorIgnoreAndNotify").Error())

	assert.Equal(t, ProtocolMessageNotCompatibleWithReceiverState, NewProtocolMessageNotCompatibleWithReceiverState("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolMessageNotCompatibleWithReceiverState", NewProtocolMessageNotCompatibleWithReceiverState("ProtocolMessageNotCompatibleWithReceiverState").Error())

	assert.Equal(t, ProtocolSemanticError, NewProtocolSemanticError("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolSemanticError", NewProtocolSemanticError("ProtocolSemanticError").Error())

	assert.Equal(t, ProtocolAbstractSyntaxErrorFalselyConstructedMessage, NewProtocolAbstractSyntaxErrorFalselyConstructedMessage("").(*TypedError).E2APType)
	assert.Equal(t, "ProtocolAbstractSyntaxErrorFalselyConstructedMessage", NewProtocolAbstractSyntaxErrorFalselyConstructedMessage("ProtocolAbstractSyntaxErrorFalselyConstructedMessage").Error())

	assert.Equal(t, MiscUnspecified, NewMiscUnspecified("").(*TypedError).E2APType)
	assert.Equal(t, "MiscUnspecified", NewMiscUnspecified("MiscUnspecified").Error())

	assert.Equal(t, MiscControlProcessingOverload, NewMiscControlProcessingOverload("").(*TypedError).E2APType)
	assert.Equal(t, "MiscControlProcessingOverload", NewMiscControlProcessingOverload("MiscControlProcessingOverload").Error())

	assert.Equal(t, MiscHardwareFailure, NewMiscHardwareFailure("").(*TypedError).E2APType)
	assert.Equal(t, "MiscHardwareFailure", NewMiscHardwareFailure("MiscHardwareFailure").Error())

	assert.Equal(t, MiscOMIntervention, NewMiscOMIntervention("").(*TypedError).E2APType)
	assert.Equal(t, "MiscOMIntervention", NewMiscOMIntervention("MiscOMIntervention").Error())

}

func TestPredicates(t *testing.T) {
	assert.False(t, IsRICUnspecified(errors.New("RICUnspecified")))
	assert.True(t, IsRICUnspecified(NewRICUnspecified("RICUnspecified")))

	assert.False(t, IsRICRANFunctionIDInvalid(errors.New("RICRANFunctionIDInvalid")))
	assert.True(t, IsRICRANFunctionIDInvalid(NewRICRANFunctionIDInvalid("RICRANFunctionIDInvalid")))

	assert.False(t, IsRICActionNotSupported(errors.New("RICActionNotSupported")))
	assert.True(t, IsRICActionNotSupported(NewRICActionNotSupported("RICActionNotSupported")))

	assert.False(t, IsRICExcessiveActions(errors.New("RICExcessiveActions")))
	assert.True(t, IsRICExcessiveActions(NewRICExcessiveActions("RICExcessiveActions")))

	assert.False(t, IsRICDuplicateAction(errors.New("RICDuplicateAction")))
	assert.True(t, IsRICDuplicateAction(NewRICDuplicateAction("RICDuplicateAction")))

	assert.False(t, IsRICDuplicateEvent(errors.New("RICDuplicateEvent")))
	assert.True(t, IsRICDuplicateEvent(NewRICDuplicateEvent("RICDuplicateEvent")))

	assert.False(t, IsRICFunctionResourceLimit(errors.New("RICFunctionResourceLimit")))
	assert.True(t, IsRICFunctionResourceLimit(NewRICFunctionResourceLimit("RICFunctionResourceLimit")))

	assert.False(t, IsRICRequestIDUnknown(errors.New("RICRequestIDUnknown")))
	assert.True(t, IsRICRequestIDUnknown(NewRICRequestIDUnknown("RICRequestIDUnknown")))

	assert.False(t, IsRICInconsistentActionSubsequentActionSequence(errors.New("RICInconsistentActionSubsequentActionSequence")))
	assert.True(t, IsRICInconsistentActionSubsequentActionSequence(NewRICInconsistentActionSubsequentActionSequence("RICInconsistentActionSubsequentActionSequence")))

	assert.False(t, IsRICControlMessageInvalid(errors.New("RICControlMessageInvalid")))
	assert.True(t, IsRICControlMessageInvalid(NewRICControlMessageInvalid("RICControlMessageInvalid")))

	assert.False(t, IsRICCallProcessIDInvalid(errors.New("RICCallProcessIDInvalid")))
	assert.True(t, IsRICCallProcessIDInvalid(NewRICCallProcessIDInvalid("RICCallProcessIDInvalid")))

	assert.False(t, IsRICServiceUnspecified(errors.New("RICServiceUnspecified")))
	assert.True(t, IsRICServiceUnspecified(NewRICServiceUnspecified("RICServiceUnspecified")))

	assert.False(t, IsRICServiceFunctionNotRequired(errors.New("RICServiceFunctionNotRequired")))
	assert.True(t, IsRICServiceFunctionNotRequired(NewRICServiceFunctionNotRequired("RICServiceFunctionNotRequired")))

	assert.False(t, IsRICServiceExcessiveFunctions(errors.New("RICServiceExcessiveFunctions")))
	assert.True(t, IsRICServiceExcessiveFunctions(NewRICServiceExcessiveFunctions("RICServiceExcessiveFunctions")))

	assert.False(t, IsRICServiceRICResourceLimit(errors.New("RICServiceRICResourceLimit")))
	assert.True(t, IsRICServiceRICResourceLimit(NewRICServiceRICResourceLimit("RICServiceRICResourceLimit")))

	assert.False(t, IsProtocolUnspecified(errors.New("ProtocolUnspecified")))
	assert.True(t, IsProtocolUnspecified(NewProtocolUnspecified("ProtocolUnspecified")))

	assert.False(t, IsProtocolTransferSyntaxError(errors.New("ProtocolTransferSyntaxError")))
	assert.True(t, IsProtocolTransferSyntaxError(NewProtocolTransferSyntaxError("ProtocolTransferSyntaxError")))

	assert.False(t, IsProtocolAbstractSyntaxErrorReject(errors.New("ProtocolAbstractSyntaxErrorReject")))
	assert.True(t, IsProtocolAbstractSyntaxErrorReject(NewProtocolAbstractSyntaxErrorReject("ProtocolAbstractSyntaxErrorReject")))

	assert.False(t, IsProtocolAbstractSyntaxErrorIgnoreAndNotify(errors.New("ProtocolAbstractSyntaxErrorIgnoreAndNotify")))
	assert.True(t, IsProtocolAbstractSyntaxErrorIgnoreAndNotify(NewProtocolAbstractSyntaxErrorIgnoreAndNotify("ProtocolAbstractSyntaxErrorIgnoreAndNotify")))

	assert.False(t, IsProtocolMessageNotCompatibleWithReceiverState(errors.New("ProtocolMessageNotCompatibleWithReceiverState")))
	assert.True(t, IsProtocolMessageNotCompatibleWithReceiverState(NewProtocolMessageNotCompatibleWithReceiverState("ProtocolMessageNotCompatibleWithReceiverState")))

	assert.False(t, IsProtocolSemanticError(errors.New("ProtocolSemanticError")))
	assert.True(t, IsProtocolSemanticError(NewProtocolSemanticError("ProtocolSemanticError")))

	assert.False(t, IsProtocolAbstractSyntaxErrorFalselyConstructedMessage(errors.New("ProtocolAbstractSyntaxErrorFalselyConstructedMessage")))
	assert.True(t, IsProtocolAbstractSyntaxErrorFalselyConstructedMessage(NewProtocolAbstractSyntaxErrorFalselyConstructedMessage("ProtocolAbstractSyntaxErrorFalselyConstructedMessage")))

	assert.False(t, IsMiscUnspecified(errors.New("MiscUnspecified")))
	assert.True(t, IsMiscUnspecified(NewMiscUnspecified("MiscUnspecified")))

	assert.False(t, IsMiscControlProcessingOverload(errors.New("MiscControlProcessingOverload")))
	assert.True(t, IsMiscControlProcessingOverload(NewMiscControlProcessingOverload("MiscControlProcessingOverload")))

	assert.False(t, IsMiscHardwareFailure(errors.New("MiscHardwareFailure")))
	assert.True(t, IsMiscHardwareFailure(NewMiscHardwareFailure("MiscHardwareFailure")))

	assert.False(t, IsMiscOMIntervention(errors.New("MiscOMIntervention")))
	assert.True(t, IsMiscOMIntervention(NewMiscOMIntervention("MiscOMIntervention")))

}

func TestGRPCToError(t *testing.T) {
	assert.Nil(t, FromGRPC(status.New(codes.OK, "").Err()))
	e2apErr := &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_UNSPECIFIED,
				},
			},
		},
	}
	stat, err := status.New(codes.Internal, "RICUnspecified").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICUnspecified(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICRANFunctionIDInvalid").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICRANFunctionIDInvalid(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICActionNotSupported").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICActionNotSupported(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICExcessiveActions").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICExcessiveActions(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_DUPLICATE_ACTION,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICDuplicateAction").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICDuplicateAction(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_DUPLICATE_EVENT,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICDuplicateEvent").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICDuplicateEvent(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICFunctionResourceLimit").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICFunctionResourceLimit(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICRequestIDUnknown").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICRequestIDUnknown(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICInconsistentActionSubsequentActionSequence").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICInconsistentActionSubsequentActionSequence(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICControlMessageInvalid").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICControlMessageInvalid(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Ric_{
				Ric: &e2api.Error_Cause_Ric{
					Type: e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICCallProcessIDInvalid").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICCallProcessIDInvalid(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_RicService_{
				RicService: &e2api.Error_Cause_RicService{
					Type: e2api.Error_Cause_RicService_UNSPECIFIED,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICServiceUnspecified").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICServiceUnspecified(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_RicService_{
				RicService: &e2api.Error_Cause_RicService{
					Type: e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICServiceFunctionNotRequired").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICServiceFunctionNotRequired(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_RicService_{
				RicService: &e2api.Error_Cause_RicService{
					Type: e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICServiceExcessiveFunctions").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICServiceExcessiveFunctions(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_RicService_{
				RicService: &e2api.Error_Cause_RicService{
					Type: e2api.Error_Cause_RicService_RIC_RESOURCE_LIMIT,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "RICServiceRICResourceLimit").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsRICServiceRICResourceLimit(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_UNSPECIFIED,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolUnspecified").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolUnspecified(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolTransferSyntaxError").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolTransferSyntaxError(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolAbstractSyntaxErrorReject").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolAbstractSyntaxErrorReject(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolAbstractSyntaxErrorIgnoreAndNotify").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolAbstractSyntaxErrorIgnoreAndNotify(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolMessageNotCompatibleWithReceiverState").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolMessageNotCompatibleWithReceiverState(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_SEMANTIC_ERROR,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolSemanticError").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolSemanticError(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Protocol_{
				Protocol: &e2api.Error_Cause_Protocol{
					Type: e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "ProtocolAbstractSyntaxErrorFalselyConstructedMessage").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsProtocolAbstractSyntaxErrorFalselyConstructedMessage(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Misc_{
				Misc: &e2api.Error_Cause_Misc{
					Type: e2api.Error_Cause_Misc_UNSPECIFIED,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "MiscUnspecified").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsMiscUnspecified(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Misc_{
				Misc: &e2api.Error_Cause_Misc{
					Type: e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "MiscControlProcessingOverload").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsMiscControlProcessingOverload(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Misc_{
				Misc: &e2api.Error_Cause_Misc{
					Type: e2api.Error_Cause_Misc_HARDWARE_FAILURE,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "MiscHardwareFailure").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsMiscHardwareFailure(FromGRPC(stat.Err())))

	e2apErr = &e2api.Error{
		Cause: &e2api.Error_Cause{
			Cause: &e2api.Error_Cause_Misc_{
				Misc: &e2api.Error_Cause_Misc{
					Type: e2api.Error_Cause_Misc_OM_INTERVENTION,
				},
			},
		},
	}
	stat, err = status.New(codes.Internal, "MiscOMIntervention").WithDetails(e2apErr)
	assert.NoError(t, err)
	assert.True(t, IsMiscOMIntervention(FromGRPC(stat.Err())))

}
