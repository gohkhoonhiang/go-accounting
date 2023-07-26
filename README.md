# go-accounting
A personal/household accounting apis app built using Go

# APIs

## Budget

A model that represents one or more budget line. Each budget should have a `StartDate` and `EndDate` to indicate the budget period.
Optionally, the budget will have a `TotalBudgeted` amount that is the sum of all the budget lines. If one or more budget line is
allocated, then the `TotalAllocated` amount will be the sum of all allocated budget line amounts. `BalanceToAllocate` amount will
be the difference between `TotalBudgeted` and `TotalAllocated` amount.

### Get all budgets

Get all existing budgets.

```bash
GET /budgets
```

<details><summary>
Request Path Params
</summary>

N/A
</details>

<details><summary>
Request Body Params
</summary>

N/A
</details>

<details><summary>
Request Example
</summary>

```bash
curl http://localhost:9090/budgets
```
</details>

<details><summary>
Response Example
</summary>

```json
{
    "data": [
        {
            "id": 2,
            "startDate": "2023-01-01T08:00:00+08:00",
            "endDate": "2023-02-01T07:59:59+08:00",
            "totalBudgeted": 0,
            "totalAllocated": 0,
            "balanceToAllocate": 0,
            "createdAt": "2023-07-26T12:48:10.316942+08:00",
            "updatedAt": "2023-07-26T12:48:10.316942+08:00"
        }
    ]
}
```
</details>

### Create new budget

Create a new budget.

```bash
POST /budgets
```

<details><summary>
Request Path Params
</summary>

N/A
</details>

<details><summary>
Request Body Params
</summary>

| name | type | example |
| --- | --- | --- |
| startDate | timestamp | 2023-01-01T00:00:00Z |
| endDate | timestamp | 2023-01-01T00:00:00Z |
</details>

<details><summary>
Request Example
</summary>

```bash
curl -d '{"startDate":"2023-01-01T00:00:00Z","endDate":"2023-01-31T23:59:59Z"}' -H "Content-Type: application/json" -X POST http://localhost:9090/budgets
```
</details>

<details><summary>
Response Example
</summary>

```json
{
    "data": {
        "id": 2,
        "startDate": "2023-01-01T08:00:00+08:00",
        "endDate": "2023-02-01T07:59:59+08:00",
        "totalBudgeted": 0,
        "totalAllocated": 0,
        "balanceToAllocate": 0,
        "createdAt": "2023-07-26T12:48:10.316942+08:00",
        "updatedAt": "2023-07-26T12:48:10.316942+08:00"
    }
}
```
</details>

### Update existing budget

Update existing budget with `ID=id`.

```bash
PATCH /budgets/:id
```

<details><summary>
Request Path Params
</summary>

| name | type | example |
| --- | --- | --- |
| id | string | 3 |
</details>

<details><summary>
Request Body Params
</summary>

| name | type | example |
| --- | --- | --- |
| startDate | timestamp | 2023-01-01T00:00:00Z |
| endDate | timestamp | 2023-01-01T00:00:00Z |
| totalBudgeted | numeric | 4000 |
| totalAllocated | numeric | 2048 |
| balanceToAllocate | numeric | 1952 |
</details>

<details><summary>
Request Example
</summary>

```bash
curl -d '{"totalBudgeted":4000}' -H "Content-Type: application/json" -X PATCH http://localhost:9090/budgets/3
```
</details>

<details><summary>
Response Example
</summary>

```json
{
    "data": {
        "id": 3,
        "startDate": "2023-01-01T08:00:00+08:00",
        "endDate": "2023-02-01T07:59:59+08:00",
        "totalBudgeted": 4000,
        "totalAllocated": 0,
        "balanceToAllocate": 0,
        "createdAt": "2023-07-26T12:48:10.316942+08:00",
        "updatedAt": "2023-07-26T12:48:10.316942+08:00"
    }
}
```
</details>

### Delete existing budget

Delete existing budget with `ID=id`.

```bash
DELETE /budgets/:id
```

<details><summary>
Request Path Params
</summary>

| name | type | example |
| --- | --- | --- |
| id | string | 3 |
</details>

<details><summary>
Request Body Params
</summary>

N/A
</details>

<details><summary>
Request Example
</summary>

```bash
curl -X DELETE http://localhost:9090/budgets/1
```
</details>

<details><summary>
Response Example
</summary>

```json
{
    "data": true
}
```
</details>