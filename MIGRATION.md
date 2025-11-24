# ScriptureBot Migration Guide

This guide documents the process for migrating ScriptureBot from an existing Google Cloud Project (Source) to a new Google Cloud Project (Destination), specifically targeting the `asia-southeast1` region.

## 1. Prerequisites

Before starting the migration, ensure the following are set up in the **Destination Project**:

### Google Cloud Setup
*   **Project Created**: You should have the new Project ID ready.
*   **APIs Enabled**:
    *   Cloud Run Admin API
    *   Secret Manager API
    *   Artifact Registry API
    *   Cloud Datastore API (Firestore in Datastore mode)
*   **Artifact Registry**:
    *   Create a Docker repository in `asia-southeast1`.
    *   Note the Repository Name (e.g., `scripturebot-repo`).
*   **Service Accounts**:
    *   **CI/CD Agent**: A service account for GitHub Actions (e.g., `cicd-agent@<project-id>.iam.gserviceaccount.com`).
        *   Grant `Cloud Run Admin`, `Service Account User`, `Artifact Registry Writer`.
        *   **Create Key**: Generate a new JSON key for this service account and download it.
    *   **Compute Service Account**: The default compute service account (or a custom one) for Cloud Run.
        *   Grant `Datastore User` (or `Cloud Datastore User`) and `Secret Manager Secret Accessor`.

### GitHub Secrets
Update the following secrets in the GitHub Repository settings:

*   `GCLOUD_PROJECT_ID`: The new Project ID.
*   `GCLOUD_SA_KEY`: The content of the JSON key file you downloaded for the CI/CD Agent.
*   `GCLOUD_ARTIFACT_REPOSITORY_ID`: The name of the repository created in Artifact Registry.
*   `TELEGRAM_ID`: The Telegram Bot Token (ensure it matches the one used in the source project if preserving identity).
*   `TELEGRAM_ADMIN_ID`: Your Telegram User ID.

## 2. Data Migration

A migration tool has been added to `cmd/migrate/main.go` to assist with transferring User data. This tool avoids complex cross-project permissions by running locally with your admin credentials.

### Step 2.1: authenticate locally
Ensure you are authenticated with `gcloud` and have access to **both** projects.

```bash
gcloud auth login
gcloud auth application-default login
```

### Step 2.2: Export Data (Backup)
Run the export command against the **Source Project**.

```bash
go run cmd/migrate/main.go -mode export -project <SOURCE_PROJECT_ID> -file backup_users.json
```
*   This will fetch all users and save them to `backup_users.json`.
*   Review the file to ensure data looks correct.

### Step 2.3: Import Data
Run the import command against the **Destination Project**.

```bash
go run cmd/migrate/main.go -mode import -project <DESTINATION_PROJECT_ID> -file backup_users.json
```
*   This will read `backup_users.json` and upload each user to the new Datastore.
*   **Note**: Ensure the Destination Project has Datastore (Firestore in Datastore mode) enabled.

## 3. Deployment

The `deployment.yml` workflow has been updated to deploy to `asia-southeast1` using Service Account Key authentication.

1.  Push the changes to the `master` branch.
2.  Monitor the "Build, Stage and Deploy Automation" action in GitHub.
3.  Once successful, verify the new Cloud Run service URL.

## 4. Verification

1.  Open the new Cloud Run service logs.
2.  Send a message to the bot (e.g., `/start` or `/verse`).
3.  Verify that the bot responds and no database errors occur in the logs.
4.  If you migrated data, check if your user preferences (e.g., translation version) are preserved.
