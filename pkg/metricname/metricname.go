package metricname

const (
	//PROCESS_FLOW_INPUT_SOURCE_DONE_JOB = "process_flow_input_source_done_job"

	//PROCESS_FLOW_INPUT_SOURCE  = "process_flow_input_source"
	//PROCESS_FLOW_OUTPUT_SOURCE = "process_flow_output_source"

	//PROCESS_FLOW_INPUT_CORE  = "process_flow_input_core"

	ProcessFlowOutputCore = "process_flow_output_core"

	ProcessFlowInputDestination  = "process_flow_input_destination"
	ProcessFlowOutputDestination = "process_flow_output_destination"

	ProcessFlowInputDestinationWorker         = "process_flow_input_destination_worker"
	ProcessFlowOutputDestinationWorkerDoneJob = "process_flow_output_destination_worker_done_job"

	DestinationInputAckError       = "destination_input_ack_error"
	DestinationInputUnmarshalError = "destination_input_unmarshal_error"

	DestinationEventPublisherNotFound = "destination_event_publisher_not_found"
	DestinationEventPublishError      = "destination_event_publisher_publish_error"

	DestinationWorkerInputAckError       = "destination_worker_input_unmarshal_error"
	DestinationWorkerInputUnmarshalError = "destination_worker_input_unmarshal_error"

	DestinationWorkerEventSendToWorker = "destination_worker_send_to_worker"
	DestinationWorkerHandleEventError  = "destination_worker_handle_event_error"
)
