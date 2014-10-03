CREATE TABLE repositories ( 
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  source VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  UNIQUE (source, name) ON CONFLICT REPLACE 
);

CREATE TABLE unique_partition_names ( 
  repository INTEGER,
  name VARCHAR(100),
  PRIMARY KEY(repository, name),
  FOREIGN KEY(repository) REFERENCES repositories(id)
);

CREATE TABLE range_partitions ( 
  repository INTEGER,
  name VARCHAR(100),
  PRIMARY KEY(repository, name),
  FOREIGN KEY(repository, name) REFERENCES unique_partition_names(repository, name)
);

CREATE TABLE set_partitions ( 
  repository INTEGER,
  name VARCHAR(100),
  value VARCHAR(255),
  PRIMARY KEY(repository, name),
  FOREIGN KEY(repository, name) REFERENCES unique_partition_names(repository, name)
);
