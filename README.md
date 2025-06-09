# Foliage SDK

<p align="center">
  <img src="./docs/pics/logo.png" width="600" alt="Foliage Logo">
</p>

[Foliage](https://www.foliage.dev/) is a collaborative application platform built upon a distributed graph database, providing a unified and extensible environment for effortless automation, cross-domain connectivity, and high-performance, edge-friendly runtimes.

[![License][License-Image]][License-Url] ![Lint][Lint-Status-Image-Url]

[License-Url]: https://www.apache.org/licenses/LICENSE-2.0
[License-Image]: https://img.shields.io/badge/License-Apache2-blue.svg
[Lint-Status-Image-Url]: https://github.com/foliagecp/sdk/actions/workflows/golangci-lint.yml/badge.svg

## Table of Contents

- [Introduction](#introduction)<!-- omit from toc --> 
- [Core Concepts](#core-concepts)
  - [Abstract](#abstract)
  - [Features](#features)
- [Getting Started](#getting-started)
  - [Minimum Requirements](#minimum-requirements)
  - [Installation](#installation)
  - [Health Status Check](#health-status-check)
  - [Running Tests](#running-tests)
  - [Customization](#customization)
- [Development](#development)
  - [Working with the SDK](#working-with-the-sdk)
- [Technology Stack](#technology-stack)
- [Roadmap](#roadmap)
- [References](#references)
- [License](#license)
- [Contribution](#contribution)

## 🚀 Introduction 

**Foliage** is an open-source, collaborative platform powered by a distributed graph database. It offers a flexible and high-performance foundation for edge computing, automation, and cross-domain integration. Ideal for IoT, process orchestration, and real-time logic execution.
## 🧠 Core Concepts

### 🧩 Abstract

Foliage reimagines complex systems by bringing their structure and logic into a shared, abstracted graph space. This enables clearer visibility and deeper understanding—blurring the line between model and implementation.

![Abstract](./docs/pics/FoliageUnification.jpg)

# System Architecture Overview

The system is conceptually divided into two main layers:

- **Data Layer** — a graph and object model describing the structure and relationships between elements.
- **Logic Layer** — a functional layer implementing the system’s behavior.

---

## Functional Layer

The functional layer is built on **serverless stateful functions**, each encapsulating the behavior of a specific type of object or a group of related objects represented in the object model. These functions can be viewed as **logic projections onto individual graph nodes**.

Each function:

- Maintains its own **state**, storing information about previous interaction steps with a specific object.
- Holds a **separate state for each object** it interacts with.
- Is responsible for executing a **clearly defined fragment of business logic**.

---

## Function Interaction

During execution, a function may:

- **Send calls (signals)** to other functions associated with neighboring objects in the graph.
- **Wait for their responses** before continuing its own logic.

A key feature is the **non-blocking execution model**:

- When sending signals, the function **completes the current execution iteration**, saving context information in its state.
- Upon reactivation (e.g., triggered by signals from previously called functions), it **resumes execution** with the received results.
- This design **avoids blocking** the function while waiting for responses and allows it to be **reused in other tasks**.

---

## Principle of Mass Asynchronous Interaction

This behavior model reflects processes typical in real-world systems (such as large organizations), where:

- Numerous **agents (functions)** simultaneously tackle different subtasks.
- They **exchange information and coordinate**, but are **not rigidly dependent** on the completion of all interactions before starting new ones.
- The result is **scalable, parallel, and asynchronous** operation without mutual blocking.

This ensures the system is **resilient to delays** and can **efficiently handle tasks** that are complex in depth and distributed in logic.

### 🔍 Features

By unifying various knowledge domains into a single shared space, Foliage simplifies relationships, uncovers hidden dependencies, and promotes consistent logic across your architecture.


#### Graph Database & Functional Graph
Foliage utilizes graph vertices for data storage and edges for connectivity between objects. These edges also serve as a means of data storage. Functions can be triggered by incoming signals within the context of vertices and utilize edges to call functions on neighboring vertices.

#### Information & Metainformation in One Graph
In Foliage, all object types, link types, type connectivity, functions, and applications are stored in the same data graph. This integration allows for complete linkage between data and metadata, facilitating graph traversal and signal distribution based on data types.

#### Distributed Event Bus
Foliage employs an asynchronous event system where all signals are represented as events in various topics within a clusterized event bus. The event bus is persistent and implements an "exactly once" method of signal processing.

#### Distributed Async Runtime
An application's runtime is composed of asynchronous functions called on objects. Business logic is defined through the declaration of call chains that implement its various use cases. Functions can be distributed across geographical and logical boundaries, providing flexibility in execution.

#### Serverless Stateful Functions
Foliage's stateful functions are distributed across numerous runtimes connected to the common network, eliminating the need for centralized management. Each stateful function instance has its own persistent context, enabling it to store data between calls.

#### Persistent Storage
Function contexts are persistent and stored asynchronously in the central core cluster. They can be restored in the event of function or application crashes or relocations. Each function has a dedicated context for each object (graph vertex).

#### Graph as Signal Path
Graphs in Foliage serve not only as data models but also as a means to propagate signals from one object to another. Signals can traverse one or many edges, depending on edge types and attributes.

#### Edge-Friendly Runtime
Functions can be directly triggered on object controllers, such as BMC, PLC, RPi, etc. Function calls are routed to edge runtimes running on corresponding controllers.

#### High Performance
Foliage boasts high performance, capable of handling up to 400,000 function calls per second on a midsize server. Scalability is nearly linear through clusterization.

#### XPath-Like Query Language for Graph Traversal and Signal Distribution
Foliage provides a query language for defining graph queries, enabling graph exploration, object discovery, and the definition of traversal routes for signal distribution, among other applications.

#### Applications Can Run Simultaneously and Extend Existing Graph Functional Models
Applications can coexist and interact using the same graph data model and communicate through graph edges. This data model reuse enhances code reuse.

#### No-Code Function and Application Construction
Foliage allows for low-code/no-code declaration of both functions (via scripts and configuration files) and applications. Applications can be visually designed or configured via declaration of call chains.

#### Flexible/Tunable No-Code Graph Observation Web UI Construction
Foliage offers a template-based UI toolkit, enabling the use of graphs as data models to construct interactive user interfaces with graph objects and links.

#### Weighted Functional Graphs for Scalar, Vector, and Tensor Signals Propagation, and ML Applications
In Foliage's graph database, links between vertices can be weighted, effectively turning the entire graph into a neural-like network. This feature makes Foliage suitable for MLOps applications and signal propagation involving scalars, vectors, and tensors.


## 🛠️ Getting Started

### ✅ Minimum Requirements

**Native Installation**

Foliage’s native setup follows the same system requirements as [NATS Jetstream](https://docs.nats.io/running-a-nats-service/introduction/installation#with-jetstream).

**Docker Installation**

To run Foliage in a containerized environment, all you need is [Docker](https://docs.docker.com/desktop/install/linux-install/).

---

### 📥 Installation

Clone the repository to get started:

```bash
git clone https://github.com/foliagecp/sdk.git
```

📚 Full setup details are available in the [documentation](https://pkg.go.dev/github.com/foliagecp/sdk).

---

### 🩺 Health Check

1. **Check that NATS server and Foliage runtime are running fine**:
```sh
% docker ps

CONTAINER ID   IMAGE                      COMMAND                  CREATED          STATUS          PORTS                                                                    NAMES
...
b5a2deb84082   foliage-sdk-tests:latest   "/usr/bin/tests basic"   11 minutes ago   Up 11 minutes                                                                            tests-runtime-1
fac8d1bfef3a   nats:latest                "/nats-server -js -s…"   11 minutes ago   Up 11 minutes   0.0.0.0:4222->4222/tcp, 0.0.0.0:6222->6222/tcp, 0.0.0.0:8222->8222/tcp   tests-nats-1
``` 

2. **Check that NATS server is running fine**:
```sh
% docker logs tests-nats-1

...
[1] 2023/10/16 09:00:43.094325 [INF] Server is ready
```

3. **Check that Foliage runtime runs without errors**:
```sh
% docker logs tests-runtime-1 | grep "error" -i
```


### 🧪 Running Tests

Foliage provides a set of test samples to help you get familiar with the platform. Follow these steps to run them:

#### 1. Navigate to `tests`:

```sh
cd tests
```

#### 2. Build the tests runtime:

```sh
docker-compose build
```

####3. Customize the test environment in `.env`:
```sh
nano ./basic/.env
```

####4. Start the tests:

```sh
docker-compose up -d
```

####5. To stop and clean everything up:
```sh
docker-compose down -v
```
💡 Use the `TEST_NAME` environment variable to run other test suites. TK!



## 🧑‍💻 Development

### Working with the SDK

Start building your own apps with Foliage by installing the SDK:

```sh
go get github.com/foliagecp/sdk
```

Helpful guides:

- [Graph CRUD operations](./docs/graph_crud.md)  
- [JPGQL: Foliage's JSON Path Graph Query Language](./docs/jpgql.md)  
- [Visual graph debugger](./docs/graph_debug.md)  
- [How to write a Foliage app](./docs/how_to_write_an_application.md)  
- [Performance measurement](./docs/performance_measures.md)

## 🧰 Technology Stack

Foliage runs on a modern and efficient tech stack:

### Backend

- NATS Jetstream  
- NATS KV Store  
- NATS WebSocket  
- GoLang  
- JavaScript (V8 Engine)

### Frontend

- React  
- TypeScript/JavaScript  
- WebSocket

### Common

- Docker  
- Docker Compose

🔍 [Learn about our technology choices](./docs/technologies_comparison.md)

---

## 🗺 Roadmap

Check out where we’re headed:  
![Roadmap](./docs/pics/Roadmap.jpg)

## 📎 References

- [Glossary](./docs/glossary.md)  
- [Code conventions](./docs/conventions.md)  
- [External API](./docs/external_api.md)

## 📄 License

Unless otherwise noted, the Foliage source files are distributed under the Apache Version 2.0 license found in the LICENSE file.

## 🤝 Contributing

We welcome your ideas, improvements, and feedback. Help us grow Foliage into a powerful tool for building intelligent, connected systems. Check the [issues](https://github.com/foliagecp/sdk/issues) and open a PR anytime!
