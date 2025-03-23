import mongoose from "mongoose";
import { DB_URI, NODE_ENV } from "../config/env.js";

if (!DB_URI) {
    throw new Error("DB_URI is not provided inside .env file");
}

const connectDB = async () => {
    try {
        await mongoose.connect(DB_URI);
        console.log(`✅ MongoDB connected in ${NODE_ENV} mode`);
    } catch (error) {
        console.error(`❌ MongoDB connection failed: ${error.message}`);
        process.exit(1); // Encerra o processo com erro
    } finally {
        console.log("🔄 Attempt to connect to MongoDB executed.");
    }
};

export default connectDB;
