# Go REST API Application

## Introducción

Esta aplicación es un servidor API REST construido con Go (Golang) que gestiona usuarios y transacciones financieras. Permite a los usuarios cargar transacciones a través de archivos CSV, almacenar datos en una base de datos PostgreSQL, y generar resúmenes de transacciones para enviar por correo electrónico. La aplicación está diseñada para ser modular y extensible.

## Características

- **Carga de Transacciones**: Permite a los usuarios cargar transacciones a través de archivos CSV.
- **Gestión de Usuarios**: Guarda la información del usuario en una base de datos PostgreSQL.
- **Generación de Resúmenes**: Calcula resúmenes de transacciones y envía informes por correo electrónico.

## Requisitos

- **Go**: Esta aplicación está construida con Go 1.22.3. Asegúrate de tener Go instalado si deseas ejecutar el código directamente.

### Clonar el Repositorio

Primero, clona el repositorio:

```bash
git clone https://github.com/juanrod1028/Transactions
cd Transactions
```
- En el archivo main.go deveras reemplazar los campos vacios de ´COMPANY_EMAIL´ y ´COMPANY_EMAIL_PASS´
- Tambien deberas ajustar el host, port, user, password y bdname para hacer la conexion con postgres

Luego Ejecuta el make file
```bash
make run
