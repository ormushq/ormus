package metricname

const (
	PROCESS_FLOW_INPUT_SOURCE_DONE_JOB = "process_flow_input_source_done_job"

	PROCESS_FLOW_INPUT_SOURCE  = "process_flow_input_source"
	PROCESS_FLOW_OUTPUT_SOURCE = "process_flow_output_source"

	PROCESS_FLOW_INPUT_CORE  = "process_flow_input_core"
	PROCESS_FLOW_OUTPUT_CORE = "process_flow_output_core"

	PROCESS_FLOW_INPUT_DESTINATION  = "process_flow_input_destination"
	PROCESS_FLOW_OUTPUT_DESTINATION = "process_flow_output_destination"

	PROCESS_FLOW_INPUT_DESTINATION_WORKER           = "process_flow_input_destination_worker"
	PROCESS_FLOW_OUTPUT_DESTINATION_WORKER_DONE_JOB = "process_flow_output_destination_worker_done_job"

	DESTINATION_INPUT_ACK_ERROR       = "destination_input_ack_error"
	DESTINATION_INPUT_UNMARSHAL_ERROR = "destination_input_unmarshal_error"

	DESTINATION_EVENT_PUBLISHER_NOT_FOUND = "destination_event_publisher_not_found"
	DESTINATION_EVENT_PUBLISH_ERROR       = "destination_event_publisher_publish_error"

	DESTINATION_WORKER_INPUT_ACK_ERROR       = "destination_worker_input_unmarshal_error"
	DESTINATION_WORKER_INPUT_UNMARSHAL_ERROR = "destination_worker_input_unmarshal_error"

	DESTINATION_WORKER_EVENT_SEND_TO_WORKER = "destination_worker_send_to_worker"
	DESTINATION_WORKER_HANDLE_EVENT_ERROR   = "destination_worker_handle_event_error"
)
