## Migration and Database

To manage migrations, you'll need to install the migrate tool, [golang-migrate](https://github.com/golang-migrate/migrate/tree/master), either locally or globally. Refer to the [installation guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for detailed instructions, you can install it globally or locally using the [curl method](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#download-pre-built-binary-windows-macos-or-linux).

If you choose to install it locally, you can execute commands directly from the [migration.sh](./migration.sh) file too.

> [!NOTE]
> **Setting up the Database:**
> This project utilizes a PostgreSQL database. Ensure you have PostgreSQL running before applying migrations. To configure the application for your database, create and customize a `.env` file based on the [.env.example](./.env.example) provided.

### Running Migrations

> [!WARNING]
> **Attention: Migration Setup:**
> It's crucial to modify the PostgreSQL URL within the `.env` file to run the migration commando on the correct database.

This will run all migration into the [database/migrations/](database/migrations/) folder, recreating the current state of the database. Run this command to build database schema ⬇️

```bash
$ ./migration.sh run
```

### Creating a Migration

If you desire to make modification into the database schema, run the command bellow ⬇️ and populate the `.up` and `.down` files accordingly inside of the [database/migrations/](database/migrations/) folder.

```bash
$ ./migration.sh create 
```
