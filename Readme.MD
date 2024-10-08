
# Follow/Unfollow GraphQL Service

This service allows users to follow and unfollow each other. The service is exposed via a GraphQL API.

## Requirements
- Go 1.18+
- MySQL
- GraphQL

## Setup
1. Clone the repository:
```
git clone https://github.com/jayendrapawar/goGraphQL
cd goGraphQL
```

2. Set Env variables
```
export DB_USER=root
export DB_PASSWORD=1234567890
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=graphql
```
3. DB Setup
```
CREATE DATABASE mydb;
USE mydb;
SOURCE scripts/db_setup.sql;
```

4. Run Server
```
go run server/main.go

```
## DB Schema

### Users Table

| Column Name | Data Type      | Constraints                    |
|-------------|----------------|--------------------------------|
| `id`        | `INT`          | `PRIMARY KEY`, `AUTO_INCREMENT` |
| `name`      | `VARCHAR(255)` | `NOT NULL`                      |

### Follows Table

| Column Name  | Data Type | Constraints                                             |
|--------------|-----------|---------------------------------------------------------|
| `follower_id`| `INT`     | `FOREIGN KEY` (references `Users(id)` `ON DELETE CASCADE`), `PRIMARY KEY` |
| `following_id`| `INT`    | `FOREIGN KEY` (references `Users(id)` `ON DELETE CASCADE`), `PRIMARY KEY` |

## Relationships
```
   (1)----< (Follows) >----(M)
   Users                      Users

```


## Sample GraphQL Queries and Mutations

Goto Chrome http://localhost:8080/graphql

1. Follow a user:
```
mutation {
  follow(follower_id: "1", following_id: "2")
}
```
1. Unfollow a user:
```
mutation {
  unfollow(follower_id: "1", following_id: "2")
}
```

2. List followers
```
query {
  followers(user_id: "2") {
    id
    name
  }
}
```

3. List following
```
query {
  following(user_id: "1") {
    id
    name
  }
}
```
