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
- Change the Compose Path to `./docker-compose.yaml`.
- If you are using your own repository, [setup a webhook under Deployments](https://docs.dokploy.com/application/overview#github).
- Add your environment configuration under Environment, you can start by copying `.env.example`.
  - Set `TRAEFIK_HOST` to your domain. For example, `amth.one`.
  - Add a Discord Bot token and public key
  - Ensure the `DATABASE_DIR` and `DATABASE_NAME` are set correctly, this will be the path **on your host**. If this is misconfigured, you will lose data on restart.
  - Add Wargaming App IDs, not that one of them will be public.
  - If you are planning to track a lot of users, add proxies for Wargaming API in the following format: `user:password@host:port?wgAppId=your_app_id&maxRps=20`
  - Read through all variables prefixed with `INIT_`. Those will allow you to create admin user accounts and etc.
- Head over to General and click on Deploy.
- Once the app is deployed and everything looks good in Logs, set the INTERACTIONS ENDPOINT URL under General Information in Discord Application settings to `https://yourdomain/discord/callback`. For example, `https://amth.one/discord/callback`.
- You can start using Discord commands now!

## On Source Code and Licensing
Aftermath is a **source-available** project, **not** open source. Please note that this means you are not permitted to freely modify and redistribute any part of this codebase. For more information, refer to the specific licensing terms provided in the LICENSE file.