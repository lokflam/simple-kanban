env "local" {
  url = "postgres://postgres@localhost:5432/postgres?search_path=public&sslmode=disable"
  dev = "docker://postgres/16/dev?search_path=public&sslmode=disable"
  src = "file://db/schemas"
  exclude = [
    "atlas_schema_revisions",
  ]
  migration {
    dir = "file://db/migrations"
  }
  diff {
    skip {
      drop_schema = true
      drop_table  = true
    }
    concurrent_index {
      create = true
      drop   = true
    }
  }
}