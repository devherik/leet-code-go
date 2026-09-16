# Software Engineering & Algorithm Practice (Go & Python)

This repository serves as a practical workbench for algorithmic problem solving, concurrency primitives, software design patterns, Clean Architecture, and domain modeling using **Go** and **Python**.

---

## 🎯 Focus Areas

- **Algorithmic Logic (LeetCode):** Solving complex data structures and algorithmic problems with an emphasis on time/space complexity and idiomatic implementations.
- **Concurrency & Thread Safety:** Synchronization mechanisms, race-condition mitigation (`sync.RWMutex`, `sync.WaitGroup`, Python `threading.Lock`, `asyncio.Lock`, and double-checked locking).
- **Software Design Patterns:** Behavioral and structural patterns in Go (e.g., Chain of Responsibility).
- **Clean Architecture & DDD:** Separation of concerns across Domain, Use Case, Interface/HTTP, and Agent layers.
- **Financial Systems:** Double-entry ledger engine maintaining balanced transactions and account aggregation across both Go and Python.

---

## 📁 Repository Structure

```
.
├── 01.go ... 49.go      # LeetCode solutions in Go
├── 3612.py              # LeetCode solutions in Python
├── mutex.go             # Go concurrency: thread-safe account & sync primitives
├── mutex.py             # Python concurrency: safe bank account & async cache (Thundering Herd protection)
│
├── chain/               # Chain of Responsibility pattern (Hospital workflow)
│   ├── department.go    # Handler interface
│   ├── reception.go     # Step 1: Patient registration
│   ├── doctor.go        # Step 2: Medical examination
│   ├── medical.go       # Step 3: Medication dispensing
│   ├── cashier.go       # Step 4: Billing & payment
│   └── main.go          # Pipeline execution
│
├── bigoproject/         # Clean Architecture & Domain-Driven Design
│   ├── domain/          # Entities and domain business rules (User)
│   ├── usecase/         # Application business rules and repositories
│   ├── http/            # HTTP transport / controllers
│   └── agent/           # Agent decision orchestrator
│
└── ledger/              # Double-entry bookkeeping engine
    ├── go/              # Go implementation with unit tests (main_test.go)
    └── python/          # Python implementation with pytest suite (test.py)
```

---

## 🧩 Modules & Implementations

### 1. LeetCode Solutions

| Problem | File | Language | Topic |
| :--- | :--- | :--- | :--- |
| **#1 Two Sum** | [`01.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/01.go) | Go | Hash Map / Complement lookup |
| **#2 Add Two Numbers** | [`02.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/02.go) | Go | Linked Lists / Elementary math |
| **#3 Longest Substring Without Repeating Characters** | [`03.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/03.go) | Go | Sliding Window / Hash Set |
| **#4 Median of Two Sorted Arrays** | [`04.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/04.go) | Go | Binary Search / Divide & Conquer |
| **#49 Group Anagrams** | [`49.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/49.go) | Go | Sorting & Hash Map grouping |
| **#3612 Process String** | [`3612.py`](file:///c:/Users/herik.rezende/Projetos/leet-code/3612.py) | Python | String manipulation & simulation |

### 2. Concurrency & Synchronization

- **Go ([`mutex.go`](file:///c:/Users/herik.rezende/Projetos/leet-code/mutex.go)):** Thread-safe bank account with read-write mutex (`sync.RWMutex`) and concurrent routine coordination via `sync.WaitGroup`.
- **Python ([`mutex.py`](file:///c:/Users/herik.rezende/Projetos/leet-code/mutex.py)):**
  - `SafeBankAccount`: Thread-safe balance management using `threading.Lock`.
  - `AsyncCache`: In-memory cache protected against the **Thundering Herd** problem using `asyncio.Lock` and double-checked locking.

### 3. Design Patterns (`chain/`)

Implements the **Chain of Responsibility** pattern modeling a hospital patient intake pipeline:
$$\text{Reception} \longrightarrow \text{Doctor} \longrightarrow \text{Medical} \longrightarrow \text{Cashier}$$

### 4. Clean Architecture (`bigoproject/`)

Enforces the Dependency Rule where outer layers depend strictly on inner abstractions:
- `domain`: Pure business entities and invariant validations (`User`).
- `usecase`: Orchestration of business actions and repository interfaces.
- `http`: Boundary translation between HTTP payloads and domain use cases.
- `agent`: Agent-driven decision routing.

### 5. Double-Entry Ledger (`ledger/`)

Financial ledger guaranteeing accounting invariants:
- **Balance Invariant:** $\sum \text{Debits} == \sum \text{Credits}$ for every transaction.
- **Implementations:** Available in both Go and Python with complete test suites.

---

## 🚀 How to Run

### LeetCode Solutions & Standalone Files

```bash
# Run a specific Go solution
go run 01.go

# Run Python solution
uv run python 3612.py
# or
python 3612.py
```

### Chain of Responsibility

```bash
cd chain
go run .
```

### Double-Entry Ledger

```bash
# Go implementation
cd ledger/go
go run main.go

# Python implementation
cd ledger/python
uv run python ledger.py
```

---

## 🧪 Running Tests & Quality Checks

### Go Tests & Quality Checks

```bash
# Run ledger unit tests
cd ledger/go
go test -v

# Code formatting & vet check across root
go fmt ./...
go vet ./...
```

### Python Tests

```bash
# Run pytest on the ledger test suite
uv run pytest ledger/python/test.py
```
