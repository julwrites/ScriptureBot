# Testing Strategy

This project employs a hybrid testing strategy to ensure code quality while minimizing external dependencies and costs.

## Test Categories

### 1. Unit Tests (Standard)
*   **Default Behavior:** By default, all tests run in "mock mode".
*   **Goal:** Fast, reliable, and cost-free verification of logic.
*   **Mechanism:** External services (Bible AI API, BibleGateway scraping) are mocked using function replacement (e.g., `SubmitQuery`, `GetPassageHTML`) or interface mocking.
*   **Execution:** these tests are run automatically on every Pull Request (MR).

### 2. Integration Tests (Live)
*   **Conditional Behavior:** Specific tests are capable of switching to "live mode" when appropriate environment variables are detected.
*   **Goal:** Verify that the application correctly interacts with real external services (Contract Testing) and that credentials/configurations are valid.
*   **Execution:** These tests should be run on a scheduled basis (e.g., nightly or weekly) or manually when verifying infrastructure changes.

## Live Tests & Configuration

The following tests support live execution:

### `TestSubmitQuery`
*   **File:** `pkg/app/api_client_test.go`
*   **Description:** Verifies connectivity to the Bible AI API.
*   **Trigger:**
    *   `BIBLE_API_URL` is set AND
    *   `BIBLE_API_URL` is NOT `https://example.com`
*   **Required Variables:**
    *   `BIBLE_API_URL`: The endpoint of the Bible AI API.
    *   `BIBLE_API_KEY`: A valid API key.
*   **Rationale:** Ensures that the client code (request marshaling, auth headers) matches the actual API expectation and that the API is reachable.

### `TestUserDatabaseIntegration`
*   **File:** `pkg/app/database_integration_test.go`
*   **Description:** Verifies Read/Write operations to Google Cloud Firestore/Datastore.
*   **Trigger:**
    *   `GCLOUD_PROJECT_ID` is set.
*   **Required Variables:**
    *   `GCLOUD_PROJECT_ID`: The Google Cloud Project ID.
    *   *Note:* Requires active Google Cloud credentials (e.g., `GOOGLE_APPLICATION_CREDENTIALS` or `gcloud auth`).
*   **Rationale:** Verifies that database permissions and client initialization are correct, preventing runtime errors in production. Uses a specific test user ID (`test-integration-user-DO-NOT-DELETE`) to avoid affecting real user data.

## Rationale for Strategy

1.  **Cost Reduction:** The Bible AI API may incur costs per call. Mocking prevents racking up bills during routine development.
2.  **Speed:** Live calls are slow. Mocked tests run instantly.
3.  **Reliability:** External services can be flaky. Mocked tests only fail if the code is broken.
4.  **Verification:** We still need to know if the API changed or if our secrets are wrong. The conditional integration tests provide this safety net without the daily cost/latency penalty.
