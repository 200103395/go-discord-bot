# Discord Bot Go

## Introduction

Discord Bot Go is a simple Discord bot written in Go. This bot serves as a foundation for creating and customizing your own Discord bot to enhance your server experience.

## Getting Started

### 1. Create a Discord Application

- Go to the [Discord Developer Portal](https://discord.com/developers/applications).
- Click on "New Application" to create a new Discord application.
- Navigate to the "Bot" tab and click "Add Bot" to add a bot to your application.

### 2. Invite the Bot to Your Server

- On the "Bot" tab, find the "OAuth2" section.
- In the "OAuth2 URL Generator" section, select the "bot" scope and the necessary permissions.
- Copy the generated URL and paste it into your browser. Follow the instructions to invite the bot to your server.

### 3. Retrieve Bot Token

- In the "Bot" tab of your application, find the "TOKEN" section. Copy the token.

### 4. Configure .env file

- Create a file named `.env` in the project root.
- Add the following information to the file:
  ```env
  BOT_TOKEN=YOUR_BOT_TOKEN
  BOT_PREFIX=YOUR_BOT_PREFIX
  ```
  Replace YOUR_BOT_TOKEN with the bot token you copied earlier.
  Replace YOUR_BOT_PREFIX with any prefix you want for you bot commands

### 5. Build and Run:
- Make sure you have Go installed on your machine.
- Open a terminal in the project directory.
- Run the bot using:
  ```bash
  go run main.go
- The bot should now be online and responsive to the specified command prefix.
### 6. Customization
Feel free to customize the bot to fit your needs. You can add new commands, implement additional features, and modify the behavior according to your server's requirements.

### 7. Contributing
If you'd like to contribute to this project, feel free to fork the repository, make your changes, and submit a pull request. Your contributions are welcome!

### 8. License
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
