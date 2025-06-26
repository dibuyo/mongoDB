const { MongoClient, ObjectId } = require('mongodb');
require("dotenv").config();

const env = process.env.NODE_ENV || 'development';
const mongoHost = process.env.MONGO_HOST || 'localhost';
const mongoPort = process.env.MONGO_PORT || '27017';
const mongoUser = process.env.MONGO_USER;
const mongoPassword = process.env.MONGO_PASSWORD;

const mongoCredential = mongoUser && mongoPassword ? `${mongoUser}:${mongoPassword}@` : '';

async function main() {
    const startTime = Date.now();
    console.log(`🔹 Proceso iniciado en ${startTime} ms`);

    const uri = `mongodb://${mongoCredential}${mongoHost}:${mongoPort}`;
    const client = new MongoClient(uri);

    try {
        await client.connect();
        console.log("✅ Conectado a MongoDB");

        const database = client.db('sample_db');
        const collection = database.collection('mycollection');

        // Crear un documento
        const doc = { name: "Martín Rivas", age: 42 };
        const insertResult = await collection.insertOne(doc);
        console.log(`📝 Documento insertado con _id: ${insertResult.insertedId}`);

        // Buscar por ID
        const findById = await collection.find({ _id: new ObjectId(`${insertResult.insertedId}`) }).next();
        console.log("🔍 Documento encontrado por ID:", findById);

        // Buscar todos los documentos
        const findResult = await collection.find({}).toArray();
        console.log("📄 Documentos encontrados:", findResult);

    } finally {
        await client.close();
        console.log("🔒 Conexión cerrada");

        const endTime = Date.now();
        console.log(`🔹 Proceso finalizado en ${endTime} ms`);
        console.log(`⏱️ Duración total: ${endTime - startTime} ms`);
    }
}

main().catch(console.error);