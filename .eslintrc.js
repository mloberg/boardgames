module.exports = {
    env: {
        browser: true,
        es2021: true,
    },
    parserOptions: {
        ecmaVersion: 13,
        sourceType: 'module',
    },
    extends: [
        'eslint:recommended',
        'plugin:import/recommended',
        'plugin:prettier/recommended',
        'plugin:unicorn/recommended',
    ],
    rules: {
        'import/order': [
            'error',
            {
                alphabetize: { order: 'desc' },
                'newlines-between': 'always',
            },
        ],
        'linebreak-style': ['error', 'unix'],
        'no-console': 'error',
        'no-else-return': 'error',
        'no-unused-vars': ['error', { ignoreRestSiblings: true, argsIgnorePattern: '^_' }],
        'no-useless-return': 'error',
    },
};
