#!/bin/bash

# Function to check and install dependencies on Linux
install_linux_dependencies() {
  echo "Updating package list..."
  sudo apt-get update

  echo "Installing GCC and libc6-dev..."
  sudo apt-get install -y gcc libc6-dev

  echo "Setting CGO_ENABLED=1..."
  export CGO_ENABLED=1
  echo 'export CGO_ENABLED=1' >> ~/.bashrc
  source ~/.bashrc
}

# Function to set environment variable on Windows (Git Bash)
install_windows_dependencies() {
  echo "Setting CGO_ENABLED=1..."
  export CGO_ENABLED=1
  echo 'export CGO_ENABLED=1' >> ~/.bash_profile
  source ~/.bash_profile
}

# Detect OS and run appropriate function
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  echo "Detected Linux OS"
  install_linux_dependencies
elif [[ "$OSTYPE" == "msys" ]]; then
  echo "Detected Windows OS"
  install_windows_dependencies
else
  echo "Unsupported OS: $OSTYPE"
  exit 1
fi

# Clean Go cache and run the application
echo "Cleaning Go cache..."
go clean -cache

echo "Running Go application..."
go run main.go
