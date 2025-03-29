import mongoose from 'mongoose';
import bcrypt from 'bcryptjs';
import User from '../models/user.model.js';
import jwt from 'jsonwebtoken';
import { JWT_SECRET, JWT_EXPIRES_IN } from '../config/env.js';


export const signUp = async (req, res, next) => {
    const session = await mongoose.startSession();
    session.startTransaction();

    try {
        const { name, email, password } = req.body;

        // Verifica se o usuário já existe
        const existingUser = await User.findOne({ email });
        if (existingUser) {
            session.endSession();
            return res.status(409).json({ message: 'User already exists' });
        }

        // Criptografa a senha
        const hashedPassword = await bcrypt.hash(password, 10);

        // Cria um novo usuário dentro da transação
        const newUser = await User.create([{ name, email, password: hashedPassword }], { session });

        // Gera um token JWT para autenticação
        const token = jwt.sign({ userID: newUser[0]._id }, JWT_SECRET, { expiresIn: JWT_EXPIRES_IN });

        // Confirma a transação e finaliza a sessão
        await session.commitTransaction();
        session.endSession();

        return res.status(201).json({
            message: 'User created successfully',
            success: true,
            data: {
                user: newUser[0],
                token,
            }
        });

    } catch (error) {
        await session.abortTransaction();
        session.endSession();
        next(error);
    }
};


export const signIn = async (req, res) => {
    try {
        // Verifica se o usuário existe 
        const { email, password } = req.body;

        const user = await User.findOne({ email });

        if (!user) {
            return res.status(404).json({ message: 'Invalid email or password' });
        }

        // Verifica se a senha está correta
        const isPasswordValid = await bcrypt.compare(password, user.password);
        if (!isPasswordValid) {
            return res.status(404).json({ message: 'Invalid email or password' });
        }
        // Gera um token JWT para autenticação
        const token = jwt.sign({ userID: user._id }, JWT_SECRET, { expiresIn: JWT_EXPIRES_IN });

        // Retorna o usuário e o token
        return res.status(200).json({
            message: 'User authenticated successfully',
            success: true,
            data: {
                user,
                token,
            }
        });

    } catch (error) {
        next(error);
    }


}

export const signOut = async (req, res) => {

}