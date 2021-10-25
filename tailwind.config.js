module.exports = {
    purge: [
        'layouts/**/*.html',
        'assets/**/*.js',
    ],
    darkMode: false, // or 'class' or 'media'
    theme: {
        extend: {
            zIndex: {
                '-1': '-1',
            },
        },
    },
    variants: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
    ],
};
