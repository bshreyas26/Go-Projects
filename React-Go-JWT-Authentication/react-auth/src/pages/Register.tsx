import React from 'react';
import { useNavigate } from 'react-router-dom';
import config from '../config';

const Register = () => {
    const [name, setName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!email.endsWith('@gmail.com')) {
            alert('Please use a Gmail address (e.g., example@gmail.com).');
            return;
        }

        const response = await fetch(`${config.API_BASE_URL}/api/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name,
                email,
                password,
            }),
        });

        const data = await response.json();
        if (response.ok) {
            alert('Registration successful!');
            navigate('/login');
        } else {
            console.log(data);
            alert(`Registration failed: ${data.error}`);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <h1 className="h3 mb-3 fw-normal">Please Register</h1>
            <input
                type="text"
                className="form-control"
                placeholder="Name"
                required
                onChange={e => setName(e.target.value)}
            />
            <input
                type="email"
                className="form-control"
                placeholder="Email Address"
                required
                onChange={e => setEmail(e.target.value)}
            />
            <input
                type="password"
                className="form-control"
                placeholder="Password"
                required
                onChange={e => setPassword(e.target.value)}
            />
            <button className="btn btn-primary w-100 py-2" type="submit">
                Submit
            </button>
        </form>
    );
};

export default Register;
