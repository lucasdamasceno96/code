import { Router } from 'express';

const subscriptionRouter = Router();

subscriptionRouter.get('/', (req, res) => { title: 'GET all subscriptions' });

subscriptionRouter.get('/:id', (req, res) => { title: 'GET subscription by id' });

subscriptionRouter.post('/', (req, res) => { title: 'CREATE new subscription' });

subscriptionRouter.put('/:id', (req, res) => { title: 'UPDATE subscription by id' });

subscriptionRouter.delete('/:id', (req, res) => { title: 'DELETE subscription by id' });

subscriptionRouter.get('/user/:id', (req, res) => { title: 'GET all user subscriptions by user id' });

subscriptionRouter.put('/:id/cancel', (req, res) => { title: 'CANCEL subscription by id' });

subscriptionRouter.get('/:id/upcoming-renewals', (req, res) => { title: 'GET upcoming  renewsals ' });



export default subscriptionRouter;