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
* **Styling:** Custom CSS (dark theme)

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

## 12. Troubleshooting & Common Issues

### Backend Issues

**Issue:** Server won't start or port already in use

```bash
# Check if port 8080 is already in use
lsof -i :8080

# Kill the process using the port
kill -9 <PID>

# Or use a different port in .env
PORT=8081
```

**Issue:** Dependencies not installing

```bash
# Clear module cache
go clean -modcache

# Re-download dependencies
go mod download
go mod tidy
```

**Issue:** CORS errors from frontend

Ensure your backend has CORS middleware enabled for `http://localhost:5173`

---

### Frontend Issues

**Issue:** Cannot connect to backend

Check that:
- Backend is running on `http://localhost:8080`
- Health endpoint responds: `curl http://localhost:8080/health`
- Browser console shows correct API URL

**Issue:** Build fails or dependencies error

```bash
# Clear node_modules and reinstall
rm -rf node_modules package-lock.json
npm install

# Or use clean install
npm ci
```

**Issue:** Blank page or routing not working

Ensure React Router is properly configured in `App.tsx` and all routes are registered

---

## 13. Testing Guide

### Manual Testing Checklist

**Backend API Testing (via Postman/cURL)**

- [ ] Health check returns `OK`
- [ ] POST analyze endpoint accepts valid JSON
- [ ] POST returns valid `analysis_id`
- [ ] GET results endpoint retrieves stored data
- [ ] Invalid `analysis_id` returns 404
- [ ] Malformed JSON returns 400

**Frontend Testing**

- [ ] Home page loads correctly
- [ ] Navigation links work
- [ ] Upload page accepts files
- [ ] Results display in table format
- [ ] Error messages show for invalid IDs
- [ ] Loading states appear during API calls

### Sample cURL Commands

**Health Check:**
```bash
curl http://localhost:8080/health
```

**Analyze Transcript:**
```bash
curl -X POST http://localhost:8080/api/entity-classification/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "transcript_id": "11111111-1111-1111-1111-111111111111",
    "transcript_text": "Interview with Elon Musk about SpaceX and Tesla",
    "creator_id": "22222222-2222-2222-2222-222222222222"
  }'
```

**Get Result:**
```bash
curl http://localhost:8080/api/entity-classification/results/{analysis_id}
```

---

## 14. Architecture Overview

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

## 16. Development Workflow

### For Backend Development

1. Make code changes in appropriate directory
2. Run `go mod tidy` to update dependencies
3. Test locally with `go run ./cmd/server`
4. Validate with Postman/cURL
5. Check logs for errors
6. Commit changes with descriptive message

### For Frontend Development

1. Make changes in `src/` directory
2. Hot reload will show changes automatically
3. Test in browser
4. Check browser console for errors
5. Validate API integration
6. Build for production: `npm run build`

### Git Workflow

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Make changes and commit
git add .
git commit -m "feat: add your feature description"

# Push to remote
git push origin feature/your-feature-name

# Create pull request for review
```

---

## 17. Performance Considerations

### Current Limitations (In-Memory Storage)

- Data lost on server restart
- Not suitable for production
- Limited concurrent request handling
- No data persistence across instances

### Future Optimizations (Post-PostgreSQL)

- Connection pooling for database
- Caching layer for frequent queries
- Pagination for large result sets
- Background job processing for large files
- Rate limiting to prevent abuse

---

## 18. Security Considerations

### Current Implementation

- Basic input validation
- No authentication (demo only)
- CORS enabled for localhost
- No rate limiting

### Production Requirements (Future)

- JWT-based authentication
- API key validation
- Input sanitization
- Rate limiting per user
- HTTPS only
- SQL injection prevention
- XSS protection
- File upload size limits
- Virus scanning for uploads

---

## 19. Contributing Guidelines

We welcome contributions! Please follow these guidelines:

### Code Style

**Go Backend:**
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Add comments for complex logic
- Write unit tests for new features

**React Frontend:**
- Use TypeScript for type safety
- Follow React best practices
- Keep components small and focused
- Use functional components with hooks

### Pull Request Process

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit PR with clear description
6. Wait for code review

### Commit Message Format

```
feat: add new feature
fix: resolve bug
docs: update documentation
refactor: improve code structure
test: add tests
```

---

## 20. License & Contact

### License

This project is proprietary and confidential. All rights reserved.

### Contact Information

For questions or support:

- **Technical Lead:** [Your Name]
- **Email:** [your.email@company.com]
- **Slack Channel:** #entity-classification-service
- **Issue Tracker:** [GitHub Issues Link]

### Support

If you encounter issues:

1. Check the Troubleshooting section
2. Search existing GitHub issues
3. Create a new issue with:
   - Clear description
   - Steps to reproduce
   - Expected vs actual behavior
   - Environment details
   - Screenshots if applicable

---

## 21. Changelog

### Version 0.2.0 (Current)
- Added frontend React application
- Implemented file upload functionality
- Added results visualization
- Enhanced error handling

### Version 0.1.0
- Initial backend implementation
- Basic REST API endpoints
- In-memory storage
- Health check endpoint

---

## 22. Roadmap

### Phase 1 (Current - Demo)
- ✅ Backend API implementation
- ✅ Frontend UI development
- ✅ In-memory storage
- ✅ Basic validation

### Phase 2 (Database Integration)
- ⏳ PostgreSQL setup
- ⏳ Data persistence
- ⏳ Migration scripts
- ⏳ Connection pooling

### Phase 3 (AI Integration)
- ⏳ Python worker setup
- ⏳ Gemini API integration
- ⏳ Real entity extraction
- ⏳ ML-based classification

### Phase 4 (Production Ready)
- ⏳ Authentication & authorization
- ⏳ Rate limiting
- ⏳ Monitoring & logging
- ⏳ Performance optimization
- ⏳ Deployment pipeline

---

## 23. Acknowledgments

- **Gorilla Mux** for routing
- **Vite** for fast frontend development
- **React** for UI framework
- **Google Gemini** (upcoming AI integration)

---

**Last Updated:** January 18, 2026  
**Maintained By:** Development Team  
**Status:** Active Development