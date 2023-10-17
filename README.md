# udf-go

## About

This is a collection of User Defined Functions (UDF) written in Go.

More information about creating UDF, see official document:

- [MySQL](https://dev.mysql.com/doc/extending-mysql/8.0/en/adding-loadable-function.html)
- [MariaDB](https://mariadb.com/kb/en/creating-user-defined-functions/)
- [PostgreSQL](https://www.postgresql.org/docs/current/xfunc-c.html)

## Functions

- `udf_convert_kana` - Convert "kana" one from another ("zen-kaku", "han-kaku" and more) for UTF-8.  
  This is inspired by [mb_convert_kana](https://www.php.net/manual/en/function.mb-convert-kana.php) function in PHP.

## Installation

### Build

To build needed C/C++ header files, you must include the header files in the build command.  
More information see #include directive in each `main.go` and each database's document.

For example, to build `udf_convert_kana` function for MySQL, run the following command:

```
CGO_ENABLED=1 CGO_CFLAGS="-O2 -g -I/usr/include/mysql" go build -buildmode=c-shared -o udf_convert_kana.so ./udf/mysql/udf_convert_kana
```

For example, to build `udf_convert_kana` function for PostgreSQL, run the following command:

```
CGO_ENABLED=1 CGO_CFLAGS="-O2 -g -I/usr/include/postgresql/${MAJOR_VERSION}/server" go build -buildmode=c-shared -o udf_convert_kana.so ./udf/postgres/udf_convert_kana
```

### Move so files

After building, you must move the so files to the plugin's directory.  
For MySQL, `mysql_config --plugindir` command shows the directory where the plugin files should be stored.  
(For example, `/usr/lib/mysql/plugin` or `/usr/lib64/mysql/plugin` is the common directory.)

For PostgreSQL, `pg_config --pkglibdir` command shows the directory where the plugin files should be stored.  
(For example, `/usr/lib/postgresql/${MAJOR_VERSION}/lib` or `/usr/pgsql-${MAJOR_VERSION}/lib` is the common directory.)

### Install

For example, to install `udf_convert_kana` function for MySQL, run the following command:

```
CREATE FUNCTION udf_convert_kana RETURNS STRING SONAME 'udf_convert_kana.so';
```

For example, to install `udf_convert_kana` function for PostgreSQL, run the following command:

```
CREATE FUNCTION udf_convert_kana(text, text) RETURNS text
  AS '/usr/lib/postgresql/11/lib/udf_convert_kana', 'udf_convert_kana'
  LANGUAGE C STRICT;
```
