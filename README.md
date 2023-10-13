# udf-go

## About

This is a collection of User Defined Functions (UDF) written in Go.

More information about creating UDF, see official document:

- [MySQL](https://dev.mysql.com/doc/extending-mysql/8.0/en/adding-loadable-function.html)
- [MariaDB](https://mariadb.com/kb/en/creating-user-defined-functions/)

## Functions

- `udf_convert_kana` - Convert "kana" one from another ("zen-kaku", "han-kaku" and more).  
  This is inspired by [mb_convert_kana](https://www.php.net/manual/en/function.mb-convert-kana.php) function in PHP.

## Installation

### Build

To build needed C/C++ header files, you must include the header files in the build command.  
More information see #include directive in each `main.go` and each database's document.

For example, to build `udf_convert_kana` function for MySQL, run the following command:

```
CGO_ENABLED=1 CGO_CFLAGS="-O2 -g -I/usr/include/mysql" go build -buildmode=c-shared -o udf_convert_kana.so ./udf/mysql/udf_convert_kana
```

### Move so files

After building, you must move the so files to the plugin's directory.  
For MySQL, `/usr/lib/mysql/plugin` or `/usr/lib64/mysql/plugin` is the common directory.  
(`mysql_config --plugindir` command shows the directory where the plugin files should be stored.)

### Install

For example, to install `udf_convert_kana` function for MySQL, run the following command:

```
CREATE FUNCTION udf_convert_kana RETURNS STRING SONAME 'udf_convert_kana.so';
```
