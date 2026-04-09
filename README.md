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

### Manual setup 
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

## PR Tracking
Example:
```bash
go run main.go trackPr -u safishamsi -r graphify --state merged
```
output
```bash
→ ✅ Add README example for using graph.json with an... gdesai23
→ ✅ docs: add Japanese README                          eltociear
→ ✅ fix: git hooks fail when graphify is installed ... Minidoracat
→ ✅ feat: add Trae and Trae CN platform support        ljinshuan
→ ✅ Fix #19: Suppress ANSI escape codes breaking Po... azizur100389
→ ✅ security: SSRF private IP blocking + Neo4j Cyph... bperkins-oss
→ ✅ feat: add Swift language support                   douglance
→ ✅ docs: add Simplified Chinese README                cyxer000
```
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

_Yeah...its that simple_

---
## Screenshots
<img width="1039" height="201" alt="Screenshot 2026-04-08 203006" src="https://github.com/user-attachments/assets/2868e25a-beab-42b2-9887-e7143c4d40b2" />

<img width="1362" height="238" alt="image" src="https://github.com/user-attachments/assets/27448ae2-8d4b-420b-bf01-2cfc31f6d6fb" />

<img width="1338" height="684" alt="image" src="https://github.com/user-attachments/assets/32cf6a91-16d1-40d9-b2aa-a4dc70ec8719" />
<sub>Used `graphify` repo as examples here.</sub>

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


<sub>uses the `MIT` LICENSE</sub>
