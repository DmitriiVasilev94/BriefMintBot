# BriefMintBot 🍃

`BriefMintBot` is a personal Telegram assistant designed to aggregate daily essentials, process them using Large Language Models (LLMs), and deliver a concise, refreshing morning digest ("the mint briefing"). In addition it serves as personal English tutor to improve grammar and vocabulary.

This repository is built as a pet project to master a modern, enterprise-grade backend stack, containerization best practices, and Infrastructure as Code (IaC) within the AWS cloud environment.

---

## 🎯 Key Features

The bot automates your morning routine by combining three data segments into a single, comprehensive message:
1. **Curated News:** A condensed summary of the past 24 hours' global events, filtered via LLM to remove clickbait and noise.
2. **Financial Overview:** Stock market portfolio tracking based on custom user tickers.
3. **English Tutor Minute:** Embedded language maintenance featuring a couple of colloquial phrases for active vocabulary and a grammar rule for quick review.

---

## 🛠️ Tech Stack

Engineered following industry production standards:
* **Backend:** Go (Clean Architecture, idiomatic code structures)
* **Infrastructure as Code:** Terraform
* **Cloud Provider:** AWS (Amazon Web Services)
* **Containerization:** Docker (Multi-stage builds optimized for minimal image sizes)
* **Architecture Style:** Planned transition to either Serverless (AWS Lambda) or Containerized Orchestration (AWS ECS/Fargate) post-MVP.

---

## 📈 Architecture Tracking (ADR)

We value conscious engineering decisions. Major architectural trade-offs (e.g., on-demand vs scheduled ingestion, database choices, cloud layouts) are systematically recorded as **Architectural Decision Records** inside the docs/adr/ directory.

---

## 🚀 Quick Start (Local MVP)

### Prerequisites
* Go (1.21+ recommended)
* Docker

### Installation & Initialization
1. Clone the repository:
```
git clone https://github.com/DmitriiVasilev94/BriefMintBot.git
cd BriefMintBot
```
2. Initialize dependencies:
```
go mod tidy
```

3. Create a .env file in the project root and add your Telegram credentials:
```
TELEGRAM_APITOKEN=your_secret_token_here
```

4. Run the application locally:
```
go run cmd/bot/main.go
```