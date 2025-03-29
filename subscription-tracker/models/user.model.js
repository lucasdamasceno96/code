import mongoose from 'mongoose';

const UserSchema = new mongoose.Schema({
    name: {
        type: String,
        required: [true, 'Please enter your name'],
        trim: true,
        minLength: 2,
        maxLength: 50,
    },
    email: {
        type: String,
        required: [true, 'Please enter your email'],
        trim: true,
        unique: true,
        lowercase: true,
        minLength: 6,
        maxLength: 50,
        match: [
            /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/,
            'Please enter a valid email',
        ]
    },
    password: {
        type: String,
        required: [true, 'Please enter your password'],
        minLength: [8, 'Password must be at least 8 characters long'],

    },
}, { timestamps: true });

const User = mongoose.model('User', UserSchema);

export default User;
