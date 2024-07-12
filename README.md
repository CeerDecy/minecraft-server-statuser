# Minecraft Server Statuser

## Description
This is a container image for `minecraft-server-operator` to get the status of a Minecraft server.

## How to use
1. Run on docker

    ```bash
    docker run -d --name minecraft-server-statuser --network host ceerdecy/minecraft-server-statuser:1.0.0
    ```
   
    the default port is 25566
    ```bash
   curl http://127.0.0.1:25566/status
    ```

2. Run on kubernetes
   
   coming soon...



