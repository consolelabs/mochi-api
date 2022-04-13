# Mochi Bot API

#### Setup dev

```bash
# install required cli and clone the .env
make setup

# reset and seed db
reset-db
```

### Migration

```bash
# create new migration
make migrate-new name=create-new-table

# up
make migration-up

# down
make migration-down
```

### Start dev

```bash
make dev
```

### Test

```bash
make test
```
