# 🚀 gitrack

> **GitHub activity, without the noise.**

A fast, minimal CLI to track GitHub user activity in real time — cleanly filtered, colorized, and actually readable.

---

## ✨ Features

* 🎯 Filter events by type (`push`, `pr`, `issues`, etc.)
* ⚡ Real-time GitHub activity tracking
* 🎨 Color-coded output for quick scanning
* 🧩 Clean, minimal CLI UX (no clutter, no nonsense)

---

## 📦 Installation
### Windows
 - Download the `.exe` file. 
 - Add to Path

### Manual setup (recommended, due to cross platform)
```bash
git clone https://github.com/yourusername/gitrack
cd gitrack
go build -o gitrack
```

(Optional: move it to your PATH)

```bash
mv gitrack /usr/local/bin
```

---

## ⚡ Usage

### Track a user

```bash
gitrack track --user torvalds
```

### Filter by event type

```bash
gitrack track --user torvalds --type PushEvent
```

### Friendly aliases 

```bash
gitrack track -u torvalds -t push
gitrack track -u torvalds -t pr
gitrack track -u torvalds -t issues
```

---

## 🎯 Supported Event Types

| CLI Input      | GitHub Event        |
| -------------- | ------------------- |
| `push`         | `PushEvent`         |
| `watch` / star | `WatchEvent`        |
| `pr` / pull    | `PullRequestEvent`  |
| `issues`       | `IssuesEvent`       |
| `comment`      | `IssueCommentEvent` |
| `create`       | `CreateEvent`       |
| `release`      | `ReleaseEvent`      |

---

## 🖥 Example Output

```text
→ Starred | torvalds/linux at 12:04 PM
→ PR Opened | microsoft/vscode at 12:06 PM
→ Issue | golang/go at 12:08 PM
```

---

## 🧠 Why gitrack?

GitHub’s default activity feeds are noisy, inconsistent, and hard to scan.

**gitrack fixes that.**

* Only see what you care about
* Clean output you can actually read
* Built for developers who like things simple

---

## 🛠 Tech Stack

* Go (Golang)
* Cobra (CLI framework)
* GitHub Events API
_ Yeah...its that simple_

---

## 📌 Roadmap


* [ ] Live streaming mode
* [ ] Config file support
* [ ] Custom UI (textual-style interface)

---

## 🤝 Contributing

PRs welcome. If you’ve got ideas to make it cleaner or faster — go for it, it's open source for a reason!

---

## ⭐ Support
If you like the project, drop a star ⭐
It helps more than you think.
If you like the project, drop a star ⭐
It helps more than you think.
