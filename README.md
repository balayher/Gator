# Gator

A command line tool for aggregating RSS feeds and viewing their posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `Gator` with:

```bash
go install github.com/balayher/Gator@latest
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
Gator register <name>
```

Add a feed:

```bash
Gator addfeed <url>
```

Start the aggregator:

```bash
Gator agg 30s
```

View the posts:

```bash
Gator browse [limit]
```

There are a few other commands you'll need as well:

- `Gator login <name>` - Log in as a user that already exists
- `Gator users` - List all users
- `Gator feeds` - List all feeds
- `Gator follow <url>` - Follow a feed that already exists in the database
- `Gator unfollow <url>` - Unfollow a feed that already exists in the database
