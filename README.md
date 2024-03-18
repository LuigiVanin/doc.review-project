## Migration and Database

To manage migrations, you'll need to install the migrate tool, [golang-migrate](https://github.com/golang-migrate/migrate/tree/master), either locally or globally. Refer to the [installation guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for detailed instructions.

If you choose to install it locally, you can execute commands directly from the [migration_tool.sh](./migration_tool.sh) file.

> ![NOTE] **Setting up the Database**
> This project utilizes a PostgreSQL database. Ensure you have PostgreSQL running before applying migrations. To configure the application for your database, create and customize a `.env` file based on the [.env.example](./.env.example) provided.

> ![WARNING] **Attention: Migration Setup**
> It's crucial to modify the PostgreSQL URL within the [migration_tool.sh](./migration_tool.sh) file (currently, this process requires manual intervention) at line 6.

### Creating a Migration

```bash
$ ./migration_tool.sh create 
```

### Running Migrations

```bash
$ ./migration_tool.sh run
```