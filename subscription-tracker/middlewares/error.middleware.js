const errorMiddleware = (err, req, res, next) => {
    try {
        let error = { ...err, message: err.message };
        if (err.name === 'ValidationError') {
            error.statusCode = 400;
            error.message = Object.values(err.errors).map((value) => value.message).join(', ');
        }
        if (err.name === 'CastError') {
            error.statusCode = 400;
            error.message = `Invalid ${err.path}: ${err.value}`;
        }
        if (err.code === 11000) {
            error.statusCode = 400;
            error.message = `Duplicate field value entered`;
        }
        if (err.name === 'JsonWebTokenError') {
            error.statusCode = 401;
            error.message = 'Invalid token. Please log in again';
        }
        if (err.name === 'TokenExpiredError') {
            error.statusCode = 401;
            error.message = 'Your token has expired. Please log in again';
        }
        res.status(error.statusCode || 500).json({
            status: 'error',
            message: error.message || 'Internal Server Error',
        });
    } catch (error) {
        next(error);
    }

};

export default errorMiddleware;