# Minecraft Server Statuser

## Description
This is a container image for `minecraft-server-operator` to get the status of a Minecraft server.

It is only used to monitor the Minecraft server locally or within the same Kubernetes Pod, and the port must be open on 25565.

[DockerHub](https://hub.docker.com/repository/docker/ceerdecy/minecraft-server-statuser)

## Features
- Monitor Minecraft server status locally or within the same Pod.
- Easy to deploy with Docker and Kubernetes.

## Prerequisites
- Docker installed on your local machine.
- (For Kubernetes deployment) Kubernetes cluster up and running.

## How to use
1. Run on docker

    ```bash
    docker run -d --name minecraft-server-statuser --network host ceerdecy/minecraft-server-statuser:1.0.0
    ```

   By default, the status server will be accessible on port 25566. You can check the Minecraft server status with:
    ```bash
   curl http://127.0.0.1:25566/status
    ```

2. Run on kubernetes
   
   coming soon...

## Response Data
```json
{
  "online": true,
  "host": "127.0.0.1",
  "port": 25565,
  "ip_address": "127.0.0.1",
  "retrieved_at": 1720770854853,
  "expires_at": 1720770914853,
  "srv_record": null,
  "version": {
    "name_raw": "1.20.3",
    "name_clean": "1.20.3",
    "name_html": "\u003cspan\u003e\u003cspan\u003e1.20.3\u003c/span\u003e\u003c/span\u003e",
    "protocol": 765
  },
  "players": {
    "online": 0,
    "max": 20,
    "list": []
  },
  "motd": {
    "raw": "A Minecraft Server",
    "clean": "A Minecraft Server",
    "html": "\u003cspan\u003e\u003cspan\u003eA Minecraft Server\u003c/span\u003e\u003c/span\u003e"
  },
  "icon": null,
  "mods": [],
  "software": null,
  "plugins": []
}
```

## License

This project is licensed under the MIT License.

