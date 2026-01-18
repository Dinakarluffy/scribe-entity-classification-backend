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
- **Storage:** In-memory (temporary)
- **Future DB:** PostgreSQL (awaiting credentials)
- **API Testing:** Postman

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

### 6.1 Health Check (Backend Validation Checkpoint)

This endpoint is used to verify that the backend server is running and able to handle **GET requests** correctly.

**Checkpoint Purpose:**
- Confirms the server has started successfully
- Confirms routing is configured correctly
- Confirms GET endpoints are accessible
- Used as a basic backend availability check

**Method:** GET  
**URL:**

```
http://localhost:8080/health
```

**Response:**

```
OK
```

**Checkpoint Status:**
- If response is `OK` → ✅ Backend is running and reachable
- If response is `404` → ❌ Route not registered
- If response is `405` → ❌ HTTP method mismatch
- If server does not respond → ❌ Backend not running

This checkpoint must pass before testing other API endpoints.

---

### 6.2 Analyze Entity & Classification (POST API Validation Checkpoint)

This endpoint analyzes a transcript and returns extracted entities, tone, style, and safety flags.

**Checkpoint Purpose:**
- Validates POST request handling
- Confirms request payload parsing
- Confirms API response structure matches PRD
- Confirms backend logic execution (in-memory storage)

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

## 7. Current Status

- ✅ Backend API endpoints implemented
- ✅ API contracts validated via Postman
- ✅ In-memory storage used for testing
- ⏳ PostgreSQL integration pending approval
- ⏳ Python worker (Gemini) integration pending

---

## 8. Next Steps

Once PostgreSQL credentials are provided:
- Replace in-memory repository with PostgreSQL
- Persist analysis results
- Enable full end-to-end testing

---

## 9. Frontend (Vite + React)

This repository also includes a **React-based frontend** used to interact with the Entity & Classification backend APIs.

The frontend provides:

* Transcript and file upload UI
* Results visualization in tabular format
* Manual result lookup using `analysis_id`
* Basic validations and error handling

> ⚠️ Note:
> The frontend currently consumes backend APIs backed by **in-memory storage**.
> Full persistence will be enabled after PostgreSQL approval.

---

## 9.1 Frontend Tech Stack

* **Framework:** React
* **Build Tool:** Vite
* **Language:** TypeScript
* **Routing:** React Router DOM
* **HTTP Client:** Axios / Fetch


---

## 9.2 Frontend Features

### Pages Implemented

| Page             | Route            | Description                                                |
| ---------------- | ---------------- | ---------------------------------------------------------- |
| Home             | `/`              | Entry page with navigation                                 |
| Upload           | `/upload`        | Upload video, audio, or text files                         |
| Get Results      | `/results`       | Displays all available analysis results (demo / list view) |
| Get Result by ID | `/results-by-id` | Fetch and display result using `analysis_id`               |

---

### Functional Capabilities

* Upload transcript files (`.mp4`, `.mp3`, `.wav`, `.txt`)
* Trigger backend analysis
* Display extracted:
  * People
  * Tools
  * Brands
  * Products
* Show tone and style classification
* Highlight sensitive domains and severity
* Display `created_at` timestamp
* Manual lookup via `analysis_id`
* Graceful error handling for invalid IDs and network failures

---

## 9.3 Frontend Project Structure

```
frontend/
│
├── src/
│   ├── pages/
│   │   ├── Home.tsx
│   │   ├── Upload.tsx
│   │   ├── Results.tsx
│   │   └── GetResultById.tsx
│   │
│   ├── App.tsx
│   ├── main.tsx
│   └── index.css
│
├── package.json
├── vite.config.ts
└── README.md
```

---

## 9.4 Frontend Prerequisites

Before running the frontend, ensure you have:

* **Node.js** (v18 or higher recommended)
* **npm** (comes with Node.js)
* Backend server running on `http://localhost:8080`

---

## 9.5 How to Run the Frontend

### Step 1: Navigate to frontend directory

```bash
cd frontend
```

### Step 2: Install dependencies

```bash
npm install
```

### Step 3: Start the development server

```bash
npm run dev
```

### Expected Output

```
VITE v5.x  ready in xxx ms

➜  Local:   http://localhost:5173/
```

Open the browser and navigate to:

```
http://localhost:5173
```

---

## 9.6 Frontend ↔ Backend Integration

| Frontend Action    | Backend API                                            |
| ------------------ | ------------------------------------------------------ |
| Upload transcript  | `POST /api/entity-classification/analyze`              |
| Fetch result by ID | `GET /api/entity-classification/results/{analysis_id}` |
| Health validation  | `GET /health`                                          |

> ⚠️ Ensure the backend server is running **before** starting the frontend.

---

## 9.7 Current Frontend Status

* ✅ UI pages implemented
* ✅ API integration working
* ✅ Table-based results view
* ✅ Error handling implemented
* ⏳ PostgreSQL-backed listing pending
* ⏳ Gemini-powered analysis pending

---

## 9.8 Known Limitations (Temporary)

* "Get Results" list page depends on backend list support
* Data persistence resets on backend restart (in-memory)
* File upload currently validates only client-side

These will be resolved after database integration.

---

## 10. Overall System Status

* Backend APIs: ✅ Complete (in-memory)
* Frontend UI: ✅ Complete
* Database: ⏳ Pending approval
* Python Worker (Gemini): ⏳ Pending

---

## 11. Demo Readiness

The system is **demo-ready** for:

* API validation
* Frontend interaction walkthrough
* End-to-end flow explanation
* Schema and architecture review

---

## 12. Architecture Overview

### System Flow Diagram

```
┌─────────────┐         ┌─────────────┐         ┌──────────────┐
│   Frontend  │────────▶│   Backend   │────────▶│  PostgreSQL  │
│  (React)    │  HTTP   │   (Go)      │  SQL    │  (Pending)   │
└─────────────┘         └─────────────┘         └──────────────┘
       │                       │
       │                       │
       │                       ▼
       │                ┌──────────────┐
       └───────────────▶│ Python Worker│
         (Future)        │  (Gemini AI) │
                        └──────────────┘
```

### Component Responsibilities

**Frontend (React + Vite)**
- User interface and interaction
- File upload handling
- API request management
- Results visualization
- Client-side validation

**Backend (Go + Gorilla Mux)**
- REST API endpoints
- Request validation
- Business logic orchestration
- Data persistence (in-memory → PostgreSQL)
- Response formatting

**Python Worker (Pending)**
- Gemini API integration
- NLP analysis
- Entity extraction
- Tone and style classification
- Safety content detection

**Database (PostgreSQL - Pending)**
- Persistent storage
- Query optimization
- Data integrity
- Indexing for performance

---

## 15. API Response Codes

| Status Code | Meaning                  | When It Occurs                        |
|-------------|--------------------------|---------------------------------------|
| 200         | OK                       | Successful GET request                |
| 201         | Created                  | Successful POST (analysis created)    |
| 400         | Bad Request              | Invalid JSON or missing required fields|
| 404         | Not Found                | Invalid analysis_id                   |
| 405         | Method Not Allowed       | Wrong HTTP method used                |
| 500         | Internal Server Error    | Backend processing error              |

---
