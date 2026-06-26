# FrontDoor

FrontDoor is a small Go application that provides a simple homepage for your homelab or server.

Instead of remembering ports, URLs, or connection commands, FrontDoor gives you a single page where you can launch web services and quickly copy connection information for things like SMB shares.

The project is intentionally lightweight:

- Single Go binary
- JSON configuration
- No database
- No external dependencies
- Easy to deploy as a systemd service

## Features

- Launch web services from a single page
- Display SMB shares with click-to-copy connection instructions
- Configured through a single JSON file
- Runs as a Linux systemd service

## Example Configuration

```json
{
  "listen": ":8080",
  "services": [
    {
      "name": "Grafana",
      "url": "http://localhost:3000"
    },
    {
      "name": "Gitea",
      "url": "http://localhost:3001"
    }
  ],
  "shares": [
    {
      "name": "Media",
      "info": "smb://MyShare/media"
    },
    {
      "name": "data2",
      "info": "Connect using: smb://MyShare/data2",
      "instruction":"mount -t cifs //MyShare/data2 /mnt/data2"
    }
  ]
}
```

## Build

```bash
go build
```

## Run

```bash
frontdoor server
```

Then open your browser:

```
http://localhost:8080
```

## Example Use Cases

- Home lab dashboard
- Development server homepage
- Small office server portal
- Personal infrastructure landing page
