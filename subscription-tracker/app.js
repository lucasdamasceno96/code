import express from 'express';
import { PORT } from './config/env.js';


const app = express();

app.get('/', (req, res) => {
    res.send('Welcome to the Subscription Tracker API');
});

app.listen(3000, () => {
    console.log(`Subscription Tracker API is running on https://localhost:${PORT}`);
});

export default app;