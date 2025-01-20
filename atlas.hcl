variable "database_name" {
  type    = string
  default = getenv("DATABASE_NAME")
}
variable "database_path" {
  type    = string
  default = getenv("DATABASE_PATH")
}
variable "sources" {
  type = list(string)
  default = [
    "file://internal/database/schema/schema.hcl",
    "file://internal/database/schema/accounts.hcl",
    "file://internal/database/schema/app.hcl",
    "file://internal/database/schema/auth.hcl",
    "file://internal/database/schema/discord.hcl",
    "file://internal/database/schema/glossary.hcl",
    "file://internal/database/schema/users.hcl",
  ]
}


env "local" {
  src = var.sources

  migration {
    dir = "file://internal/database/migrations?format=golang-migrate"
  }

  url = "sqlite://${var.database_path}/${var.database_name}?_fk=1"
  dev = "sqlite://file?mode=memory&_fk=1"
}

env "migrate" {
  src = var.sources

  migration {
    dir = "file:///migrations"
  }
  tx-mode = "all"

  url = "sqlite:///data/${var.database_name}?_fk=1"
}