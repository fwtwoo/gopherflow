# GopherFlow 🛠️

![GopherFlow Usage Example](/public/example.png)

A fast and intuitive CLI tool that generates **high-quality, industry-standard commit messages** using the Groq API with Llama 3.3 70B 🚀

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square) [![Go](https://img.shields.io/badge/Go-1.13-blue?style=flat-square&logo=go)](https://go.dev/)

## ✨ Features

- 🧠 **AI-Powered Commit Messages** – Get professional-grade commits based on your description
- ⚡ **Lightning Fast** – Powered by Groq's ultra-fast inference (300+ tokens/sec)
- 🎯 **Minimal & Focused UX** – Built for speed, clarity, and efficiency
- 💯 **100% Free** – No credit card required, no usage fees

## 🚀 Getting Started

### Prerequisites

- Go 1.13+
- Internet connection (for API access)

### Installation

#### Clone the repository
```bash
git clone https://github.com/ebbekarlstad/gopherflow.git
```

#### Navigate to project directory
```bash
cd gopherflow
```

#### Install GitHub dependencies
```bash
go install .
```

#### Build the project
```bash
go build -o gopherflow
```

#### Create an .env file with your Groq API key

1. Get your free API key from [Groq Console](https://console.groq.com/keys)
2. Create a `.env` file in the project root:

```bash
API_KEY="your-groq-api-key-here"

Prompt_prefix="You are a Conventional Commit expert. Generate a commit message that follows strict conventional commit standards..."
```

**Resources:**
- [Get Groq API Key (Free)](https://console.groq.com/keys)
- [Prompt Engineering Guide](https://www.atlassian.com/blog/artificial-intelligence/ultimate-guide-writing-ai-prompts)

### Usage

Start the tool by running:
```bash
gopherflow generate
```

---

## 🛠️ Tech Stack

- 🐹 **Go (Golang)** – Fast, efficient, compiled language
- 🖥️ **Cobra** – Modern Go library for creating CLI applications
- 🧠 **Groq API** – Ultra-fast AI inference using Llama 3.3 70B
- ⚡ **300+ tokens/sec** – Industry-leading inference speed

---

#### ⭐ Star this repo if you find it helpful!
