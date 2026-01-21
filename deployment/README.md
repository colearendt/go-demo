# Deployment Guide

This guide shows how to deploy go-demo as a systemd service on Linux.

## Installation

1. **Download the binary** from the [releases page](https://github.com/colearendt/go-demo/releases) or build it yourself with embedding enabled:
   ```bash
   # Build with version info
   VERSION=$(git describe --tags --always --dirty)
   go build -tags embed -ldflags "-X main.version=${VERSION}"
   ```

   Check the version:
   ```bash
   ./go-demo -version
   ```

2. **Create the application directory**:
   ```bash
   sudo mkdir -p /opt/go-demo
   sudo mkdir -p /var/log/go-demo
   ```

3. **Copy the binary**:
   ```bash
   sudo cp go-demo /opt/go-demo/
   sudo chmod +x /opt/go-demo/go-demo
   ```

4. **Set ownership** (using www-data user, adjust as needed):
   ```bash
   sudo chown -R www-data:www-data /opt/go-demo
   sudo chown -R www-data:www-data /var/log/go-demo
   ```

5. **Install the systemd service**:
   ```bash
   sudo cp systemd/go-demo.service /etc/systemd/system/
   sudo systemctl daemon-reload
   ```

6. **Enable and start the service**:
   ```bash
   sudo systemctl enable go-demo
   sudo systemctl start go-demo
   ```

## Management

Check service status:
```bash
sudo systemctl status go-demo
```

View logs:
```bash
sudo journalctl -u go-demo -f
```

Restart the service:
```bash
sudo systemctl restart go-demo
```

Stop the service:
```bash
sudo systemctl stop go-demo
```

## Configuration

The service file is configured to:
- Run on port 8000 at 127.0.0.1
- Use embedded files (`-embedded` flag)
- Auto-restart on failure
- Run as `www-data` user
- Include security hardening options

To modify the configuration, edit `/etc/systemd/system/go-demo.service` and reload:
```bash
sudo systemctl daemon-reload
sudo systemctl restart go-demo
```

## Reverse Proxy

For production use, put this behind a reverse proxy like nginx or caddy:

**nginx example** (`/etc/nginx/sites-available/go-demo`):
```nginx
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```
