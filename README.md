# Summit 🛠️

A simple CLI tool that generates conventional commit messages using AI. Built because we've all stared at the `git commit` trying to figure out what to write.

**Contributions and Issues are welcome!**

<img width="1092" height="500" alt="carbon(1)" src="https://github.com/user-attachments/assets/9b707006-b732-4c59-aa89-fe4ff61949c4" />

## What it does

You type what you changed, it gives you a proper commit message, copied to your clipboard.

```bash
summit fixed login bug
⚡ fix(auth): resolve login validation bug
[Copied to clipboard]
```

## Installation

### Option 1: Quick Install (if you already have Go installed)

```bash
go install github.com/ekrlstd/summit@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/ekrlstd/summit.git
cd summit
go mod tidy
go build -o summit
./summit fixed login bug  # Use "./summit" when running locally
```

**Optional:** Now move to PATH for global access:

```bash
sudo mv summit /usr/local/bin/ # Or other directory that's in your $PATH
summit fixed login bug  # Now "summit" works from anywhere
```

## Dependencies

```
Go 1.13+
Internet connection (for API)
```

The tool uses a shared Groq API key that's baked in in the code, so you don't have to deal with any custom API keys.

## Usage

Just describe what you did:

```bash
summit added user authentication
⚡ feat(auth): add user authentication

summit updated docs
⚡ docs: update documentation

summit refactored the entire codebase
⚡ refactor: simplify codebase structure
```

The commit message gets copied to your clipboard automatically.

## Advanced (Optional)

### If you want to create your own API key you need to:

Set an environment variable:

```bash
export API_KEY="your-groq-api-key"
```

Or create a `.env` file:

```bash
API_KEY="your-groq-api-key"
```

Get a free key from [console.groq.com/keys](https://console.groq.com/keys) (no credit card needed).

### Want to customize how it writes commits?

Add `Prompt_prefix` to your `.env` file and customize to your liking.

## Tech Stack

- **Go** - Fast, and compiles to a single binary
- **Cobra** - CLI library
- **Groq API** - ultra-fast AI (Llama 3.3 70B, ~300 tokens/sec)

## Notes

- The tool uses a shared free API key. Please don't abuse it. If you're using this 100+ times a day, you might want to get your own key.
- Works best with actual code changes. Weird inputs might give weird results.
- The "generate" subcommand from older versions still works but is hidden. Just use `summit <description>` now.

## Why?

Because I've sat for way too long trying to create a conventional and "proffesional" commit message. Now, I can type whatever comes to mind, and let the AI do the rest.

---

**Please star this repo if you liked it!** ⭐
