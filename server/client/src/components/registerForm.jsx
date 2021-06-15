import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import * as userService from '../services/userServise'

class RegisterForm extends Form {
    state = {
        data: {username: '', password: ''},
        errors: {}
    };

    schema = {
        username: Joi.string()
            .required()
            .label('Username'),
        password: Joi.string()
            .required()
            .min(5)
            .label('Password'),
    };

    doSubmit = async () => {
        try {
            const response = await userService.register(this.state.data);
            console.log(response);
            localStorage.setItem("token", response.headers['authorization'])
            this.props.history.push('/');
        }
        catch (ex) {
            if (ex.response && ex.response.status === 422) {
               const errors = {...this.state.errors};
               errors.username = ex.response.data;
               this.setState({errors});
            }
        }
    };

    render() {
        return (
            <div className="container">
                <h1 className="title center-text">Register</h1>
                <form onSubmit={this.handleSubmit}>
                    {this.renderInput('username', 'Username')}
                    {this.renderInput('password', 'Password', 'password')}
                    {this.renderButton("Register")}
                </form>
            </div>
        );
    }
}

export default RegisterForm;