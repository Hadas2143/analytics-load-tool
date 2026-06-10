# Analytics Load Simulation Tool 📊

A high-performance, production-grade traffic simulation infrastructure designed to replicate historical data load patterns with **93% precision** from production environments for active backend testing.

## 🎯 Project Overview
This tool automates the process of reading massive amounts of historical raw logs, storing them efficiently, parsing data structures into normalized entities, and executing controlled traffic simulation workloads onto production-like cloud databases.

## 🏗️ System Architecture & Workflow

<!-- גררי לפה את תמונת הארכיטקטורה הכחולה עם הריבועים הלבנים -->


1.  **Logs Source:** Ingests raw query data files (CSV, JSON, LOG).
2.  **File Reader:** Parses and streams raw records sequentially.
3.  **MongoDB Storage:** Stores unstructured raw records for fast staging.
4.  **Parser Pipeline:** Normalizes staging entities into predefined structural formats (`ParsedQuery`).
5.  **Traffic Simulator:** Controls precise load timing to replay traffic patterns realistically.
6.  **Formatter & Runner:** Builds and streams finalized data blocks directly into **Google BigQuery**.
7.  **Consul Integration:** Manages active system configurations and process checkpoint states.

## 📈 Observability & Monitoring

The system includes built-in metric collection pipelines connected directly to **Datadog** for full end-to-end operational visibility.

<!-- גררי לפה את תמונת הגרפים הלבנה של הדאשבורד -->


Metrics monitored include:
*   Real-time processing latency and query median duration (`loadtool.simulated.realtime.median`).
*   Pipeline ingest success rates (`loadtool.log.success`).
*   Total record throughput successfully delivered to target endpoints (`loadtool.records.sent`).

## 🛠️ Built With
*   **Language:** Go (Golang)
*   **Databases:** MongoDB, Google BigQuery
*   **Infrastructure & DevOps:** Docker, Kubernetes, Consul
*   **Monitoring:** Datadog
