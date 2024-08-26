# Aftermath
This is the most recent revision/rewrite of Aftermath. The main goal of this project is to simplify deployment and make it significantly cheaper to host.

## Deploying
Deploying Aftermath is relatively simple. You can build the binary manually, use Docker Compose, or build a Docker image.

### Docker Compose with [Dokploy](https://dokploy.com/)
- Get a VPS instalnce. Purchase a domain and point it to this instance.
- Create a new Bot Application in Discord
- [Install Dokploy and Docker](https://docs.dokploy.com/get-started/installation).
- Create a new Project and add a new Compose service.
- In Service settings > General > Provider, select your forked repository or Git > `git@github.com:Cufee/aftermath.git`. The branch should be `master`, but you can pick something different.
- Change the Compose Path to `./docker-compose.dokploy.yaml`.
- Add your environment configuration under Environment, you can start by copying `.env.example`.
  - Set `TRAEFIK_HOST` to your domain. For example, `amth.one`.
  - Add a Discord Bot token and public key
  - Ensure the `DATABASE_DIR` and `DATABASE_NAME` are set correctly, this will be the path **on your host**. If this is misconfigured, you will lose data on restart.
  - Add Wargaming App IDs, note that one of them will be public.
  - If you are planning to track a lot of users, add proxies for Wargaming API in the following format: `user:password@host:port?wgAppId=your_app_id&maxRps=20`
  - Read through all variables prefixed with `INIT_`. Those will allow you to create admin user accounts and etc.
- Head over to General and click on Deploy.
- You can start using Discord commands now!

### Locally with Docker Compose
- Setup a reverse proxy for your machine, or use something like [ngrok](https://ngrok.com/docs/getting-started/)
- Create a new Bot Application in Discord
- [Install Docker](https://docs.docker.com/engine/install/)
- Add your environment configuration under Environment, you can start by copying `.env.example`.
  - Add a Discord Bot token and public key
  - Ensure the `DATABASE_DIR` and `DATABASE_NAME` are set correctly, this will be the path **on your host**. If this is misconfigured, you will lose data on restart.
  - Add Wargaming App IDs, note that one of them will be public.
  - If you are planning to track a lot of users, add proxies for Wargaming API in the following format: `user:password@host:port?wgAppId=your_app_id&maxRps=20`
  - Read through all variables prefixed with `INIT_`. Those will allow you to create admin user accounts and etc.
- Start all services with `docker compose up -d`
- You can start using Discord commands now!

## Licensing
Aftermath is licensed under a dual licensing model:

1. **Open Source License**: The project is available under the GNU Affero General Public License version 3 (AGPL-3.0), making it free and open source for non-commercial and some commercial uses.

2. **Commercial License**: For uses not compatible with the AGPL-3.0, such as using the software for commercial hosting without disclosing source code, a separate commercial license is available.

This dual approach allows for open collaboration while also providing options for commercial use cases. For more details, please refer to the LICENSE.md file in this repository.

For commercial licensing inquiries, please contact license@byvko.dev.