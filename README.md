## Golang-based Simple Note-taking API

This project is a simple note-taking API built using Golang. It supports CRUD operations for notes (Create, Read, Update, and Delete). The API is containerized using Docker and deployed using Kubernetes and Terraform.


## Features

- Basic CRUD operations for notes (Create, Read, Update, Delete)
- Each note contains an ID, title, and content
- In-memory data store using Go map
- Containerization using Docker
- Deployment with Kubernetes and Terraform

## Goals

The Golang-based Simple Note-taking API project aims to achieve the following goals:

- Develop practical experience with Golang, Docker, Kubernetes, and Terraform by developing a RESTful API.
- Understand CRUD operations for notes, each of which contains an ID, title, and content.
- Learn how to use an in-memory data store using a Go map for note storage, rather than needing to set up a complex database system.
- Practice containerization of an application using Docker, making it easier to package and deploy.
- Gain hands-on experience with managing the deployment of the application using Kubernetes and Terraform.
- Showcase proficiency in these technologies to potential employers.

### API Endpoints

The following API endpoints are available:

- Create a new note: `POST /notes`
- Get all notes: `GET /notes`
- Get a specific note by ID: `GET /notes/:id`
- Update a note: `PUT /notes/:id`
- Delete a note: `DELETE /notes/:id`

### Usage

Once the application is deployed, you can use tools like `curl` or Postman to interact with the API. Here are some examples using `curl`:

1. **Create a new note**:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "My Note", "content": "This is a test note."}' http://localhost/notes
```
2. **Get all notes**:

```bash
curl -X GET http://localhost/notes
```
3. Get a specific note by ID (replace NOTE_ID with the actual ID):
```bash
curl -X GET http://localhost/notes/NOTE_ID
```
4. Update a note (replace NOTE_ID with the actual ID):
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Note", "content": "This note has been updated."}' http://localhost/notes/NOTE_ID
```
5. Delete a note (replace NOTE_ID with the actual ID):
```bash
curl -X DELETE http://localhost/notes/NOTE_ID
```
Remember to replace NOTE_ID with the actual note ID in the API calls where it's required. By following these examples, users can interact with the API to create, read, update, and delete notes.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
