# DAppNode bot for announcements

This bot will get all new packages and new versions from the Smart Contracts Repository and Registry correspondingly. When a new package or version is published it will detect and do an announcement int he announcement channel of the dappnode discord server

This bot will be subscribed to:

- Registry Smart Contracts (this is all the packages published): this suscription will return all new versions published
- Repository Smart Contract: this suscription will return all new packages published

# Development

To start the application is needed to configure the `test.env` file located at `/build/src/test.env`
The values needed are:

- ANNOUNCEMENTS_CHANNEL_ID= // A discord channel ID
- GETH_RPC= // An ethereum RPC with websockets
- DISCORD_TOKEN= // A discord bot token

**Build**
`docker-compose build`

**Up**
`docker-compose up`
