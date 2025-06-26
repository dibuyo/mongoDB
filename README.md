# mongoDB

Los ejemplos provisto por MongoDB Inc. los podemos encontrar en los siguientes Links Oficiales:
- https://atlas-education.s3.amazonaws.com/sampledata.archive
- https://download.restheart.com/sampledata.archive

Puedes utilizar el siguiente comando para descargarlo.
```console
curl https://atlas-education.s3.amazonaws.com/sampledata.archive -o examples/sampledata.archive
```

Para poder realizar un restore sobre la base de datos local utilizar el siguiente comando (sin autenticación)
```console
mongorestore --archive=sampledata.archive
```

Con Autenticación
```console
mongorestore --archive=sampledata.archive --username ${MONGO_USER} --password ${MONGO_PASSWORD}
```
> [!IMPORTANT]
> Asegúrese de que MongoDB se esté ejecutando en localhost:27017 (este comando asume que no se está ejecutando con autenticación requerida).

## Explorador

Para poder navegar las bases de datos ingresar a la siguiente URL http://localhost:8081/

## Software Recomendado

[Atlas Cli](https://www.mongodb.com/try/download/atlascli)

Opciones pagas:
[DataGrip - Jetbrains](https://www.jetbrains.com/datagrip/features/mongodb/)

## Mostrar y Filtros colleciones

Se pueden realizar filtros simples utilizando la función find() llamando a la collección.
```console
db.getCollection('users').find();
db.users.find();
```
> [!NOTE]
> db.getCollection es Útil si el nombre de la colección contiene caracteres especiales, espacios, o empieza con un número.

Ejemplo de find:
```console
db.getCollection('users').find({
    /*
    * Filtro
    * campoA: valor o expresión
      */
      },
      {
      /*
    * Proyección
    * _id: 0, // excluir _id
    * campoA: 1 // incluir campo
      */
      }
      )
      .sort({
      /*
    * campoA: 1 // ascendente
    * campoB: -1 // descendente
      */
     });
```

## Filtros y Proyecciones

<details>
    <summary>Filtros Simples</summary>
    
    Filtros con operadores logicos mayor que ( > $gt)
    ```console
       db.accounts.find({ account_id: { $gt: 51080 } });
    ```
    
    Filtros con operadores logicos mayor que ( >= $gte)
    ```console
       db.accounts.find({ account_id: { $gte: 51080 } });
    ```
    
    Filtro con IN
    ```console
       db.accounts.find({ products: {$in: [ "CurrencyService" ] } });
    ```
</details>

<details>
    <summary>Limitar Resultado</summary>

    ### Limitar el resultado

    Se utiliza la función Limit

    Como por ejemplo Limitar el resultado a 20 documentos.
    ```console
    db.accounts.find().limit(20);
    ```

    > [!IMPORTANT]
    > Completar con un número positivo
</details>

<details>
    <summary>### Ordenamiento</summary>

    Para poder ordenar el resultado de documentos

    
    ```javascript
        { $sort: { <field1>: <sort order>, <field2>: <sort order> ... } }
    ```

    Ordenar de forma creciente o ascendente por numero de cuenta.
    ```console
    db.accounts.find().sort({account_id: 1});
    ```
</details>

<details>
    <summary>Contar Registros</summary>

    ### Contar cantidad de Registros

    El count() está deprecado en las versiones más recientes de MongoDB, incluyendo MongoDB 5.0+ y MongoDB 8.0.

    Para poder contar con Filtros.
    ```console
    db.accounts.countDocuments({ products: {$in: [ "CurrencyService" ] } });
    ```

    Para poder contar todos los documentos sin importar el filtro.
    ```console
    db.accounts.estimatedDocumentCount();
    ```
</details>