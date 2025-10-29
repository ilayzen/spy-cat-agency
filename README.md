# spy-cat-agency

to start the program use 
```
make up
```

✅ Ability to create a spy cat in the system

✅ A cat is described as Name, Years of Experience, Breed, and Salary

✅ Breed must be validated (via TheCatAPI cache)

✅ Ability to remove spy cats from the system

✅ Ability to update spy cats’ information 
(Salary)

✅ Ability to list spy cats

✅ Ability to get a single spy cat

✅ Middeware implemented
Missions / Targets

✅ Ability to create a mission in the system along with targets

✅ A target is described as Name, Country, Notes, and Complete state

✅ Ability to delete a mission

✅ A mission contains information about Cat, Targets, and Complete state

✅ Each target is unique to a mission, so the endpoint accepts an object describing targets

✅ A mission cannot be deleted if it is already assigned to a cat

✅ Ability to update mission

✅ Ability to mark it as completed

✅ Ability to update mission targets
✅ Ability to mark them as completed

❌ Ability to update Notes

❌ Notes cannot be updated if either the target 
or the mission is completed

❌ Ability to delete targets from an existing 
mission

❌ A target cannot be deleted if it is already 
completed

❌ Ability to add targets to an existing mission

❌ A target cannot be added if the mission is already completed

❌ Ability to assign a cat to a mission

❌ Ability to list missions

❌ Ability to get a single mission

❌ No validation fields 

# 

Base URL:

```
/api/v1
```

---

### `GET /cats`

Retrieve a list of all spy cats.

**Response 200**

```json
[
  {
    "id": 5,
    "name": "Whiskers",
    "years_experience": 5,
    "breed": "Siamese",
    "salary": 120000,
    "created_at": "2025-10-29T10:00:00Z"
  }
]

```

---

### `POST /cats`

Create a new spy cat.

> This endpoint passes through the BreedCacheMiddleware which validates the cat’s breed via TheCatAPI.
> 

**Request**

```json
{
  "name": "Whiskers",
  "years_experience": 5,
  "breed": "Siamese",
  "salary": 120000
}

```

**Response 200** — empty body (can be extended to return created cat object).

---

### `GET /cats/{id}`

Retrieve a single cat by ID.

**Response 200**

```json
{
  "id": 5,
  "name": "Whiskers",
  "years_experience": 5,
  "breed": "Siamese",
  "salary": 120000,
  "created_at": "2025-10-29T10:00:00Z"
}

```

---

### `DELETE /cats/{id}`

Delete a spy cat.

**Response 200** — empty body.

---

### `PUT /cats/{id}`

Update a cat’s information (currently only salary).

**Request**

```json
{
  "salary": 130000
}

```

**Response 200** — empty body.

---

## Missions

### `GET /missions`

Retrieve all missions with their related cats and targets.

**Response 200**

```json
[
  {
    "mission": {
      "id": 10,
      "cat_id": 5,
      "completed": false,
      "completed_at": null,
      "created_at": "2025-10-29T10:00:00Z"
    },
    "cat": {
      "id": 5,
      "name": "Whiskers",
      "years_experience": 5,
      "breed": "Siamese",
      "salary": 120000,
      "created_at": "2025-10-01T12:00:00Z"
    },
    "targets": [
      {
        "id": 101,
        "mission_id": 10,
        "name": "Retrieve secret cheese",
        "country": "France",
        "notes": "Guarded by dogs",
        "completed": false,
        "completed_at": null,
        "created_at": "2025-10-29T10:30:00Z"
      }
    ]
  }
]

```

---

### `POST /missions`

Create a mission with one or more targets.

**Request**

```json
{
  "mission": {
    "cat_id": 5,
    "completed": false,
    "completed_at": null},
  "targets": [
    {
      "name": "Retrieve secret cheese",
      "country": "France",
      "notes": "Guarded by dogs, proceed at night.",
      "completed": false}
  ]
}

```

**Response 200** — empty body.

---

### `GET /missions/{id}`

Retrieve a specific mission with its cat and targets.

**Response 200**

```json
{
  "mission": {
    "id": 10,
    "cat_id": 5,
    "completed": false,
    "completed_at": null,
    "created_at": "2025-10-29T10:00:00Z"
  },
  "cat": {
    "id": 5,
    "name": "Whiskers",
    "years_experience": 5,
    "breed": "Siamese",
    "salary": 120000,
    "created_at": "2025-10-01T12:00:00Z"
  },
  "targets": [
    {
      "id": 101,
      "mission_id": 10,
      "name": "Retrieve secret cheese",
      "country": "France",
      "notes": "Guarded by dogs",
      "completed": false,
      "completed_at": null,
      "created_at": "2025-10-29T10:30:00Z"
    }
  ]
}

```

---

### `DELETE /missions/{id}`

Delete a mission.

**Response 200** — empty body.

---

### `PUT /missions/{id}`

Update a mission (e.g., mark it as completed).

**Request**

```json
{
  "cat_id": 5,
  "completed": true,
  "completed_at": "2025-10-29T12:00:00Z"
}

```

**Response 200** — empty body.

---

## Targets (Mission-specific)

### `POST /missions/{id}/targets`

Add a new target to an existing mission.

**Request**

```json
{
  "name": "Exfiltrate blueprints",
  "country": "Germany",
  "notes": "Use courier route B",
  "completed": false}

```

**Response 200** — empty body.

---

### `PUT /missions/{id}/targets/{targetId}`

Update an existing target.

**Request**

```json
{
  "name": "Exfiltrate blueprints",
  "country": "Germany",
  "notes": "Switch to route C",
  "completed": true,
  "completed_at": "2025-10-29T13:00:00Z"
}

```

**Response 200** — empty body.

---

### `DELETE /missions/{id}/targets/{targetId}`

Delete a target from a mission.

**Response 200** — empty body.

---

## Assignment

### `POST /missions/{id}/assign`

Assign a cat to a mission.

**Request**

```json
{
  "cat_id": 5
}

```

**Response 200** — empty body.


todo:

add graceful shutdown

add migrations down

refactor transactions

Slightly improve the cache with middleware
