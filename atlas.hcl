variable "database_name" {
  type    = string
  default = getenv("DATABASE_NAME")
}
variable "database_user" {
  type    = string
  default = getenv("DATABASE_USER")
}
variable "database_password" {
  type    = string
  default = getenv("DATABASE_PASSWORD")
}
variable "database_host" {
  type    = string
  default = getenv("DATABASE_HOST")
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
    dir = "file://internal/database/migrations"
  }

  url = "postgresql://${var.database_user}:${var.database_password}@${var.database_host}/${var.database_name}?sslmode=disable"
  dev = "docker://postgres/17/dev?search_path=public"
}

env "migrate" {
  src = [
    "file:///schema/schema.hcl",
    "file:///schema/accounts.hcl",
    "file:///schema/app.hcl",
    "file:///schema/auth.hcl",
    "file:///schema/discord.hcl",
    "file:///schema/glossary.hcl",
    "file:///schema/users.hcl",
  ]

  migration {
    dir = "file:///migrations"
  }
  tx-mode = "all"

  url = "postgresql://${var.database_user}:${var.database_password}@${var.database_host}/${var.database_name}?sslmode=disable"
}