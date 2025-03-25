import express from 'express';
import { PORT } from './config/env.js';

import authRouter from './routes/auth.route.js';
import useRouter from './routes/user.route.js';
import subscriptionRouter from './routes/subscription.route.js';
import connectDB from './database/mongodb.js';
import errorMiddleware from './middlewares/error.mindleware.js';



const app = express();

app.use('/api/v1/auth', authRouter);
app.use('/api//v1/users', useRouter);
app.use('/api//v1/subscriptions', subscriptionRouter);

app.use(errorMiddleware);

app.get('/', (req, res) => {
    res.send('Welcome to the Subscription Tracker API');
});

app.listen(PORT, async () => {
    console.log(`Subscription Tracker API is running on https://localhost:${PORT}`);

    await connectDB();
});

export default app;