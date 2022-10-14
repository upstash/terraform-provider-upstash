data "upstash_qstash_endpoint_data" "exampleQstashEndpointData" {
    endpoint_id = resource.upstash_qstash_endpoint.exampleQstashEndpoint.endpoint_id
}