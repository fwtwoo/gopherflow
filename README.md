# GopherFlow 🛠️

![GopherFlow Usage Example](/public/example.png)

A fast and intuitive CLI tool that generates **high-quality, industry-standard commit messages** using AI ⚡

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square) [![Go](https://img.shields.io/badge/Go-1.13-blue?style=flat-square&logo=go)](https://go.dev/)

## ✨ Features

-   🧠 **AI-Powered Commit Messages** – Get professional-grade commits in seconds
-   ⚡ **Lightning Fast** – Powered by Groq's ultra-fast inference (300+ tokens/sec)
-   🎯 **Minimal UX** – Just type your description and go
-   💯 **100% Free** – No credit card or API key setup required
-   📋 **Auto-Copy** – Commit message copied to clipboard automatically
-   ✨ **Zero Setup** – Works out of the box, no configuration needed

## 🚀 Getting Started

### Prerequisites

-   Go 1.13+
-   Internet connection (for API access)

### Installation

#### Install via Go (Easiest)
```bash
go install github.com/fwtwoo/gopherflow@latest
```

#### Or Clone and Build
```bash
git clone https://github.com/fwtwoo/gopherflow.git
cd gopherflow
go mod tidy
go build -o gopherflow
```

**✨ That's it! No API key setup, no configuration files needed.**

### Usage

Simply describe your changes:
```bash
gopherflow fixed login bug
```

Output:
```
⚡ fix(auth): resolve login validation bug
[Copied to clipboard]
```

More examples:
```bash
gopherflow added user profile page
⚡ feat(profile): add user profile page
[Copied to clipboard]

gopherflow updated README documentation
⚡ docs(readme): update documentation
[Copied to clipboard]

gopherflow refactored database connection logic
⚡ refactor(database): simplify connection logic
[Copied to clipboard]
```

---

## 🔧 Advanced Usage (Optional)

### Use Your Own API Key

Want to use your own Groq API key? Set the environment variable:
```bash
export API_KEY="your-groq-api-key"
gopherflow fixed login bug
```

Or create a `.env` file:
```bash
API_KEY="your-groq-api-key"
```

Get your free API key from [console.groq.com](https://console.groq.com/keys)

### Custom Prompt Prefix

You can customize the AI's behavior by setting the `Prompt_prefix` environment variable in a `.env` file.

---

## 🛠️ Tech Stack

-   🐹 **Go (Golang)** – Fast, efficient, compiled language
-   🖥️ **Cobra** – Modern Go library for creating CLI applications
-   🧠 **Groq API** – Ultra-fast AI inference using Llama 3.3 70B
-   ⚡ **300+ tokens/sec** – Industry-leading inference speed

---

## 📝 Note

This tool uses a shared Groq API key for convenience. Please be respectful with usage. The free tier has generous rate limits, but if you plan to use it very heavily, consider using your own free API key from [console.groq.com](https://console.groq.com/keys).

---

## 📄 License

MIT License - see LICENSE file for details

---

#### ⭐ Star this repo if you find it helpful!
