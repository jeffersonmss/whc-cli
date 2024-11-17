# WHC CLI
This project is a single cli to read the whc-sites.csv file and filter by state and category

## Usage

### Docker

Printing all sites:
```bash
docker container run jeffersonmss/whc
```

Filtering by Country
```bash
docker container run jeffersonmss/whc -s brazil
```

Filtering by category
```bash
docker container run jeffersonmss/whc -c Cultural
```

Filtering by Country and category
```bash
docker container run jeffersonmss/whc -s brazil -s cultural
```

The filter aren't case sensitive.
