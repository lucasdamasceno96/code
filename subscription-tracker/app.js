import express from 'express';
import { PORT } from './config/env.js';

import authRouter from './routes/auth.route.js';
import useRouter from './routes/user.route.js';
import subscriptionRouter from './routes/subscription.route.js';



const app = express();

app.use('/api/v1/auth', authRouter);
app.use('/api//v1/users', useRouter);
app.use('/api//v1/subscriptions', subscriptionRouter);

app.get('/', (req, res) => {
    res.send('Welcome to the Subscription Tracker API');
});

app.listen(3000, () => {
    console.log(`Subscription Tracker API is running on https://localhost:${PORT}`);
});

export default app;