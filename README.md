# OpenKommander

OpenKommander is a command line utility and admin UI for Apache Kafka compatible brokers.

## Prerequisites

- [Podman](https://podman.io/getting-started/installation) (required for running the development environment)
- [Make](https://www.gnu.org/software/make/) (required for running development commands)

## Development Environment Setup

1. **Install Podman**
   Follow the installation instructions for your operating system on the [Podman website](https://podman.io/getting-started/installation).

2. **Clone the repository**
   ```bash
   git clone https://github.com/IBM/openkommander.git
   cd openkommander
   ```

3. **Start the development environment**
   ```bash
   make dev
   ```

4. **Execute into the container**
    ```bash
    podman exec -it openkommander-app-1 bash
    ```

    Note: Replace `openkommander-app-1` in case your is different

5. **Build and install cli**
    ```bash
    make dev-run
    ```

## Configuration

The application uses a configuration file located at `config/config.yaml`. By default, it is configured for the development environment:

```yaml
kafka:
  broker: kafka:9093
```

### Custom Configuration

You can modify `config/config.yaml` to connect to different Kafka clusters:

```yaml
# Development environment (default)
kafka:
  broker: kafka:9093

# Custom environment example
kafka:
  broker: localhost:9092  # For local Kafka installation
  # broker: kafka-cluster.example.com:9093  # For remote cluster
```

The configuration file is loaded when the application starts. If you need to connect to a different broker after starting the application, you can use the `ok login` command with a custom broker address:

```bash
$ ok login
Enter broker address [kafka:9093]: localhost:9092
```

## CLI Usage

After running the application, you can use the following commands:


### Commands

All commands start with prefix `ok`

| Command | Description | Arguments |
|---------|-------------|-----------|
| `login` | Connect to a Kafka cluster | None | ok login |
| `logout` | End the current session | None | ok logout |
| `session` | Display current session information | None | ok session |
| `metadata` | Display cluster information | None | ok metadata |
| `rest` | Display rest server information | brokers | ok rest --brokers kafka:9093 |
| `help` | Display available commands | None | 
ßßß

### Example Workflow

1. Build the CLI:
   ```bash
   make build
   ```

2. Install the CLI:
   ```bash
   make install
   ```

   Note: It may fail due to permission if needed add `sudo` for example
   ```bash
   sudo make install
   ```

3. Connect to the cluster:
   ```bash
   $ ok login
   Connected to Kafka cluster at kafka:9093
   ```

4. View session information:
   ```bash
   $ ok session
   Current session: Brokers: [kafka:9093], Authenticated: true
   ```

5. View cluster information:
   ```bash
   $ ok metadata
   Cluster Brokers:
    - kafka:9093 (ID: 1)
   ```

6. Start REST server:
   ```bash
   $ ok rest --brokers kafka:9093
   ```

7. End session and exit:
   ```bash
   $ ok logout
   Logged out successfully!
   $ ok exit
   Exiting application.
   ```

### Rest Endpoints

# List topics
curl -X GET http://localhost:8080/topics

# Create a topic
curl -X POST http://localhost:8080/topics \
  -H "Content-Type: application/json" \
  -d '{"name":"my-topic","partitions":2,"replication_factor":1}'

# Delete a topic
curl -X DELETE http://localhost:8080/topics/my-topic

# Broker status
curl -X GET http://localhost:8080/api/v1/status

# Broker management
curl -X GET http://localhost:8080/api/v1/brokers