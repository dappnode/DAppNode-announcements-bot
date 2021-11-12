# DAppNode bot for announcements

This bot will get all new packages and new versions from the Smart Contracts Repository and Registry correspondingly. When a new package or version is published it will detect and do an announcement int he announcement channel of the dappnode discord server

This bot will be subscribed to:

- Registry Smart Contracts (this is all the packages published): this suscription will return all new versions published
- Repository Smart Contract: this suscription will return all new packages published

# ENV

The application has two modes deppending on the `GO_ENV` that can be set on the `docker-compose.yml`

- Development: `GO_ENV=development`
- Production: `GO_ENV=production`

## Development

To start the application in development mode is needed to configure the `test.env` file located at `/build/src/test.env`
The values needed are:

- ANNOUNCEMENTS_CHANNEL_ID= // A discord channel ID
- GETH_RPC= // An ethereum RPC with websockets
- DISCORD_TOKEN= // A discord bot token

## Production

To start the application in production mode is needed to configure a `.env` file at `/build/src/test.env`
Same values as in the test file

**Build**
`docker-compose build`

**Up**
`docker-compose up`
