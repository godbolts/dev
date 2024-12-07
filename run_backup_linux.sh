#!/bin/bash
#Starting file for only Linux/Mac systems.

echo "Detecting OS..."

# Start the server
cd server
if [ -f start_server.sh ]; then
    bash start_server.sh &
else
    echo "Error: Missing start_server.sh"
    exit 1
fi
cd ..

# Start the client
cd client
if [ -f start_client.sh ]; then
    bash start_client.sh &
else
    echo "Error: Missing start_client.sh"
    exit 1
fi
cd ..
