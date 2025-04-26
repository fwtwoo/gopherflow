# GopherFlow 🛠️

![GopherFlow Usage Example](/public/example.png)

A fast and intuitive CLI tool that generates **high-quality, industry-standard commit messages** using the DeepSeek API from OpenRouter.ai 🚀

![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square) [![Go](https://img.shields.io/badge/Go-1.13-blue?style=flat-square&logo=go)](https://go.dev/)

## ✨ Features

-   🧠 **AI-Powered Commit Messages** – Get professional-grade commits based on your description
-   🎯 **Minimal & Focused UX** – Built for speed, clarity, and efficiency
    

## 🚀 Getting Started

### Prerequisites

-   Go 1.13+
-   Internet connection (for API access)
    

### Installation
#### Clone the repository
```bash
git clone https://github.com/ebbekarlstad/gopherflow.git
```

#### Navigate to project directory

```bash
cd gopherflow
```

####  Install GitHub dependencies

```bash
go install .
```

####  Build the project

```bash
go build -o gopherflow
```

####  Create an .env file that contains a Prompt Prefix and a Deepseek API key:
Here are some resources for this
```bash
https://www.atlassian.com/blog/artificial-intelligence/ultimate-guide-writing-ai-prompts
https://openrouter.ai/deepseek/deepseek-v3-base:free
```

### Usage
Start the tool by running
```bash
gopherflow generate
```
---



## 🛠️ Tech Stack

-   🐹 **Go (Golang)** – Fast, efficient, compiled language
- 🖥️ **Cobra** – Modern Go library for creating CLI applications
    
-   🧠 **LLM API Integration** – AI text generation using OpenRouter.ai
    

----------

#### ⭐ Star this repo if you find it helpful!