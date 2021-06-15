import http from './httpService';

const  apiEndpointUsers = '/users';
const  apiEndpointAuth = '/auth';

export function register(user) {
    return http.post(apiEndpointUsers, {
        username: user.username,
        password: user.password
    });
}

export function login(username, password) {
    return http.post(apiEndpointAuth, { username, password });
}

