# 🚇 Tunnix

A lightweight, self-hostable alternative to NGROK written in Go. Tunnix creates secure SSH tunnels to expose your local services to the internet through your own server.

**🔗 Project URL**: [https://github.com/yashraj-n/tunnix](https://github.com/yashraj-n/tunnix)

[![Go Version](https://img.shields.io/badge/Go-1.23.4+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](docker/)
[![GitHub](https://img.shields.io/badge/GitHub-yashraj--n%2Ftunnix-blue.svg)](https://github.com/yashraj-n/tunnix)

## ✨ Features

- **Self-hosted**: Complete control over your tunnel infrastructure
- **SSH-based**: Secure tunneling using SSH protocol
- **Lightweight**: Minimal resource footprint
- **Docker-ready**: Easy deployment with Docker Compose
- **Cross-platform**: Works on Windows, macOS, and Linux
- **Simple CLI**: Easy-to-use command-line interface
- **Real-time logging**: Colored, structured logging with slog

## 🏗️ Architecture

Tunnix consists of two main components:

1. **Client** (`main.go`): Connects to your remote server and creates SSH tunnels
2. **Server** (Docker container): SSH server that accepts connections and forwards traffic

```
┌─────────────┐    SSH Tunnel    ┌─────────────┐    HTTP/HTTPS    ┌─────────────┐
│   Your App  │ ───────────────► │ Tunnix      │ ───────────────► │   Internet  │
│  (localhost)│                  │ Server      │                  │             │
└─────────────┘                  └─────────────┘                  └─────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23.4 or higher (for building from source)
- Docker and Docker Compose (for server deployment)
- A VPS or cloud server with public IP

### 📥 Downloads

Pre-built binaries are available for all major platforms:
Download from the [releases page](https://github.com/yashraj-n/tunnix/releases).

### 1. Deploy the Server

```bash
# Clone the repository
git clone https://github.com/yashraj-n/tunnix.git
cd tunnix

# Navigate to docker directory
cd docker

# Create environment file
cat > .env << EOF
SSH_API_KEY=your-secure-password-here
TUNNIX_FQDN=your-domain.com
EOF

# Start the server
docker-compose up -d
```

### 2. Run the Client

#### Option A: Using Pre-built Binary (Recommended)

```bash
# Download and run (Linux/macOS)
chmod +x tunnix-linux-amd64
./tunnix-linux-amd64 -host your-server-ip -password your-secure-password-here -port 8080

# Windows
tunnix-windows-amd64.exe -host your-server-ip -password your-secure-password-here -port 8080
```

#### Option B: Build from Source

```bash
# Build the client
go build -o tunnix main.go

# Run the client
./tunnix -host your-server-ip -password your-secure-password-here -port 8080
```

### 3. Access Your Service

Your local service running on port 8080 will now be accessible at:
```
http://your-domain.com:12001
```

## 📖 Usage

### Command Line Options

```bash
./tunnix [flags]

Flags:
  -host string
        Address to remote Server (required)
  -password string
        Password to use for SSH authentication (required)
  -port int
        Local port to forward (default 8080)
```

### Examples

```bash
# Forward local web server
./tunnix -host 192.168.1.100 -password mypassword -port 3000

# Forward local API
./tunnix -host my-server.com -password mypassword -port 8000

# Forward local database (not recommended for production)
./tunnix -host my-server.com -password mypassword -port 5432
```

## 🔧 Configuration

### Server Configuration

The server runs on:
- **SSH Port**: 12000 (for client connections)
- **Tunnel Port**: 12001 (for forwarded traffic)

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `SSH_API_KEY` | Password for SSH authentication | Yes |
| `TUNNIX_FQDN` | Domain name for your tunnel | Yes |

### SSH Configuration

The server uses a custom SSH configuration (`docker/sshd_config`) with:
- Password authentication enabled
- TCP forwarding allowed
- Gateway ports enabled
- Security restrictions applied

## 🐳 Docker Deployment

### Using Docker Compose

```yaml
services:
  tunnix:
    build: .
    container_name: tunnix-ssh-server
    ports:
      - "12000:12000"  # SSH port
      - "12001:12001"  # Tunnel port
    environment:
      - SSH_API_KEY=${SSH_API_KEY}
      - TUNNIX_FQDN=${TUNNIX_FQDN}
    restart: unless-stopped
```

### Manual Docker Build

```bash
# Build the image
docker build -t tunnix .

# Run the container
docker run -d \
  --name tunnix-server \
  -p 12000:12000 \
  -p 12001:12001 \
  -e SSH_API_KEY=your-password \
  -e TUNNIX_FQDN=your-domain.com \
  tunnix
```

## 🔒 Security Considerations

⚠️ **Important Security Notes:**

1. **Change Default Credentials**: Always use a strong password for `SSH_API_KEY`
2. **Firewall Configuration**: Only expose necessary ports (12000, 12001)
3. **HTTPS**: Consider using a reverse proxy (nginx/traefik) for HTTPS termination
4. **Network Isolation**: Run the server in a private network when possible
5. **Regular Updates**: Keep the server and dependencies updated

### Recommended Security Setup

```bash
# Use UFW firewall
ufw allow 12000/tcp
ufw allow 12001/tcp
ufw enable

# Use strong passwords
SSH_API_KEY=$(openssl rand -base64 32)
```

## 📦 Releases

Pre-built binaries are automatically generated for each release. You can:

- **Download the latest release**: [Latest Release](https://github.com/yashraj-n/tunnix/releases/latest)
- **View all releases**: [Releases Page](https://github.com/yashraj-n/tunnix/releases)
- **Build from source**: See Development section below

### Supported Platforms

- ✅ Windows (amd64)
- ✅ macOS (amd64)
- ✅ Linux (amd64)

## 🛠️ Development

### Project Structure

```
tunnix/
├── config/          # Configuration management
├── docker/          # Docker server setup
├── network/         # Network connection handling
├── ssh/            # SSH client and tunnel logic
├── main.go         # Client entry point
├── go.mod          # Go module file
└── README.md       # This file
```

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yashraj-n/tunnix.git
cd tunnix

# Install dependencies
go mod download

# Build the client
go build -o tunnix main.go

# Run tests (if available)
go test ./...
```

### Dependencies

- `golang.org/x/crypto/ssh`: SSH client implementation
- `github.com/lmittmann/tint`: Colored logging
- `github.com/elliotchance/sshtunnel`: SSH tunnel utilities

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by [NGROK](https://ngrok.com/) and similar tunneling services
- Built with Go's excellent SSH library
- Docker configuration based on Alpine Linux for minimal footprint

## ⚠️ Disclaimer

This tool is provided as-is for educational and development purposes. Use at your own risk in production environments. Always follow security best practices when exposing services to the internet.

---

**Made with ❤️ by [yashraj-n](https://github.com/yashraj-n)**

If you find this project useful, please consider giving it a ⭐ on GitHub!
