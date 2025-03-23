# 🗨️ Go WebSocket Chat Server

A WebSocket-based chat server written in Go that supports **private**, **group**, and **broadcast** messaging.

## 🚀 Features
- **Private Chat**: One-on-one messaging.
- **Group Chat**: Chat within a group.
- **Broadcast Messages**: Send messages to all connected users.
- **Auto Restart**: Recovers from crashes automatically.

---

## 📦 Installation

Make sure you have **Go** installed. Then, clone this repository:

```sh
git clone https://github.com/quydmfl/go-translate-chat.git
cd go-translate-chat
```

## 🔧 Build & Run

### ✅ **Build the Server**
For different operating systems:

```sh
# Linux/macOS
make build

# Windows
go build -o bin/chat-server.exe main.go
```

### ▶ **Run the Server**
```sh
make run
```

If port `8080` is in use, free it before running:
```sh
make clean
```

---

## 🛠️ Makefile Commands

| Command  | Description |
|----------|------------|
| `make build` | Builds the server binary. |
| `make run` | Runs the server. |
| `make clean` | Stops the server and clears port `8080`. |

---

## 📱 WebSocket API

### 1⃣ Connect to WebSocket
Clients can connect using:

```
ws://localhost:8080/ws
```

### 2⃣ Sending Messages
Messages must be JSON-formatted:

#### **🔹 Private Chat**
Send a message to a specific user:
```json
{
  "type": "private",
  "sender": "user1",
  "target": "user2",
  "message": "Hello, user2!"
}
```

#### **🔸 Group Chat**
Send a message to a group:
```json
{
  "type": "group",
  "sender": "user1",
  "target": "group1",
  "message": "Hello group1!"
}
```

#### **📢 Broadcast Message**
Send a message to all users:
```json
{
  "type": "broadcast",
  "sender": "user1",
  "message": "Hello everyone!"
}
```

### 3⃣ Receiving Messages
Clients will receive messages in the same JSON format.

---

## 🔄 Auto Recovery on Crash
If the server crashes, it will automatically restart after 5 seconds.

You can manually stop it using:
```sh
make clean
```

---

## 🛠️ Development

- Run the server manually:
  ```sh
  go run main.go server
  ```
- Modify and test changes.

---

## 🐜 License
This project is licensed under the MIT License.

🚀 **Happy Coding!**
