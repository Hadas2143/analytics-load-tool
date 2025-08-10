Analytics Load Tool
Analytics Load Tool is a Go-based system designed to replay historical traffic from log files into modern analytics pipelines.
It is primarily used for simulation, load testing, and data backfill operations.

The tool processes events from stored logs, sends them through MongoDB, and finally writes them into Google BigQuery.
It works locally or in cloud environments such as Kubernetes, with full observability through Datadog.

Main Features
Historical Log Replay – Process large log datasets and simulate original traffic patterns

MongoDB Integration – Insert events into one or more MongoDB instances with configurable write guarantees

BigQuery Writing – Supports streaming and batch inserts for performance and cost control

Query Manipulator – Modify or parameterize SQL dynamically

Real-Time Metrics & Health Checks – Monitor processing rates, errors, and latency

Configurable Playback – Adjust replay rate, time shift, and jitter for realistic scenarios

Cloud-Native Deployment – Kubernetes-ready with resource controls and readiness checks

Datadog Integration – Send custom application metrics for monitoring and alerting

Architecture Overview
Log Source – Reads from local files or object storage

Simulator – Parses events, adjusts timestamps, applies jitter, and controls replay speed

MongoDB – Stores intermediate or enriched data for downstream processing

Consumers – Read and transform data from MongoDB

BigQuery Loader – Writes processed data into BigQuery tables

Observability Layer – Datadog metrics, /healthz and /ready endpoints for monitoring

Data Flow:
[Logs] → [Simulator] → [MongoDB] → [Consumers] → [BigQuery]

Installation & Deployment
Prerequisites
Go 1.22+

GCP project with BigQuery enabled

Service Account JSON with correct BigQuery permissions

Deployment Modes
Local Execution – For development or small runs

Docker – Portable containerized environment

Kubernetes – Scalable, production-grade deployment with auto-restart and probes

Monitoring & Observability
Health Endpoints – /healthz and /ready for readiness and liveness checks

Datadog Integration – Environment, service, and version tags for filtering metrics

Performance Tracking – Monitor throughput, error rates, and latency over time
