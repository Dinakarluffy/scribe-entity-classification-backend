# Entity & Classification Service – Backend (Go)

This repository contains the **Go backend implementation** for the **Entity & Classification Service**.

The service exposes REST APIs to analyze transcript text and return:
- Named entities (people, tools, brands, products, companies)
- Tone classification
- Style detection
- Safety / sensitive content flags

> ⚠️ Note: PostgreSQL integration is pending approval.  
> Current implementation uses **in-memory storage** for API validation and testing.

---

## 1. Tech Stack

- **Language:** Go
- **Framework:** net/http + Gorilla Mux
- **Storage:** In-memory (temporary)
- **Future DB:** PostgreSQL (awaiting credentials)
- **API Testing:** Postman / curl

---

## 2. Project Structure

```
backend/
│
├── cmd/server/              # Application entry point
│   └── main.go
│
├── config/                  # Environment configuration
│   └── config.go
│
├── handlers/                # HTTP handlers (controllers)
│   └── scribe_entity_classification_handler.go
│
├── middleware/              # Middleware (logging)
│   └── logger.go
│
├── models/                  # Data models
│   └── scribe_entity_classification.go
│
├── repository/              # In-memory repository (temporary)
│   └── scribe_entity_classification_repo.go
│
├── routes/                  # API routes
│   └── routes.go
│
├── go.mod
├── go.sum
└── README.md
```

---

## 3. Prerequisites (Requirements)

Before running the project, ensure you have:

### Required
- **Go** (version 1.20 or higher recommended)
- **Git**
- **Postman** (for API testing)

### Optional (for future steps)
- PostgreSQL (credentials will be provided later)
- Gemini API key (for Python worker integration)

---

## 4. Environment Variables

Create a `.env` file (optional for now).

Example `.env.example`:

```env
PORT=8080
GEMINI_API_KEY=your_gemini_api_key_here
DATABASE_URL=postgresql://user:password@host:port/dbname
```

> ⚠️ Do NOT commit `.env` files to GitHub.

---

## 5. How to Run the Backend

### Step 1: Navigate to backend directory

```bash
cd backend
```

### Step 2: Install dependencies

```bash
go mod tidy
```

### Step 3: Start the server

```bash
go run ./cmd/server
```

### Expected Output

```
Server running on port 8080
```

---

## 6. API Endpoints

### 6.1 Health Check

**Method:** GET  
**URL:**

```
http://localhost:8080/health
```

**Response:**

```
OK
```

---

### 6.2 Analyze Entity & Classification

**Method:** POST  
**URL:**

```
http://localhost:8080/api/entity-classification/analyze
```

**Headers:**

```
Content-Type: application/json
```

**Request Body (JSON):**

```json
{
  "transcript_id": "11111111-1111-1111-1111-111111111111",
  "transcript_text": "This is an interview with Tim Ferriss about productivity and tools like Notion.",
  "creator_id": "22222222-2222-2222-2222-222222222222"
}
```

**Response Example:**

```json
{
  "status": "success",
  "analysis_id": "generated-uuid",
  "entities": {
    "people": ["Tim Ferriss"],
    "tools": ["Notion"],
    "brands": [],
    "products": ["The 4-Hour Workweek"],
    "companies": []
  },
  "tone": {
    "primary": "conversational",
    "secondary": "educational",
    "confidence": 0.88
  },
  "style": {
    "primary": "interview",
    "confidence": 0.92
  },
  "safety_flags": {
    "sensitive_domains": ["mental health"],
    "severity": "low",
    "requires_review": false
  }
}
```

---

### 6.3 Get Analysis Result

**Method:** GET  
**URL:**

```
http://localhost:8080/api/entity-classification/results/{analysis_id}
```

Replace `{analysis_id}` with the value returned from the POST API.

**Response Example:**

```json
{
  "analysis_id": "generated-uuid",
  "transcript_id": "11111111-1111-1111-1111-111111111111",
  "creator_id": "22222222-2222-2222-2222-222222222222",
  "entities": {},
  "tone": {},
  "style": {},
  "safety_flags": {},
  "created_at": "2026-01-17T16:45:00Z"
}
```

---

## 7. How to Paste Postman Outputs in README / Docs

When documenting Postman outputs:

1. Copy the **raw JSON response**
2. Paste it inside triple backticks
3. Specify `json` for formatting

Example:

````md
```json
{
  "status": "success",
  "analysis_id": "uuid",
  "entities": {}
}
```
````

This keeps the output clean and readable in GitHub.

---

## 8. Current Status

- ✅ Backend API endpoints implemented
- ✅ API contracts validated via Postman
- ✅ In-memory storage used for testing
- ⏳ PostgreSQL integration pending approval
- ⏳ Python worker (Gemini) integration pending

---

## 9. Next Steps

Once PostgreSQL credentials are provided:
- Replace in-memory repository with PostgreSQL
- Persist analysis results
- Enable full end-to-end testing

---

## 10. Maintainer

**Dinakaran S**  
Backend – Entity & Classification Service
