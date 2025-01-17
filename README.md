# File reader web service

A brief description of the project.

## Table of Contents

- [Description](#description)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Project Structure](#project-structure)

## Description

This project is a service that comunicates with several API's, retrevie data and saves it into a data base.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-repo/project.git
    cd project
    ```

2. Build the Docker containers:

    ```sh
    docker-compose build
    ```

## Configuration

Configure the environment variables by editing the [config.env](http://_vscodecontentref_/0) file.

## Usage

To start the application using Docker Compose, run:

```sh
docker-compose up -d
```
# Desafio Teorico

## Procesos, hilos y corrutinas

1. Un caso en el que usarías procesos para resolver un problema y por qué:
Utilizaria procesos para analizar el espectro de onda de una señal, para hacer calculos de color en una imagen, o para hacer operaciones griptograficas; esto por que los procesos tienen su propia memoria, y son ideales para tareas que consumen mucha CPU

2. Un caso en el que usarías threads para resolver un problema y por qué:
Utilizaria hilos para leer y escribir varios archivos a la vez, para descargar o subir varios archivos a un repositorio o para hacer varios calculos matematicos de uso intermedio.
Se manejarian de forma eficiente al compartir los recursos de un solo proceso y al no ser operaciones tan complejas es menos probable los hilos que se afecten entre ellos mismos 

3. Un caso en el que usarías corrutinas para resolver un problema y por qué:
Las corrutinas las utilizaria para manejar varias solicitudes de una API, administrar las conexiones a una base de datos, servidores o algun recurso especifico; es la opcion mas ligera de las 3 y es especialmente util para aplicaciones de entrada/salida no bloqueante

## Optimización de recursos del sistema operativo

1.  Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de ellos información en una API HTTP. ¿Cómo lo harías?
La mejor forma seria utilizar corrutinas o solicitudes en paralelo, para realizar varias llamadas a la API y obtener la informacion con mayor velocidad. En caso de que fuera muy pesado para la API recibir tanta informacion se puede enviar la informacion por lotes y finalmente en caso de que se ejecute en un entorno de servidor deberia de considerarse mas de una instancia para tener un balanceo de cargas eficiente.

## Análisis de complejidad

1. Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con complejidades O(n^2), O(n^3), O(2^n) y O(n log n), respectivamente, ¿Cuál de los algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.
Favoreceria el algoritmo O(n log n) ya que en terminos de Big O es medianamente eficiente y en caso de que tenga una buena optizacion resultaria util dependiendo del problema que se vaya a resolver.
El algoritmo que descartaria al principio seria el que tenga la complejidad O(2^n) ya que es sumamente lento e ineficiente; lo que demoraria a la operacion en terminos de tiempo y recursos de CPU

2. Asume que dispones de dos bases de datos para utilizar en diferentes problemas a resolver. La primera llamada AlfaDB tiene una complejidad de O(1) en consulta y O(n^2) en escritura. La segunda llamada BetaDB que tiene una complejidad de O(log n) tanto para consulta, como para escritura. ¿Describe en forma sucinta, qué casos de uso podrías atacar con cada una?
La base de datos AlfaDB podria utilizarse para almacenar datos historicos que se consultan con frecuencia. Por ejemplo un inventario donde no se guardan cosas nuevas muy a menudo pero la informacion se consulta de manera muy recurrente.
Para la base de datos BetaDB al ser una opcion más balanceada seria bastante util para usarse como una base de datos donde se procesa mucha informacion de gran volumen, como procesos ETL, a lo largo de pipelines en un ciclo de analisis de datos e incluso para vaciar los resultados de operaciones complejas como informacion de señales, imagenes y geolocalizacion por ejemplo.